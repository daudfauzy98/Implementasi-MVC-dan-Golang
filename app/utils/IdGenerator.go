package utils

import "math/rand"

// RangeInt untuk membuat id secara otomatis
func RangeInt(low, hi int) int {
	return low + rand.Intn(hi-low)
}
