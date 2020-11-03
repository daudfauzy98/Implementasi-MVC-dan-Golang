package utils

import "math/rand"

func RangeInt(low, hi int) int {
	return low + rand.Intn(hi-low)
}
