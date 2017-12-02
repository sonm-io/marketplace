package srv

import (
	"github.com/sonm-io/marketplace/usecase/marketplace/query"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/report"

	//"github.com/grpc-ecosystem/go-grpc-middleware/tags/zap"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
	//"go.uber.org/zap"
	"golang.org/x/net/context"
	"log"
)

// GetOrderByID retrieves order information by order id.
func (m *Marketplace) GetOrderByID(ctx context.Context, req *pb.ID) (*pb.Order, error) {
	// this.loggerFactory.FromCtx(ctx)
	//l := ctx_zap.Extract(ctx)
	//l.Info("Getting order", zap.Any("req", req))
	log.Printf("Getting order %+v", req.GetId())

	order := &report.GetOrderReport{}
	if err := m.orderByID.Handle(query.GetOrder{ID: req.GetId()}, order); err != nil {
		log.Printf("cannot retreive order: %v\n", err)
		return nil, err
	}

	resp := &pb.Order{}
	bindGetOrderReport(order, resp)

	log.Printf("order %#+v retrieved\n", resp)
	//l.Info("Order retrieved", zap.Any("order", resp))

	return resp, nil
}

func bindGetOrderReport(r *report.GetOrderReport, pbOrder *pb.Order) {
	// build result
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
