package geometry

import (
	"math"
)

func Abs32(a float32) float32 {
	if a < 0 {
		return -a
	}
	if a == 0 {
		return 0
	}
	return a
}

func Min32(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func FuzzyEqual32(a, b float32) bool {
	return Abs32(a-b) <= 0.00001*Min32(Abs32(a), Abs32(b))
}

func FuzzyEqual64(a, b float64) bool {
	return math.Abs(a-b) <= 0.000000000001*math.Min(math.Abs(a), math.Abs(b))
}

func Clamp64(a, min, max float64) float64 {
	if a < min {
		return min
	}
	if a > max {
		return max
	}
	return a
}

func Mix64(x, y, a float64) float64 {
	return x*(1-a) + y*a
}
