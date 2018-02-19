package srv

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	"github.com/sonm-io/marketplace/infra/grpc/interceptor"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
)

func TestCheckPermissions_EmptyContextGiven_ErrorReturned(t *testing.T) {
	// act
	err := CheckPermissions(context.Background(), &pb.Order{})

	// assert
	assert.EqualError(t, err,
		"rpc error: code = Unauthenticated desc = auth failed: cannot get eth address from context")
}

func TestCheckPermissions_CredentialsDoNotMatch_ErrorReturned(t *testing.T) {
	// arrange
	order := &pb.Order{
		OrderType: pb.OrderType_BID,
		ByuerID:   "0xAAA568CD389580B6737FF56b61BE4F4eE802E2Db",
	}
	ctx := interceptor.EthAddrToContext(context.Background(),
		common.HexToAddress("0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db"))

	// act
	err := CheckPermissions(ctx, order)

	// assert
	assert.EqualError(t, err,
		"rpc error: code = PermissionDenied desc = auth failed: SupplierID/BuyerID and ethereum address differ")
}
