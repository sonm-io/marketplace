package srv_test

import (
	"context"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
)

func (s *MarketplaceTestSuite) getOrderByID() {
	// arrange
	expected := &pb.Order{
		Id:        "1b5dfa00-af3c-4e2d-b64b-c5d62e89430b",
		OrderType: pb.OrderType_BID,
		Price:     "777",
		ByuerID:   "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db",

		Slot: &pb.Slot{
			BuyerRating: 555,
			Resources: &pb.Resources{
				CpuCores: 1,
				GpuCount: pb.GPUCount_SINGLE_GPU,
				RamBytes: 100000000,
				Storage:  1000000000,

				NetworkType:  pb.NetworkType_INCOMING,
				NetTrafficIn: 500000,

				Properties: map[string]float64{
					"hash_rate": 105.7,
				},
			},
		},
	}

	// act
	obtained, err := s.client.GetOrderByID(context.Background(), &pb.ID{Id: "1b5dfa00-af3c-4e2d-b64b-c5d62e89430b"})

	// assert
	s.NoError(err)
	s.Equal(expected, obtained)
}
