package srv

import (
	"github.com/sonm-io/marketplace/usecase/marketplace/command"

	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
	"golang.org/x/net/context"
)

func (m *Marketplace) CancelOrder(_ context.Context, req *pb.Order) (*pb.Empty, error) {
	if err := m.commandBus.Handle(command.CancelOrder{ID: req.Id}); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
