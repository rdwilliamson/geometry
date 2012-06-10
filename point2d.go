package geometry

import (
	"math"
)

// 2D point.
type Point2D struct {
	X, Y float64
}

// Returns the addition of two points.
func (p1 Point2D) Plus(p2 Point2D) Point2D {
	return Point2D{p1.X + p2.X, p1.Y + p2.Y}
}

// Add a point.
func (p1 *Point2D) Add(p2 Point2D) {
	p1.X += p2.X
	p1.Y += p2.Y
}

// Returns the difference between the first and second point.
func (p1 Point2D) Minus(p2 Point2D) Point2D {
	return Point2D{p1.X - p2.X, p1.Y - p2.Y}
}

// Subtract a point.
func (p1 *Point2D) Subtract(p2 Point2D) {
	p1.X -= p2.X
	p1.Y -= p2.Y
}

// Returns true if the two points are the same.
func (p1 Point2D) Equal(p2 Point2D) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}

// Returns true if the two points are very close.
func (p1 Point2D) FuzzyEqual(p2 Point2D) bool {
	return FuzzyEqual(p1.X, p2.X) && FuzzyEqual(p1.Y, p2.Y)
}

// Returns the distance between the two points.
func (p1 Point2D) DistanceTo(p2 Point2D) float64 {
	return math.Hypot(p2.X-p1.X, p2.Y-p1.Y)
}

// Returns the squared distance between the two points.
func (p1 Point2D) SquaredDistanceTo(p2 Point2D) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	return dx*dx + dy*dy
}
