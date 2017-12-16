package srv

import (
	"context"
	"go.uber.org/zap"

	"github.com/grpc-ecosystem/go-grpc-middleware/tags/zap"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"

	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/marketplace/query"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/report"
)

const (
	defaultResultsCount = 100
)

// GetOrders retrieves orders by matching the given order options against the storage.
func (m *Marketplace) GetOrders(ctx context.Context, req *pb.GetOrdersRequest) (*pb.GetOrdersReply, error) {

	logger := ctx_zap.Extract(ctx)
	logger.Sugar().Infof("Getting orders %s", req)

	limit := req.GetCount()
	if limit == 0 {
		limit = defaultResultsCount
	}

	q := query.GetOrders{}
	bindGetOrdersQuery(req, &q)

	orders := report.GetOrdersReport{}
	if err := m.ordersBySpec.Handle(q, &orders); err != nil {
		logger.Sugar().Infof("Cannot retrieve orders: %v\n", err)
		return nil, err
	}
	logger.Info("Orders retrieved\n", zap.Any("orders", orders))

	var resp []*pb.Order
	for idx := range orders {
		pbOrder := &pb.Order{}
		bindGetOrderReport(&orders[idx], pbOrder)
		resp = append(resp, pbOrder)
	}

	return &pb.GetOrdersReply{Orders: resp}, nil
}

func bindGetOrdersQuery(req *pb.GetOrdersRequest, q *query.GetOrders) {
	if q == nil {
		return
	}

	q.Limit = req.GetCount()
	q.Order.OrderType = ds.OrderType(req.GetOrderType())

	if req.Slot == nil {
		return
	}

	if q.Order.Slot == nil {
		q.Order.Slot = &ds.Slot{}
	}

	q.Order.Slot.SupplierRating = req.GetSlot().GetSupplierRating()
	q.Order.Slot.BuyerRating = req.GetSlot().GetBuyerRating()
	if req.Slot.Resources != nil {
		res := req.GetSlot().GetResources()
		q.Order.Slot.Resources = ds.Resources{
			CPUCores:      res.GetCpuCores(),
			RAMBytes:      res.GetRamBytes(),
			GPUCount:      ds.GPUCount(res.GetGpuCount()),
			Storage:       res.GetStorage(),
			NetworkType:   ds.NetworkType(res.GetNetworkType()),
			NetTrafficIn:  res.GetNetTrafficIn(),
			NetTrafficOut: res.GetNetTrafficOut(),
			Properties:    res.GetProperties(),
		}
	}
}
