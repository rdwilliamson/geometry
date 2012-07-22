package geometry

import (
	"math"
)

// A line represented by two points.
type Line2D struct {
	P1, P2 Point2D
}

// Converts the line to a vector from the first point to the second.
func (l *Line2D) ToVector() Vector2D {
	return Vector2D{l.P2.X - l.P1.X, l.P2.Y - l.P1.X}
}

// Returns the length of the line.
func (l *Line2D) Length() float64 {
	dx := l.P2.X - l.P1.X
	dy := l.P2.Y - l.P1.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// Returns the squared length of the line.
func (l *Line2D) LengthSquared() float64 {
	dx, dy := l.P2.X-l.P1.X, l.P2.Y-l.P1.Y
	return dx*dx + dy*dy
}

// Returns the midpoint of the line.
func (l *Line2D) Midpoint() Point2D {
	return Point2D{(l.P1.X + l.P2.X) * 0.5, (l.P1.Y + l.P2.Y) * 0.5}
}

// Returns a normal vector with the same length as the line.
func (l *Line2D) Normal() Vector2D {
	return Vector2D{l.P2.Y - l.P1.Y, l.P1.X - l.P2.X}
}

// Returns true if the lines are equal.
func (l1 *Line2D) Equal(l2 *Line2D) bool {
	return (l1.P1 == l2.P1 && l1.P2 == l2.P2) || (l1.P1 == l2.P2 && l1.P2 == l2.P1)
}

// Returns true if the line are very close.
func (l1 *Line2D) FuzzyEqual(l2 *Line2D) bool {
	dx1, dy1 := l1.P1.X-l2.P1.X, l1.P1.Y-l2.P1.Y
	dx2, dy2 := l1.P2.X-l2.P2.X, l1.P2.Y-l2.P2.Y
	if dx1*dx1+dy1*dy1 < 0.000000000001*0.000000000001 &&
		dx2*dx2+dy2*dy2 < 0.000000000001*0.000000000001 {
		return true
	}
	dx1, dy1 = l1.P1.X-l2.P2.X, l1.P1.Y-l2.P2.Y
	dx2, dy2 = l1.P2.X-l2.P1.X, l1.P2.Y-l2.P1.Y
	return dx1*dx1+dy1*dy1 < 0.000000000001*0.000000000001 &&
		dx2*dx2+dy2*dy2 < 0.000000000001*0.000000000001
}

// Return the distance between a point and a line segment. See
// http://softsurfer.com/Archive/algorithm_0102/algorithm_0102.htm
func (l *Line2D) SegmentPointDistance(p *Point2D) float64 {
	ldx, ldy := l.P2.X-l.P1.X, l.P2.Y-l.P1.Y
	c1 := ldx*(p.X-l.P1.X) + ldy*(p.Y-l.P1.Y)
	if c1 <= 0 {
		x, y := p.X-l.P1.X, p.Y-l.P1.Y
		return math.Sqrt(x*x + y*y)
	}
	c2 := ldx*ldx + ldy*ldy
	if c2 <= c1 {
		x, y := p.X-l.P2.X, p.Y-l.P2.Y
		return math.Sqrt(x*x + y*y)
	}
	c1 /= c2
	x, y := p.X-(l.P1.X+ldx*c1), p.Y-(l.P1.Y+ldy*c1)
	return math.Sqrt(x*x + y*y)
}

// Returns the distance between a point and a line. See
// http://softsurfer.com/Archive/algorithm_0102/algorithm_0102.htm
func (l *Line2D) PointDistance(p *Point2D) float64 {
	ldx, ldy := l.P2.X-l.P1.X, l.P2.Y-l.P1.Y
	u := (ldx*(p.X-l.P1.X) + ldy*(p.Y-l.P1.Y)) / (ldx*ldx + ldy*ldy)
	x, y := p.X-(l.P1.X+ldx*u), p.Y-(l.P1.Y+ldy*u)
	return math.Sqrt(x*x + y*y)
}

// Returns the intersection of two lines. See
// http://local.wasp.uwa.edu.au/~pbourke/geometry/lineline2d/
func (l1 *Line2D) Intersection(l2 *Line2D) Point2D {
	l1dx, l1dy := l1.P2.X-l1.P1.X, l1.P2.Y-l1.P1.Y
	l2dx, l2dy := l2.P2.X-l2.P1.X, l2.P2.Y-l2.P1.Y
	d := l2dy*l1dx - l2dx*l1dy
	if d == 0 {
		return Point2D{math.Inf(1), math.Inf(1)}
	}
	ua := (l2dx*l1.P1.Y - l2.P1.Y - l2dy*l1.P1.X - l2.P1.X) / d
	return Point2D{l1.P1.X + ua*l1dx, l1.P1.Y + ua*l1dy}
}

// Returns the intersection of two lines and if the intersection occurs
// between. See http://local.wasp.uwa.edu.au/~pbourke/geometry/lineline2d/
func (l1 *Line2D) SegmentIntersection(l2 *Line2D) (Point2D, bool) {
	l1dx, l1dy := l1.P2.X-l1.P1.X, l1.P2.Y-l1.P1.Y
	l2dx, l2dy := l2.P2.X-l2.P1.X, l2.P2.Y-l2.P1.Y
	d := l2dy*l1dx - l2dx*l1dy
	if d == 0 {
		return Point2D{math.Inf(1), math.Inf(1)}, false
	}
	d = 1 / d
	dx, dy := l1.P1.X-l2.P1.X, l1.P1.Y-l2.P1.Y
	ua := l2dx*dy - l2dy*dx
	ub := l1dx*dy - l1dy*dx
	ua *= d
	ub *= d
	var seg bool
	if 0 <= ua && ua <= 1 && 0 <= ub && ub <= 1 {
		seg = true
	}
	return Point2D{l1.P1.X + ua*l1dx, l1.P1.Y + ua*l1dy}, seg
}
