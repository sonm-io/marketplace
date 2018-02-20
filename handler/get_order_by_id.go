package handler

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/tags/zap"
	"golang.org/x/net/context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/sonm-io/marketplace/proto"
)

// GetOrderByID retrieves order information by order id.
func (m *Marketplace) GetOrderByID(ctx context.Context, req *pb.ID) (*pb.Order, error) {
	logger := ctx_zap.Extract(ctx)
	logger.Sugar().Infof("Getting order %s", req.GetId())

	resp := &pb.Order{}
	if err := m.marketService.OrderByID(req.GetId(), resp); err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get order: %v", err)
	}

	logger.Sugar().Infof("Order %+v retrieved\n", resp)
	return resp, nil
}
