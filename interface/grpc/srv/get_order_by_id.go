package srv

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/tags/zap"
	"golang.org/x/net/context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/sonm-io/marketplace/interface/grpc/proto"

	"github.com/sonm-io/marketplace/usecase/marketplace/query"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/report"
)

// GetOrderByID retrieves order information by order id.
func (m *Marketplace) GetOrderByID(ctx context.Context, req *pb.ID) (*pb.Order, error) {
	logger := ctx_zap.Extract(ctx)
	logger.Sugar().Infof("Getting order %s", req.GetId())

	order := &report.GetOrderReport{}
	if err := m.orderByID.Handle(query.GetOrder{ID: req.GetId()}, order); err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get order: %v", err)
	}

	resp := &pb.Order{}
	bindGetOrderReport(order, resp)

	logger.Sugar().Infof("Order %+v retrieved\n", resp)
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

	pbSlot.Duration = rSlot.Duration
	pbSlot.SupplierRating = rSlot.SupplierRating
	pbSlot.BuyerRating = rSlot.BuyerRating

	if pbSlot.Resources == nil {
		pbSlot.Resources = &pb.Resources{}
	}

	pbRes := pbSlot.Resources
	rRes := rSlot.Resources

	pbRes.CpuCores = rRes.CPUCores
	pbRes.RamBytes = rRes.RAMBytes
	pbRes.GpuCount = pb.GPUCount(rRes.GPUCount)
	pbRes.Storage = rRes.Storage

	pbRes.NetworkType = pb.NetworkType(rRes.NetworkType)
	pbRes.NetTrafficIn = rRes.NetTrafficIn
	pbRes.NetTrafficOut = rRes.NetTrafficOut

	pbRes.Properties = rRes.Properties
}
