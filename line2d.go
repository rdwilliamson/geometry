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
		return PointDistance2D(p, l.P1)
	}
	c2 := DotProduct2D(Vector2D(v), Vector2D(v))
	if segment && c2 <= c1 {
		return PointDistance2D(p, l.P2)
	}
	pp := l.P1.Plus(Point2D(Vector2D(v).Scaled(c1 / c2)))
	return PointDistance2D(p, pp)
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

// Returns the angle of the line from -Pi to Pi.
func (l Line2D) Angle() float64 {
	return math.Atan2(l.P2.Y-l.P1.Y, l.P2.X-l.P1.X)
}

// Returns a normal vector with the same length as the line.
func (l Line2D) Normal() Vector2D {
	return Vector2D{l.P2.Y - l.P1.Y, l.P1.X - l.P2.X}
}
