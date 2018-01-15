package cli

import (
	"context"
	"crypto/ecdsa"
	"database/sql"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zapgrpc"

	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
	gRPC "google.golang.org/grpc"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	// register sqlite3 driver
	_ "github.com/mattn/go-sqlite3"

	"github.com/sonm-io/marketplace/infra/cqrs"
	"github.com/sonm-io/marketplace/infra/grpc"
	"github.com/sonm-io/marketplace/infra/grpc/interceptor"
	infraStorage "github.com/sonm-io/marketplace/infra/storage/sqllite"
	"github.com/sonm-io/marketplace/infra/util"

	"github.com/sonm-io/marketplace/interface/adaptor"
	"github.com/sonm-io/marketplace/interface/grpc/srv"
	"github.com/sonm-io/marketplace/interface/storage/sqllite"

	report "github.com/sonm-io/marketplace/interface/reporting/sqllite"
	"github.com/sonm-io/marketplace/usecase/marketplace/command"
)

// App application root.
type App struct {
	sync.RWMutex

	db     *sql.DB
	conf   *Config
	logger *zap.Logger

	privateKey *ecdsa.PrivateKey
	server     *gRPC.Server
	creds      credentials.TransportCredentials
	rotator    util.HitlessCertRotator

	schedulerQuitCh chan bool
}

// NewApp creates a new App instance.
func NewApp(opts ...Option) *App {
	return &App{conf: NewConfig(opts...)}
}

// Init initialize the application.
func (a *App) Init() error {
	if err := a.initConfig(); err != nil {
		return err
	}

	if err := a.initLogger(); err != nil {
		return err
	}

	a.logger.Info("Config loaded form", zap.String("path", a.conf.CfgPath))
	a.logger.Info("Public key", zap.String("address", util.PubKeyToAddr(a.privateKey.PublicKey).Hex()))

	if err := a.initStorage(); err != nil {
		return err
	}

	db := infraStorage.NewStorage(a.db)
	repo := sqllite.NewOrderStorage(db)

	getOrderHandler := report.NewOrderByIDHandler(db)
	getOrdersHandler := report.NewMatchOrdersHandler(db)

	createBidOrderHandler := command.NewCreateBidOrderHandler(repo)
	createAskOrderHandler := command.NewCreateAskOrderHandler(repo)
	cancelOrderHandler := command.NewCancelOrderHandler(repo)

	commandBus := cqrs.NewCommandBus()
	commandBus.RegisterHandler("CreateBidOrder", adaptor.FromDomain(createBidOrderHandler))
	commandBus.RegisterHandler("CreateAskOrder", adaptor.FromDomain(createAskOrderHandler))
	commandBus.RegisterHandler("CancelOrder", adaptor.FromDomain(cancelOrderHandler))

	return a.initServer(
		srv.NewMarketplace(adaptor.ToDomain(commandBus), getOrderHandler, getOrdersHandler))
}

func (a *App) initConfig() error {
	a.Lock()
	defer a.Unlock()

	if a.conf.CfgPath == "" {
		absCfgPath, err := filepath.Abs(filepath.Dir(os.Args[0]) + "/etc/market.yaml")
		if err != nil {
			return fmt.Errorf("cannot get absolute path of config file: %v", err)
		}
		a.conf.CfgPath = absCfgPath
	}

	cfgFileExists, err := a.pathExists(a.conf.CfgPath)
	if err != nil {
		return fmt.Errorf("cannot check if config file exists: %v", err)
	}

	if !cfgFileExists {
		return fmt.Errorf("config file %q does not exist", a.conf.CfgPath)
	}
	a.conf.FromFile(a.conf.CfgPath)

	key, err := a.conf.EthCfg.LoadKey()
	if err != nil {
		return fmt.Errorf("cannot load private key: %v", err)
	}

	a.privateKey = key

	return nil
}

func (a *App) initLogger() error {
	a.Lock()
	defer a.Unlock()

	atom := zap.NewAtomicLevel()
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	logger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		//zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atom,
	))

	// TODO: (screwyprof) Read log level from settings file.
	atom.SetLevel(zap.DebugLevel)

	// grpclog will only log messages if zap log level is DebugLevel.
	// it's needed to avoid grpclog flood at Info level.
	//
	// grpclog log level can also be set via GRPC_GO_LOG_SEVERITY_LEVEL ENV variable.
	// Possible values (INFO, WARNING, ERROR, FATAL)
	// Default is ERROR.
	// See also GRPC_GO_LOG_VERBOSITY_LEVEL.
	grpclog.SetLogger(zapgrpc.NewLogger(logger, zapgrpc.WithDebug()))

	a.logger = logger
	return nil
}

func (a *App) initStorage() error {
	dataDirExists, err := a.pathExists(a.conf.DataDir)
	if err != nil {
		return fmt.Errorf("cannot check if data dir exists: %v", err)
	}

	if !dataDirExists {
		a.conf.DataDir = filepath.Dir(os.Args[0]) + "/data"
	}

	dataDir, err := filepath.Abs(a.conf.DataDir)
	if err != nil {
		return fmt.Errorf("cannot get absolute path of database directory: %v", err)
	}
	a.conf.DataDir = dataDir
	a.logger.Info("Data dir", zap.String("path", a.conf.DataDir))

	a.logger.Info("Importing database schema", zap.String("schema", a.conf.DataDir+"/schema.sql"))
	schema, err := ioutil.ReadFile(a.conf.DataDir + "/schema.sql")
	if err != nil {
		return fmt.Errorf("cannot read database schema file: %v", err)
	}
	a.logger.Info("Database schema successfully imported")

	db, err := sql.Open("sqlite3", a.conf.DataDir+"/data.db")
	if err != nil {
		return fmt.Errorf("cannot open database: %v", err)
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return fmt.Errorf("cannot import database schema: %v", err)
	}
	a.db = db

	return nil
}

func (a *App) initServer(mp *srv.Marketplace) error {
	a.RLock()
	logger := a.logger
	key := a.privateKey
	a.RUnlock()

	if key == nil {
		return fmt.Errorf("private key is not loaded")
	}

	rotator, tlsConfig, err := util.NewHitlessCertRotator(context.Background(), key)
	if err != nil {
		return err
	}

	creds := util.NewTLS(tlsConfig)

	opts := []grpc.ServerOption{
		grpc.WithGRPCOptions(gRPC.Creds(creds)),

		grpc.WithUnaryInterceptor(interceptor.NewUnaryZapLogger(logger)),
		grpc.WithUnaryInterceptor(interceptor.NewUnaryAuthenticator(interceptor.AuthFunc)),
		grpc.WithUnaryInterceptor(interceptor.NewUnarySimpleTracer()),
		grpc.WithUnaryInterceptor(interceptor.NewUnaryPanic()),

		grpc.WithStreamInterceptor(interceptor.NewStreamZapLogger(logger)),
		grpc.WithStreamInterceptor(interceptor.NewStreamAuthenticator(interceptor.AuthFunc)),
		grpc.WithStreamInterceptor(interceptor.NewStreamSimpleTracer()),
		grpc.WithStreamInterceptor(interceptor.NewStreamPanic()),
	}

	a.Lock()
	a.rotator = rotator
	a.creds = creds
	a.server = grpc.NewServer(opts...)
	a.Unlock()

	pb.RegisterMarketServer(a.server, mp)
	return nil
}

// Run runs the application.
func (a *App) Run() error {
	a.RLock()
	if a.server == nil {
		a.RUnlock()
		return fmt.Errorf("application is not initialized")
	}
	a.RUnlock()

	a.runScheduler()

	lis, err := net.Listen("tcp", a.conf.ListenAddr)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	a.logger.Info("Ready to serve", zap.String("addr", a.conf.ListenAddr))
	//a.logger.Debug("Application config", zap.Any("conf", a.conf))
	return a.server.Serve(lis)
}

// Stop gracefully stops the application.
func (a *App) Stop() {
	a.RLock()
	defer a.RUnlock()

	a.logger.Info("Application stopping...")

	time.AfterFunc(a.shutDownTimeOut(), func() {
		fmt.Printf("Application killed after timeout %s", a.shutDownTimeOut().String())
		os.Exit(2)
	})

	a.server.GracefulStop()
	a.rotator.Close()
	a.stopScheduler()
	a.db.Close()
	a.logger.Sync() // nolint

	a.logger.Info("Done")
}

// shutDownTimeOut the application will be terminated forcefully after time is up
// used during application termination.
func (a *App) shutDownTimeOut() time.Duration {
	return 15 * time.Second
}

// Creds a kludge to ease integration testing.
// TODO: (screwyprof) read keys on client side.
func (a *App) Creds() credentials.TransportCredentials {
	a.RLock()
	defer a.RUnlock()
	return a.creds
}

// PublicKey used in integration tests.
func (a *App) PublicKey() ecdsa.PublicKey {
	a.RLock()
	defer a.RUnlock()
	return a.privateKey.PublicKey
}

func (a *App) pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
