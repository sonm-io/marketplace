package srv_test

import (
	"context"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
)

func (s *MarketplaceTestSuite) getOrders() {
	// arrange
	req := &pb.GetOrdersRequest{
		OrderType: pb.OrderType_BID,
		Count:     100,
		Slot: &pb.Slot{
			Resources: &pb.Resources{
				NetTrafficIn: 500000,
			},
		},
	}

	expected := &pb.GetOrdersReply{
		Orders: []*pb.Order{
			{
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
			},
		},
	}

	// act
	obtained, err := s.client.GetOrders(context.Background(), req)

	// assert
	s.NoError(err)
	s.Equal(expected, obtained)
}
