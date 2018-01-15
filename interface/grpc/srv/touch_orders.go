package srv

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/tags/zap"
	"go.uber.org/zap"

	pb "github.com/sonm-io/marketplace/interface/grpc/proto"

	"github.com/sonm-io/marketplace/usecase/marketplace/command"
)

// CancelOrder removes the given order from the storage.
func (m *Marketplace) TouchOrders(ctx context.Context, req *pb.TouchOrdersRequest) (*pb.Empty, error) {
	logger := ctx_zap.Extract(ctx)
	logger.Info("Touching orders", zap.Strings("ids", req.GetIDs()))

	return &pb.Empty{}, m.commandBus.Handle(command.TouchOrders{IDs: req.GetIDs()})
}
