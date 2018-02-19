package srv

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pb "github.com/sonm-io/marketplace/handler/proto"
	"github.com/sonm-io/marketplace/handler/srv/mocks"
)

//
//func TestMarketplaceGetOrdersByID_BuyerIDGiven_CorrespondentOrdersReturned(t *testing.T) {
//	// arrange
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	pricePerSecond, err := pb.NewBigIntFromString("100")
//	require.NoError(t, err)
//
//	buyerID := "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db"
//	req := &pb.GetOrdersRequest{
//		Order: &pb.Order{
//			ByuerID: buyerID,
//		},
//		Count: 100,
//	}
//
//	q := query.GetOrders{}
//	bindGetOrdersQuery(req, &q)
//
//	expected := &pb.GetOrdersReply{
//		Orders: []*pb.Order{
//			{
//				Id:             "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
//				ByuerID:        buyerID,
//				PricePerSecond: pricePerSecond,
//				Slot: &pb.Slot{
//					Duration: 900,
//					Resources: &pb.Resources{
//						CpuCores: 4,
//						RamBytes: 10000,
//					},
//				},
//			},
//		},
//	}
//
//	ordersReport := []report.GetOrderReport{
//		{
//			Order: ds.Order{
//				ID:             "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
//				BuyerID:        "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db",
//				PricePerSecond: "100",
//				Slot: &ds.Slot{
//					Duration: 900,
//					Resources: ds.Resources{
//						CPUCores: 4,
//						RAMBytes: 10000,
//					},
//				},
//			},
//		},
//	}
//
//	ordersBySpec := mocks.NewMockQueryHandler(ctrl)
//	ordersBySpec.EXPECT().Handle(q, &report.GetOrdersReport{}).
//		SetArg(1, ordersReport).
//		Return(nil)
//
//	m := NewMarketplace(nil, nil, ordersBySpec)
//
//	obtained, err := m.GetOrders(context.Background(), req)
//
//	assert.NoError(t, err)
//	assert.Equal(t, expected, obtained)
//}
//
//func TestMarketplaceGetOrdersByID_SlotGiven_CorrespondentOrdersReturned(t *testing.T) {
//	// arrange
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	pricePerSecond, err := pb.NewBigIntFromString("100")
//	require.NoError(t, err)
//
//	buyerID := "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db"
//	req := &pb.GetOrdersRequest{
//		Order: &pb.Order{
//			OrderType: pb.OrderType_BID,
//			ByuerID:   buyerID,
//			Slot: &pb.Slot{
//				Resources: &pb.Resources{
//					CpuCores: 2,
//				},
//			},
//		},
//		Count: 100,
//	}
//
//	q := query.GetOrders{}
//	bindGetOrdersQuery(req, &q)
//
//	expected := &pb.GetOrdersReply{
//		Orders: []*pb.Order{
//			{
//				Id:             "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
//				ByuerID:        buyerID,
//				PricePerSecond: pricePerSecond,
//				Slot: &pb.Slot{
//					Duration: 900,
//					Resources: &pb.Resources{
//						CpuCores: 4,
//						RamBytes: 10000,
//					},
//				},
//			},
//		},
//	}
//
//	ordersReport := []report.GetOrderReport{
//		{
//			Order: ds.Order{
//				ID:             "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
//				BuyerID:        "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db",
//				PricePerSecond: "100",
//				Slot: &ds.Slot{
//					Duration: 900,
//					Resources: ds.Resources{
//						CPUCores: 4,
//						RAMBytes: 10000,
//					},
//				},
//			},
//		},
//	}
//
//	ordersBySpec := mocks.NewMockQueryHandler(ctrl)
//	ordersBySpec.EXPECT().Handle(q, &report.GetOrdersReport{}).
//		SetArg(1, ordersReport).
//		Return(nil)
//
//	m := NewMarketplace(nil, nil, ordersBySpec)
//
//	obtained, err := m.GetOrders(context.Background(), req)
//
//	assert.NoError(t, err)
//	assert.Equal(t, expected, obtained)
//}

func TestMarketplace_GetOrders_ValidRequestGiven_ValidResponse(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	order := &pb.Order{}
	req := &pb.GetOrdersRequest{
		Order: order,
		Count: 10,
	}

	serviceMock := mocks.NewMockMarketService(ctrl)
	serviceMock.EXPECT().
		MatchOrders(order, uint64(10), &pb.GetOrdersReply{})

	m := NewMarketplace(serviceMock)
	// act
	_, err := m.GetOrders(context.Background(), req)

	// assert
	assert.NoError(t, err)
}

func TestMarketplaceGetOrders_AnErrorOccurredWhileGettingOrders_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	order := &pb.Order{}
	req := &pb.GetOrdersRequest{
		Order: order,
		Count: 10,
	}

	serviceMock := mocks.NewMockMarketService(ctrl)
	serviceMock.EXPECT().
		MatchOrders(order, uint64(10), &pb.GetOrdersReply{}).
		Return(fmt.Errorf("some error"))

	m := NewMarketplace(serviceMock)

	// act
	_, err := m.GetOrders(context.Background(), req)

	// assert
	assert.EqualError(t, err, "rpc error: code = Internal desc = cannot get orders: some error")
}

func TestMarketplaceGetOrders_CountIsNotSet_DefaultLimitApplied(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	order := &pb.Order{}
	req := &pb.GetOrdersRequest{
		Order: order,
		Count: 0,
	}

	serviceMock := mocks.NewMockMarketService(ctrl)
	serviceMock.EXPECT().
		MatchOrders(order, uint64(defaultResultsCount), &pb.GetOrdersReply{}).Times(1)

	m := NewMarketplace(serviceMock)

	// act
	_, err := m.GetOrders(context.Background(), req)

	// assert
	require.NoError(t, err)
}
