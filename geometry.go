// Package geometry contains geometric primitives.
//
// It is safe for the output of methods to overlap with input.
// All angles are in radians.
package geometry

// Check if a and b are very close.

// Check if a and b are very close.
func FuzzyEqual(a, b float64) bool {
	// handle case when a and b are near zero and on opposite sides of it
	if a*b < 0 {
		a += 1
		b += 1
	}

	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	// if values are small compare around one instead of zero
	if a < 1 || b < 1 {
		a += 1
		b += 1
	}

	if a < b {
		return b-a <= 1e-12*a
	}
	return a-b <= 1e-12*b
}
