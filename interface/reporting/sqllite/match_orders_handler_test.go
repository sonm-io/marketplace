package sqllite

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ds "github.com/sonm-io/marketplace/datastruct"
	sds "github.com/sonm-io/marketplace/infra/storage/sqllite/datastruct"

	"github.com/sonm-io/marketplace/interface/reporting/sqllite/mocks"

	"github.com/sonm-io/marketplace/usecase/marketplace/query"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/report"
)

func TestMatchOrdersHandlerHandle_ValidQueryGiven_OrdersReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orders := []ds.Order{
		{
			ID:         "test_order_101",
			OrderType:  ds.Ask,
			Price:      "101",
			SupplierID: "TestSupplier",
			Slot: &ds.Slot{
				Duration: 900,
			},
		},
		{
			ID:         "test_order_105",
			OrderType:  ds.Ask,
			Price:      "105",
			SupplierID: "TestSupplier2",
			Slot: &ds.Slot{
				Duration: 600,
			},
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

	q := query.GetOrders{
		Order: ds.Order{
			OrderType: ds.Ask,
			Slot:      &ds.Slot{},
		},
		Limit: 10,
	}

	var (
		orderRow  sds.OrderRow
		orderRows sds.OrderRows
	)

	for idx := range orders {
		orderRow = sds.OrderRow{}
		orderToRow(&orders[idx], &orderRow)
		orderRows = append(orderRows, orderRow)
	}

	stmt, err := MatchOrdersStmt(q.Order, q.Limit)
	require.NoError(t, err)

	sql, args, err := ToSQL(stmt)
	require.NoError(t, err)

	var rows sds.OrderRows
	storage := mocks.NewMockOrderRowsFetcher(ctrl)
	storage.EXPECT().FetchRows(&rows, sql, args...).
		SetArg(0, orderRows).
		Return(nil)

	h := NewMatchOrdersHandler(storage)

	// act
	var obtained report.GetOrdersReport
	err = h.Handle(q, &obtained)

	// assert
	require.NoError(t, err)
	assert.Equal(t, expected, obtained)
}

func TestMatchOrdersHandlerHandle_BuyerIDGiven_OrdersReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buyerID := "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db"
	q := query.GetOrders{
		Order: ds.Order{
			OrderType: ds.Bid,
			BuyerID:   buyerID,
		},
		Limit: 10,
	}

	orders := []ds.Order{
		{
			ID:        "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
			BuyerID:   "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db",
			OrderType: ds.Bid,
			Price:     "100",
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

	var (
		orderRow  sds.OrderRow
		orderRows sds.OrderRows
	)

	for idx := range orders {
		orderRow = sds.OrderRow{}
		orderToRow(&orders[idx], &orderRow)
		orderRows = append(orderRows, orderRow)
	}

	stmt, err := MatchOrdersStmt(q.Order, q.Limit)
	require.NoError(t, err)

	sql, args, err := ToSQL(stmt)
	require.NoError(t, err)

	var rows sds.OrderRows
	storage := mocks.NewMockOrderRowsFetcher(ctrl)
	storage.EXPECT().FetchRows(&rows, sql, args...).
		SetArg(0, orderRows).
		Return(nil)

	h := NewMatchOrdersHandler(storage)

	// act
	var obtained report.GetOrdersReport
	err = h.Handle(q, &obtained)

	// assert
	require.NoError(t, err)
	assert.Equal(t, expected, obtained)
}

func TestMatchOrdersHandlerHandle_IncorrectQueryResultGiven_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockOrderRowsFetcher(ctrl)
	h := NewMatchOrdersHandler(storage)

	// act
	result := &struct{}{}
	err := h.Handle(query.GetOrders{}, result)

	// assert
	assert.EqualError(t, err, fmt.Sprintf("invalid result %v given", result))
}

func TestMatchOrdersHandlerHandle_IncorrectQueryGiven_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockOrderRowsFetcher(ctrl)
	h := NewMatchOrdersHandler(storage)

	// act
	q := unknownQuery{}
	err := h.Handle(q, &report.GetOrdersReport{})

	// assert
	assert.EqualError(t, err, fmt.Sprintf("invalid query %v given", q))
}
