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
func (u *netUtil) GetLocalIPs() []string {
	ips := u.GetLocalNetIPs()
	ret := make([]string, len(ips))
	for i := range ips {
		ret = append(ret, ips[i].String())
	}
	return ret
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

func (u *netUtil) GetLocalNetIPs() []net.IP {
	var ips []net.IP
	// 获取所有网卡接口信息
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil
	}

	for _, iface := range ifaces {
		// 排除本地回环接口和隧道接口
		if iface.Flags&net.FlagLoopback != 0 || iface.Flags&net.FlagPointToPoint != 0 {
			continue
		}

		// 获取当前网卡接口的所有 IP 地址信息
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		// 遍历当前网卡接口的所有 IP 地址
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}

			// 排除本地 IP 地址和无效 IP 地址
			if ipNet.IP.IsLoopback() || ipNet.IP.IsLinkLocalUnicast() || ipNet.IP.IsLinkLocalMulticast() ||
				ipNet.IP.IsInterfaceLocalMulticast() || ipNet.IP.IsMulticast() || ipNet.IP.IsUnspecified() {
				continue
			}

			// 输出符合条件的 IP 地址
			ips = append(ips, ipNet.IP)
		}
	}
	return ips
}
