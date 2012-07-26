package geometry

import "math"

// A Vector2D is a 2D vector or 2D point depending on how it's used.
type Vector2D struct {
	X, Y float64
}

// NewVector2D returns a new Vector2D.
func NewVector2D(x, y float64) *Vector2D {
	return &Vector2D{x, y}
}

// Add sets z to the piecewise sum a+b and returns z.
func (z *Vector2D) Add(a, b *Vector2D) *Vector2D {
	z.X = a.X + b.X
	z.Y = a.Y + b.Y
	return z
}

// AngleDifference returns the angle between a and b.
func (a *Vector2D) AngleDifference(b *Vector2D) float64 {
	return math.Acos((a.X*b.X + a.Y*b.Y) / math.Sqrt((a.X*a.X+a.Y*a.Y)*(b.X*b.X+b.Y*b.Y)))
}

// AngleCosDifference returns the cos of the angle between a and b.
func (a *Vector2D) AngleCosDifference(b *Vector2D) float64 {
	return (a.X*b.X + a.Y*b.Y) / math.Sqrt((a.X*a.X+a.Y*a.Y)*(b.X*b.X+b.Y*b.Y))
}

// AngleCosDifference returns the squared cos of the angle between a and b.
func (a *Vector2D) AngleCosSquaredDifference(b *Vector2D) float64 {
	dot := (a.X*b.X + a.Y*b.Y)
	return dot * dot / ((a.X*a.X + a.Y*a.Y) * (b.X*b.X + b.Y*b.Y))
}

// Dist returns the distance between a and b.
func (a *Vector2D) Distance(b *Vector2D) float64 {
	dx, dy := b.X-a.X, b.Y-a.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// DistSq returns the squared distance between a and b.
func (a *Vector2D) DistanceSquared(b *Vector2D) float64 {
	dx, dy := b.X-a.X, b.Y-a.Y
	return dx*dx + dy*dy
}

// Divide sets z to the piecewise quotient a/b and returns z.
func (z *Vector2D) Divide(a, b *Vector2D) *Vector2D {
	z.X = a.X / b.X
	z.Y = a.Y / b.Y
	return z
}

// DotProduct returns the dot product of a and b.
func (a *Vector2D) DotProduct(b *Vector2D) float64 {
	return a.X*b.X + a.Y*b.Y
}

// Equal compares a and b and returns a boolean indicating if they are equal.
func (a *Vector2D) Equal(b *Vector2D) bool {
	return *a == *b
}

// FuzzyEqual compares a and b and returns a boolean indicating if they are
// very close.
func (a *Vector2D) FuzzyEqual(b *Vector2D) bool {
	dx, dy := b.X-a.X, b.Y-a.Y
	return dx*dx+dy*dy < 0.000000000001*0.000000000001
}

// Magnitude returns the magnitude of a.
func (a *Vector2D) Magnitude() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y)
}

// MagnitudeSquared returns the squared magnitude of a.
func (a *Vector2D) MagnitudeSquared() float64 {
	return a.X*a.X + a.Y*a.Y
}

// Multiply sets z to the piecewise multiplication of a*b and returns z.
func (z *Vector2D) Multiply(a, b *Vector2D) *Vector2D {
	z.X = a.X * b.X
	z.Y = a.Y * b.Y
	return z
}

// Normalize sets z to a unit vector in the same direction as x and returns z.
func (z *Vector2D) Normalize(x *Vector2D) *Vector2D {
	l := 1 / math.Sqrt(x.X*x.X+x.Y*x.Y)
	z.X = x.X * l
	z.Y = x.Y * l
	return z
}

// Projection sets z to the projection of a onto b and returns z.
func (z *Vector2D) Projection(a, b *Vector2D) *Vector2D {
	s := (a.X*b.X + a.Y*b.Y) / (b.X*b.X + b.Y*b.Y)
	z.X = b.X * s
	z.Y = b.Y * s
	return z
}

// Set sets z to x and returns z.
func (z *Vector2D) Set(x *Vector2D) *Vector2D {
	z.X = x.X
	z.Y = x.Y
	return z
}

// ScalarProjection returns the scalar projection of a onto b.
func (a *Vector2D) ScalarProjection(b *Vector2D) float64 {
	return (a.X*b.X + a.Y*b.Y) / (b.X*b.X + b.Y*b.Y)
}

// Scale sets z to scalar multiplication n*x and returns z.
func (z *Vector2D) Scale(x *Vector2D, n float64) *Vector2D {
	z.X = x.X * n
	z.Y = x.Y * n
	return z
}

// Subtract Sets z to the piecewise difference a-b and returns z;
func (z *Vector2D) Subtract(a, b *Vector2D) *Vector2D {
	z.X = a.X - b.X
	z.Y = a.Y - b.Y
	return z
}
