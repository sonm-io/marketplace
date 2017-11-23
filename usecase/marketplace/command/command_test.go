package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type unknownCommand struct{}

func (c unknownCommand) CommandID() string {
	return "UnknownCommand"
}

func TestCancelOrder_CommandID(t *testing.T) {
	assert.Equal(t, "CancelOrder", CancelOrder{}.CommandID())
}

func TestCreateOrder_CommandID(t *testing.T) {
	assert.Equal(t, "CreateOrder", CreateOrder{}.CommandID())
}
