package tools

import (
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
