package cli

import (
	"net"

	"github.com/sonm-io/marketplace/infra/cqrs"
	"github.com/sonm-io/marketplace/interface/storage"
	"github.com/sonm-io/marketplace/usecase/marketplace/query"

	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
	"google.golang.org/grpc"

	"github.com/sonm-io/marketplace/infra/storage/inmemory"
	"github.com/sonm-io/marketplace/interface/adaptor"
	"github.com/sonm-io/marketplace/interface/grpc/srv"

	"github.com/sonm-io/marketplace/usecase/marketplace/command"
)

type Config struct {
	ListenAddr string
}

type Option func(f *Config)

func WithListenAddr(addr string) Option {
	return func(c *Config) {
		c.ListenAddr = addr
	}
}

type App struct {
	conf *Config
}

func NewApp(opts ...Option) *App {

	conf := &Config{ListenAddr: ":9090"}
	for _, option := range opts {
		option(conf)
	}

	return &App{conf: conf}
}

func (l *App) Run() error {

	repo := storage.NewOrderStorage(inmemory.NewStorage())

	getOrderHandler := query.NewGetOrderHandler(repo)
	getOrdersHandler := query.NewGetOrdersHandler(repo)
	createOrderHandler := command.NewCreateOrderValidator(command.NewCreateOrderHandler(repo))
	cancelOrderHandler := command.NewCancelOrderHandler(repo)

	commandBus := cqrs.NewCommandBus()
	commandBus.RegisterHandler("CreateOrder", adaptor.FromDomain(createOrderHandler))
	commandBus.RegisterHandler("CancelOrder", adaptor.FromDomain(cancelOrderHandler))

	lis, err := net.Listen("tcp", l.conf.ListenAddr)
	if err != nil {
		return err
	}

	s := grpc.NewServer(
		grpc.RPCCompressor(grpc.NewGZIPCompressor()),
		grpc.RPCDecompressor(grpc.NewGZIPDecompressor()),
	)

	pb.RegisterMarketServer(s, srv.NewMarketplace(adaptor.ToDomain(commandBus), getOrderHandler, getOrdersHandler))

	return s.Serve(lis)
}
