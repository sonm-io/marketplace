package srv

import (
	pb "github.com/sonm-io/marketplace/handler/proto"
)

type MarketService interface {
	CreateAskOrder(o pb.Order) error
	CreateBidOrder(o pb.Order) error
	TouchOrders(IDs []string) error
	CancelOrder(ID string) error

	OrderByID(ID string, result interface{}) error
	MatchOrders(req *pb.Order, limit uint64, result interface{}) error
}

// Marketplace a GRPC Server implementing Marketplace API.
type Marketplace struct {
	marketService MarketService
}

// NewMarketplace creates a new instance of Marketplace.
func NewMarketplace(marketService MarketService) *Marketplace {
	return &Marketplace{
		marketService: marketService,
	}
}
