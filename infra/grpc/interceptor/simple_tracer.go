package interceptor

import (
	"context"
	"github.com/pborman/uuid"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type ctxKey struct{}

var (
	requestIDKey = &ctxKey{}
)

// NewUnarySimpleTracer Creates new unary interceptor for tracing requests.
func NewUnarySimpleTracer() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (
		interface{}, error) {
		return handler(ToContext(ctx, uuid.New()), req)
	}
}

// NewStreamSimpleTracer Creates new stream interceptor for tracing requests.
func NewStreamSimpleTracer() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo,
		handler grpc.StreamHandler) error {
		wrapped := grpc_middleware.WrapServerStream(stream)
		wrapped.WrappedContext = ToContext(stream.Context(), uuid.New())

		return handler(srv, wrapped)
	}
}

// ToContext returns a context which knows its request ID.
func ToContext(ctx context.Context, requestID string) context.Context {
	ctx_zap.AddFields(ctx, zap.String("reqID", requestID))
	return context.WithValue(ctx, requestIDKey, requestID)
}

// Extract takes a requestID from context.
func Extract(ctx context.Context) string {
	reqID, ok := ctx.Value(requestIDKey).(string)
	if !ok || reqID == "" {
		reqID = "n/a"
	}
	return reqID
}
