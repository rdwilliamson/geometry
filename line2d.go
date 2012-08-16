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

// Copy sets z to x and returns z.
func (z *Line2D) Copy(x *Line2D) *Line2D {
	z.P1.X = x.P1.X
	z.P1.Y = x.P1.Y
	z.P2.X = x.P2.X
	z.P2.Y = x.P2.Y
	return z
}

// Equal compares a and b and returns true if they are exactly equal or false
// otherwise.
func (a *Line2D) Equal(b *Line2D) bool {
	am, bm := (a.P2.Y-a.P1.Y)/(a.P2.X-a.P1.X), (b.P2.Y-b.P1.Y)/(b.P2.X-b.P1.X)
	return am == bm && a.P1.Y-am*a.P1.X == b.P1.Y-bm*b.P1.X
}

// FuzzyEqual compares a and b and returns true of they are very close or false
// otherwise.
func (a *Line2D) FuzzyEqual(b *Line2D) bool {
	am, bm := (a.P2.Y-a.P1.Y)/(a.P2.X-a.P1.X), (b.P2.Y-b.P1.Y)/(b.P2.X-b.P1.X)
	return FuzzyEqual(am, bm) && FuzzyEqual(a.P1.Y-am*a.P1.X, b.P1.Y-bm*b.P1.X)
}

// Length returns the length of line segment x.
func (x *Line2D) Length() float64 {
	dx, dy := x.P2.X-x.P1.X, x.P2.Y-x.P1.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// LengthSquared returns the squared length of line segment x.
func (x *Line2D) LengthSquared() float64 {
	dx, dy := x.P2.X-x.P1.X, x.P2.Y-x.P1.Y
	return dx*dx + dy*dy
}

// Midpoint sets point z to the midpoint of line segment x, then returns z.
func (x *Line2D) Midpoint(z *Vector2D) *Vector2D {
	z.X = (x.P1.X + x.P2.X) * 0.5
	z.Y = (x.P1.Y + x.P2.Y) * 0.5
	return z
}

// PointAngularDistance returns the angle the line segment a would have to
// rotate about its midpoint to pass through point b.
func (a *Line2D) PointAngularDistance(b *Vector2D) float64 {
	mpx, mpy := (a.P1.X+a.P2.X)*0.5, (a.P1.Y+a.P2.Y)*0.5
	l1dx, l1dy := a.P1.X-mpx, a.P1.Y-mpy
	l2dx, l2dy := b.X-mpx, b.Y-mpy
	return math.Abs(math.Acos((l1dx*l2dx+l1dy*l2dy)/
		math.Sqrt((l1dx*l1dx+l1dy*l1dy)*(l2dx*l2dx+l2dy*l2dy))) - math.Pi/2)
}

// PointAngularDistanceCosSquared returns the cos of the squared angle the line
// segment a would have to rotate about its midpoint to pass through point b.
func (a *Line2D) PointAngularDistanceCosSquared(b *Vector2D) float64 {
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

// SegmentEqual compares a and b and returns true if the line segments are
// exactly equal and false otherwise.
func (a *Line2D) SegmentEqual(b *Line2D) bool {
	return (a.P1 == b.P1 && a.P2 == b.P2) || (a.P1 == b.P2 && a.P2 == b.P1)
}

// SegmentFuzzyEqual compares a and b as line segments and returns true if they
// are very close and false otherwise.
func (a *Line2D) SegmentFuzzyEqual(b *Line2D) bool {
	return (a.P1.FuzzyEqual(&b.P1) && a.P2.FuzzyEqual(&b.P2)) ||
		(a.P1.FuzzyEqual(&b.P2) && a.P2.FuzzyEqual(&b.P1))
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

// Slope returns the slope of x.
func (x *Line2D) Slope() float64 {
	return (x.P2.Y - x.P1.Y) / (x.P2.X - x.P1.X)
}
