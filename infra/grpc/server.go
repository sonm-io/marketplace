package grpc

import (
	"google.golang.org/grpc"
	// registers grpc gzip encoder/decoder
	"github.com/grpc-ecosystem/go-grpc-middleware"
	_ "google.golang.org/grpc/encoding/gzip"
)

type options struct {
	grpcOptions []grpc.ServerOption

	unary  []grpc.UnaryServerInterceptor
	stream []grpc.StreamServerInterceptor
}

type ServerOption func(options *options)

func WithGrpcOptions(opt ...grpc.ServerOption) ServerOption {
	return func(o *options) {
		o.grpcOptions = opt
	}
}

func WithUnaryInterceptor(u grpc.UnaryServerInterceptor) ServerOption {
	return func(o *options) {
		o.unary = append(o.unary, u)
	}
}

func WithStreamInterceptor(s grpc.StreamServerInterceptor) ServerOption {
	return func(o *options) {
		o.stream = append(o.stream, s)
	}
}

// NewServer creates a new instance of gRPC server with default interceptors.
func NewServer(opt ...ServerOption) *grpc.Server {
	opts := options{}
	for _, o := range opt {
		o(&opts)
	}

	return grpc.NewServer(
		append([]grpc.ServerOption{
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(opts.unary...)),
			grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(opts.stream...)),
		}, opts.grpcOptions...)...,
	)
}
