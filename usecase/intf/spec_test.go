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
	s := ratingGreaterThanSpec{CompositeSpecification: &BaseSpecification{}, rating: rating}
	s.Relate(s)
	return s
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
	s := categoryEqualsSpec{CompositeSpecification: &BaseSpecification{}, category: category}
	s.Relate(s)
	return s
}

func (s categoryEqualsSpec) IsSatisfiedBy(object interface{}) bool {
	o, ok := object.(testObject)
	if !ok {
		return false
	}

	return o.category == s.category
}
