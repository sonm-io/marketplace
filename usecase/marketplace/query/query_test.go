package query

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetOrder_QueryID(t *testing.T) {
	assert.Equal(t, "GetOrder", GetOrder{}.QueryID())
}

func TestGetOrders_QueryID(t *testing.T) {
	assert.Equal(t, "GetOrders", GetOrders{}.QueryID())
}
