package gohelper

import (
	"context"
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"time"

	"gitee.com/linakesi/lzc-sdk/lang/go/localdevice"
	"google.golang.org/grpc"
)

type metadataCredentials struct {
	authToken string
}

func (c *metadataCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	if len(uri) != 1 {
		fmt.Println("BUG: unexpected uri", uri)
	}
	return map[string]string{
		"lzc_dapi_auth_token": c.authToken,
	}, nil
}

func (c *metadataCredentials) RequireTransportSecurity() bool {
	return true
}

func newMetadataCredentials(authToken string) (*metadataCredentials, error) {
	return &metadataCredentials{authToken: authToken}, nil
}

type AuthToken struct {
	Token    string
	Deadline time.Time
}

func RequestAuthToken(conn *grpc.ClientConn) (*AuthToken, error) {
	perm := localdevice.NewPermissionManagerClient(conn)
	atr, err := genRequestAuthTokenRequest()
	if err != nil {
		return nil, err
	}
	resp, err := perm.RequestAuthToken(context.Background(), atr)
	if err != nil {
		return nil, err
	}
	return &AuthToken{
		Token: resp.Token,
		Deadline: resp.Deadline.AsTime(),
	}, nil
}

func genRequestAuthTokenRequest() (*localdevice.RequestAuthTokenRequest, error) {
	boxCertBytes, err := os.ReadFile(CAPath)
	if err != nil {
		return nil, err
	}
	appCertBytes, err := os.ReadFile(APPCertPath)
	if err != nil {
		return nil, err
	}
	appKeyBytes, err := os.ReadFile(APPKeyPath)
	if err != nil {
		return nil, err
	}

	appCert, err := certFromPEM(appCertBytes)
	if err != nil {
		return nil, err
	}
	appKey, err := privateKeyFromPEM(appKeyBytes)
	if err != nil {
		return nil, err
	}
	signer, ok := appKey.(crypto.Signer)
	if !ok {
		return nil, fmt.Errorf("type %T doesn't implement crypto.Signer", appKey)
	}
	signature, err := signer.Sign(
		rand.Reader, []byte(appCert.Subject.SerialNumber), crypto.Hash(0),
	)
	if err != nil {
		return nil, err
	}
	return &localdevice.RequestAuthTokenRequest{
		BoxCert:   boxCertBytes,
		AppCert:   appCertBytes,
		Signature: signature,
	}, nil
}

func privateKeyFromPEM(b []byte) (crypto.PrivateKey, error) {
	block, _ := pem.Decode(b)
	if block == nil {
		return nil, errors.New("cannot parse private key")
	}

	switch block.Type {
	case "PRIVATE KEY":
		x, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		if k, ok := x.(ed25519.PrivateKey); ok {
			return &k, nil
		}
		if k, ok := x.(rsa.PrivateKey); ok {
			return &k, nil
		}
		return nil, fmt.Errorf("unsupported private key type %T", x)
	case "RSA PRIVATE KEY":
		return x509.ParsePKCS1PrivateKey(block.Bytes)
	default:
		return nil, fmt.Errorf("Not support Private Key type: %s", block.Type)
	}
}

func certFromPEM(b []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode([]byte(b))
	if block == nil {
		return nil, errors.New("cannot load pem cert")
	}
	return x509.ParseCertificate(block.Bytes)
}
