package srv

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	pb "github.com/sonm-io/marketplace/handler/proto"
	"github.com/sonm-io/marketplace/handler/srv/mocks"
)

func TestMarketplaceTouchOrders_AnErrorOccurredWhileTouchingOrders_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orderIDs := []string{"cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7"}

	serviceMock := mocks.NewMockMarketService(ctrl)
	serviceMock.EXPECT().
		TouchOrders(orderIDs).
		Return(fmt.Errorf("some error"))

	m := NewMarketplace(serviceMock)

	// act
	_, err := m.TouchOrders(context.Background(), &pb.TouchOrdersRequest{IDs: orderIDs})

	// assert
	assert.EqualError(t, err, "rpc error: code = Internal desc = cannot touch orders: some error")
}

func TestMarketplaceTouchOrders_ValidIDsGiven_ValidResponse(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orderIDs := []string{"cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7"}

	serviceMock := mocks.NewMockMarketService(ctrl)
	serviceMock.EXPECT().TouchOrders(orderIDs)

	m := NewMarketplace(serviceMock)

	// act
	_, err := m.TouchOrders(context.Background(), &pb.TouchOrdersRequest{IDs: orderIDs})

	// assert
	assert.NoError(t, err)
}
