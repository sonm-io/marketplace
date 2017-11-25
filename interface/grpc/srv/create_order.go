package srv

import (
	"fmt"
	"log"

	"github.com/pborman/uuid"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
	"golang.org/x/net/context"

	"github.com/sonm-io/marketplace/usecase/marketplace/command"
)

func (m *Marketplace) CreateOrder(ctx context.Context, req *pb.Order) (*pb.Order, error) {

	ID := uuid.New()
	if err := m.createBidOrder(ID, req); err != nil {
		return nil, fmt.Errorf("cannot create bid order: %v", err)
	}

	// return response
	return m.GetOrderByID(ctx, &pb.ID{Id: ID})
}

func (m *Marketplace) createBidOrder(ID string, req *pb.Order) error {

	// map request to command
	// TODO: (screwyprof) move to smth like cmd := model.Bind(req), or model.bind(&req, &cmd)
	cmd := command.CreateBidOrder{
		ID:        ID,
		Price:     req.GetPrice(),
		OrderType: int(req.GetOrderType()),
		BuyerID:   req.GetByuerID(),
	}

	log.Printf("Creating bid order %+v", cmd)

	// handle command
	if err := m.commandBus.Handle(cmd); err != nil {
		log.Printf("cannot create bid order: %v\n", err)
		return err
	}

	log.Printf("bid order %s created\n", cmd.ID)
	return nil
}
