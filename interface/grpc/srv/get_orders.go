package srv

import (
	"github.com/sonm-io/marketplace/report"
	"github.com/sonm-io/marketplace/usecase/marketplace/query"

	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
	"golang.org/x/net/context"
)

const (
	defaultResultsCount = 100
)

func (m *Marketplace) GetOrders(_ context.Context, req *pb.GetOrdersRequest) (*pb.GetOrdersReply, error) {

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

	orders := &report.Order{}
	if err := m.orderByID.Handle(q, orders); err != nil {
		return nil, err
	}

	return &pb.GetOrdersReply{}, nil
}
