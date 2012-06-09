package geometry

import (
	"math"
)

func FuzzyEqual(a, b float64) bool {
	return math.Abs(a-b) <= 0.000000000001*math.Min(math.Abs(a), math.Abs(b))
}

func Clamp(a, min, max float64) float64 {
	if a < min {
		return min
	}
	if a > max {
		return max
	}
	return a
}

func Mix(x, y, a float64) float64 {
	return x*(1-a) + y*a
}
