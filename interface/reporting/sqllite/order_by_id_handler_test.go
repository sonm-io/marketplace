package sqllite

import (
	"testing"
	"fmt"
	"database/sql"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	ds "github.com/sonm-io/marketplace/datastruct"
	sds "github.com/sonm-io/marketplace/infra/storage/sqllite/datastruct"

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
			ID:        "test_order",
			BuyerID:   "TestBuyer",
			OrderType: ds.Ask,
			Price:     "555",
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
	orderToRow(&expected.Order, &orderRow)
	orderRow.Status = Active

	storage := mocks.NewMockOrderRowFetcher(ctrl)
	storage.EXPECT().FetchRow("test_order", &sds.OrderRow{}).
		SetArg(1, orderRow).
		Return(nil)

	h := NewOrderByIDHandler(storage)

	// act
	var obtained report.GetOrderReport
	err := h.Handle(query.GetOrder{ID: "test_order"}, &obtained)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, expected, obtained)
}

func TestOrderByIDHandlerHandle_InExistentOrderIDGiven_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockOrderRowFetcher(ctrl)
	storage.EXPECT().FetchRow("test_order", &sds.OrderRow{}).
		Return(sql.ErrNoRows)

	h := NewOrderByIDHandler(storage)

	// act
	var obtained report.GetOrderReport
	err := h.Handle(query.GetOrder{ID: "test_order"}, &obtained)

	// assert
	assert.EqualError(t, err, "order test_order is not found")
}

func TestOrderByIDHandlerHandle_StorageErrorOccurred_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockOrderRowFetcher(ctrl)
	storage.EXPECT().FetchRow("test_order", &sds.OrderRow{}).
		Return(fmt.Errorf("sql: some error"))

	h := NewOrderByIDHandler(storage)

	// act
	var obtained report.GetOrderReport
	err := h.Handle(query.GetOrder{ID: "test_order"}, &obtained)

	// assert
	assert.EqualError(t, err, "an error occured: sql: some error")
}

func TestOrderByIDHandlerHandle_InactiveOrderIDGiven_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var orderRow sds.OrderRow
	orderRow.Status = uint8(InActive)

	storage := mocks.NewMockOrderRowFetcher(ctrl)
	storage.EXPECT().FetchRow("inactive_order", &sds.OrderRow{}).
		SetArg(1, orderRow).
		Return(nil)

	h := NewOrderByIDHandler(storage)

	// act
	var obtained report.GetOrderReport
	err := h.Handle(query.GetOrder{ID: "inactive_order"}, &obtained)

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

