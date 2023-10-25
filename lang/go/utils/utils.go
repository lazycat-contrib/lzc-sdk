package utils

import (
	"net"

	"github.com/vishvananda/netlink"
)

// 返回当前容器的IP地址或内部域名
//
// 需要保证运行环境的dns server为lzc-dns
func QueryServiceAddress() (string, error) {
	ips, err := net.LookupIP("host.lzcapp")
	if err != nil {
		return "", err
	}
	var lastErr error
	for _, hostIP := range ips {
		routes, err := netlink.RouteGet(hostIP)
		if err != nil {
			lastErr = err
			continue
		}
		return routes[0].Src.String(), nil
	}
	return "", lastErr
}
