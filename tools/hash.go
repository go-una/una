package tools

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(txt string) string {
	h := md5.New()
	h.Write([]byte(txt))
	return hex.EncodeToString(h.Sum(nil))
}
