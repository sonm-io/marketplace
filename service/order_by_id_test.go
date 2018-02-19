package service

import (
	"fmt"
	"testing"

	"github.com/gocraft/dbr"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sonm-io/marketplace/ds"
	pb "github.com/sonm-io/marketplace/handler/proto"
	mds "github.com/sonm-io/marketplace/mapper/datastruct"

	"github.com/sonm-io/marketplace/mapper"
	"github.com/sonm-io/marketplace/service/mocks"
)

func TestOrderByIDHandlerHandle_ExistingIDGiven_OrderReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	price, err := pb.NewBigIntFromString("555")
	require.NoError(t, err)

	expected := pb.Order{
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
	}

	var orderRow mds.OrderRow
	mapper.OrderToRow(&ds.Order{Order: &expected}, &orderRow)
	orderRow.Status = mds.Active

	sql, args, err := ToSQL(OrderByIDStmt("test_order"))
	require.NoError(t, err)

	var row mds.OrderRow
	storage := mocks.NewMockStorage(ctrl)
	storage.EXPECT().FetchRow(&row, sql, args...).
		SetArg(0, orderRow).
		Return(nil)

	srv := NewMarketService(storage)

	// act
	var obtained pb.Order
	err = srv.OrderByID("test_order", &obtained)

	// assert
	require.NoError(t, err)
	assert.Equal(t, expected, obtained)
}

func TestOrderByIDHandlerHandle_InExistentOrderIDGiven_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sql, args, err := ToSQL(OrderByIDStmt("test_order"))
	require.NoError(t, err)

	var row mds.OrderRow
	storage := mocks.NewMockStorage(ctrl)
	storage.EXPECT().FetchRow(&row, sql, args...).
		Return(dbr.ErrNotFound)

	srv := NewMarketService(storage)

	// act
	var obtained pb.Order
	err = srv.OrderByID("test_order", &obtained)

	// assert
	assert.EqualError(t, err, "order test_order is not found")
}

func TestOrderByIDHandlerHandle_StorageErrorOccurred_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sql, args, err := ToSQL(OrderByIDStmt("test_order"))
	require.NoError(t, err)

	var row mds.OrderRow
	storage := mocks.NewMockStorage(ctrl)
	storage.EXPECT().FetchRow(&row, sql, args...).
		Return(fmt.Errorf("sql: some error"))

	srv := NewMarketService(storage)

	// act
	var obtained pb.Order
	err = srv.OrderByID("test_order", &obtained)

	// assert
	assert.EqualError(t, err, "cannot retrieve order: sql: some error")
}

func TestOrderByIDHandlerHandle_InactiveOrderIDGiven_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var orderRow mds.OrderRow
	orderRow.Status = uint8(mds.Cancelled)

	sql, args, err := ToSQL(OrderByIDStmt("inactive_order"))
	require.NoError(t, err)

	var row mds.OrderRow
	storage := mocks.NewMockStorage(ctrl)
	storage.EXPECT().FetchRow(&row, sql, args...).
		SetArg(0, orderRow).
		Return(nil)

	srv := NewMarketService(storage)

	// act
	var obtained pb.Order
	err = srv.OrderByID("inactive_order", &obtained)

	// assert
	assert.EqualError(t, err, "order inactive_order is inactive")
}

func TestOrderByIDHandlerHandle_IncorrectQueryResultGiven_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockStorage(ctrl)
	srv := NewMarketService(storage)

	// act
	result := &struct{}{}
	err := srv.OrderByID("123", result)

	// assert
	assert.EqualError(t, err, fmt.Sprintf("invalid result %v given", result))
}
