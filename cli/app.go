package cli

import (
	"fmt"
	"net"
	"os"
	"sync"

	"go.uber.org/zap"

	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
	gRPC "google.golang.org/grpc"

	"github.com/sonm-io/marketplace/infra/cqrs"
	"github.com/sonm-io/marketplace/infra/grpc"
	"github.com/sonm-io/marketplace/infra/grpc/interceptor"
	"github.com/sonm-io/marketplace/infra/storage/inmemory"

	"github.com/sonm-io/marketplace/interface/adaptor"
	"github.com/sonm-io/marketplace/interface/grpc/srv"
	"github.com/sonm-io/marketplace/interface/storage"

	"github.com/sonm-io/marketplace/usecase/marketplace/command"
	"github.com/sonm-io/marketplace/usecase/marketplace/query"
	"go.uber.org/zap/zapcore"
)

// Config application configuration object.
type Config struct {
	ListenAddr string
}

// Option is a configuration parameter.
type Option func(f *Config)

// WithListenAddr sets listen address.
func WithListenAddr(addr string) Option {
	return func(c *Config) {
		c.ListenAddr = addr
	}
}

// App application root.
type App struct {
	sync.RWMutex

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

	repo := storage.NewOrderStorage(inmemory.NewStorage())

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

	// a fallback/root logger for events without context
	//logger, err := zap.NewProduction(
	//	zap.Fields(zap.Int("pid", os.Getpid()), zap.String("exe", path.Base(os.Args[0]))),
	//)
	//defer l.Sync()
	//
	//if err != nil {
	//	return fmt.Errorf("cannot init logger: %v", err)
	//}

	//gRPC log level is set via env. Possible values (INFO, ERROR, WARNING)
	//see also gRPC_GO_LOG_VERBOSITY_LEVEL
	//os.Setenv("GRPC_GO_LOG_SEVERITY_LEVEL", "ERROR")

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

	// use in destructor
	//defer logger.Sync()

	atom.SetLevel(zap.InfoLevel)

	a.logger = logger
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

	if a.server == nil {
		return
	}
	a.server.GracefulStop()

	if a.logger == nil {
		return
	}
	a.logger.Sync() // nolint
}
