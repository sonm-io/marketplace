package srv

import (
	"golang.org/x/net/context"

	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
	"github.com/sonm-io/marketplace/usecase/marketplace/command"
)

// CancelOrder removes the given order from the storage.
func (m *Marketplace) CancelOrder(ctx context.Context, req *pb.Order) (*pb.Empty, error) {
	if err := CheckPermissions(ctx, req); err != nil {
		return nil, err
	}

	if err := m.commandBus.Handle(command.CancelOrder{ID: req.Id}); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
