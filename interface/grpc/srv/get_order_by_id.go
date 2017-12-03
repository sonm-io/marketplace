package srv

import (
	"golang.org/x/net/context"

	"github.com/grpc-ecosystem/go-grpc-middleware/tags/zap"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"

	"github.com/sonm-io/marketplace/usecase/marketplace/query"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/report"
)

// GetOrderByID retrieves order information by order id.
func (m *Marketplace) GetOrderByID(ctx context.Context, req *pb.ID) (*pb.Order, error) {
	logger := ctx_zap.Extract(ctx)
	logger.Sugar().Infof("Getting order", req.GetId())

	order := &report.GetOrderReport{}
	if err := m.orderByID.Handle(query.GetOrder{ID: req.GetId()}, order); err != nil {
		logger.Sugar().Infof("cannot retreive order: %v\n", err)
		return nil, err
	}

	resp := &pb.Order{}
	bindGetOrderReport(order, resp)

	logger.Sugar().Infof("order %#+v retrieved\n", resp)
	return resp, nil
}

func bindGetOrderReport(r *report.GetOrderReport, pbOrder *pb.Order) {
	if r == nil {
		return
	}

	pbOrder.Id = r.ID
	pbOrder.ByuerID = r.BuyerID
	pbOrder.SupplierID = r.SupplierID
	pbOrder.Price = r.Price
	pbOrder.OrderType = pb.OrderType(r.OrderType)

	if r.Slot == nil {
		return
	}

	if pbOrder.Slot == nil {
		pbOrder.Slot = &pb.Slot{}
	}

	pbSlot := pbOrder.Slot
	rSlot := r.Slot

	pbSlot.SupplierRating = rSlot.SupplierRating
	pbSlot.BuyerRating = rSlot.BuyerRating

	if pbSlot.Resources == nil {
		pbSlot.Resources = &pb.Resources{}
	}

	pbRes := pbSlot.Resources
	rRes := rSlot.Resources

	pbRes.CpuCores = rRes.CPUCores
	pbRes.RamBytes = rRes.RAMBytes
	pbRes.Storage = rRes.Storage
}
