package srv

import (
	pb "github.com/sonm-io/marketplace/handler/proto"
	"golang.org/x/net/context"
)

// GetProcessing method exists just to match the Marketplace interface.
// The Market service itself is unable to know anything about processing orders.
// This method is implemented for Node in `insonmnia/node/market.go:348`
func (m *Marketplace) GetProcessing(ctx context.Context, req *pb.Empty) (*pb.GetProcessingReply, error) {
	return nil, nil
}
