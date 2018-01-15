package srv_test

import (
	"context"
	"testing"

	pb "github.com/sonm-io/marketplace/interface/grpc/proto"

	"github.com/stretchr/testify/require"
)

func (s *MarketplaceTestSuite) TouchOrders(t *testing.T) {
	ID := "1b5dfa00-af3c-4e2d-b64b-c5d62e89430b"

	_, err := s.client.TouchOrders(context.Background(), &pb.TouchOrdersRequest{IDs: []string{ID}})
	require.NoError(t, err)

	// act
	_, err = s.client.GetOrderByID(context.Background(), &pb.ID{Id: ID})

	// assert
	require.NoError(t, err)
}
