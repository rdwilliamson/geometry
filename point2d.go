// Package geometry implements some basic euclidean geometric primitives. All
// angles are measured in radians.
package geometry

import (
	"fmt"
	"math"
)

// Point2D can represent either a 2D point or a 2D vector.
type Point2D struct {
	X, Y float64
}

// DistTo returns the distance between the two points.
func (p1 Point2D) DistTo(p2 Point2D) float64 {
	dx, dy := p2.X-p1.X, p2.Y-p1.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// DistToSq returns the squared distance between the two points.
func (p1 Point2D) DistToSq(p2 Point2D) float64 {
	dx, dy := p2.X-p1.X, p2.Y-p1.Y
	return dx*dx + dy*dy
}

// Length returns the distance the point is from the origin, or the length of
// the vector.
func (p Point2D) Length() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// LengthSq returns the squared distance the point is from the origin, or the
// squared length of the vector.
func (p Point2D) LengthSq() float64 {
	return p.X*p.X + p.Y*p.Y
}

// Normalized returns a unit vector in the same direction.
func (p Point2D) Normalized() Point2D {
	s := 1.0 / math.Sqrt(p.X*p.X+p.Y*p.Y)
	return Point2D{p.X * s, p.Y * s}
}

// Plus returns the addition of the two vectors.
func (p1 Point2D) Plus(p2 Point2D) Point2D {
	return Point2D{p1.X + p2.X, p1.Y + p2.Y}
}

func (p Point2D) String() string {
	return fmt.Sprintf("(%g, %g)", p.X, p.Y)
}

// Minus returns the first vector minus the second.
func (p1 Point2D) Minus(p2 Point2D) Point2D {
	return Point2D{p1.X - p2.X, p1.Y - p2.Y}
}

// Scaled returns the scaled vector in the same direction.
func (p Point2D) Scaled(s float64) Point2D {
	return Point2D{p.X * s, p.Y * s}
}

// ToPoint3D returns a 3D copy of the point with passed z value.
func (p Point2D) ToPoint3D(z float64) Point3D {
	return Point3D{p.X, p.Y, z}
}

// DotProduct returns the dot product of the two vector.
func DotProduct2D(p1, p2 Point2D) float64 {
	return p1.X*p2.X + p1.Y*p2.Y
}
