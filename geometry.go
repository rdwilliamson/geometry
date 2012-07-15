// Package geometry contains geometric primitives.
//
// Angles are in radians.
package geometry

import (
	"math"
)

// Check if a and b are very close.
func FuzzyEqual(a, b float64) bool {
	return math.Abs(a-b) <= 0.000000000001*math.Min(math.Abs(a), math.Abs(b))
}
