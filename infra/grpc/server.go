package grpc

import (
	"google.golang.org/grpc"
	// registers grpc gzip encoder/decoder
	_ "google.golang.org/grpc/encoding/gzip"
	"github.com/grpc-ecosystem/go-grpc-middleware"

	"github.com/sonm-io/marketplace/infra/grpc/interceptor"
)

// NewServer creates a new instance of gRPC server with default interceptors.
func NewServer(opts ...grpc.ServerOption) *grpc.Server {
	return grpc.NewServer(
		append([]grpc.ServerOption{
			grpc.UnaryInterceptor(DefaultServerInterceptor()),
			grpc.StreamInterceptor(DefaultServerStreamInterceptor()),
		}, opts...)...)
}

func DefaultServerInterceptor() grpc.UnaryServerInterceptor {
	return grpc_middleware.ChainUnaryServer(
		interceptor.NewUnaryPanic(), // should be the last
	)
}

func DefaultServerStreamInterceptor() grpc.StreamServerInterceptor {
	return grpc_middleware.ChainStreamServer(
		interceptor.NewStreamPanic(),
	)
}
