package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPropertiesLessOrEqual_LesserOrEqualPropertiesGiven_TrueReturned(t *testing.T) {
	var p1, p2 Properties

	p1 = map[string]float64{
		"foo": 777,
		"bar": 555,
	}

	p2 = map[string]float64{
		"foo": 777.1,
		"bar": 555.0001,
	}

	obtained := p1.LessOrEqual(p2)
	assert.True(t, obtained)
}

func TestPropertiesLessOrEqual_InExistentPropertyGiven_FalseReturned(t *testing.T) {
	var p1, p2 Properties

	p1 = map[string]float64{
		"foo": 777.1,
	}

	p2 = map[string]float64{
		"foo": 777.1,
		"bar": 555,
	}

	obtained := p1.LessOrEqual(p2)
	assert.False(t, obtained)
}

func TestPropertiesLessOrEqual_ALesserPropertyGiven_FalseReturned(t *testing.T) {
	var p1, p2 Properties

	p1 = map[string]float64{
		"foo": 777.1,
		"bar": 555,
	}

	p2 = map[string]float64{
		"foo": 777.1,
		"bar": 444,
	}

	obtained := p1.LessOrEqual(p2)
	assert.False(t, obtained)
}

func TestPropertiesGreaterOrEqual_GreaterOrEqualPropertiesGiven_TrueReturned(t *testing.T) {
	var p1, p2 Properties

	p1 = map[string]float64{
		"foo": 777.1,
		"bar": 555.0001,
	}

	p2 = map[string]float64{
		"foo": 777,
		"bar": 555,
	}

	obtained := p1.GreaterOrEqual(p2)
	assert.True(t, obtained)
}

func TestPropertiesGreaterOrEqual_InExistentPropertyGiven_FalseReturned(t *testing.T) {
	var p1, p2 Properties

	p1 = map[string]float64{
		"foo": 777.1,
	}

	p2 = map[string]float64{
		"foo": 777.1,
		"bar": 555,
	}

	obtained := p1.GreaterOrEqual(p2)
	assert.False(t, obtained)
}

func TestProperties_GreaterOrEqualOrEqual_ABiggerPropertyGiven_FalseReturned(t *testing.T) {
	var p1, p2 Properties

	p1 = map[string]float64{
		"foo": 777.1,
		"bar": 444,
	}

	p2 = map[string]float64{
		"foo": 777.1,
		"bar": 555,
	}

	obtained := p1.GreaterOrEqual(p2)
	assert.False(t, obtained)
}
