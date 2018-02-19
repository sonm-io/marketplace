package handler

import (
	"context"
	"go.uber.org/zap"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/grpc-ecosystem/go-grpc-middleware/tags/zap"
	pb "github.com/sonm-io/marketplace/proto"
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

	resp := &pb.GetOrdersReply{}
	if err := m.marketService.MatchOrders(req.Order, limit, resp); err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get orders: %v", err)
	}

	logger.Info("Orders retrieved", zap.Any("orders", resp.Orders))
	return resp, nil
}
