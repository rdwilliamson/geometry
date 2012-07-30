package geometry

import "math"

// A Vector3D is a 3D vector or 3D point depending on how it's used.
type Vector3D struct {
	X, Y, Z float64
}

// NewVector2D returns a new Vector2D.

// Add sets z to the piecewise sum a+b and returns z.

// AngleDifference returns the angle between a and b.

// AngleCosDifference returns the squared cos of the angle between a and b.

// DirectionEqual compares the direction of a and b and returns a boolean
// indicating if they are equal.

// DirectionFuzzyEqual compares the direction of a and b and returns a boolean
// indicating if they are equal.

// Distance returns the distance between a and b.

// DistanceSquared returns the squared distance between a and b.

// Divide sets z to the piecewise quotient a/b and returns z.

// DotProduct returns the dot product of a and b.

// Equal compares a and b and returns a boolean indicating if they are equal.

// FuzzyEqual compares a and b and returns a boolean indicating if they are
// very close.

// Magnitude returns the magnitude of a.

// MagnitudeSquared returns the squared magnitude of a.

// Multiply sets z to the piecewise multiplication of a*b and returns z.

// Normalize sets z to a unit vector in the same direction as x and returns z.

// Projection sets z to the projection of a onto b and returns z.

// Set sets z to x and returns z.

// ScalarProjection returns the scalar projection of a onto b.

// Scale sets z to scalar multiplication n*x and returns z.

// Subtract Sets z to the piecewise difference a-b and returns z.
