package inmemory

import (
	"testing"

	"github.com/sonm-io/marketplace/entity"
	"github.com/stretchr/testify/assert"
)

func TestOrderStorageStore_ValidObjectGiven_ObjectStored(t *testing.T) {
	// arrange
	s := NewStorage()

	// act
	err := s.Store(&entity.Order{
		ID: "test_obj",
	})

	// assert
	assert.NoError(t, err, "non-error result expected")
	assert.Equal(t, 1, len(s.db))
}

func TestOrderStorageByID_ExistentIDGiven_ValidResultReturned(t *testing.T) {
	// arrange

	expected := &entity.Order{
		ID: "test_obj",
	}

	s := NewStorage()
	s.Store(expected)



	// act
	obtained := &entity.Order{}
	err := s.ByID("test_obj", obtained)

	// assert
	assert.NoError(t, err, "non-error result expected")
	assert.Equal(t, expected, obtained)
}

func TestOrderStorageByID_InExistentIDGiven_ErrorReturned(t *testing.T) {
	// arrange
	s := NewStorage()

	// act
	obtained := &entity.Order{}
	err := s.ByID("n/a", obtained)

	// assert
	assert.EqualError(t, err, errOrderNotFound.Error())
}

func TestOrderStorageRemove_ValidIDGiven_ObjectRemoved(t *testing.T) {
	// arrange
	s := NewStorage()
	s.Store(&entity.Order{
		ID: "test_obj",
	})

	// act
	err := s.Remove("test_obj")

	// assert
	assert.NoError(t, err, "non-error result expected")
	assert.Equal(t, 0, len(s.db))
}

func TestOrderStorageRemove_InExistentIDGiven_ErrorReturned(t *testing.T) {
	// arrange
	s := NewStorage()

	// act
	err := s.Remove("n/a")

	// assert
	assert.EqualError(t, err, errOrderNotFound.Error())
}

func TestOrderStorageMatch_SpecCriteriaGiven_CollectionReturned(t *testing.T) {
	// arrange
	s := NewStorage()
	s.Store(&entity.Order{
		ID: "test_obj_100",
		Price: 100,
	})
	s.Store(&entity.Order{
		ID: "test_obj_101",
		Price: 101,
	})
	s.Store(&entity.Order{
		ID: "test_obj_105",
		Price: 105,
	})
	s.Store(&entity.Order{
		ID: "test_obj_110",
		Price: 110,
	})

	expected := []*entity.Order{
		{
			ID: "test_obj_101",
			Price: 101,
		},
		{
			ID: "test_obj_105",
			Price: 105,
		},
	}


	// act
	q := ConcreteCriteria{
		Limit:100,
		Spec:PriceIsBetweenTestSpec{From:101, To:106},
	}

	var obtained []*entity.Order
	err := s.Match(q, &obtained)

	// assert
	assert.NoError(t, err, "non-error result expected")
	assert.EqualValues(t, expected, obtained)
}

type PriceIsBetweenTestSpec struct {
	From int64
	To int64
}

func (s PriceIsBetweenTestSpec) IsSatisfiedBy(object interface{}) bool {
	order := object.(*entity.Order)
	return order.Price >= s.From && order.Price < s.To
}