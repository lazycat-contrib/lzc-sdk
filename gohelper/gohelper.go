package gohelper

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

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

var (
	CAPath        = "/lzcapp/certs/box.crt"
	APPCertPath   = "/lzcapp/certs/app.crt"
	APPKeyPath    = "/lzcapp/certs/app.key"
	APISocketPath = "/lzcapp/lzc-apis.socket"
)

func DialLZCAPI() (*grpc.ClientConn, error) {
	cred, err := buildClientCredOption(CAPath, APPKeyPath, APPCertPath)
	if err != nil {
		return nil, err
	}
	return grpc.Dial("unix:"+APISocketPath, cred)
}
