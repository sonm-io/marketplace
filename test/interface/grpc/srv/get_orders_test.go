package srv_test

import (
	"context"
	"testing"

	"github.com/sonm-io/marketplace/infra/util"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (s *MarketplaceTestSuite) getBidOrdersByBuyerID(t *testing.T) {
	// arrange

	// e.g 0x9B27D3C3571731deDb23EaFEa34a3a6E05daE159
	buyerID := util.PubKeyToAddr(s.App.PublicKey()).Hex()

	pricePerSecond, err := pb.NewBigIntFromString("777")
	require.NoError(t, err)

	req := &pb.GetOrdersRequest{
		Order: &pb.Order{
			ByuerID:   buyerID,
			OrderType: pb.OrderType_BID,
		},
		Count: 100,
	}

	expected := &pb.GetOrdersReply{
		Orders: []*pb.Order{
			{
				Id:             "1b5dfa00-af3c-4e2d-b64b-c5d62e89430b",
				OrderType:      pb.OrderType_BID,
				PricePerSecond: pricePerSecond,
				ByuerID:        buyerID,

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
			},
		},
	}

	// act
	obtained, err := s.client.GetOrders(context.Background(), req)

	// assert
	require.NoError(t, err)
	assert.Equal(t, expected, obtained)
}

func (s *MarketplaceTestSuite) getBidOrders(t *testing.T) {
	// arrange
	req := &pb.GetOrdersRequest{
		Order: &pb.Order{
			OrderType: pb.OrderType_BID,
			Slot: &pb.Slot{
				Resources: &pb.Resources{
					CpuCores:     4,
					GpuCount:     pb.GPUCount_MULTIPLE_GPU,
					NetworkType:  pb.NetworkType_INCOMING,
					NetTrafficIn: 500000,
				},
			},
		},
		Count: 100,
	}

	// smth like "0x9B27D3C3571731deDb23EaFEa34a3a6E05daE159"
	buyerID := util.PubKeyToAddr(s.App.PublicKey()).Hex()

	pricePerSecond, err := pb.NewBigIntFromString("777")
	require.NoError(t, err)

	expected := &pb.GetOrdersReply{
		Orders: []*pb.Order{
			{
				Id:             "1b5dfa00-af3c-4e2d-b64b-c5d62e89430b",
				OrderType:      pb.OrderType_BID,
				PricePerSecond: pricePerSecond,
				ByuerID:        buyerID,

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
			},
		},
	}

	// act
	obtained, err := s.client.GetOrders(context.Background(), req)

	// assert
	require.NoError(t, err)
	assert.Equal(t, expected, obtained)
}

func (s *MarketplaceTestSuite) getBidOrdersMatchingProperties(t *testing.T) {
	// arrange
	req := &pb.GetOrdersRequest{
		Order: &pb.Order{
			OrderType: pb.OrderType_BID,
			Slot: &pb.Slot{
				Resources: &pb.Resources{
					CpuCores:     4,
					GpuCount:     pb.GPUCount_SINGLE_GPU,
					NetworkType:  pb.NetworkType_INCOMING,
					NetTrafficIn: 500000,
					Properties: map[string]float64{
						"hash_rate": 777,
					},
				},
			},
		},
		Count: 100,
	}

	// smth like "0x9B27D3C3571731deDb23EaFEa34a3a6E05daE159"
	buyerID := util.PubKeyToAddr(s.App.PublicKey()).Hex()

	pricePerSecond, err := pb.NewBigIntFromString("777")
	require.NoError(t, err)

	expected := &pb.GetOrdersReply{
		Orders: []*pb.Order{
			{
				Id:             "1b5dfa00-af3c-4e2d-b64b-c5d62e89430b",
				OrderType:      pb.OrderType_BID,
				PricePerSecond: pricePerSecond,
				ByuerID:        buyerID,

				Slot: &pb.Slot{
					Duration:       900,
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
			},
		},
	}

	// act
	obtained, err := s.client.GetOrders(context.Background(), req)

	// assert
	require.NoError(t, err)
	assert.Equal(t, expected, obtained)
}

func (s *MarketplaceTestSuite) getBidOrdersWithInExistentProperties(t *testing.T) {
	// arrange
	req := &pb.GetOrdersRequest{
		Order: &pb.Order{
			OrderType: pb.OrderType_BID,
			Slot: &pb.Slot{
				Resources: &pb.Resources{
					CpuCores:     4,
					GpuCount:     pb.GPUCount_SINGLE_GPU,
					NetworkType:  pb.NetworkType_INCOMING,
					NetTrafficIn: 500000,
					Properties: map[string]float64{
						"some_unknown_property": 42,
					},
				},
			},
		},
		Count: 100,
	}

	expected := &pb.GetOrdersReply{}

	// act
	obtained, err := s.client.GetOrders(context.Background(), req)

	// assert
	require.NoError(t, err)
	assert.Equal(t, expected, obtained)
}

func (s *MarketplaceTestSuite) getAskOrders(t *testing.T) {
	// arrange
	req := &pb.GetOrdersRequest{
		Order: &pb.Order{
			OrderType: pb.OrderType_ASK,
			Slot: &pb.Slot{
				Duration:       900,
				SupplierRating: 555,
				Resources: &pb.Resources{
					CpuCores:     1,
					RamBytes:     100000000,
					Storage:      1000000000,
					NetworkType:  pb.NetworkType_INCOMING,
					NetTrafficIn: 500000,
					Properties: map[string]float64{
						"cycles": 42,
						"foo":    1101,
					},
				},
			},
		},
		Count: 100,
	}

	// smth like "0x8125721C2413d99a33E351e1F6Bb4e56b6b633FD"
	supplierID := util.PubKeyToAddr(s.App.PublicKey()).Hex()

	pricePerSecond, err := pb.NewBigIntFromString("777")
	require.NoError(t, err)

	expected := &pb.GetOrdersReply{
		Orders: []*pb.Order{
			{
				Id:             "fc018acd-d9a9-4b8a-a45f-f90456a469c1",
				OrderType:      pb.OrderType_ASK,
				PricePerSecond: pricePerSecond,
				SupplierID:     supplierID,

				Slot: &pb.Slot{
					Duration:       600,
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
			},
		},
	}

	// act
	obtained, err := s.client.GetOrders(context.Background(), req)

	// assert
	require.NoError(t, err)
	assert.Equal(t, expected, obtained)
}

func (s *MarketplaceTestSuite) getAskOrdersMatchingProperties(t *testing.T) {
	// arrange
	req := &pb.GetOrdersRequest{
		Order: &pb.Order{
			OrderType: pb.OrderType_ASK,
			Slot: &pb.Slot{
				Duration:       900,
				SupplierRating: 555,
				Resources: &pb.Resources{
					CpuCores:     1,
					RamBytes:     100000000,
					Storage:      1000000000,
					NetworkType:  pb.NetworkType_INCOMING,
					NetTrafficIn: 500000,
					Properties: map[string]float64{
						"cycles": 42,
						"foo":    777,
					},
				},
			},
		},
		Count: 100,
	}

	// smth like "0x8125721C2413d99a33E351e1F6Bb4e56b6b633FD"
	supplierID := util.PubKeyToAddr(s.App.PublicKey()).Hex()

	pricePerSecond, err := pb.NewBigIntFromString("777")
	require.NoError(t, err)

	expected := &pb.GetOrdersReply{
		Orders: []*pb.Order{
			{
				Id:             "fc018acd-d9a9-4b8a-a45f-f90456a469c1",
				OrderType:      pb.OrderType_ASK,
				PricePerSecond: pricePerSecond,
				SupplierID:     supplierID,

				Slot: &pb.Slot{
					Duration:       600,
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
			},
		},
	}

	// act
	obtained, err := s.client.GetOrders(context.Background(), req)

	// assert
	require.NoError(t, err)
	assert.Equal(t, expected, obtained)
}
