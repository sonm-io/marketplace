package query

import (
	"testing"

	"fmt"
	"github.com/golang/mock/gomock"
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/mocks"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/report"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/spec"
	"github.com/stretchr/testify/assert"
)

func TestGetOrdersHandlerHandle_ValidCommandGiven_OrderReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orders := []ds.Order{
		{
			ID:        "test_order_101",
			OrderType: ds.Ask,
			Price:     101,
			BuyerID:   "TestBuyer",
		},
		{
			ID:        "test_order_105",
			OrderType: ds.Ask,
			Price:     105,
			BuyerID:   "TestBuyer",
		},
	}

	expected := report.GetOrdersReport{
		{
			Order: orders[0],
		},
		{
			Order: orders[1],
		},
	}

	q := GetOrders{
		OrderType: int(ds.Ask),
		Limit:     10,
	}

	s, _ := spec.OrdersBySlot(ds.OrderType(q.OrderType), ds.Slot{
		BuyerRating:    q.Slot.BuyerRating,
		SupplierRating: q.Slot.SupplierRating,
	})

	storage := mocks.NewMockOrderBySpecStorage(ctrl)
	storage.EXPECT().
		BySpecWithLimit(s, uint64(10)).
		Return(orders, nil)

	h := NewGetOrdersHandler(storage)

	// act
	var obtained report.GetOrdersReport

	err := h.Handle(q, &obtained)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, expected, obtained)
}

func TestGetOrdersHandlerHandle_IncorrectQueryResultGivenErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockOrderBySpecStorage(ctrl)
	h := NewGetOrdersHandler(storage)

	// act
	result := &struct{}{}
	err := h.Handle(GetOrders{}, result)

	// assert
	assert.EqualError(t, err, fmt.Sprintf("invalid result %v given", result))
}

func TestGetOrdersHandlerHandle_IncorrectQueryGivenErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockOrderBySpecStorage(ctrl)
	h := NewGetOrdersHandler(storage)

	// act
	q := unknownQuery{}
	err := h.Handle(q, &report.GetOrdersReport{})

	// assert
	assert.EqualError(t, err, fmt.Sprintf("invalid query %v given", q))
}
