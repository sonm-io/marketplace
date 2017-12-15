package inmemory

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/infra/storage/inmemory"
	"github.com/sonm-io/marketplace/interface/storage/inmemory/mocks"
)

func TestOrderStorageByID_ExistingIDGiven_OrderReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	order := &ds.Order{}
	order.ID = "test_order"

	engineMock := mocks.NewMockEngine(ctrl)
	engineMock.EXPECT().Get("test_order").Return(order, nil)

	s := NewOrderStorage(engineMock)

	// act
	obtained, err := s.ByID("test_order")

	// assert
	assert.NoError(t, err, "non-error result expected")
	assert.Equal(t, order, &obtained)
}

func TestOrderStorageAdd_ValidOrderGiven_OrderStored(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := ds.Order{}
	expected.ID = "test_order"

	engineMock := mocks.NewMockEngine(ctrl)
	engineMock.EXPECT().Add(&expected, "test_order").Return(nil)

	s := NewOrderStorage(engineMock)

	// act
	err := s.Add(&expected)

	// assert
	assert.NoError(t, err, "non-error result expected")
}

func TestOrderStorageRemove_ExistingIDGiven_OrderRemoved(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	engineMock := mocks.NewMockEngine(ctrl)
	engineMock.EXPECT().Remove("test_order").Return(nil)

	s := NewOrderStorage(engineMock)

	// act
	err := s.Remove("test_order")

	// assert
	assert.NoError(t, err, "non-error result expected")
}

func TestOrderStorageBySpecWithLimit_ValidSpecGiven_OrdersReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []ds.Order{
		{
			ID:    "test_obj_101",
			Price: 101,
		},
		{
			ID:    "test_obj_105",
			Price: 105,
		},
	}

	var ordersIntf []interface{}
	for idx := range expected {
		ordersIntf = append(ordersIntf, &expected[idx])
	}

	//var orders report.GetOrdersReport
	spec := priceIsBetweenTestSpec{From: 101, To: 106}
	q := inmemory.ConcreteCriteria{
		Limit: 10,
		Spec:  spec,
	}

	engineMock := mocks.NewMockEngine(ctrl)
	engineMock.EXPECT().Match(q).Return(ordersIntf, nil)

	s := NewOrderStorage(engineMock)

	// act
	obtained, err := s.BySpecWithLimit(spec, 10)

	// assert
	assert.NoError(t, err, "non-error result expected")
	assert.Equal(t, expected, obtained)

}

type priceIsBetweenTestSpec struct {
	From int64
	To   int64
}

func (s priceIsBetweenTestSpec) IsSatisfiedBy(object interface{}) bool {
	order := object.(*ds.Order)
	return order.Price >= s.From && order.Price < s.To
}
