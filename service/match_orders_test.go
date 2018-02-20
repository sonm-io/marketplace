package service

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sonm-io/marketplace/ds"
	mds "github.com/sonm-io/marketplace/mapper/datastruct"
	pb "github.com/sonm-io/marketplace/proto"

	"github.com/sonm-io/marketplace/mapper"
	"github.com/sonm-io/marketplace/service/mocks"
)

func TestMarketServiceMatchOrders_ValidMatchingCriteriaWithPropertiesGiven_OrdersReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	price, err := pb.NewBigIntFromString("555")
	require.NoError(t, err)

	expected := &pb.GetOrdersReply{
		Orders: []*pb.Order{
			{
				Id:             "test_order",
				ByuerID:        "TestBuyer",
				OrderType:      pb.OrderType_ASK,
				PricePerSecond: price,
				Slot: &pb.Slot{
					BuyerRating:    444,
					SupplierRating: 0,
					Resources: &pb.Resources{
						CpuCores:      4,
						RamBytes:      12000,
						GpuCount:      pb.GPUCount_MULTIPLE_GPU,
						Storage:       120000,
						NetworkType:   pb.NetworkType_INCOMING,
						NetTrafficIn:  100000,
						NetTrafficOut: 10000,
						Properties: map[string]float64{
							"foo": 33.7,
						},
					},
				},
			},
		},
	}

	req := &pb.Order{
		OrderType: pb.OrderType_BID,
		Slot: &pb.Slot{
			Resources: &pb.Resources{
				Properties: map[string]float64{
					"foo": 42.7,
				},
			},
		},
	}

	orderRows := rowsFromOrders(expected.Orders)

	var rows mds.OrderRows
	storage := mocks.NewMockStorage(ctrl)
	storage.EXPECT().FetchRows(&rows, gomock.Any(), gomock.Any()).
		SetArg(0, orderRows).
		Return(nil)

	srv := NewMarketService(storage)

	// act
	var obtained pb.GetOrdersReply
	err = srv.MatchOrders(req, 123, &obtained)

	// assert
	require.NoError(t, err)
	assert.Equal(t, expected, &obtained)
}

func TestMarketServiceMatchOrders_ValidMatchingCriteriaGiven_OrdersReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	price, err := pb.NewBigIntFromString("555")
	require.NoError(t, err)

	expected := &pb.GetOrdersReply{
		Orders: []*pb.Order{
			{
				Id:             "test_order",
				ByuerID:        "TestBuyer",
				OrderType:      pb.OrderType_ASK,
				PricePerSecond: price,
				Slot: &pb.Slot{
					BuyerRating:    444,
					SupplierRating: 0,
					Resources: &pb.Resources{
						CpuCores:      4,
						RamBytes:      12000,
						GpuCount:      pb.GPUCount_MULTIPLE_GPU,
						Storage:       120000,
						NetworkType:   pb.NetworkType_INCOMING,
						NetTrafficIn:  100000,
						NetTrafficOut: 10000,
					},
				},
			},
		},
	}

	req := &pb.Order{
		OrderType: pb.OrderType_BID,
	}

	orderRows := rowsFromOrders(expected.Orders)

	var rows mds.OrderRows
	storage := mocks.NewMockStorage(ctrl)
	storage.EXPECT().FetchRows(&rows, gomock.Any(), gomock.Any()).
		SetArg(0, orderRows).
		Return(nil)

	srv := NewMarketService(storage)

	// act
	var obtained pb.GetOrdersReply
	err = srv.MatchOrders(req, 123, &obtained)

	// assert
	require.NoError(t, err)
	assert.Equal(t, expected, &obtained)
}

func rowsFromOrders(orders []*pb.Order) mds.OrderRows {
	var (
		order     ds.Order
		orderRow  mds.OrderRow
		orderRows mds.OrderRows
	)
	for idx := range orders {
		order = ds.Order{Order: orders[idx]}
		mapper.OrderToRow(&order, &orderRow)
		orderRows = append(orderRows, orderRow)
	}
	return orderRows
}

func TestMarketServiceMatchOrders_StorageErrorOccurred_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockStorage(ctrl)
	storage.EXPECT().FetchRows(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(fmt.Errorf("sql: some error"))

	srv := NewMarketService(storage)

	// act
	err := srv.MatchOrders(&pb.Order{OrderType: pb.OrderType_BID}, 123, &pb.GetOrdersReply{})

	// assert
	assert.EqualError(t, err, "cannot match orders: sql: some error")
}

func TestMarketServiceMatchOrders_IncorrectResultGiven_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockStorage(ctrl)
	srv := NewMarketService(storage)

	// act
	result := &struct{}{}
	err := srv.MatchOrders(&pb.Order{}, 123, result)

	// assert
	assert.EqualError(t, err, fmt.Sprintf("invalid result %v given", result))
}

func TestMarketServiceMatchOrders_IncorrectOrderTypeGiven_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockStorage(ctrl)
	srv := NewMarketService(storage)

	// act
	err := srv.MatchOrders(&pb.Order{}, 123, &pb.GetOrdersReply{})

	// assert
	assert.EqualError(t, err, "cannot build conditions: searching by any type is not supported")
}
