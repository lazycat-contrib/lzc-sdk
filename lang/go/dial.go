package gohelper

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	CAPath            = "/lzcapp/run/certs/box.crt"
	APPCertPath       = "/lzcapp/run/certs/app.crt"
	APPKeyPath        = "/lzcapp/run/certs/app.key"
	RuntimeSocketPath = "unix:///lzcapp/run/sys/lzc-apis.socket"
)

func BuildClientCredOption(caCrt string, appKey string, appCrt string) (grpc.DialOption, error) {
	cert, err := tls.LoadX509KeyPair(appCrt, appKey)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	ca, err := os.ReadFile(caCrt)
	if err != nil {
		return nil, err
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return nil, err
	}
	cred := credentials.NewTLS(&tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
		RootCAs:            certPool,
	})
	return grpc.WithTransportCredentials(cred), nil
}

func dialConn(ctx context.Context, cred grpc.DialOption) (*grpc.ClientConn, error) {
	x := os.Getenv("LZCAPP_API_GATEWAY_ADDRESS")
	if x == "" {
		return grpc.DialContext(ctx, RuntimeSocketPath, grpc.WithBlock(), cred)
	} else {
		return grpc.DialContext(ctx, x, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
}
