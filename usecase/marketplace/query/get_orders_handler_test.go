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

func TestGetOrdersHandlerHandle_ValidQueryGiven_OrdersReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orders := []ds.Order{
		{
			ID:        "test_order_101",
			OrderType: ds.Ask,
			Price:     "101",
			BuyerID:   "TestBuyer",
		},
		{
			ID:        "test_order_105",
			OrderType: ds.Ask,
			Price:     "105",
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
		Order: ds.Order{
			OrderType: ds.Ask,
			Slot:      &ds.Slot{},
		},
		Limit: 10,
	}

	s, _ := spec.MatchOrders(q.Order)

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

func TestGetOrdersHandlerHandle_BuyerIDGiven_OrdersReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buyerID := "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db"
	q := GetOrders{
		Order: ds.Order{
			BuyerID: buyerID,
		},
		Limit: 10,
	}

	orders := []ds.Order{
		{
			ID:      "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
			BuyerID: "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db",
			Price:   "100",
			Slot: &ds.Slot{
				Duration: 900,
				Resources: ds.Resources{
					CPUCores: 4,
					RAMBytes: 10000,
				},
			},
		},
	}

	expected := report.GetOrdersReport{
		{
			Order: orders[0],
		},
	}

	s, _ := spec.MatchOrders(q.Order)

	storage := mocks.NewMockOrderBySpecStorage(ctrl)
	storage.EXPECT().
		BySpecWithLimit(s, uint64(10)).
		Return(orders, nil)

	h := NewGetOrdersHandler(storage)

	// act
	var obtained report.GetOrdersReport

	err := h.Handle(q, &obtained)

	assert.NoError(t, err)
	assert.Equal(t, expected, obtained)
}
