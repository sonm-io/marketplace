package srv

import (
	"github.com/sonm-io/marketplace/report"
	"github.com/sonm-io/marketplace/usecase/marketplace/query"

	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
	"golang.org/x/net/context"
)

func (m *Marketplace) GetOrderByID(_ context.Context, req *pb.ID) (*pb.Order, error) {
	order := &report.GetOrderReport{}
	if err := m.orderByID.Handle(query.GetOrder{ID: req.GetId()}, order); err != nil {
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
		},
	}

	return resp, nil
}
