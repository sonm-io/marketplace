package handler

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/tags/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/sonm-io/marketplace/proto"
)

// CancelOrder removes the given order from the storage.
func (m *Marketplace) CancelOrder(ctx context.Context, req *pb.Order) (*pb.Empty, error) {
	logger := ctx_zap.Extract(ctx)
	logger.Sugar().Infof("Canceling order %s", req.GetId())

	if err := m.checkPermissions(ctx, req); err != nil {
		return nil, err
	}

	if err := m.marketService.CancelOrder(req.Id); err != nil {
		return nil, err
	}

	logger.Sugar().Infof("Order %s successfully canceled", req.GetId())

	return &pb.Empty{}, nil
}

func (m *Marketplace) checkPermissions(ctx context.Context, req *pb.Order) error {
	// CheckPermissions requires OrderType, Supplier/BuyerID, that's why order is retrieved.
	order, err := m.orderByID(req.GetId())
	if err != nil {
		return err
	}

	req.OrderType = order.OrderType
	req.ByuerID = order.ByuerID
	req.SupplierID = order.SupplierID

	return CheckPermissions(ctx, req)
}

func (m *Marketplace) orderByID(ID string) (*pb.Order, error) {
	order := &pb.Order{}
	if err := m.marketService.OrderByID(ID, order); err != nil {
		return nil, status.Errorf(codes.Internal, "cannot cancel order: %v", err)
	}

	return order, nil
}
