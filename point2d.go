// Package geometry implements some basic euclidean geometric primitives. All
// angles are measured in radians.
package geometry

import (
	"math"
)

// Point2D can represent either a homogeneous 2D point or a 2D vector.
type Point2D struct {
	X, Y, W float64
}

// Creates a new 2D point or vector.
func NewPoint2D(x, y float64) Point2D {
	return Point2D{x, y, 1}
}

// Direction returns a unit vector pointing the same direction.
func (p Point2D) Direction() Point2D {
	s := 1 / p.Length()
	return Point2D{p.X * s, p.Y * s, 1}
}

// DistTo returns the distance between the two points.
func (p1 Point2D) DistTo(p2 Point2D) float64 {
	if p1.IsInf() || p2.IsInf() {
		return math.Inf(1)
	}
	p1 = p1.Normalized()
	p2 = p2.Normalized()
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// DistToSq returns the squared distance between the two points.
func (p1 Point2D) DistToSq(p2 Point2D) float64 {
	if p1.IsInf() || p2.IsInf() {
		return math.Inf(1)
	}
	p1 = p1.Normalized()
	p2 = p2.Normalized()
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	return dx*dx + dy*dy
}

// Checks if a point is at infinity
func (p Point2D) IsInf() bool {
	return p.W == 0 || math.IsNaN(p.X) || math.IsNaN(p.Y)
}

// Length returns the distance a normalized point is from the origin, or the
// length of the vector.
func (p Point2D) Length() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// LengthSq returns the squared distance the normalized point is from the
// origin, or the squared length of the vector.
func (p Point2D) LengthSq() float64 {
	return p.X*p.X + p.Y*p.Y
}

// Normalized returns a normalized point if possible.
func (p Point2D) Normalized() Point2D {
	if p.W == 1 {
		return p
	}
	if p.W == 0 {
		return Point2D{math.Inf(1), math.Inf(1), 0}
	}
	s := 1 / p.W
	return Point2D{p.X * s, p.Y * s, 1}
}

// Plus returns the addition of the two vectors.
func (p1 Point2D) Plus(p2 Point2D) Point2D {
	return Point2D{p1.X + p2.X, p1.Y + p2.Y, 1}
}

// Minus returns the first vector minus the second.
func (p1 Point2D) Minus(p2 Point2D) Point2D {
	return Point2D{p1.X - p2.X, p1.Y - p2.Y, 1}
}

// Scaled returns the scaled vector in the same direction.
func (p Point2D) Scaled(s float64) Point2D {
	return Point2D{p.X * s, p.Y * s, 1}
}

// ToPoint3D returns a 3D copy of the point with passed z value.
func (p Point2D) ToPoint3D(z float64) Point3D {
	return Point3D{p.X, p.Y, z, 1}
}

// DotProduct returns the dot product of the two vector.
func DotProduct2D(p1, p2 Point2D) float64 {
	return p1.X*p2.X + p1.Y*p2.Y
}
