package geometry

import (
	"math"
)

// A line represented by two points.
type Line2D struct {
	P1, P2 Point2D
}

// Converts the line to a vector from the first point to the second.
func (l *Line2D) ToVector2D() *Vector2D {
	return &Vector2D{l.P2.X - l.P1.X, l.P2.Y - l.P1.X}
}

// Returns the length of the line.
func (l *Line2D) Length() float64 {
	return math.Hypot(l.P2.X-l.P1.X, l.P2.Y-l.P1.Y)
}

// Returns the squared length of the line.
func (l *Line2D) LengthSquared() float64 {
	dx := l.P2.X - l.P1.X
	dy := l.P2.Y - l.P1.Y
	return dx*dx + dy*dy
}

// Returns the midpoint of the line.
func (l *Line2D) Midpoint() *Point2D {
	return &Point2D{(l.P1.X + l.P2.X) * 0.5, (l.P1.Y + l.P2.Y) * 0.5}
}

// Returns a normal vector with the same length as the line.
func (l *Line2D) Normal() *Vector2D {
	return &Vector2D{l.P2.Y - l.P1.Y, l.P1.X - l.P2.X}
}

// Find the distance between a point and a line segment. See
// http://softsurfer.com/Archive/algorithm_0102/algorithm_0102.htm
func LineSegmentPointDistance2D(l *Line2D, p *Point2D) float64 {
	ldx, ldy := l.P2.X-l.P1.X, l.P2.Y-l.P1.Y
	c1 := ldx*(p.X-l.P1.X) + ldy*(p.Y-l.P1.Y)
	if c1 <= 0 {
		return math.Hypot(p.X-l.P1.X, p.Y-l.P1.Y)
	}
	c2 := ldx*ldx + ldy*ldy
	if c2 <= c1 {
		return math.Hypot(p.X-l.P2.X, p.Y-l.P2.Y)
	}
	c1 /= c2
	return math.Hypot(p.X-(l.P1.X+ldx*c1), p.Y-(l.P1.Y+ldy*c1))
}

// Find the distance between a point and a line. See
// http://softsurfer.com/Archive/algorithm_0102/algorithm_0102.htm
func LinePointDistance2D(l *Line2D, p *Point2D) float64 {
	ldx, ldy := l.P2.X-l.P1.X, l.P2.Y-l.P1.Y
	u := (ldx*(p.X-l.P1.X) + ldy*(p.Y-l.P1.Y)) / (ldx*ldx + ldy*ldy)
	return math.Hypot(p.X-(l.P1.X+ldx*u), p.Y-(l.P1.Y+ldy*u))
}

// Calculates the intersection point of two lines and determines if it occurred
// on both. From Graphics Gems III, Faster Line Segment Intersection.
// TODO break into seperate functions
func (l1 *Line2D) Intersection(l2 *Line2D) (*Point2D, bool) {
	a := l1.P2.Copy()
	a.Subtract(&l1.P1)
	b := l2.P2.Copy()
	b.Subtract(&l2.P1)
	denominator := a.Y*b.X - a.X*b.Y
	if denominator == 0.0 {
		return &Point2D{math.Inf(1), math.Inf(1)}, false
	}
	denominator = 1.0 / denominator
	c := l1.P1.Copy()
	c.Subtract(&l2.P1)
	A := (b.Y*c.X - b.X*c.Y) * denominator
	intersection := &Point2D{l1.P1.X + a.X*A, l1.P1.Y + a.Y*A}
	if A < 0.0 || A > 1.0 {
		return intersection, false
	}
	B := (a.X*c.Y - a.Y*c.X) * denominator
	if B < 0.0 || B > 1.0 {
		return intersection, false
	}
	return intersection, true
}
