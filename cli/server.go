package cli

import (
	"log"
	"net"

	"github.com/sonm-io/marketplace/infra/cqrs"
	"github.com/sonm-io/marketplace/infra/errors"
	"github.com/sonm-io/marketplace/infra/storage/inmemory"

	"github.com/sonm-io/marketplace/interface/adaptor"
	"github.com/sonm-io/marketplace/interface/grpc/srv"
	"github.com/sonm-io/marketplace/interface/storage"

	"github.com/sonm-io/marketplace/usecase/marketplace/command"
	"github.com/sonm-io/marketplace/usecase/marketplace/query"

	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	// registers grpc gzip encoder/decoder
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/status"

	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"sync"
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
	conf *Config

	sync.RWMutex
	server *grpc.Server
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
	repo := storage.NewOrderStorage(inmemory.NewStorage())

	getOrderHandler := query.NewGetOrderHandler(repo)
	getOrdersHandler := query.NewGetOrdersHandler(repo)

	createOrderHandler := command.NewCreateBidOrderHandler(repo)
	cancelOrderHandler := command.NewCancelOrderHandler(repo)

	commandBus := cqrs.NewCommandBus()
	commandBus.RegisterHandler("CreateBidOrder", adaptor.FromDomain(createOrderHandler))
	commandBus.RegisterHandler("CancelOrder", adaptor.FromDomain(cancelOrderHandler))

	panicHandler := grpc_recovery.RecoveryHandlerFunc(func(p interface{}) (err error) {
		log.Printf("%+v", errors.Callers())

		err = status.Errorf(codes.Internal, "%s", p)
		return
	})

	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(panicHandler),
	}

	a.Lock()
	a.server = grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(grpc_recovery.UnaryServerInterceptor(opts...)),
		grpc_middleware.WithStreamServerChain(grpc_recovery.StreamServerInterceptor(opts...)),
	)
	a.Unlock()

	pb.RegisterMarketServer(a.server, srv.NewMarketplace(adaptor.ToDomain(commandBus), getOrderHandler, getOrdersHandler))

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
}
