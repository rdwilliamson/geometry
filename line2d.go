package geometry

import (
	"math"
)

// A line represented by two points.
type Line2D struct {
	P1, P2 Point2D
}

// Find the distance between a point and a line. Segment determines if the
// line is treated as a line segment or infinite line.
//
// For details see:
// http://softsurfer.com/Archive/algorithm_0102/algorithm_0102.htm
func LinePointDistance2D(l Line2D, p Point2D, segment bool) float64 {
	v := l.P2.Minus(l.P1)
	w := p.Minus(l.P1)
	c1 := DotProduct2D(Vector2D(v), Vector2D(w))
	if segment && c1 <= 0 {
		return p.DistanceTo(l.P1)
	}
	c2 := DotProduct2D(Vector2D(v), Vector2D(v))
	if segment && c2 <= c1 {
		return p.DistanceTo(l.P2)
	}
	return p.DistanceTo(l.P1.Plus(Point2D(Vector2D(v).Scaled(c1 / c2))))
}

// Converts the line to a vector from the first point to the second.
func (l Line2D) ToVector2D() Vector2D {
	return Vector2D{l.P2.X - l.P1.X, l.P2.Y - l.P1.X}
}

// Returns the horizontal component of the line.
func (l Line2D) Dx() float64 {
	return l.P2.X - l.P1.X
}

// Returns the vertical component of the line.
func (l Line2D) Dy() float64 {
	return l.P2.Y - l.P1.Y
}

// Returns the length of the line.
func (l Line2D) Length() float64 {
	return math.Hypot(l.P2.X-l.P1.X, l.P2.Y-l.P1.Y)
}

// Returns the squared length of the line.
func (l Line2D) LengthSquared() float64 {
	dx := l.P2.X - l.P1.X
	dy := l.P2.Y - l.P1.Y
	return dx*dx + dy*dy
}

// Returns the midpoint of the line.
func (l Line2D) Midpoint() Point2D {
	return Point2D{(l.P1.X + l.P2.X) * 0.5, (l.P1.Y + l.P2.Y) * 0.5}
}

// Returns the angle of the line.
func (l Line2D) Angle() float64 {
	return math.Atan2(l.P2.Y-l.P1.Y, l.P2.X-l.P1.X)
}

// Returns a normal vector with the same length as the line.
func (l Line2D) Normal() Vector2D {
	return Vector2D{l.P2.Y - l.P1.Y, l.P1.X - l.P2.X}
}

// Calculates the intersection point of two lines and determines if it occurred
// on both. From Graphics Gems III, Faster Line Segment Intersection.
func (l1 Line2D) Intersection(l2 Line2D) (Point2D, bool) {
	a := l1.P2.Minus(l1.P1)
	b := l2.P1.Minus(l2.P2)
	denominator := a.Y*b.X - a.X*b.Y
	if denominator == 0.0 {
		// TODO determine where on line at infinity they intersect
		return Point2D{math.Inf(1), math.Inf(1)}, false
	}
	denominator = 1.0 / denominator
	c := l1.P1.Minus(l2.P1)
	A := (b.Y*c.X - b.X*c.Y) * denominator
	intersection := l1.P1.Plus(Point2D{a.X * A, a.Y * A})
	if A < 0.0 || A > 1.0 {
		return intersection, false
	}
	B := (a.X*c.Y - a.Y*c.X) * denominator
	if B < 0.0 || B > 1.0 {
		return intersection, false
	}
	return intersection, true
}
