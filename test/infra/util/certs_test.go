package util

import (
	"context"
	"crypto/x509"
	"testing"
	"time"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/sonm-io/marketplace/infra/util"
	"github.com/stretchr/testify/require"
)

func TestHitlessRotator(t *testing.T) {
	req := require.New(t)
	key, err := ethcrypto.GenerateKey()
	if err != nil {
		t.Fatalf("%v", err)
	}
	ctx := context.Background()
	r, cfg, err := util.NewHitlessCertRotator(ctx, key)
	req.NoError(err)
	defer r.Close()

	deadline := time.Now().Add(time.Second * 6)
	for time.Now().Before(deadline) {
		tCfg, _ := cfg.GetCertificate(nil)
		x509Cert, err := x509.ParseCertificate(tCfg.Certificate[0])
		req.NoError(err)
		_, err = util.CheckCert(x509Cert)
		req.NoError(err)

		tCfgCl, _ := cfg.GetClientCertificate(nil)
		x509CertCl, err := x509.ParseCertificate(tCfgCl.Certificate[0])
		req.NoError(err)
		_, err = util.CheckCert(x509CertCl)
		req.NoError(err)

		time.Sleep(time.Second)
	}
}
