package geometry

import "math"

// A Point2D represents a 64 bit precision 2D point.
type Point2D struct {
	X, Y float64
}

// Add sets z to the piecewise sum a+b and returns z.
func (z *Point2D) Add(a, b *Point2D) *Point2D {
	z.X = a.X + b.X
	z.Y = a.Y + b.Y
	return z
}

// Dist calculates and returns the distance between a and b.
func (a *Point2D) Dist(b *Point2D) float64 {
	dx, dy := b.X-a.X, b.Y-a.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// Equal compares a and b and returns a boolean indicating if they are equal.
func (a *Point2D) Equal(b *Point2D) bool {
	return *a == *b
}

// Equal compares a and b and returns a boolean indicating if they are very close.
func (a *Point2D) FuzzyEqual(b *Point2D) bool {
	dx, dy := b.X-a.X, b.Y-a.Y
	return dx*dx+dy*dy < 0.000000000001*0.000000000001
}

// Set sets z to x and returns z.
func (z *Point2D) Set(x *Point2D) *Point2D {
	*z = *x
	return z
}

// SqDist calculates and returns the squared distance between a and b.
func (a *Point2D) SqDist(b *Point2D) float64 {
	dx, dy := b.X-a.X, b.Y-a.Y
	return dx*dx + dy*dy
}

// Sets z to the piecewise difference a-b and returns z;
func (z *Point2D) Sub(a, b *Point2D) *Point2D {
	z.X = a.X - b.X
	z.Y = a.Y - b.Y
	return z
}
