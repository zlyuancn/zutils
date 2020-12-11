/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/12/11
   Description :
-------------------------------------------------
*/

package zutils

import (
	"net"
)

var Net = new(netUtil)

type netUtil struct{}

// 获取所有能检测到的本地ip, 不会返回 127.*
func (*netUtil) GetLocalIPs() []string {
	var ips []string
	address, err := net.InterfaceAddrs()
	if err != nil {
		return ips
	}

	for _, addr := range address {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips
}

// 获取所有能检测到的本地ip, 返回第一个匹配 prefix 的ip, 否则返回空字符串
func (u *netUtil) GetLocalIPMatchPrefix(prefix string) string {
	ips := u.GetLocalIPs()
	le := len(prefix)
	if le == 0 {
		return ips[0]
	}

	for _, ip := range ips {
		if len(ip) >= le && ip[:le] == prefix {
			return ip
		}
	}
	return ""
}
