package query

import (
	"testing"

	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/mocks"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/report"
	"github.com/stretchr/testify/assert"
)

func TestGetOrderHandlerHandle_ExistingIDGiven_OrderReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := report.GetOrderReport{
		ID:        "test_order",
		BuyerID:   "TestBuyer",
		OrderType: int(report.ASK),
		Price:     555,
		Slot: report.Slot{
			BuyerRating:    0,
			SupplierRating: 0,
		},
	}

	storage := mocks.NewMockOrderByIDStorage(ctrl)
	storage.EXPECT().ByID("test_order").Return(expected, nil)

	h := NewGetOrderHandler(storage)

	// act
	var obtained report.GetOrderReport
	err := h.Handle(GetOrder{ID: "test_order"}, &obtained)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, expected, obtained)
}

func TestGetOrderHandlerHandle_IncorrectQueryResultGivenErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockOrderByIDStorage(ctrl)
	h := NewGetOrderHandler(storage)

	// act
	result := &struct{}{}
	err := h.Handle(GetOrder{}, result)

	// assert
	assert.EqualError(t, err, fmt.Sprintf("invalid result %v given", result))
}

func TestGetOrderHandlerHandle_IncorrectQueryGivenErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockOrderByIDStorage(ctrl)
	h := NewGetOrderHandler(storage)

	// act
	q := unknownQuery{}
	err := h.Handle(q, &report.GetOrderReport{})

	// assert
	assert.EqualError(t, err, fmt.Sprintf("invalid query %v given", q))
}
