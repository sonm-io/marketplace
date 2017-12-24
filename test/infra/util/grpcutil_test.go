package util

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/stretchr/testify/require"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/sonm-io/marketplace/infra/util"
)

func TestTLSGenCerts(t *testing.T) {
	priv, err := ethcrypto.GenerateKey()
	if err != nil {
		t.Fatalf("%v", err)
	}
	certPEM, keyPEM, err := util.GenerateCert(priv, time.Second*20)
	if err != nil {
		t.Fatalf("%v", err)
	}
	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		t.Fatalf("%v", err)
	}
	x509Cert, err := x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		t.Fatalf("%v", err)
	}
	_, err = util.CheckCert(x509Cert)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSecureGRPCConnect(t *testing.T) {
	req := require.New(t)
	ctx := context.Background()
	serPriv, err := ethcrypto.GenerateKey()
	req.NoError(err)
	rot, serTLS, err := util.NewHitlessCertRotator(ctx, serPriv)
	req.NoError(err)
	defer rot.Close()
	serCreds := util.NewTLS(serTLS)
	server := util.MakeGrpcServer(serCreds)
	lis, err := net.Listen("tcp", "localhost:0")
	req.NoError(err)
	defer lis.Close()
	go func() {
		server.Serve(lis)
	}()

	t.Run("ClientWithTLS", func(t *testing.T) {
		clientPriv, err := ethcrypto.GenerateKey()
		req.NoError(err)
		rot, clientTLS, err := util.NewHitlessCertRotator(ctx, clientPriv)
		req.NoError(err)
		defer rot.Close()
		var clientCreds = util.NewTLS(clientTLS)
		conn, err := util.MakeGrpcClient(ctx, lis.Addr().String(), clientCreds, grpc.WithTimeout(time.Second), grpc.WithBlock())
		req.NoError(err)
		defer conn.Close()

		err = grpc.Invoke(ctx, "/DummyService/dummyMethod", nil, nil, conn)
		req.NotNil(err)

		st, ok := status.FromError(err)
		req.True(ok)
		req.True(st.Code() == codes.Unimplemented)
	})

	t.Run("ClientWithoutTLS", func(t *testing.T) {
		conn, err := util.MakeGrpcClient(ctx, lis.Addr().String(), nil, grpc.WithBlock(), grpc.WithTimeout(time.Second))
		if err != nil {
			// On Linux we can have an error here due to failed TLS Handshake
			// It's expected behavior
			return
		}
		// If we passed here, error must occur after the first call
		req.NotNil(conn)
		defer conn.Close()
		err = grpc.Invoke(ctx, "/DummyService/dummyMethod", nil, nil, conn)
		req.NotNil(err)
		st, ok := status.FromError(err)
		req.True(ok)
		req.True(st.Code() == codes.Internal || st.Code() == codes.Unavailable)
	})
}
