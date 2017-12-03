package interceptor

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/grpclog"

	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
)

var panicHandler = grpc_recovery.RecoveryHandlerFunc(func(p interface{}) (err error) {
	grpclog.Errorf("PANIC recovered: %+v", Callers())
	err = status.Errorf(codes.Internal, "%s", p)
	return
})

// NewUnaryPanic Creates new unary interceptor to recover from panics.
func NewUnaryPanic() grpc.UnaryServerInterceptor {
	return grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(panicHandler))
}

// NewStreamPanic Creates new stream interceptor to recover from panics.
func NewStreamPanic() grpc.StreamServerInterceptor {
	return grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandler(panicHandler))
}