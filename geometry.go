// Package geometry contains geometric primitives.
//
// It is safe for the output of methods to overlap with input.
// All angles are in radians.
package geometry

// Check if a and b are very close.
func FuzzyEqual(a, b float64) bool {
	// TODO handle case when a and b are near zero and on oposite sides of it
	absA := a
	if absA < 0 {
		absA = -absA
	}
	absB := b
	if absB < 0 {
		absB = -absB
	}
	if absA < 1 || absB < 1 {
		absA += 1
		absB += 1
	}
	if absA < absB {
		return absB-absA <= 1e-12*absA
	}
	return absA-absB <= 1e-12*absB
}
