package tools

import "testing"

func TestMD5(t *testing.T) {
	actual1 := MD5("abc")
	expect1 := "900150983cd24fb0d6963f7d28e17f72"
	if actual1 != expect1 {
		t.Fatalf("MD5(\"acb\") error, expect: %s, actual: %s", expect1, actual1)
	}

	actual2 := MD5("")
	expect2 := "d41d8cd98f00b204e9800998ecf8427e"
	if actual2 != expect2 {
		t.Fatalf("MD5(\"\") error, expect: %s, actual: %s", expect2, actual2)
	}
}
