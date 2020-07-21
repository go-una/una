package tools

import (
	"log"
	"net"
	"testing"
)

func TestIP2Long(t *testing.T) {
	l1 := IP2Long("127.0.0.1")
	if l1 != 2130706433 {
		t.Fatalf("ip2long expect %d, actual %d", 2130706433, l1)
	}

	l2 := IP2Long("192.168.0.1")
	if l2 != 3232235521 {
		t.Fatalf("ip2long expect %d, actual %d", 3232235521, l2)
	}
}

func TestIP2LongInvalid(t *testing.T) {
	l1 := IP2Long("")
	if l1 != 0 {
		t.Fatalf("ip2long expect %d, actual %d", 0, l1)
	}

	l2 := IP2Long("192.168")
	if l2 != 0 {
		t.Fatalf("ip2long expect %d, actual %d", 0, l2)
	}

	l3 := IP2Long("192.168.0.1.1")
	if l3 != 0 {
		t.Fatalf("ip2long expect %d, actual %d", 0, l3)
	}

	l4 := IP2Long("192.168.0.1x")
	if l4 != 0 {
		t.Fatalf("ip2long expect %d, actual %d", 0, l4)
	}

	l5 := IP2Long("192.168.0.x")
	if l5 != 0 {
		t.Fatalf("ip2long expect %d, actual %d", 0, l5)
	}
}

func TestIPInRange(t *testing.T) {
	ipRangeList := []string{
		"127.0.0.1",
		"192.168.11.0/24",
	}

	ipStr1 := "127.0.0.1"
	if !IPInRange(net.ParseIP(ipStr1), ipRangeList) {
		t.Fatalf("IPInRange: %s expect in range, actual not", ipStr1)
	}

	ipStr2 := "192.168.11.95"
	if !IPInRange(net.ParseIP(ipStr2), ipRangeList) {
		t.Fatalf("IPInRange: %s expect in range, actual not", ipStr2)
	}

	ipStr3 := "192.168.0.1"
	if IPInRange(net.ParseIP(ipStr3), ipRangeList) {
		t.Fatalf("IPInRange: %s expect not in range, actual in range", ipStr3)
	}

	ipStr4 := "223.5.5.5"
	if IPInRange(net.ParseIP(ipStr4), ipRangeList) {
		t.Fatalf("IPInRange: %s expect not in range, actual in range", ipStr4)
	}
}

func TestGetInternalIPv4(t *testing.T) {
	ip, err := GetInternalIPv4()
	if err != nil {
		t.Fatal(err)
	}
	log.Println("internal ip:", ip)
}
