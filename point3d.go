package geometry

import (
	"math"
)

// 3D point.
type Point3D struct {
	X, Y, Z float64
}

// Returns the addition of two points.
func (p1 Point3D) Plus(p2 Point3D) Point3D {
	return Point3D{p1.X + p2.X, p1.Y + p2.Y, p1.Z + p2.Z}
}

// Add a point.
func (p1 *Point3D) Add(p2 Point3D) {
	p1.X += p2.X
	p1.Y += p2.Y
	p1.Z += p2.Z
}

// Returns the difference between the first and second point.
func (p1 Point3D) Minus(p2 Point3D) Point3D {
	return Point3D{p1.X - p2.X, p1.Y - p2.Y, p1.Z - p2.Z}
}

// Subtract a point.
func (p1 *Point3D) Subtract(p2 Point3D) {
	p1.X -= p2.X
	p1.Y -= p2.Y
	p1.Z -= p2.Z
}

// Returns true if the two points are the same.
func (p1 Point3D) Equal(p2 Point3D) bool {
	return p1.X == p2.X && p1.Y == p2.Y && p1.Z == p2.Z
}

// Returns true if the two points are very close.
func (p1 Point3D) FuzzyEqual(p2 Point3D) bool {
	return FuzzyEqual(p1.X, p2.X) && FuzzyEqual(p1.Y, p2.Y) && FuzzyEqual(p1.Z, p2.Z)
}

// Returns the distance between the two points.
func (p1 Point3D) DistanceTo(p2 Point3D) float64 {
	dx, dy, dz := p2.X-p1.X, p2.Y-p1.Y, p2.Z-p1.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

// Returns the squared distance between the two points.
func (p1 Point3D) SquaredDistanceTo(p2 Point3D) float64 {
	dx, dy, dz := p2.X-p1.X, p2.Y-p1.Y, p2.Z-p1.Z
	return dx*dx + dy*dy + dz*dz
}
