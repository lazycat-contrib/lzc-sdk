package gohelper

import (
	"context"
	"errors"

	grpcpeer "google.golang.org/grpc/peer"

	"google.golang.org/grpc/credentials"
)

type Application struct {
	AppId     string
	BoxId     string
	AppDomain string
}

// 通过client TLS查询对端实际的APPID
func PeerApplication(ctx context.Context) (*Application, error) {
	p, ok := grpcpeer.FromContext(ctx)
	if !ok {
		return nil, errors.New("NOT AUTH")
	}
	crts := p.AuthInfo.(credentials.TLSInfo).State.PeerCertificates
	if len(crts) != 1 {
		return nil, errors.New("invalid client")
	}
	return &Application{
		AppId:     crts[0].Subject.SerialNumber,
		BoxId:     crts[0].Issuer.SerialNumber,
		AppDomain: crts[0].Subject.CommonName,
	}, nil

}
