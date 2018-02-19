package srv

import (
	"context"
	"testing"

	"github.com/sonm-io/marketplace/infra/util"
	pb "github.com/sonm-io/marketplace/proto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (s *MarketplaceTestSuite) getBidOrderByID(t *testing.T) {
	// arrange

	// smth like "0x9B27D3C3571731deDb23EaFEa34a3a6E05daE159"
	BuyerID := util.PubKeyToAddr(s.App.PublicKey()).Hex()

	pricePerSecond, err := pb.NewBigIntFromString("777")
	require.NoError(t, err)

	expected := &pb.Order{
		Id:             "1b5dfa00-af3c-4e2d-b64b-c5d62e89430b",
		OrderType:      pb.OrderType_BID,
		PricePerSecond: pricePerSecond,
		ByuerID:        BuyerID,

		Slot: &pb.Slot{
			Duration:       900,
			BuyerRating:    555,
			SupplierRating: 666,
			Resources: &pb.Resources{
				CpuCores: 1,
				GpuCount: pb.GPUCount_MULTIPLE_GPU,
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
	require.NoError(t, err)
	assert.Equal(t, expected, obtained)
}

func (s *MarketplaceTestSuite) getInExistentOrder(t *testing.T) {
	// act
	_, err := s.client.GetOrderByID(context.Background(), &pb.ID{Id: "non-existent-order"})

	// assert
	assert.EqualError(t, err,
		`rpc error: code = Internal desc = cannot get order: order non-existent-order is not found`)
}
