package gohelper

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	CAPath      = "/lzcapp/run/certs/box.crt"
	APPCertPath = "/lzcapp/run/certs/app.crt"
	APPKeyPath  = "/lzcapp/run/certs/app.key"
)

func apiSocketPath() string {
	x := os.Getenv("LZCAPP_API_GATEWAY_ADDRESS")
	if x == "" {
		return "unix:/lzcapp/run/sys/lzc-apis.socket"
	}
	return x
}

func buildClientCredOption(caCrt string, appKey string, appCrt string) (grpc.DialOption, error) {
	cert, err := tls.LoadX509KeyPair(appCrt, appKey)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caCrt)
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

func DialCred() (grpc.DialOption, error) {
	return buildClientCredOption(CAPath, APPKeyPath, APPCertPath)
}
