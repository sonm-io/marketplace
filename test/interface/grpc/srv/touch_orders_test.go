package srv_test

import (
	"context"
	//"fmt"
	"testing"

	//"github.com/sonm-io/marketplace/infra/util"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"

	//"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (s *MarketplaceTestSuite) TouchOrders(t *testing.T) {
	// smth like "0x9B27D3C3571731deDb23EaFEa34a3a6E05daE159"
	ID := "1b5dfa00-af3c-4e2d-b64b-c5d62e89430b"
	//buyerID := util.PubKeyToAddr(s.App.PublicKey()).Hex()

	_, err := s.client.TouchOrders(context.Background(), &pb.TouchOrdersRequest{IDs: []string{ID}})
	require.NoError(t, err)

	// act
	_, err = s.client.GetOrderByID(context.Background(), &pb.ID{Id: ID})

	// assert
	require.NoError(t, err)
	//assert.Equal(obtained.)
	//assert.EqualError(t, err,
	//fmt.Sprintf(`rpc error: code = Internal desc = cannot get order: order %s is inactive`, ID))
}
