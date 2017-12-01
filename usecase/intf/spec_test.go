package intf

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAndSpecification_IsSatisfiedBy(t *testing.T) {

	object := testObject{
		rating:   777,
		category: "any",
	}

	s := NewRatingGreaterThanSpec(5).And(NewCategoryEqualsSpec("any"))
	obtained := s.IsSatisfiedBy(object)

	assert.True(t, obtained)

}

type testObject struct {
	ID       int
	category string
	rating   int
}

type ratingGreaterThanSpec struct {
	CompositeSpecification
	rating int
}

func NewRatingGreaterThanSpec(rating int) CompositeSpecification {
	return BaseSpecification{ratingGreaterThanSpec{rating: rating}}
}

func (s ratingGreaterThanSpec) IsSatisfiedBy(object interface{}) bool {
	o, ok := object.(testObject)
	if !ok {
		return false
	}

	return o.rating > s.rating
}

type categoryEqualsSpec struct {
	CompositeSpecification
	category string
}

func NewCategoryEqualsSpec(category string) CompositeSpecification {
	return BaseSpecification{categoryEqualsSpec{category: category}}
}

func (s categoryEqualsSpec) IsSatisfiedBy(object interface{}) bool {
	o, ok := object.(testObject)
	if !ok {
		return false
	}

	return o.category == s.category
}
