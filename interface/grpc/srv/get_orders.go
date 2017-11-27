package srv

import (
	"github.com/sonm-io/marketplace/usecase/marketplace/query"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/report"

	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
	"golang.org/x/net/context"
	"log"
)

const (
	defaultResultsCount = 100
)

func (m *Marketplace) GetOrders(_ context.Context, req *pb.GetOrdersRequest) (*pb.GetOrdersReply, error) {

	log.Printf("Getting orders %+v", req)

	limit := req.GetCount()
	if limit == 0 {
		limit = defaultResultsCount
	}

	slot := query.Slot{
		BuyerRating:    req.Slot.GetBuyerRating(),
		SupplierRating: req.Slot.GetSupplierRating(),
	}

	q := query.GetOrders{
		OrderType: int(req.OrderType),
		Slot:      slot,
		Limit:     limit,
	}

	orders := report.GetOrdersReport{}
	if err := m.ordersBySpec.Handle(q, &orders); err != nil {
		log.Printf("cannot retreive orders: %v\n", err)
		return nil, err
	}

	log.Printf("orders %#+v retreived\n", orders)

	var resp []*pb.Order
	return &pb.GetOrdersReply{Orders: resp}, nil
}
