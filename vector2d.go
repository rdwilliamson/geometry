package geometry

import (
	"math"
)

// A Vector2D is a 2D vector or 2D point depending on how it's used.
type Vector2D struct {
	X, Y float64
}

// NewVector2D returns a new Vector2D.
func NewVector2D(x, y float64) *Vector2D {
	return &Vector2D{x, y}
}

// Add sets z to the piecewise sum a+b then returns z.
func (z *Vector2D) Add(a, b *Vector2D) *Vector2D {
	z.X = a.X + b.X
	z.Y = a.Y + b.Y
	return z
}

// Copy sets z to x then returns z.
func (z *Vector2D) Copy(x *Vector2D) *Vector2D {
	z.X = x.X
	z.Y = x.Y
	return z
}

// DirectionEqual compares the direction of a and b then returns true if they
// are exactly equal or false otherwise.
func (a *Vector2D) DirectionEqual(b *Vector2D) bool {
	if a.X == 0 && b.X == 0 {
		return a.Y*b.Y > 0
	}
	s := a.X / b.X
	return s > 0 && a.Y == s*b.Y
}

// DirectionFuzzyEqual compares the direction of a and b then returns true if
// they are very close or false otherwise.
func (a *Vector2D) DirectionFuzzyEqual(b *Vector2D) bool {
	if FuzzyEqual(math.Abs(a.X)+math.Abs(b.X), 0) {
		return a.Y*b.Y > 0
	}
	if a.X > b.X {
		s := a.X / b.X
		return s > 0 && FuzzyEqual(a.Y, s*b.Y)
	}
	s := b.X / a.X
	return s > 0 && FuzzyEqual(s*a.Y, b.Y)
}

// Divide sets z to the piecewise quotient a/b then returns z.
func (z *Vector2D) Divide(a, b *Vector2D) *Vector2D {
	z.X = a.X / b.X
	z.Y = a.Y / b.Y
	return z
}

// DotProduct returns the dot product of a and b.
func (a *Vector2D) DotProduct(b *Vector2D) float64 {
	return a.X*b.X + a.Y*b.Y
}

// Equal compares a and b then returns true of they are exactly equal or false
// otherwise.
func (a *Vector2D) Equal(b *Vector2D) bool {
	return *a == *b
}

// FuzzyEqual compares a and b then returns true if they are very close or
// false otherwise.
func (a *Vector2D) FuzzyEqual(b *Vector2D) bool {
	return FuzzyEqual(a.X, b.X) && FuzzyEqual(a.Y, b.Y)
}

// Magnitude returns the magnitude of x.
func (x *Vector2D) Magnitude() float64 {
	return math.Sqrt(x.X*x.X + x.Y*x.Y)
}

// MagnitudeSquared returns the squared magnitude of x.
func (x *Vector2D) MagnitudeSquared() float64 {
	return x.X*x.X + x.Y*x.Y
}

// Multiply sets z to the piecewise multiplication of a*b then returns z.
func (z *Vector2D) Multiply(a, b *Vector2D) *Vector2D {
	z.X = a.X * b.X
	z.Y = a.Y * b.Y
	return z
}

// Normalize sets z to a unit vector in the same direction as x then returns z.
func (z *Vector2D) Normalize(x *Vector2D) *Vector2D {
	l := 1 / math.Sqrt(x.X*x.X+x.Y*x.Y)
	z.X = x.X * l
	z.Y = x.Y * l
	return z
}

// Projection sets z to the projection of a onto b then returns z.
func (z *Vector2D) Projection(a, b *Vector2D) *Vector2D {
	s := (a.X*b.X + a.Y*b.Y) / (b.X*b.X + b.Y*b.Y)
	z.X = b.X * s
	z.Y = b.Y * s
	return z
}

// ScalarProjection returns the scalar projection of a onto b.
func (a *Vector2D) ScalarProjection(b *Vector2D) float64 {
	return (a.X*b.X + a.Y*b.Y) / (b.X*b.X + b.Y*b.Y)
}

// Scale sets z to scalar multiplication n*x then returns z.
func (z *Vector2D) Scale(x *Vector2D, n float64) *Vector2D {
	z.X = x.X * n
	z.Y = x.Y * n
	return z
}

// Subtract Sets z to the piecewise difference a-b then returns z.
func (z *Vector2D) Subtract(a, b *Vector2D) *Vector2D {
	z.X = a.X - b.X
	z.Y = a.Y - b.Y
	return z
}
