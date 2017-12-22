package interceptor

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"github.com/ethereum/go-ethereum/common"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
)

type authCtx struct{}

var (
	authCtxKey = &authCtx{}
)

// EthAuthInfo holds auth info containing ethereum address.
type EthAuthInfo interface {
	credentials.AuthInfo
	WalletAddress() common.Address
}

// AuthFunc default auth by ehtereum address.
var AuthFunc = func(ctx context.Context) (context.Context, error) {
	authInfo, err := AuthInfoFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth info: %v", err)
	}

	key, err := EthAddrFromAuthInfo(authInfo)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid key: %v", err)
	}
	return EthAddrToContext(ctx, key), nil
}

// EthAddrToContext puts the given eth address to the context.
func EthAddrToContext(ctx context.Context, address common.Address) context.Context {
	return context.WithValue(ctx, authCtxKey, address)
}

// EthAddrFromContext extracts eth address from the given context.
func EthAddrFromContext(ctx context.Context) (common.Address, error) {
	address, ok := ctx.Value(authCtxKey).(common.Address)
	if !ok {
		return common.Address{}, fmt.Errorf("cannot get eth address from context")
	}
	return address, nil
}

// EthAddrFromAuthInfo extracts eth addr from auth info.
func EthAddrFromAuthInfo(authInfo EthAuthInfo) (common.Address, error) {
	if authInfo == nil {
		return common.Address{}, fmt.Errorf("auth info is not set")
	}
	return authInfo.WalletAddress(), nil
}

// AuthInfoFromContext extracts eth address from the context.
func AuthInfoFromContext(ctx context.Context) (EthAuthInfo, error) {
	pr, ok := peer.FromContext(ctx)
	if !ok || pr == nil {
		return nil, status.Error(codes.Unauthenticated, "failed to get peer from ctx")
	}

	addr, ok := pr.AuthInfo.(EthAuthInfo)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "wrong AuthInfo type")
	}
	return addr, nil
}

// NewUnaryAuthenticator Creates new unary interceptor to authenticate requests.
func NewUnaryAuthenticator(authFunc grpc_auth.AuthFunc) grpc.UnaryServerInterceptor {
	if authFunc == nil {
		authFunc = AuthFunc
	}
	return grpc_auth.UnaryServerInterceptor(authFunc)
}

// NewStreamAuthenticator Creates new stream interceptor to authenticate requests.
func NewStreamAuthenticator(authFunc grpc_auth.AuthFunc) grpc.StreamServerInterceptor {
	if authFunc == nil {
		authFunc = AuthFunc
	}
	return grpc_auth.StreamServerInterceptor(authFunc)
}
