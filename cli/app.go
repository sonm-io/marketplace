package cli

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zapgrpc"

	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
	gRPC "google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	// register sqlite3 driver
	_ "github.com/mattn/go-sqlite3"

	"github.com/sonm-io/marketplace/infra/cqrs"
	"github.com/sonm-io/marketplace/infra/grpc"
	"github.com/sonm-io/marketplace/infra/grpc/interceptor"
	engine "github.com/sonm-io/marketplace/infra/storage/sqllite"

	"github.com/sonm-io/marketplace/interface/adaptor"
	"github.com/sonm-io/marketplace/interface/grpc/srv"
	"github.com/sonm-io/marketplace/interface/storage/sqllite"

	"github.com/sonm-io/marketplace/usecase/marketplace/command"
	"github.com/sonm-io/marketplace/usecase/marketplace/query"
)

// Config application configuration object.
type Config struct {
	ListenAddr string
	DataDir    string
}

// Option is a configuration parameter.
type Option func(f *Config)

// WithListenAddr sets listen address.
func WithListenAddr(addr string) Option {
	return func(c *Config) {
		c.ListenAddr = addr
	}
}

// WithDataDir sets the database path.
func WithDataDir(dirPath string) Option {
	return func(c *Config) {
		c.DataDir = dirPath
	}
}

// App application root.
type App struct {
	sync.RWMutex

	db     *sql.DB
	conf   *Config
	logger *zap.Logger
	server *gRPC.Server
}

// NewApp creates a new App instance.
func NewApp(opts ...Option) *App {
	conf := &Config{ListenAddr: ":9090"}
	for _, option := range opts {
		option(conf)
	}

	return &App{conf: conf}
}

// Init initialize the application.
func (a *App) Init() error {
	if err := a.initLogger(); err != nil {
		return err
	}

	if err := a.initStorage(); err != nil {
		return err
	}

	repo := sqllite.NewOrderStorage(engine.NewOrderStorage(a.db))

	getOrderHandler := query.NewGetOrderHandler(repo)
	getOrdersHandler := query.NewGetOrdersHandler(repo)

	createOrderHandler := command.NewCreateBidOrderHandler(repo)
	cancelOrderHandler := command.NewCancelOrderHandler(repo)

	commandBus := cqrs.NewCommandBus()
	commandBus.RegisterHandler("CreateBidOrder", adaptor.FromDomain(createOrderHandler))
	commandBus.RegisterHandler("CancelOrder", adaptor.FromDomain(cancelOrderHandler))

	a.initServer(srv.NewMarketplace(adaptor.ToDomain(commandBus), getOrderHandler, getOrdersHandler))
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

func (a *App) initServer(mp *srv.Marketplace) {

	a.RLock()
	opts := []grpc.ServerOption{
		grpc.WithUnaryInterceptor(interceptor.NewUnaryZapLogger(a.logger)),
		grpc.WithUnaryInterceptor(interceptor.NewUnarySimpleTracer()),
		grpc.WithUnaryInterceptor(interceptor.NewUnaryPanic()),

		grpc.WithStreamInterceptor(interceptor.NewStreamZapLogger(a.logger)),
		grpc.WithStreamInterceptor(interceptor.NewStreamSimpleTracer()),
		grpc.WithStreamInterceptor(interceptor.NewStreamPanic()),
	}
	a.RUnlock()

	a.Lock()
	a.server = grpc.NewServer(opts...)
	a.Unlock()

	pb.RegisterMarketServer(a.server, mp)
}

// Run runs the application.
func (a *App) Run() error {
	a.RLock()
	if a.server == nil {
		a.RUnlock()
		return fmt.Errorf("application is not initialized")
	}
	a.RUnlock()

	lis, err := net.Listen("tcp", a.conf.ListenAddr)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	return a.server.Serve(lis)
}

// Stop gracefully stops the application.
func (a *App) Stop() {
	a.RLock()
	defer a.RUnlock()

	if a.server != nil {
		a.server.GracefulStop()
	}

	if a.db != nil {
		a.db.Close()
	}

	if a.logger != nil {
		a.logger.Sync() // nolint
	}
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
