package srv

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	pb "github.com/sonm-io/marketplace/interface/grpc/proto"

	"github.com/sonm-io/marketplace/usecase/intf/mocks"
	"github.com/sonm-io/marketplace/usecase/marketplace/command"
)

func TestMarketplaceTouchOrders_ValidParamsGiven_ValidResponse(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orderIDs := []string{"cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7"}

	touchOrdersMock := mocks.NewMockCommandHandler(ctrl)
	touchOrdersMock.EXPECT().
		Handle(command.TouchOrders{IDs: orderIDs}).
		Return(nil)

	m := NewMarketplace(touchOrdersMock, nil, nil)

	// act
	_, err := m.TouchOrders(context.Background(), &pb.TouchOrdersRequest{IDs: orderIDs})

	// assert
	assert.NoError(t, err)
}

func TestMarketplaceTouchOrders_AnErrorOccurredWhileTouchingOrders_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orderIDs := []string{"cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7"}

	touchOrdersMock := mocks.NewMockCommandHandler(ctrl)
	touchOrdersMock.EXPECT().
		Handle(command.TouchOrders{IDs: orderIDs}).
		Return(fmt.Errorf("cannot touch orders"))

	m := NewMarketplace(touchOrdersMock, nil, nil)

	// act
	_, err := m.TouchOrders(context.Background(), &pb.TouchOrdersRequest{IDs: orderIDs})

	// assert
	assert.EqualError(t, err, "cannot touch orders")
}
