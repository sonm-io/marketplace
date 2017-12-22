package srv_test

import (
	"context"
	"github.com/sonm-io/marketplace/infra/util"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
)

func (s *MarketplaceTestSuite) getBidOrderByID() {

	// smth like "0x9B27D3C3571731deDb23EaFEa34a3a6E05daE159"
	BuyerID := util.PubKeyToAddr(s.App.PublicKey()).Hex()

	expected := &pb.Order{
		Id:        "1b5dfa00-af3c-4e2d-b64b-c5d62e89430b",
		OrderType: pb.OrderType_BID,
		Price:     "777",
		ByuerID:   BuyerID,

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
	obtained, err := s.client.GetOrderByID(context.Background(), &pb.ID{Id: "1b5dfa00-af3c-4e2d-b64b-c5d62e89430b"})

	// assert
	s.NoError(err)
	s.Equal(expected, obtained)
}

func (s *MarketplaceTestSuite) getInExistentOrder() {

	// act
	_, err := s.client.GetOrderByID(context.Background(), &pb.ID{Id: "non-existent-order"})

	// assert
	s.EqualError(err,
		`rpc error: code = Internal desc = cannot get order: order non-existent-order is not found`)
}
