package geometry

import (
	"math"
)

// Point3D can represent either a homogeneous 3D point or a 3D vector.
type Point3D struct {
	X, Y, Z, W float64
}

// Creates a new 3D point or vector.
func NewPoint3D(x, y, z float64) Point3D {
	return Point3D{x, y, z, 1}
}

// Length returns the distance a normalized point is from the origin, or the
// length of the vector.
func (p Point3D) Length() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y + p.Z*p.Z)
}

// LengthSq returns the squared distance a normalized point is from the origin,
// or the squared length of the vector.
func (p Point3D) LengthSq() float64 {
	return p.X*p.X + p.Y*p.Y + p.Z*p.Z
}

// Direction returns a unit vector in the same direction.
func (p Point3D) Direction() Point3D {
	s := 1 / p.Length()
	return Point3D{p.X * s, p.Y * s, p.Z * s, 1}
}

// Scaled returns the scaled vector.
func (p Point3D) Scaled(s float64) Point3D {
	return Point3D{p.X * s, p.Y * s, p.Z * s, 1}
}

// ToPoint2D returns a 2D copy of the point throwing away the z value.
func (p Point3D) ToPoint2D() Point2D {
	return Point2D{p.X, p.Y, 1}
}
