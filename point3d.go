package geometry

import (
	"math"
)

// Point3D can represent either a 3D point or a 3D vector.
type Point3D struct {
	X, Y, Z float64
}

// Length returns the distance the point is from the origin, or the length of
// the vector.
func (p Point3D) Length() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y + p.Z*p.Z)
}

// LengthSq returns the squared distance the point is from the origin, or the
// squared length of the vector.
func (p Point3D) LengthSq() float64 {
	return p.X*p.X + p.Y*p.Y + p.Z*p.Z
}

// Normalized returns a unit vector in the same direction.
func (p Point3D) Normalized() Point3D {
	s := 1.0 / math.Sqrt(p.X*p.X+p.Y*p.Y+p.Z*p.Z)
	return Point3D{p.X * s, p.Y * s, p.Z * s}
}

// Scaled returns the scaled vector.
func (p Point3D) Scaled(s float64) Point3D {
	return Point3D{p.X * s, p.Y * s, p.Z * s}
}

// ToPoint2D returns a 2D copy of the point throwing away the z value.
func (p Point3D) ToPoint2D() Point2D {
	return Point2D{p.X, p.Y}
}
