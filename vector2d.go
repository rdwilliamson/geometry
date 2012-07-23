package geometry

import "math"

// A Vector2D is a 2D vector or 2D point depending on how it's used.
type Vector2D struct {
	X, Y float64
}

func NewVector2D(x, y float64) *Vector2D {
	return &Vector2D{x, y}
}

// Add sets z to the piecewise sum a+b and returns z.
func (z *Vector2D) Add(a, b *Vector2D) *Vector2D {
	z.X = a.X + b.X
	z.Y = a.Y + b.Y
	return z
}

// AngleDif returns the angle between a and b.

// CosAngleDif returns the cos of the angle between a and b.

// Dist returns the distance between a and b.
func (a *Vector2D) Dist(b *Vector2D) float64 {
	dx, dy := b.X-a.X, b.Y-a.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// DistSq returns the squared distance between a and b.
func (a *Vector2D) DistSq(b *Vector2D) float64 {
	dx, dy := b.X-a.X, b.Y-a.Y
	return dx*dx + dy*dy
}

// Div sets z to the piecewise quotient a/b and returns z.

// Dot

// Equal compares a and b and returns a boolean indicating if they are equal.
func (a *Vector2D) Equal(b *Vector2D) bool {
	return *a == *b
}

// Equal compares a and b and returns a boolean indicating if they are very close.
func (a *Vector2D) FuzzyEqual(b *Vector2D) bool {
	dx, dy := b.X-a.X, b.Y-a.Y
	return dx*dx+dy*dy < 0.000000000001*0.000000000001
}

// Mag returns the magnitude of the vector.

// MagSq

// Set sets z to x and returns z.
func (z *Vector2D) Set(x *Vector2D) *Vector2D {
	z.X = x.X
	z.Y = x.Y
	return z
}

// Mul

// Norm

// ProjOnto

// ScalarProj

// Scale

// SetLine sets z to the vector from l.P1 to l.P2 and returns z.

// SetLineNorm sets z to the normal of l with magnitude equal to l's segment
// length and 

// SetMidpoint sets z to the midpoint of l and returns z.

// Sets z to the piecewise difference a-b and returns z;
func (z *Vector2D) Sub(a, b *Vector2D) *Vector2D {
	z.X = a.X - b.X
	z.Y = a.Y - b.Y
	return z
}
