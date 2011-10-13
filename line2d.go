package geometry

import (
	"fmt"
	"math"
)

type Line2D struct {
	P1, P2 Point2D
}

func (l Line2D) Angle() float64 {
	dx, dy := l.P2.X-l.P1.X, l.P2.Y-l.P1.Y
	return math.Atan(dy / dx)
}

// Angular difference between the line l and a line with endpoints at line l's
// midpoint and point p.
func (l Line2D) AngDistPt(p Point2D) float64 {
	rl := Line2D{l.Midpoint(), p}
	a := math.Fabs(rl.Angle() - l.Angle())
	if a > math.Pi*0.5 {
		a = math.Pi - a
	}
	return a
}

// From Dan Sunday,
// http://softsurfer.com/Archive/algorithm_0102/algorithm_0102.htm
func (l Line2D) LinDistPt(p Point2D, segment bool) float64 {
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

func (l Line2D) Dx() float64 {
	return l.P2.X - l.P1.X
}

func (l Line2D) DxDy() (float64, float64) {
	return l.P2.X - l.P1.X, l.P2.Y - l.P1.Y
}

func (l Line2D) Dy() float64 {
	return l.P2.Y - l.P1.Y
}

// Distance that each endpoint moves when l is rotated around it's midpoint so
// that it passes through p.
func (l Line2D) EndPtDistSqPt(p Point2D) float64 {
	rl := Line2D{l.Midpoint(), p}
	s := math.Sqrt(l.LengthSquared() * 0.25 / rl.LengthSquared())
	rp := Point2D{rl.P1.X + s*rl.Dx(), rl.P1.Y + s*rl.Dy()}
	d := rp.DistToSq(l.P1)
	if td := rp.DistToSq(l.P2); td < d {
		d = td
	}
	return d
}

// Returns the intersection point and if said point occurs on both lines.
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
	dx, dy := l.P2.X-l.P1.X, l.P2.Y-l.P1.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func (l Line2D) LengthSquared() float64 {
	dx, dy := l.P2.X-l.P1.X, l.P2.Y-l.P1.Y
	return dx*dx + dy*dy
}

func (l Line2D) Midpoint() Point2D {
	return Point2D{(l.P1.X + l.P2.X) * 0.5, (l.P1.Y + l.P2.Y) * 0.5}
}

// Rotates a line around its midpoint
func (l Line2D) Rotated(t float64) Line2D {
	m := Point2D{(l.P1.X + l.P2.X) * 0.5, (l.P1.Y + l.P2.Y) * 0.5}
	x1 := l.P1.X - m.X
	y1 := l.P1.Y - m.Y
	x2 := l.P2.X - m.X
	y2 := l.P2.Y - m.Y
	cos := math.Cos(t)
	sin := math.Sin(t)
	return Line2D{
		Point2D{x1*cos - y1*sin + m.X, x1*sin + y1*cos + m.Y},
		Point2D{x2*cos - y2*sin + m.X, x2*sin + y2*cos + m.Y}}
}

func (l Line2D) String() string {
	return fmt.Sprintf("{%v %v}", l.P1, l.P2)
}

func (l Line2D) ToVector() Point2D {
	return Point2D{l.P2.X - l.P1.X, l.P2.Y - l.P1.Y}
}
