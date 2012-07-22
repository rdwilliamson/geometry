package geometry

import (
	"math"
)

// 2D point.
type Point2D struct {
	X, Y float64
}

// Add a point.
func (p1 *Point2D) Add(p2 *Point2D) {
	p1.X += p2.X
	p1.Y += p2.Y
}

// Subtract a point.
func (p1 *Point2D) Subtract(p2 *Point2D) {
	p1.X -= p2.X
	p1.Y -= p2.Y
}

// Returns true if the two points are the same.
func (p1 *Point2D) Equal(p2 *Point2D) bool {
	return *p1 == *p2
}

// Returns true if the two points are very close.
func (p1 *Point2D) FuzzyEqual(p2 *Point2D) bool {
	dx, dy := p2.X-p1.X, p2.Y-p1.Y
	return dx*dx+dy*dy < 0.000000000001*0.000000000001
}

// Returns the distance between the two points.
func (p1 *Point2D) DistanceTo(p2 *Point2D) float64 {
	dx, dy := p2.X-p1.X, p2.Y-p1.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// Returns the squared distance between the two points.
func (p1 *Point2D) SquaredDistanceTo(p2 *Point2D) float64 {
	dx, dy := p2.X-p1.X, p2.Y-p1.Y
	return dx*dx + dy*dy
}
