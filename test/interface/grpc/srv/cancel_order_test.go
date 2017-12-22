package srv_test

import (
	"context"
	"fmt"
	"github.com/sonm-io/marketplace/infra/util"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
)

func (s *MarketplaceTestSuite) CancelOrder() {
	// smth like "0x9B27D3C3571731deDb23EaFEa34a3a6E05daE159"
	ID := "1b5dfa00-af3c-4e2d-b64b-c5d62e89430b"
	buyerID := util.PubKeyToAddr(s.App.PublicKey()).Hex()

	req := &pb.Order{
		OrderType: pb.OrderType_BID,
		Id:        ID,
		ByuerID:   buyerID,
	}

	_, err := s.client.CancelOrder(context.Background(), req)
	s.Require().NoError(err)

	// act
	_, err = s.client.GetOrderByID(context.Background(), &pb.ID{Id: ID})

	// assert
	s.EqualError(err,
		fmt.Sprintf(`rpc error: code = Internal desc = cannot get order: order %s is inactive`, ID))
}
