package query

import (
	"testing"

	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/sonm-io/marketplace/entity"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewGetOrderHandlerHandle_ExistingIDGiven_OrderReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := GetOrderResult{
		ID:        "test_order",
		BuyerID:   "TestBuyer",
		OrderType: int(entity.ASK),
		Price:     555,
	}

	order := entity.Order{
		ID:        "test_order",
		BuyerID:   "TestBuyer",
		OrderType: entity.ASK,
		Price:     555,
	}

	storage := mocks.NewMockOrderByIDStorage(ctrl)
	storage.EXPECT().ByID("test_order").Return(&order, nil)

	h := NewGetOrderHandler(storage)

	// act
	var obtained GetOrderResult
	err := h.Handle(GetOrder{ID: "test_order"}, &obtained)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, expected, obtained)
}

func TestCancelOrderHandlerHandle_IncorrectQueryResultGivenErrorReturned(t *testing.T) {
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

func TestCancelOrderHandlerHandle_IncorrectQueryGivenErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockOrderByIDStorage(ctrl)
	h := NewGetOrderHandler(storage)

	// act
	q := unknownQuery{}
	err := h.Handle(q, &GetOrderResult{})

	// assert
	assert.EqualError(t, err, fmt.Sprintf("invalid query %v given", q))
}

type unknownQuery struct{}

func (c unknownQuery) QueryID() string {
	return "UnknownQuery"
}
