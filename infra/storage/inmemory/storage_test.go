package inmemory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testElement struct {
	ID    string
	Price int64
}

func TestStorageAdd_ValidElementGiven_ElementAddedToTheStorage(t *testing.T) {
	// arrange
	s := NewStorage()

	// act
	err := s.Add(&testElement{ID: "test_obj"}, "test_obj")

	// assert
	assert.NoError(t, err, "non-error result expected")
	assert.Equal(t, 1, len(s.elements))
}

func TestStorageGet_ExistentIDGiven_ValidResultReturned(t *testing.T) {
	// arrange

	expected := &testElement{
		ID: "test_obj",
	}

	s := NewStorage()
	s.Add(expected, "test_obj")

	// act
	obtained, err := s.Get("test_obj")

	// assert
	assert.NoError(t, err, "non-error result expected")
	assert.Equal(t, expected, obtained)
}

func TestStorageGet_InExistentIDGiven_ErrorReturned(t *testing.T) {
	// arrange
	s := NewStorage()

	// act
	_, err := s.Get("n/a")

	// assert
	assert.EqualError(t, err, errNotFound.Error())
}

func TestStorageRemove_ValidIDGiven_ElementRemoved(t *testing.T) {
	// arrange
	s := NewStorage()
	s.Add(&testElement{
		ID: "test_obj",
	}, "test_obj")

	// act
	err := s.Remove("test_obj")

	// assert
	assert.NoError(t, err, "non-error result expected")
	assert.Equal(t, 0, len(s.elements))
}

func TestStorageRemove_InExistentIDGiven_ErrorReturned(t *testing.T) {
	// arrange
	s := NewStorage()

	// act
	err := s.Remove("n/a")

	// assert
	assert.EqualError(t, err, errNotFound.Error())
}

func TestStorageMatch_CriteriaGiven_CollectionReturned(t *testing.T) {
	// arrange
	s := NewStorage()
	s.Add(&testElement{
		ID:    "test_obj_100",
		Price: 100,
	}, "test_obj_100")
	s.Add(&testElement{
		ID:    "test_obj_101",
		Price: 101,
	}, "test_obj_101")
	s.Add(&testElement{
		ID:    "test_obj_105",
		Price: 105,
	}, "test_obj_105")
	s.Add(&testElement{
		ID:    "test_obj_110",
		Price: 110,
	}, "test_obj_110")

	elements := []*testElement{
		{
			ID:    "test_obj_101",
			Price: 101,
		},
		{
			ID:    "test_obj_105",
			Price: 105,
		},
	}

	var expected []interface{}
	for _, el := range elements {
		expected = append(expected, el)
	}

	// act
	q := ConcreteCriteria{
		Limit: 100,
		Spec:  PriceIsBetweenTestSpec{From: 101, To: 106},
	}

	obtained, err := s.Match(q)

	// assert
	assert.NoError(t, err, "non-error result expected")
	assert.EqualValues(t, expected, obtained)
}

type PriceIsBetweenTestSpec struct {
	From int64
	To   int64
}

func (s PriceIsBetweenTestSpec) IsSatisfiedBy(object interface{}) bool {
	order := object.(*testElement)
	return order.Price >= s.From && order.Price < s.To
}
