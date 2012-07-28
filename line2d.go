package geometry

import "math"

// A Line2D representes a 2D line by two points P1 and P2 (represented by
// vectors) on the line. The line is treated as an infinite line unless a
// method explicitly says otherwise. If treated as a segment then P1 and P2 are
// the end points of the line segment.
type Line2D struct {
	P1, P2 Vector2D
}

// Should rays have P1 be the end point and P2 treated as a vector or P2 as a
// point on the ray? I never use rays so I'm not sure which is more convenient.

// AngleDistance returns the amount the line l would have to rotate about its
// midpoint (as if it were a segment) to pass through point p.

// AngleCosDistance returns the cos of the amount the line l would have to
// rotate about its midpoint (as if it were a segment) to pass through point p.

// Equal compares a and b and returns a boolean indicating if they are equal.
func (a *Line2D) Equal(b *Line2D) bool {
	ldx, ldy := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y
	d := 1.0 / (ldx*ldx + ldy*ldy)
	u := (ldx*(b.P1.X-a.P1.X) + ldy*(b.P1.Y-a.P1.Y)) * d
	x, y := b.P1.X-(a.P1.X+ldx*u), b.P1.Y-(a.P1.Y+ldy*u)
	if x*x+y*y != 0 {
		return false
	}
	u = (ldx*(b.P2.X-a.P1.X) + ldy*(b.P2.Y-a.P1.Y)) * d
	x, y = b.P2.X-(a.P1.X+ldx*u), b.P2.Y-(a.P1.Y+ldy*u)
	return x*x+y*y == 0
}

// FuzzyEqual compares a and b and returns a boolean indicating if they are
// very close.
func (a *Line2D) FuzzyEqual(b *Line2D) bool {
	ldx, ldy := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y
	d := 1.0 / (ldx*ldx + ldy*ldy)
	u := (ldx*(b.P1.X-a.P1.X) + ldy*(b.P1.Y-a.P1.Y)) * d
	x, y := b.P1.X-(a.P1.X+ldx*u), b.P1.Y-(a.P1.Y+ldy*u)
	if x*x+y*y >= 0.000000000001*0.000000000001 {
		return false
	}
	u = (ldx*(b.P2.X-a.P1.X) + ldy*(b.P2.Y-a.P1.Y)) * d
	x, y = b.P2.X-(a.P1.X+ldx*u), b.P2.Y-(a.P1.Y+ldy*u)
	return x*x+y*y < 0.000000000001*0.000000000001
}

// Intersection sets z to the intersection of l1 and l2 and returns z.

// Length returns the length of l as if is a line segment.
func (a *Line2D) Length() float64 {
	dx, dy := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// LengthSquared returns the length squared of l as if is a line segment.
func (a *Line2D) LengthSquared() float64 {
	dx, dy := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y
	return dx*dx + dy*dy
}

// Midpoint sets z to the segment l's midpoint and returns z.
func (x *Line2D) Midpoint(z *Vector2D) *Vector2D {
	z.X = (x.P1.X + x.P2.X) * 0.5
	z.Y = (x.P1.Y + x.P2.Y) * 0.5
	return z
}

// Normal sets z to a vector perpendicular to a and with a magnitude equal to
// the length of a as if it was a segement then returns z.
func (a *Line2D) Normal(z *Vector2D) *Vector2D {
	z.X = a.P2.Y - a.P1.Y
	z.Y = a.P1.X - a.P2.X
	return z
}

// PointDistance returns the distance point b is from line a.
func (a *Line2D) PointDistance(b *Vector2D) float64 {
	// http://local.wasp.uwa.edu.au/~pbourke/geometry/pointline/
	ldx, ldy := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y
	u := (ldx*(b.X-a.P1.X) + ldy*(b.Y-a.P1.Y)) / (ldx*ldx + ldy*ldy)
	x, y := b.X-(a.P1.X+ldx*u), b.Y-(a.P1.Y+ldy*u)
	return math.Sqrt(x*x + y*y)
}

// PointDistanceSquared returns the squared distance point b is from line a.
func (a *Line2D) PointDistanceSquared(b *Vector2D) float64 {
	// http://local.wasp.uwa.edu.au/~pbourke/geometry/pointline/
	ldx, ldy := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y
	u := (ldx*(b.X-a.P1.X) + ldy*(b.Y-a.P1.Y)) / (ldx*ldx + ldy*ldy)
	x, y := b.X-(a.P1.X+ldx*u), b.Y-(a.P1.Y+ldy*u)
	return x*x + y*y
}

// SegmentEqual compares a and b as line segments and returns a boolean
// indicating if they are equal.
func (a *Line2D) SegmentEqual(b *Line2D) bool {
	return *a == *b
}

// SegmentFuzzyEqual compares a and b as line segments and returns a boolean
// indicating if they are very close.
func (a *Line2D) SegmentFuzzyEqual(b *Line2D) bool {
	dx, dy := b.P1.X-a.P1.X, b.P1.Y-a.P1.Y
	if dx*dx+dy*dy >= 0.000000000001*0.000000000001 {
		return false
	}
	dx, dy = b.P2.X-a.P2.X, b.P2.Y-a.P2.Y
	return dx*dx+dy*dy < 0.000000000001*0.000000000001
}

// SegmentIntersection sets z to the intersection of l1 and l2 and returns a
// boolean indicating if the intersection occured on l1 and l2 as if they were
// segments.
// SegmentPointDistance
// SegmentPointDistanceSquared

// Set sets z to x and returns z.
func (z *Line2D) Set(x *Line2D) *Line2D {
	z.P1.X = x.P1.X
	z.P1.Y = x.P1.Y
	z.P2.X = x.P2.X
	z.P2.Y = x.P2.Y
	return z
}

// ToVector sets z to the vector from l.P1 to l.P2 and returns z.
func (x *Line2D) ToVector(z *Vector2D) *Vector2D {
	z.X = x.P2.X - x.P1.X
	z.Y = x.P2.Y - x.P1.Y
	return z
}

///////////////////////////////////////////////////////////////////////////////
// OLD

// // Returns the length of the line.
// func (l *Line2D) Length() float64 {
// 	dx := l.P2.X - l.P1.X
// 	dy := l.P2.Y - l.P1.Y
// 	return math.Sqrt(dx*dx + dy*dy)
// }

// // Returns the squared length of the line.
// func (l *Line2D) LengthSquared() float64 {
// 	dx, dy := l.P2.X-l.P1.X, l.P2.Y-l.P1.Y
// 	return dx*dx + dy*dy
// }

// // Returns true if the lines are equal.
// func (l1 *Line2D) Equal(l2 *Line2D) bool {
// 	return (l1.P1 == l2.P1 && l1.P2 == l2.P2) || (l1.P1 == l2.P2 && l1.P2 == l2.P1)
// }

// // Returns true if the line are very close.
// func (l1 *Line2D) FuzzyEqual(l2 *Line2D) bool {
// 	dx1, dy1 := l1.P1.X-l2.P1.X, l1.P1.Y-l2.P1.Y
// 	dx2, dy2 := l1.P2.X-l2.P2.X, l1.P2.Y-l2.P2.Y
// 	if dx1*dx1+dy1*dy1 < 0.000000000001*0.000000000001 &&
// 		dx2*dx2+dy2*dy2 < 0.000000000001*0.000000000001 {
// 		return true
// 	}
// 	dx1, dy1 = l1.P1.X-l2.P2.X, l1.P1.Y-l2.P2.Y
// 	dx2, dy2 = l1.P2.X-l2.P1.X, l1.P2.Y-l2.P1.Y
// 	return dx1*dx1+dy1*dy1 < 0.000000000001*0.000000000001 &&
// 		dx2*dx2+dy2*dy2 < 0.000000000001*0.000000000001
// }

// // Return the distance between a point and a line segment.
// func (l *Line2D) SegmentPointDistance(p *Point2D) float64 {
// 	// http://softsurfer.com/Archive/algorithm_0102/algorithm_0102.htm
// 	ldx, ldy := l.P2.X-l.P1.X, l.P2.Y-l.P1.Y
// 	c1 := ldx*(p.X-l.P1.X) + ldy*(p.Y-l.P1.Y)
// 	if c1 <= 0 {
// 		x, y := p.X-l.P1.X, p.Y-l.P1.Y
// 		return math.Sqrt(x*x + y*y)
// 	}
// 	c2 := ldx*ldx + ldy*ldy
// 	if c2 <= c1 {
// 		x, y := p.X-l.P2.X, p.Y-l.P2.Y
// 		return math.Sqrt(x*x + y*y)
// 	}
// 	c1 /= c2
// 	x, y := p.X-(l.P1.X+ldx*c1), p.Y-(l.P1.Y+ldy*c1)
// 	return math.Sqrt(x*x + y*y)
// }

// // Returns the distance between a point and a line.
// func (l *Line2D) PointDistance(p *Point2D) float64 {
// 	// http://local.wasp.uwa.edu.au/~pbourke/geometry/pointline/
// 	ldx, ldy := l.P2.X-l.P1.X, l.P2.Y-l.P1.Y
// 	u := (ldx*(p.X-l.P1.X) + ldy*(p.Y-l.P1.Y)) / (ldx*ldx + ldy*ldy)
// 	x, y := p.X-(l.P1.X+ldx*u), p.Y-(l.P1.Y+ldy*u)
// 	return math.Sqrt(x*x + y*y)
// }

// // Returns the squared distance between a point and a line.
// func (l *Line2D) PointSquaredDistance(p *Point2D) float64 {
// 	// http://local.wasp.uwa.edu.au/~pbourke/geometry/pointline/
// 	ldx, ldy := l.P2.X-l.P1.X, l.P2.Y-l.P1.Y
// 	u := (ldx*(p.X-l.P1.X) + ldy*(p.Y-l.P1.Y)) / (ldx*ldx + ldy*ldy)
// 	x, y := p.X-(l.P1.X+ldx*u), p.Y-(l.P1.Y+ldy*u)
// 	return x*x + y*y
// }

// // Returns the intersection of two lines.
// func (l1 *Line2D) Intersection(l2 *Line2D) Point2D {
// 	// http://local.wasp.uwa.edu.au/~pbourke/geometry/lineline2d/
// 	l1dx, l1dy := l1.P2.X-l1.P1.X, l1.P2.Y-l1.P1.Y
// 	l2dx, l2dy := l2.P2.X-l2.P1.X, l2.P2.Y-l2.P1.Y
// 	d := l2dy*l1dx - l2dx*l1dy
// 	if d == 0 {
// 		return Point2D{math.Inf(1), math.Inf(1)}
// 	}
// 	ua := (l2dx*l1.P1.Y - l2.P1.Y - l2dy*l1.P1.X - l2.P1.X) / d
// 	return Point2D{l1.P1.X + ua*l1dx, l1.P1.Y + ua*l1dy}
// }

// // Returns the intersection of two lines and if the intersection occurs between.
// func (l1 *Line2D) SegmentIntersection(l2 *Line2D) (Point2D, bool) {
// 	// http://local.wasp.uwa.edu.au/~pbourke/geometry/lineline2d/
// 	l1dx, l1dy := l1.P2.X-l1.P1.X, l1.P2.Y-l1.P1.Y
// 	l2dx, l2dy := l2.P2.X-l2.P1.X, l2.P2.Y-l2.P1.Y
// 	d := l2dy*l1dx - l2dx*l1dy
// 	if d == 0 {
// 		return Point2D{math.Inf(1), math.Inf(1)}, false
// 	}
// 	d = 1 / d
// 	dx, dy := l1.P1.X-l2.P1.X, l1.P1.Y-l2.P1.Y
// 	ua := l2dx*dy - l2dy*dx
// 	ub := l1dx*dy - l1dy*dx
// 	ua *= d
// 	ub *= d
// 	var seg bool
// 	if 0 <= ua && ua <= 1 && 0 <= ub && ub <= 1 {
// 		seg = true
// 	}
// 	return Point2D{l1.P1.X + ua*l1dx, l1.P1.Y + ua*l1dy}, seg
// }
