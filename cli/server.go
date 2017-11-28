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
	"google.golang.org/grpc/status"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
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

	createOrderHandler := command.NewCreateBidOrderHandler(repo)
	cancelOrderHandler := command.NewCancelOrderHandler(repo)

	commandBus := cqrs.NewCommandBus()
	commandBus.RegisterHandler("CreateBidOrder", adaptor.FromDomain(createOrderHandler))
	commandBus.RegisterHandler("CancelOrder", adaptor.FromDomain(cancelOrderHandler))

	lis, err := net.Listen("tcp", l.conf.ListenAddr)
	if err != nil {
		return err
	}

	panicHandler := grpc_recovery.RecoveryHandlerFunc(func(p interface{}) (err error) {
		log.Printf("%+v", errors.Callers())

		err = status.Errorf(codes.Internal, "%s", p)
		return
	})

	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(panicHandler),
	}

	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(grpc_recovery.UnaryServerInterceptor(opts...)),
		grpc_middleware.WithStreamServerChain(grpc_recovery.StreamServerInterceptor(opts...)),

		grpc.RPCCompressor(grpc.NewGZIPCompressor()),
		grpc.RPCDecompressor(grpc.NewGZIPDecompressor()),
	)

	pb.RegisterMarketServer(s, srv.NewMarketplace(adaptor.ToDomain(commandBus), getOrderHandler, getOrdersHandler))

	return s.Serve(lis)
}
