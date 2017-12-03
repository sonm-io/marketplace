package interceptor

import (
	"sync/atomic"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
)

var (
	loggerReplaced uint32
)

// NewUnaryZapLogger Creates new unary interceptor for logging using zap.
func NewUnaryZapLogger(logger *zap.Logger) grpc.UnaryServerInterceptor {
	if logger == nil {
		return nil
	}

	// replace gRPC logger if it's not yet replaced
	if atomic.CompareAndSwapUint32(&loggerReplaced, 0, 1) {
		grpc_zap.ReplaceGrpcLogger(logger)
	}

	return grpc_zap.UnaryServerInterceptor(logger)
}


// NewStreamZapLogger Creates new stream interceptor for logging using zap.
func NewStreamZapLogger(logger *zap.Logger) grpc.StreamServerInterceptor {
	if logger == nil {
		return nil
	}

	// replace gRPC logger if it's not yet replaced
	if atomic.CompareAndSwapUint32(&loggerReplaced, 0, 1) {
		grpc_zap.ReplaceGrpcLogger(logger)
	}

	return grpc_zap.StreamServerInterceptor(logger)
}