package geometry

import (
	"math"
)

// LineSeg2D represents a 2D line segment by the two end points.
type LineSeg2D struct {
	P1, P2 Point2D
}

// Angle calculate the line's angle from P1 to P2. Returns radians in the
// interval [-pi/2 pi/2].
func (l LineSeg2D) Angle() float64 {
	dx, dy := l.P2.X-l.P1.X, l.P2.Y-l.P1.Y
	return math.Atan(dy / dx)
}

// How far the line would have to be rotated around its midpoint to pass
// through the point. Returns radians in the interval [0 pi/2].
func (l LineSeg2D) AngDistPt(p Point2D) float64 {
	rl := LineSeg2D{l.Midpoint(), p}
	a := math.Abs(rl.Angle() - l.Angle())
	if a > math.Pi*0.5 {
		a = math.Pi - a
	}
	return a
}

// Linear distance from the line (segment) to a point.
// From Dan Sunday,
// http://softsurfer.com/Archive/algorithm_0102/algorithm_0102.htm
func (l LineSeg2D) LinDistPt(p Point2D, segment bool) float64 {
	v := l.ToVector()
	w := p.Minus(l.P1)
	c1 := DotProduct2D(w, v)
	if segment && c1 <= 0 {
		return p.DistTo(l.P1)
	}
	c2 := DotProduct2D(v, v)
	if segment && c2 <= c1 {
		return p.DistTo(l.P2)
	}
	b := c1 / c2
	a := l.P1.Plus(v.Scaled(b))
	return p.DistTo(a)
}

// Dx returns the line's horizontal distance.
func (l LineSeg2D) Dx() float64 {
	return l.P2.X - l.P1.X
}

// Dy returns the line's vertical distance.
func (l LineSeg2D) Dy() float64 {
	return l.P2.Y - l.P1.Y
}

// Intersection calculates the intersection of two lines and whether or not the
// intersection occurred on both lines.
// From Graphics Gems III, Faster Line Segment Intersection.
func (l1 LineSeg2D) Intersection(l2 LineSeg2D) (Point2D, bool) {
	a := l1.P2.Minus(l1.P1)
	b := l2.P1.Minus(l2.P2)
	denominator := a.Y*b.X - a.X*b.Y
	if denominator == 0.0 {
		// TODO determine where on line at infinity they intersect
		return Point2D{math.Inf(1), math.Inf(1), 0}, false
	}
	denominator = 1.0 / denominator
	c := l1.P1.Minus(l2.P1)
	A := (b.Y*c.X - b.X*c.Y) * denominator
	intersection := l1.P1.Plus(a.Scaled(A))
	if A < 0.0 || A > 1.0 {
		return intersection, false
	}
	B := (a.X*c.Y - a.Y*c.X) * denominator
	if B < 0.0 || B > 1.0 {
		return intersection, false
	}
	return intersection, true
}

// Intersects determines if two line segments intersect.
// From Graphics Gems III, Faster Line Segment Intersection.
func (l1 LineSeg2D) Intersects(l2 LineSeg2D) bool {
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

// Length returns the length of the line.
func (l LineSeg2D) Length() float64 {
	dx := l.P2.X-l.P1.X
	dy := l.P2.Y-l.P1.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// LengthSq returns the squared length of the line.
func (l LineSeg2D) LengthSq() float64 {
	dx := l.P2.X-l.P1.X
	dy := l.P2.Y-l.P1.Y
	return dx*dx + dy*dy
}

// Midpoint returns the midpoint of the line.
func (l LineSeg2D) Midpoint() Point2D {
	return Point2D{(l.P1.X + l.P2.X) * 0.5, (l.P1.Y + l.P2.Y) * 0.5, 1}
}

// Rotated rotates a line around its midpoint in radians.
func (l LineSeg2D) Rotated(t float64) LineSeg2D {
	m := l.Midpoint()
	x1 := l.P1.X - m.X
	y1 := l.P1.Y - m.Y
	x2 := l.P2.X - m.X
	y2 := l.P2.Y - m.Y
	cos := math.Cos(t)
	sin := math.Sin(t)
	return LineSeg2D{
		Point2D{x1*cos - y1*sin + m.X, x1*sin + y1*cos + m.Y, 1},
		Point2D{x2*cos - y2*sin + m.X, x2*sin + y2*cos + m.Y, 1}}
}

// ToVector converts the line into a vector from P1 to P2.
func (l LineSeg2D) ToVector() Point2D {
	return Point2D{l.P2.X - l.P1.X, l.P2.Y - l.P1.Y, 1}
}
