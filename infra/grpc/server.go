package grpc

import (
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	// registers grpc gzip encoder/decoder
	_ "google.golang.org/grpc/encoding/gzip"
)

type options struct {
	grpcOptions []grpc.ServerOption

	unary  []grpc.UnaryServerInterceptor
	stream []grpc.StreamServerInterceptor
}

// ServerOption sets server options.
type ServerOption func(options *options)

// WithGRPCOptions sets gRPC server options.
func WithGRPCOptions(opt ...grpc.ServerOption) ServerOption {
	return func(o *options) {
		o.grpcOptions = opt
	}
}

// WithUnaryInterceptor adds an unary interceptor to the chain.
func WithUnaryInterceptor(u grpc.UnaryServerInterceptor) ServerOption {
	return func(o *options) {
		o.unary = append(o.unary, u)
	}
}

// WithStreamInterceptor adds a stream interceptor to the chain.
func WithStreamInterceptor(s grpc.StreamServerInterceptor) ServerOption {
	return func(o *options) {
		o.stream = append(o.stream, s)
	}
}

// NewServer creates a new instance of gRPC server.
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
