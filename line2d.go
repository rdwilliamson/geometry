package geometry

import (
	"fmt"
	"math"
)

type Line2D struct {
	P1, P2 Point2D
}

func (l *Line2D) Angle() float64 {
	dx, dy := l.P2.X-l.P1.X, l.P2.Y-l.P1.Y
	return math.Atan(dy / dx)
}

func (l *Line2D) Length() float64 {
	dx, dy := l.P2.X-l.P1.X, l.P2.Y-l.P1.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func (l Line2D) String() string {
	return fmt.Sprintf("{%v %v}", l.P1, l.P2)
}

// Check if two line segments intersect.
// From Graphics Gems III, Faster Line Segment Intersection.
func LinesIntersect2D(l1, l2 Line2D) bool {
	a := l1.P2.Minus(l1.P1)
	b := l2.P1.Minus(l2.P2)
	denominator := a.Y*b.X - a.X*b.Y
	if denominator == 0.0 {
		return false
	}
	c := l1.P1.Minus(l2.P1)
	A := b.Y*c.X - b.X*c.Y
	if denominator > 0 && (A < 0.0 || A > denominator) {
		return false
	}
	B := a.X*c.Y - a.Y*c.X
	if denominator > 0 && (B > 0.0 || B < denominator) {
		return false
	}
	return true
}

// Returns the intersection point and if said point occures on both lines.
// From Graphics Gems III, Faster Line Segment Intersection.
func LinesIntersection2D(l1, l2 Line2D) (Point2D, bool) {
	a := l1.P2.Minus(l1.P1)
	b := l2.P1.Minus(l2.P2)
	denominator := a.Y*b.X - a.X*b.Y
	if denominator == 0.0 {
		return Point2D{math.Inf(1), math.Inf(1)}, false
	}
	denominator = 1.0 / denominator
	c := l1.P1.Minus(l2.P1)
	A := (b.Y*c.X - b.X*c.Y) * denominator
	intersection := l1.P1.Plus(a.Times(A))
	if A < 0.0 || A > 1.0 {
		return intersection, false
	}
	B := (a.X*c.Y - a.Y*c.X) * denominator
	if B < 0.0 || B > 1.0 {
		return intersection, false
	}
	return intersection, true
}
