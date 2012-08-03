package geometry

import "math"

// A Line2D representes a 2D line by two points P1 and P2 (represented by
// vectors) on the line. The line is treated as an infinite line unless a
// method explicitly says otherwise. If treated as a segment then P1 and P2 are
// the end points of the line segment.
type Line2D struct {
	P1, P2 Vector2D
}

// NewLine2D returns a new Line2D.
func NewLine2D(x1, y1, x2, y2 float64) *Line2D {
	return &Line2D{Vector2D{x1, y1}, Vector2D{x2, y2}}
}

// Should rays have P1 be the end point and P2 treated as a vector or P2 as a
// point on the ray? I never use rays so I'm not sure which is more convenient.

// Equal compares a and b and returns a boolean indicating if they are equal.
func (a *Line2D) Equal(b *Line2D) bool {
	ldx, ldy := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y
	d := 1.0 / (ldx*ldx + ldy*ldy)
	u := (ldx*(b.P1.X-a.P1.X) + ldy*(b.P1.Y-a.P1.Y)) * d
	x, y := b.P1.X-(a.P1.X+ldx*u), b.P1.Y-(a.P1.Y+ldy*u)
	if x*x+y*y != 0 {
		return false
	}
	return ldy/ldx == (b.P2.Y-b.P1.Y)/(b.P2.X-b.P1.X)
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
	dr := (ldy / ldx) - ((b.P2.Y - b.P1.Y) / (b.P2.X - b.P1.X))
	return dr*dr < 0.000000000001*0.000000000001
}

// Intersection sets z to the intersection of l1 and l2 and returns z.
func (a *Line2D) Intersection(b *Line2D, z *Vector2D) *Vector2D {
	// http://local.wasp.uwa.edu.au/~pbourke/geometry/lineline2d/
	l1dx, l1dy := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y
	l2dx, l2dy := b.P2.X-b.P1.X, b.P2.Y-b.P1.Y
	d := l2dy*l1dx - l2dx*l1dy
	if d == 0 {
		z.X = math.Inf(1)
		z.Y = math.Inf(1)
		return z
	}
	ua := (l2dx*a.P1.Y - b.P1.Y - l2dy*a.P1.X - b.P1.X) / d
	z.X = a.P1.X + ua*l1dx
	z.Y = a.P1.Y + ua*l1dy
	return z
}

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

// PointAngleDistance returns the amount the line l would have to rotate about
// its midpoint (as if it were a segment) to pass through point p.
func (a *Line2D) PointAngleDistance(b *Vector2D) float64 {
	mpx, mpy := (a.P1.X+a.P2.X)*0.5, (a.P1.Y+a.P2.Y)*0.5
	l1dx, l1dy := a.P1.X-mpx, a.P1.Y-mpy
	l2dx, l2dy := b.X-mpx, b.Y-mpy
	return math.Abs(math.Acos((l1dx*l2dx+l1dy*l2dy)/
		math.Sqrt((l1dx*l1dx+l1dy*l1dy)*(l2dx*l2dx+l2dy*l2dy))) - math.Pi/2)
}

// PointAngleCosSquaredDistance returns the cos of the amount the line l would
// have to rotate about its midpoint (as if it were a segment) to pass through
// point p.
func (a *Line2D) PointAngleCosSquaredDistance(b *Vector2D) float64 {
	mpx, mpy := (a.P1.X+a.P2.X)*0.5, (a.P1.Y+a.P2.Y)*0.5
	l1dx, l1dy := a.P1.X-mpx, a.P1.Y-mpy
	l2dx, l2dy := b.X-mpx, b.Y-mpy
	dot := l1dx*l2dx + l1dy*l2dy
	return dot * dot / ((l1dx*l1dx + l1dy*l1dy) * (l2dx*l2dx + l2dy*l2dy))
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
	return (a.P1 == b.P1 && a.P2 == b.P2) || (a.P1 == b.P2 && a.P2 == b.P1)
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
func (a *Line2D) SegmentIntersection(b *Line2D, z *Vector2D) bool {
	// http://local.wasp.uwa.edu.au/~pbourke/geometry/lineline2d/
	l1dx, l1dy := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y
	l2dx, l2dy := b.P2.X-b.P1.X, b.P2.Y-b.P1.Y
	d := l2dy*l1dx - l2dx*l1dy
	if d == 0 {
		z.X = math.Inf(1)
		z.Y = math.Inf(1)
		return false
	}
	d = 1 / d
	dx, dy := a.P1.X-b.P1.X, a.P1.Y-b.P1.Y
	ua := (l2dx*dy - l2dy*dx) * d
	ub := (l1dx*dy - l1dy*dx) * d
	z.X = a.P1.X + ua*l1dx
	z.Y = a.P1.Y + ua*l1dy
	return 0 <= ua && ua <= 1 && 0 <= ub && ub <= 1
}

// SegmentPointDistance returns the distance between line segment a and point
// b.
func (a *Line2D) SegmentPointDistance(b *Vector2D) float64 {
	// http://softsurfer.com/Archive/algorithm_0102/algorithm_0102.htm
	ldx, ldy := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y
	c1 := ldx*(b.X-a.P1.X) + ldy*(b.Y-a.P1.Y)
	if c1 <= 0 {
		x, y := b.X-a.P1.X, b.Y-a.P1.Y
		return math.Sqrt(x*x + y*y)
	}
	c2 := ldx*ldx + ldy*ldy
	if c2 <= c1 {
		x, y := b.X-a.P2.X, b.Y-a.P2.Y
		return math.Sqrt(x*x + y*y)
	}
	c1 /= c2
	x, y := b.X-(a.P1.X+ldx*c1), b.Y-(a.P1.Y+ldy*c1)
	return math.Sqrt(x*x + y*y)
}

// SegmentPointDistanceSquared returns the squared distance between line
// segment a and point b.
func (a *Line2D) SegmentPointDistanceSquared(b *Vector2D) float64 {
	// http://softsurfer.com/Archive/algorithm_0102/algorithm_0102.htm
	ldx, ldy := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y
	c1 := ldx*(b.X-a.P1.X) + ldy*(b.Y-a.P1.Y)
	if c1 <= 0 {
		x, y := b.X-a.P1.X, b.Y-a.P1.Y
		return x*x + y*y
	}
	c2 := ldx*ldx + ldy*ldy
	if c2 <= c1 {
		x, y := b.X-a.P2.X, b.Y-a.P2.Y
		return x*x + y*y
	}
	c1 /= c2
	x, y := b.X-(a.P1.X+ldx*c1), b.Y-(a.P1.Y+ldy*c1)
	return x*x + y*y
}

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
