package tools

import "time"

func CstZone() *time.Location {
	return time.FixedZone("CST", 8*3600)
}
