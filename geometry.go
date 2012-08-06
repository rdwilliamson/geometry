// Package geometry contains geometric primitives.
//
// It is safe for the output of methods to overlap with input.
// All angles are in radians.
package geometry

import (
	"math"
)

// Check if a and b are very close.
func FuzzyEqual(a, b float64) bool {
	return math.Abs(a-b) <= 1e-12*math.Min(math.Abs(a), math.Abs(b))
}
