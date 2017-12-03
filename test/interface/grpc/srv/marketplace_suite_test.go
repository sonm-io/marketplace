package srv_test

import (
	"fmt"
	"testing"

	"google.golang.org/grpc"
	// registers grpc gzip encoder/decoder
	_ "google.golang.org/grpc/encoding/gzip"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
	"github.com/sonm-io/marketplace/test/cli"
)

// TestMarketplaceTestSuite initialize test suit.
func TestMarketplaceTestSuite(t *testing.T) {
	suite.Run(t, new(MarketplaceTestSuite))
}

// MarketplaceTestSuite tests Marketplace service.
type MarketplaceTestSuite struct {
	cli.AppTestSuite

	conn   *grpc.ClientConn
	client pb.MarketClient
}

// SetupTest prepare state for the test.
func (s *MarketplaceTestSuite) SetupTest() {
	s.AppTestSuite.SetupTest()

	conn, err := GrpcClient()
	require.NoError(s.T(), err, "cannot get grpc client")

	//defer conn.Close()
	s.conn = conn
	s.client = pb.NewMarketClient(conn)
}

func (s *MarketplaceTestSuite) TearDownTest() {
	s.conn.Close()
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (s *MarketplaceTestSuite) TestMarketPlace() {
	s.StartApp()
	defer s.StopApp()

	s.T().Run("CreateOrder", func(t *testing.T) {
		s.createOrder()
	})

	s.T().Run("GetOrderByID", func(t *testing.T) {
		s.getOrderByID()
	})
}

func GrpcClient() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(cli.ListenAddr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("did not connect: %v", err.Error())
	}
	return conn, err
}
