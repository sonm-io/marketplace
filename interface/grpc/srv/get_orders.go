package srv

import (
	"context"
	"go.uber.org/zap"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	logger.Info("Getting orders", zap.Any("req", req))

	limit := req.GetCount()
	if limit == 0 {
		limit = defaultResultsCount
	}

	q := query.GetOrders{}
	bindGetOrdersQuery(req, &q)

	orders := report.GetOrdersReport{}
	if err := m.ordersBySpec.Handle(q, &orders); err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get orders: %v", err)
	}
	logger.Info("Orders retrieved", zap.Any("orders", orders))

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
	q.Order.OrderType = ds.OrderType(req.GetOrder().GetOrderType())

	if req.GetOrder().GetSlot() == nil {
		return
	}

	if q.Order.Slot == nil {
		q.Order.Slot = &ds.Slot{}
	}

	reqSlot := req.GetOrder().GetSlot()
	q.Order.Slot.SupplierRating = reqSlot.GetSupplierRating()
	q.Order.Slot.BuyerRating = reqSlot.GetBuyerRating()

	if reqSlot.Resources == nil {
		return
	}

	res := reqSlot.GetResources()
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
