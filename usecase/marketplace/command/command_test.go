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

func TestCreateBidOrder_CommandID(t *testing.T) {
	assert.Equal(t, "CreateBidOrder", CreateBidOrder{}.CommandID())
}

func TestCreateAskOrder_CommandID(t *testing.T) {
	assert.Equal(t, "CreateAskOrder", CreateAskOrder{}.CommandID())
}
