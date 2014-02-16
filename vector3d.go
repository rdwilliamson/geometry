package geometry

import (
	"math"
)

// A Vector3D is a 2D vector or 2D point depending on how it's used.
type Vector3D struct {
	X, Y, Z float64
}

// NewVector3D returns a new Vector3D.
func NewVector3D(x, y, z float64) *Vector3D {
	return &Vector3D{x, y, z}
}

// Add sets z to the piecewise sum a+b then returns z.
func (z *Vector3D) Add(a, b *Vector3D) *Vector3D {
	z.X = a.X + b.X
	z.Y = a.Y + b.Y
	z.Z = a.Z + b.Z
	return z
}

// Copy sets z to x then returns z.
func (z *Vector3D) Copy(x *Vector3D) *Vector3D {
	z.X = x.X
	z.Y = x.Y
	z.Z = x.Z
	return z
}

// CrossProduct sets z to the cross product of a and b then returns z.
func (z *Vector3D) CrossProduct(a, b *Vector3D) *Vector3D {
	z.X = a.Y*b.Z - a.Z*b.Y
	z.Y = a.Z*b.X - a.X*b.Z
	z.Z = a.X*b.Y - a.Y*b.X
	return z
}

// DirectionEqual compares the direction of a and b then returns true if they
// are exactly equal or false otherwise.
func (a *Vector3D) DirectionEqual(b *Vector3D) bool {
	if a.X == 0 && b.X == 0 {
		if a.Y == 0 && b.Y == 0 {
			return a.Z*b.Z > 0
		} else {
			s := a.Y / b.Y
			return s > 0 && a.Z == s*b.Z
		}
	}
	s := a.X / b.X
	if s < 0 || a.Y != s*b.Y {
		return false
	}
	return a.Z == s*b.Z
}

// DirectionFuzzyEqual compares the direction of a and b then returns true if
//they are very close or false otherwise.
func (a *Vector3D) DirectionFuzzyEqual(b *Vector3D) bool {
	if FuzzyEqual(math.Abs(a.X)+math.Abs(b.X), 0) {
		if FuzzyEqual(math.Abs(a.Y)+math.Abs(b.Y), 0) {
			return a.Z*b.Z > 0
		} else {
			if a.Y > b.Y {
				s := a.Y / b.Y
				return s > 0 && FuzzyEqual(a.Z, s*b.Z)
			}
			s := b.Y / a.Y
			return s > 0 && FuzzyEqual(s*a.Z, b.Z)
		}
	}
	if a.X > b.X {
		s := a.X / b.X
		return s > 0 && FuzzyEqual(a.Y, s*b.Y) && FuzzyEqual(a.Z, s*b.Z)
	}
	s := b.X / a.X
	return s > 0 && FuzzyEqual(s*a.Y, b.Y) && FuzzyEqual(s*a.Z, b.Z)
}

// Divide sets z to the piecewise quotient a/b then returns z.
func (z *Vector3D) Divide(a, b *Vector3D) *Vector3D {
	z.X = a.X / b.X
	z.Y = a.Y / b.Y
	z.Z = a.Z / b.Z
	return z
}

// DotProduct returns the dot product of a and b.
func (a *Vector3D) DotProduct(b *Vector3D) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// Equal compares a and b then returns true if they are exactly equal or false
// otherwise.
func (a *Vector3D) Equal(b *Vector3D) bool {
	return *a == *b
}

// FuzzyEqual compares a and b and returns true if they are very close or false
// otherwise.
func (a *Vector3D) FuzzyEqual(b *Vector3D) bool {
	return FuzzyEqual(a.X, b.X) && FuzzyEqual(a.Y, b.Y) && FuzzyEqual(a.Z, b.Z)
}

// Magnitude returns the magnitude of x.
func (x *Vector3D) Magnitude() float64 {
	return math.Sqrt(x.X*x.X + x.Y*x.Y + x.Z*x.Z)
}

// MagnitudeSquared returns the squared magnitude of x.
func (x *Vector3D) MagnitudeSquared() float64 {
	return x.X*x.X + x.Y*x.Y + x.Z*x.Z
}

// Multiply sets z to the piecewise multiplication of a*b then returns z
func (z *Vector3D) Multiply(a, b *Vector3D) *Vector3D {
	z.X = a.X * b.X
	z.Y = a.Y * b.Y
	z.Z = a.Z * b.Z
	return z
}

// Normalize sets x to a unit vector in the same direction as x then returns x.
func (x *Vector3D) Normalize() *Vector3D {
	l := 1 / math.Sqrt(x.X*x.X+x.Y*x.Y+x.Z*x.Z)
	x.X *= l
	x.Y *= l
	x.Z *= l
	return x
}

// Normalized sets z to a unit vector in the same direction as x then returns z.
func (z *Vector3D) Normalized(x *Vector3D) *Vector3D {
	l := 1 / math.Sqrt(x.X*x.X+x.Y*x.Y+x.Z*x.Z)
	z.X = x.X * l
	z.Y = x.Y * l
	z.Z = x.Z * l
	return z
}

// Projection sets z to the projection of a onto b then returns z.
func (z *Vector3D) Projection(a, b *Vector3D) *Vector3D {
	s := (a.X*b.X + a.Y*b.Y + a.Z*b.Z) / (b.X*b.X + b.Y*b.Y + b.Z*b.Z)
	z.X = b.X * s
	z.Y = b.Y * s
	z.Z = b.Z * s
	return z
}

// ScalarProjection returns the scalar projection of a onto b.
func (a *Vector3D) ScalarProjection(b *Vector3D) float64 {
	return (a.X*b.X + a.Y*b.Y + a.Z*b.Z) / (b.X*b.X + b.Y*b.Y + b.Z*b.Z)
}

// Scale sets z to scalar multiplication n*x then returns z.
func (z *Vector3D) Scale(x *Vector3D, n float64) *Vector3D {
	z.X = x.X * n
	z.Y = x.Y * n
	z.Z = x.Z * n
	return z
}

// Subtract Sets z to the piecewise difference a-b then returns z.
func (z *Vector3D) Subtract(a, b *Vector3D) *Vector3D {
	z.X = a.X - b.X
	z.Y = a.Y - b.Y
	z.Z = a.Z - b.Z
	return z
}
