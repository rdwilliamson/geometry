package geometry

import (
	"fmt"
	"math"
)

type Line2D struct {
	P1, P2 Point2D
}

func (l Line2D) Angle() float64 {
	dx, dy := l.DxDy()
	return math.Atan(dy / dx)
}

// Angular difference between the line l and a line with endpoints at line l's
// midpoint and point p.
func (l Line2D) AngularDistanceToPoint(p Point2D) float64 {
	rl := Line2D{l.Midpoint(), p}
	a := math.Fabs(rl.Angle() - l.Angle())
	if a > math.Pi * 0.5 {
		a = math.Pi - a
	}
	return a
}

// From Dan Sunday,
// http://softsurfer.com/Archive/algorithm_0102/algorithm_0102.htm
func (l Line2D) DistanceToPoint(p Point2D, segment bool) float64 {
	v := l.ToVector()
	w := p.Minus(l.P1)
	c1 := DotProduct2D(w, v)
	if segment && c1 <= 0 {
		return p.DistanceTo(l.P1)
	}
	c2 := DotProduct2D(v, v)
	if segment && c2 <= c1 {
		return p.DistanceTo(l.P2)
	}
	b := c1 / c2
	a := l.P1.Plus(v.Scaled(b))
	return p.DistanceTo(a)
}

func (l Line2D) Dx() float64 {
	return l.P2.X - l.P1.X
}

func (l Line2D) DxDy() (float64, float64) {
	return l.P2.X - l.P1.X, l.P2.Y - l.P1.Y
}

func (l Line2D) Dy() float64 {
	return l.P2.Y - l.P1.Y
}

// Distance of (either) end point of line l to a line with endpoints at line
// l's midpoint and point p.
func (l Line2D) EndpointDistanceToPoint(p Point2D) float64 {
	rl := Line2D{l.Midpoint(), p}
	return rl.DistanceToPoint(l.P1, false)
}

func (l Line2D) EndpointDistanceToPoint2(p Point2D) float64 {
	mp := l.Midpoint()
	t := Line2D{mp, p}.Angle() - l.Angle()
	a := Line2D{mp, l.P1}.Length()
	return a * math.Tan(t)
}

// Returns the intersection point and if said point occures on both lines.
// From Graphics Gems III, Faster Line Segment Intersection.
func (l1 Line2D) Intersection(l2 Line2D) (Point2D, bool) {
	a := l1.P2.Minus(l1.P1)
	b := l2.P1.Minus(l2.P2)
	denominator := a.Y*b.X - a.X*b.Y
	if denominator == 0.0 {
		return Point2D{math.Inf(1), math.Inf(1)}, false
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

// From Graphics Gems III, Faster Line Segment Intersection.
func (l1 Line2D) Intersects(l2 Line2D) bool {
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

func (l Line2D) Length() float64 {
	dx, dy := l.DxDy()
	return math.Sqrt(dx*dx + dy*dy)
}

func (l Line2D) Midpoint() Point2D {
	return Point2D{(l.P1.X + l.P2.X) * 0.5, (l.P1.Y + l.P2.Y) * 0.5}
}

func (l Line2D) String() string {
	return fmt.Sprintf("{%v %v}", l.P1, l.P2)
}

func (l Line2D) ToVector() Point2D {
	return Point2D{l.P2.X - l.P1.X, l.P2.Y - l.P1.Y}
}
