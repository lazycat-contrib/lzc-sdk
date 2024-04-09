package gohelper

import (
	"context"

	papi "gitee.com/linakesi/lzc-baseos-protos/lang/go/baseos"
	"gitee.com/linakesi/remotesocks/client"
	"gitee.com/linakesi/remotesocks/spec"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type hserverClientWrap struct {
	papi.HPortalSysClient
	conn *grpc.ClientConn
}

func (c hserverClientWrap) Close() error { return c.conn.Close() }

func NewHServerClient() (*hserverClientWrap, error) {
	const LzcAppSocketPath = "/lzcapp/run/sys/portal-server.socket"
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial("unix://"+LzcAppSocketPath, opts...)
	if err != nil {
		return nil, err
	}
	return &hserverClientWrap{HPortalSysClient: papi.NewHPortalSysClient(conn), conn: conn}, nil
}

func GetNetstack(target string) spec.Netstack {
	//暂时直接调用hserver的grpc,后续rs稳定完善后，再改到lzc-runtime的接口上
	addrFn := func() (string, error) {
		c, err := NewHServerClient()
		if err != nil {
			return "", err
		}
		defer c.Close()

		var ltype = papi.RemoteSocksRequest_Local
		if target != "" {
			ltype = papi.RemoteSocksRequest_Remote
		}
		cinfo, err := c.RemoteSocks(context.Background(), &papi.RemoteSocksRequest{
			LocationType: ltype,
			Target:       target,
		})
		if err != nil {
			return "", err
		}
		return cinfo.ServerUrl, nil
	}
	return client.NewClient(addrFn)
}
