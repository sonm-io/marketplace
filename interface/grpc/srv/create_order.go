package srv

import (
	"github.com/sonm-io/marketplace/usecase/marketplace/command"

	"github.com/pborman/uuid"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
	"golang.org/x/net/context"
)

func (m *Marketplace) CreateOrder(_ context.Context, req *pb.Order) (*pb.Order, error) {

	// map request to command
	// TODO: (screwyprof) move to smth like cmd := model.Bind(req), or model.bind(&req, &cmd)
	cmd := command.CreateOrder{
		ID:         uuid.New(),
		SupplierID: req.SupplierID,
		BuyerID:    req.ByuerID,
		Price:      req.Price,
	}

	// handle command
	if err := m.commandBus.Handle(cmd); err != nil {
		return nil, err
	}

	// return response
	return &pb.Order{Id: cmd.ID}, nil
}
