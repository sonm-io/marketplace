package srv_test

import (
	"context"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
)

func (s *MarketplaceTestSuite) createOrder() {
	// arrange
	order := &pb.Order{
		Id:        "1b5dfa00-af3c-4e2d-b64b-c5d62e89430b",
		OrderType: pb.OrderType_BID,
		ByuerID:   "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db",
		Price:     100,
		Slot: &pb.Slot{
			Resources: &pb.Resources{
				CpuCores: 1,
				RamBytes: 100000000,
				Storage:  1000000000,
			},
		},
	}

	// act
	obtained, err := s.client.CreateOrder(context.Background(), order)

	// assert
	s.NoError(err)
	s.Equal(order.Id, obtained.Id)
}
