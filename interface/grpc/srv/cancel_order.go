package srv

import (
	"golang.org/x/net/context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags/zap"

	pb "github.com/sonm-io/marketplace/interface/grpc/proto"

	"github.com/sonm-io/marketplace/usecase/marketplace/command"
	"github.com/sonm-io/marketplace/usecase/marketplace/query"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/report"

)

// CancelOrder removes the given order from the storage.
func (m *Marketplace) CancelOrder(ctx context.Context, req *pb.Order) (*pb.Empty, error) {

	logger := ctx_zap.Extract(ctx)
	logger.Sugar().Infof("Canceling order %s", req.GetId())

	// used by CheckPermissions bellow
	order, err := m.getOrderByID(req.GetId())
	if err != nil {
		return nil, err
	}
	req.OrderType = pb.OrderType(order.OrderType)
	req.ByuerID = order.BuyerID
	req.SupplierID = order.SupplierID

	if err := CheckPermissions(ctx, req); err != nil {
		return nil, err
	}

	if err := m.commandBus.Handle(command.CancelOrder{ID: req.Id}); err != nil {
		return nil, err
	}

	logger.Sugar().Infof("Order %s successfully canceled", req.GetId())

	return &pb.Empty{}, nil
}

func (m *Marketplace) getOrderByID(ID string) (*report.GetOrderReport, error) {
	var order report.GetOrderReport
	if err := m.orderByID.Handle(query.GetOrder{ID: ID}, &order); err != nil {
		return nil, status.Errorf(codes.Internal, "cannot cancel order: %v", err)
	}
	return &order, nil
}
