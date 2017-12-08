package intf

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBaseSpecification_IsSatisfiedBy_EmptySpecGiven_FalseReturned(t *testing.T) {
	object := testObject{
		rating:   777,
		category: "any",
	}

	s := BaseSpecification{}
	obtained := s.IsSatisfiedBy(object)

	assert.False(t, obtained)
}

func TestBaseSpecification_IsSatisfiedBy_SpecGiven_CorrespondingResultReturned(t *testing.T) {
	object := testObject{
		rating:   777,
		category: "any",
	}

	s := BaseSpecification{newCategoryEqualsSpec("any")}
	obtained := s.IsSatisfiedBy(object)

	assert.True(t, obtained)
}

func TestAndSpecification_IsSatisfiedBy(t *testing.T) {
	object := testObject{
		rating:   777,
		category: "any",
	}

	s := newRatingGreaterThanSpec(5).And(newCategoryEqualsSpec("any"))
	obtained := s.IsSatisfiedBy(object)

	assert.True(t, obtained)
}

func TestOrSpecification_IsSatisfiedBy(t *testing.T) {
	object := testObject{
		rating:   0,
		category: "any",
	}

	s := newRatingGreaterThanSpec(5).Or(newCategoryEqualsSpec("any"))
	obtained := s.IsSatisfiedBy(object)

	assert.True(t, obtained)
}

func TestNotSpecification_IsSatisfiedBy(t *testing.T) {
	object := testObject{
		category: "some_category",
	}

	s := newCategoryEqualsSpec("any").Not()
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

func newRatingGreaterThanSpec(rating int) CompositeSpecification {
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

func newCategoryEqualsSpec(category string) CompositeSpecification {
	return BaseSpecification{categoryEqualsSpec{category: category}}
}

func (s categoryEqualsSpec) IsSatisfiedBy(object interface{}) bool {
	o, ok := object.(testObject)
	if !ok {
		return false
	}
	return o.category == s.category
}
