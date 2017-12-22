package srv

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sonm-io/marketplace/infra/grpc/interceptor"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
)

// CheckPermissions checks that the request is authorized and the sender has the proper permissions.
func CheckPermissions(ctx context.Context, req *pb.Order) error {
	key, err := interceptor.EthAddrFromContext(ctx)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "auth failed: %v", err)
	}

	if req.OrderType != pb.OrderType_ASK && req.OrderType != pb.OrderType_BID {
		return status.Errorf(
			codes.InvalidArgument, "auth failed: incorrect order type given: %v", req.OrderType.String())
	}

	if (req.OrderType == pb.OrderType_BID && req.ByuerID != key.Hex()) ||
		(req.OrderType == pb.OrderType_ASK && req.SupplierID != key.Hex()) {
		return status.Errorf(
			codes.PermissionDenied, "auth failed: SupplierID/BuyerID and ethereum address differ")
	}
	return nil
}
