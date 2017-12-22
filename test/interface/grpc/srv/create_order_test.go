package srv_test

import (
	"context"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
	"github.com/stretchr/testify/require"

	"github.com/sonm-io/marketplace/infra/util"
)

func (s *MarketplaceTestSuite) createBidOrder() {
	// arrange
	buyerID := util.PubKeyToAddr(s.App.PublicKey()).Hex()
	order := &pb.Order{
		Id:        "1b5dfa00-af3c-4e2d-b64b-c5d62e89430b",
		OrderType: pb.OrderType_BID,
		Price:     "777",
		ByuerID:   buyerID,

		Slot: &pb.Slot{
			BuyerRating:    555,
			SupplierRating: 666,
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
	obtained, err := s.client.CreateOrder(context.Background(), order)

	// assert
	require.NoError(s.T(), err, "cannot create order")
	s.Equal(order.Id, obtained.Id)
}

func (s *MarketplaceTestSuite) createAskOrder() {
	// arrange
	supplierID := util.PubKeyToAddr(s.App.PublicKey()).Hex()
	order := &pb.Order{
		Id:         "fc018acd-d9a9-4b8a-a45f-f90456a469c1",
		OrderType:  pb.OrderType_ASK,
		Price:      "777",
		SupplierID: supplierID,

		Slot: &pb.Slot{
			SupplierRating: 555,
			Resources: &pb.Resources{
				CpuCores: 1,
				RamBytes: 100000000,
				Storage:  1000000000,

				NetworkType:  pb.NetworkType_INCOMING,
				NetTrafficIn: 500000,

				Properties: map[string]float64{
					"cycles": 42,
					"foo":    1101,
				},
			},
		},
	}

	// act
	obtained, err := s.client.CreateOrder(context.Background(), order)

	// assert
	require.NoError(s.T(), err, "cannot create order")
	s.Equal(order.Id, obtained.Id)
}
