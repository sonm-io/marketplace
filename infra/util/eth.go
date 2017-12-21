package util

import (
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

// EthAddrExtractor extracts eth address from the context.
type EthAddrExtractor func(ctx context.Context) (common.Address, error)

// ExtractEthAddr extracts eth address from the context.
func ExtractEthAddr(ctx context.Context) (common.Address, error) {
	pr, ok := peer.FromContext(ctx)
	if !ok {
		return common.Address{}, status.Error(codes.Unauthenticated, "failed to get peer from ctx")
	}

	switch info := pr.AuthInfo.(type) {
	case EthAuthInfo:
		return info.Wallet, nil
	default:
		return common.Address{}, status.Error(codes.Unauthenticated, "wrong AuthInfo type")
	}
}
