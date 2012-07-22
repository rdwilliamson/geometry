package geometry

import (
	"math"
)

// A 3D point.
type Point3D struct {
	X, Y, Z float64
}

// Add a point.
func (p1 *Point3D) Add(p2 *Point3D) {
	p1.X += p2.X
	p1.Y += p2.Y
	p1.Z += p2.Z
}

// Subtract a point.
func (p1 *Point3D) Subtract(p2 *Point3D) {
	p1.X -= p2.X
	p1.Y -= p2.Y
	p1.Z -= p2.Z
}

// Returns true if the two points are the same.
func (p1 *Point3D) Equal(p2 *Point3D) bool {
	return *p1 == *p2
}

// Returns true if the two points are very close.
func (p1 *Point3D) FuzzyEqual(p2 *Point3D) bool {
	dx, dy, dz := p2.X-p1.X, p2.Y-p1.Y, p2.Z-p1.Z
	return dx*dx+dy*dy+dz*dz < 0.000000000001*0.000000000001
}

// Returns the distance between the two points.
func (p1 *Point3D) DistanceTo(p2 *Point3D) float64 {
	dx, dy, dz := p2.X-p1.X, p2.Y-p1.Y, p2.Z-p1.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

// Returns the squared distance between the two points.
func (p1 *Point3D) SquaredDistanceTo(p2 *Point3D) float64 {
	dx, dy, dz := p2.X-p1.X, p2.Y-p1.Y, p2.Z-p1.Z
	return dx*dx + dy*dy + dz*dz
}
