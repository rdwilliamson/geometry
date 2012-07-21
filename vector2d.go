package geometry

import (
	"math"
)

// 2D vector.
type Vector2D struct {
	X, Y float64
}

// Converts the vector to a line from the origin.
func (v *Vector2D) ToLine2D() Line2D {
	return Line2D{Point2D{0, 0}, Point2D(*v)}
}

// Add a vector.
func (v1 *Vector2D) Add(v2 *Vector2D) {
	v1.X += v2.X
	v1.Y += v2.Y
}

// Subtract a vector.
func (v1 *Vector2D) Subtract(v2 *Vector2D) {
	v1.X -= v2.X
	v1.Y -= v2.Y
}

// Multiplies the vector element by element.
func (v1 *Vector2D) Multiply(v2 *Vector2D) {
	v1.X *= v2.X
	v1.Y *= v2.X
}

// Divides the vector element by element.
func (v1 *Vector2D) Divide(v2 *Vector2D) {
	v1.X /= v2.X
	v1.Y /= v2.Y
}

// Scales the vector.
func (v *Vector2D) Scale(s float64) {
	v.X *= s
	v.Y *= s
}

// Returns the length of the vector.
func (v *Vector2D) Length() float64 {
	return math.Hypot(v.X, v.Y)
}

// Returns the squared length of the vector.
func (v *Vector2D) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y
}

// Normalizes the vector. The zero vector remains unchanged.
func (v *Vector2D) Normalize() {
	if v.X == 0 && v.Y == 0 {
		return
	}
	l := 1 / math.Hypot(v.X, v.Y)
	v.X *= l
	v.Y *= l
}

// Returns true if the two vectors are the same.
func (v1 *Vector2D) Equal(v2 *Vector2D) bool {
	return v1.X == v2.X && v1.Y == v2.Y
}

// Returns true if the two vectors are close.
func (v1 *Vector2D) FuzzyEqual(v2 *Vector2D) bool {
	return FuzzyEqual(v1.X, v2.X) && FuzzyEqual(v1.Y, v2.Y)
}

// Returns the scalar projection.
func (v1 *Vector2D) ScalarProjection(v2 *Vector2D) float64 {
	return (v1.X*v2.X + v1.Y*v2.Y) / (v2.X*v2.X + v2.Y*v2.Y)
}

// Projects v1 onto v2.
func (v1 *Vector2D) ProjectedOnto(v2 *Vector2D) {
	s := (v1.X*v2.X + v1.Y*v2.Y) / (v2.X*v2.X + v2.Y*v2.Y)
	v1.X, v1.Y = v2.X*s, v2.Y*s
}

// Returns the angle between two vectors.
func (v1 *Vector2D) AngleBetween(v2 *Vector2D) float64 {
	return math.Abs(math.Atan2(v1.Y, v1.X) - math.Atan2(v2.Y, v2.X))
}

// Returns the dot product of two vectors.
func DotProduct2D(v1, v2 *Vector2D) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}
