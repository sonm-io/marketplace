package cli

import (
	"fmt"
	"net"
	"sync"

	"github.com/sonm-io/marketplace/infra/grpc"
	"github.com/sonm-io/marketplace/infra/cqrs"
	"github.com/sonm-io/marketplace/infra/storage/inmemory"

	"github.com/sonm-io/marketplace/interface/adaptor"
	"github.com/sonm-io/marketplace/interface/storage"
	"github.com/sonm-io/marketplace/interface/grpc/srv"

	"github.com/sonm-io/marketplace/usecase/marketplace/query"
	"github.com/sonm-io/marketplace/usecase/marketplace/command"

	gRPC "google.golang.org/grpc"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
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

func (a *App) initServer(mp *srv.Marketplace) {
	a.Lock()
	a.server = grpc.NewServer()
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
}
