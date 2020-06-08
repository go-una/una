package tools

import (
	"math/rand"
	"time"
)

func Rand(min, max int) int {
	if min >= max {
		return max
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min+1) + min
}
