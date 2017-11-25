package srv

import (
	"github.com/sonm-io/marketplace/report"
	"github.com/sonm-io/marketplace/usecase/marketplace/query"

	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
	"golang.org/x/net/context"
	"log"
)

func (m *Marketplace) GetOrderByID(_ context.Context, req *pb.ID) (*pb.Order, error) {

	log.Printf("Getting order %+v", req.GetId())

	order := &report.GetOrderReport{}
	if err := m.orderByID.Handle(query.GetOrder{ID: req.GetId()}, order); err != nil {
		log.Printf("cannot retreive order: %v\n", err)
		return nil, err
	}

	// build result
	resp := &pb.Order{
		Id:         order.ID,
		ByuerID:    order.BuyerID,
		SupplierID: order.SupplierID,
		Price:      order.Price,
		OrderType:  pb.OrderType(order.OrderType),
		Slot: &pb.Slot{
			SupplierRating: order.Slot.SupplierRating,
			BuyerRating:    order.Slot.BuyerRating,
			Resources: &pb.Resources{
				CpuCores: order.Slot.Resources.CpuCores,
				RamBytes: order.Slot.Resources.RamBytes,
				Storage:  order.Slot.Resources.Storage,
			},
		},
	}

	log.Printf("order %#+v retreived\n", resp)

	return resp, nil
}
