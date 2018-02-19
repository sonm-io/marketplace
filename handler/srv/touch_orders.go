package srv

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/tags/zap"
	"go.uber.org/zap"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/sonm-io/marketplace/handler/proto"
)

// CancelOrder removes the given order from the storage.
func (m *Marketplace) TouchOrders(ctx context.Context, req *pb.TouchOrdersRequest) (*pb.Empty, error) {
	logger := ctx_zap.Extract(ctx)
	logger.Info("Touching orders", zap.Strings("ids", req.GetIDs()))

	if err := m.marketService.TouchOrders(req.IDs); err != nil {
		return nil, status.Errorf(codes.Internal, "cannot touch orders: %v", err)
	}

	return &pb.Empty{}, nil
}
