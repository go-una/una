package tools

import (
	"net"
	"strings"
)

func IP2Long(ipStr string) uint32 {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return 0
	}

	b1 := uint32(ip[12])
	b2 := uint32(ip[13])
	b3 := uint32(ip[14])
	b4 := uint32(ip[15])
	return b1<<24 | b2<<16 | b3<<8 | b4
}

func IPInRange(ip net.IP, ipRangeList []string) bool {
	inRange := false
	for _, ipRange := range ipRangeList {
		if strings.Contains(ipRange, "/") {
			// IP段
			_, ipnet, err := net.ParseCIDR(ipRange)
			if err != nil {
				continue
			}

			if ipnet.Contains(ip) {
				inRange = true
				break
			}
		} else {
			// 单个IP
			if ipRange == ip.String() {
				inRange = true
				break
			}
		}
	}
	return inRange
}

// 获取本机内网 IP
func GetInternalIPv4() (net.IP, error) {
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	internalIpRangeList := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
	}

	ip := net.ParseIP("127.0.0.1")
	for _, a := range addrList {
		ipNet, ok := a.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil && IPInRange(ipNet.IP, internalIpRangeList) {
			ip = ipNet.IP
		}
	}
	return ip, nil
}
