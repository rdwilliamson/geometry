// Geometry package implementing points, vectors, lines, planes, and triangles.
package geometry

import (
	"math"
)

// Check if a and b are very close.
func FuzzyEqual(a, b float64) bool {
	return math.Abs(a-b) <= 0.000000000001*math.Min(math.Abs(a), math.Abs(b))
}

// Clamp a such that it is at least min and no more than max.
func Clamp(a, min, max float64) float64 {
	if a < min {
		return min
	}
	if a > max {
		return max
	}
	return a
}

// Linear interpolation between x and y. With a == 0 corresponding to x and
// a == 1 corresponding to y.
func Mix(x, y, a float64) float64 {
	return x*(1-a) + y*a
}
