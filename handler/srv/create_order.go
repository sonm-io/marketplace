package srv

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/tags/zap"
	"golang.org/x/net/context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/sonm-io/marketplace/handler/proto"
	"go.uber.org/zap"
)

// CreateOrder creates an order.
func (m *Marketplace) CreateOrder(ctx context.Context, req *pb.Order) (*pb.Order, error) {
	if err := CheckPermissions(ctx, req); err != nil {
		return nil, err
	}

	if err := m.validate(req); err != nil {
		return nil, err
	}

	order := *req
	// generate a unique ID if it's empty
	if order.Id == "" {
		order.Id = IDGenerator()
	}

	logger := ctx_zap.Extract(ctx)
	logger.Info("Creating order", zap.Any("order", order))

	if err := m.createOrder(order); err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create order: %v", err)
	}

	logger.Info("Order created", zap.String("id", order.Id))
	return m.GetOrderByID(ctx, &pb.ID{Id: order.Id})
}

func (m *Marketplace) validate(req *pb.Order) error {
	if req == nil || req.Slot == nil || req.Slot.Resources == nil {
		return nil
	}

	if req.Slot.Resources.GpuCount == pb.GPUCount_SINGLE_GPU {
		return status.Errorf(codes.Internal,
			"SINGLE_GPU has been deprecated, only NO_GPU and MULTIPLE_GPU are allowed")
	}

	return nil
}

func (m *Marketplace) createOrder(order pb.Order) error {
	orderCreator := m.marketService.CreateAskOrder
	if order.OrderType == pb.OrderType_BID {
		orderCreator = m.marketService.CreateBidOrder
	}
	return orderCreator(order)
}
