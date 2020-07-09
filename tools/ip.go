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
