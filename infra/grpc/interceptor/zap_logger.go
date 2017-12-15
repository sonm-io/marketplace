package interceptor

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// NewUnaryZapLogger Creates new unary interceptor for logging using zap.
func NewUnaryZapLogger(logger *zap.Logger) grpc.UnaryServerInterceptor {
	if logger == nil {
		return nil
	}
	return grpc_zap.UnaryServerInterceptor(logger)
}

// NewStreamZapLogger Creates new stream interceptor for logging using zap.
func NewStreamZapLogger(logger *zap.Logger) grpc.StreamServerInterceptor {
	if logger == nil {
		return nil
	}
	return grpc_zap.StreamServerInterceptor(logger)
}
