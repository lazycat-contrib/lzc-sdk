package gohelper

import (
	"context"

	papi "gitee.com/linakesi/lzc-sdk/lang/go/sys/portal-server"
	"gitee.com/linakesi/remotesocks/client"
	"gitee.com/linakesi/remotesocks/spec"
)

func GetNetstack(target string) spec.Netstack {
	//暂时直接调用hserver的grpc,后续rs稳定完善后，再改到lzc-runtime的接口上
	addrFn := func() (string, error) {
		c, err := papi.NewClient()
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
