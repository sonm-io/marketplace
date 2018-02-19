package srv

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
	"github.com/sonm-io/marketplace/interface/grpc/srv/mocks"
)

func TestMarketplace_GetOrderByID_ValidIDGiven_ValidResponse(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := &pb.ID{Id: "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7"}

	serviceMock := mocks.NewMockMarketService(ctrl)
	serviceMock.EXPECT().OrderByID(req.Id, &pb.Order{}).Times(1)

	m := NewMarketplace(serviceMock)

	// act
	_, err := m.GetOrderByID(context.Background(), req)

	// assert
	assert.NoError(t, err)
}

func TestMarketplaceGetOrderByID_AnErrorOccurredWhileGettingOrder_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orderID := "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7"

	serviceMock := mocks.NewMockMarketService(ctrl)
	serviceMock.EXPECT().
		OrderByID(orderID, &pb.Order{}).
		Return(fmt.Errorf("some error"))

	m := NewMarketplace(serviceMock)

	// act
	_, err := m.GetOrderByID(context.Background(), &pb.ID{Id: orderID})

	// assert
	assert.EqualError(t, err, "rpc error: code = Internal desc = cannot get order: some error")
}
