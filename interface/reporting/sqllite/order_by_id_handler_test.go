package sqllite

import (
	"fmt"
	"testing"

	"github.com/gocraft/dbr"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ds "github.com/sonm-io/marketplace/datastruct"
	sds "github.com/sonm-io/marketplace/interface/mapper/datastruct"

	"github.com/sonm-io/marketplace/interface/mapper"
	"github.com/sonm-io/marketplace/interface/reporting/sqllite/mocks"

	"github.com/sonm-io/marketplace/usecase/marketplace/query"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/report"
)

func TestOrderByIDHandlerHandle_ExistingIDGiven_OrderReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := report.GetOrderReport{
		Order: ds.Order{
			ID:             "test_order",
			BuyerID:        "TestBuyer",
			OrderType:      ds.Ask,
			PricePerSecond: "555",
			Slot: &ds.Slot{
				BuyerRating:    444,
				SupplierRating: 0,
				Resources: ds.Resources{
					CPUCores:      4,
					RAMBytes:      12000,
					GPUCount:      ds.SingleGPU,
					Storage:       120000,
					NetworkType:   ds.Inbound,
					NetTrafficIn:  100000,
					NetTrafficOut: 10000,
				},
			},
		},
	}

	var orderRow sds.OrderRow
	mapper.OrderToRow(&expected.Order, &orderRow)
	orderRow.Status = Active

	stmt, err := GetOrderByIDStmt("test_order")
	require.NoError(t, err)

	sql, args, err := ToSQL(stmt)
	require.NoError(t, err)

	var row sds.OrderRow
	storage := mocks.NewMockOrderRowFetcher(ctrl)
	storage.EXPECT().FetchRow(&row, sql, args...).
		SetArg(0, orderRow).
		Return(nil)

	h := NewOrderByIDHandler(storage)

	// act
	var obtained report.GetOrderReport
	err = h.Handle(query.GetOrder{ID: "test_order"}, &obtained)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, expected, obtained)
}

func TestOrderByIDHandlerHandle_InExistentOrderIDGiven_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	stmt, err := GetOrderByIDStmt("test_order")
	require.NoError(t, err)

	sql, args, err := ToSQL(stmt)
	require.NoError(t, err)

	var row sds.OrderRow
	storage := mocks.NewMockOrderRowFetcher(ctrl)
	storage.EXPECT().FetchRow(&row, sql, args...).
		Return(dbr.ErrNotFound)

	h := NewOrderByIDHandler(storage)

	// act
	var obtained report.GetOrderReport
	err = h.Handle(query.GetOrder{ID: "test_order"}, &obtained)

	// assert
	assert.EqualError(t, err, "order test_order is not found")
}

func TestOrderByIDHandlerHandle_StorageErrorOccurred_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	stmt, err := GetOrderByIDStmt("test_order")
	require.NoError(t, err)

	sql, args, err := ToSQL(stmt)
	require.NoError(t, err)

	var row sds.OrderRow
	storage := mocks.NewMockOrderRowFetcher(ctrl)
	storage.EXPECT().FetchRow(&row, sql, args...).
		Return(fmt.Errorf("sql: some error"))

	h := NewOrderByIDHandler(storage)

	// act
	var obtained report.GetOrderReport
	err = h.Handle(query.GetOrder{ID: "test_order"}, &obtained)

	// assert
	assert.EqualError(t, err, "an error occured: sql: some error")
}

func TestOrderByIDHandlerHandle_InactiveOrderIDGiven_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var orderRow sds.OrderRow
	orderRow.Status = uint8(InActive)

	stmt, err := GetOrderByIDStmt("inactive_order")
	require.NoError(t, err)

	sql, args, err := ToSQL(stmt)
	require.NoError(t, err)

	var row sds.OrderRow
	storage := mocks.NewMockOrderRowFetcher(ctrl)
	storage.EXPECT().FetchRow(&row, sql, args...).
		SetArg(0, orderRow).
		Return(nil)

	h := NewOrderByIDHandler(storage)

	// act
	var obtained report.GetOrderReport
	err = h.Handle(query.GetOrder{ID: "inactive_order"}, &obtained)

	// assert
	assert.EqualError(t, err, "order inactive_order is inactive")
}

func TestOrderByIDHandlerHandle_IncorrectQueryGiven_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockOrderRowFetcher(ctrl)
	h := NewOrderByIDHandler(storage)

	// act
	q := unknownQuery{}
	err := h.Handle(q, &report.GetOrderReport{})

	// assert
	assert.EqualError(t, err, fmt.Sprintf("invalid query %v given", q))
}

func TestOrderByIDHandlerHandle_IncorrectQueryResultGiven_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockOrderRowFetcher(ctrl)
	h := NewOrderByIDHandler(storage)

	// act
	result := &struct{}{}
	err := h.Handle(query.GetOrder{}, result)

	// assert
	assert.EqualError(t, err, fmt.Sprintf("invalid result %v given", result))
}
