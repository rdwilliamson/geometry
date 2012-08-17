package geometry

import (
	"math"
)

// Distance2DLinePointAngular returns the angle the line segment a would have
// to rotate about its midpoint to pass through point b.
func Distance2DLinePointAngular(a *Line2D, b *Vector2D) float64 {
	mpx, mpy := (a.P1.X+a.P2.X)*0.5, (a.P1.Y+a.P2.Y)*0.5
	l1dx, l1dy := a.P1.X-mpx, a.P1.Y-mpy
	l2dx, l2dy := b.X-mpx, b.Y-mpy
	return math.Abs(math.Acos((l1dx*l2dx+l1dy*l2dy)/
		math.Sqrt((l1dx*l1dx+l1dy*l1dy)*(l2dx*l2dx+l2dy*l2dy))) - math.Pi/2)
}

// Distance2DLinePointAngularCosSquared returns the cos of the squared angle
// the line segment a would have to rotate about its midpoint to pass through
// point b.
func Distance2DLinePointAngularCosSquared(a *Line2D, b *Vector2D) float64 {
	mpx, mpy := (a.P1.X+a.P2.X)*0.5, (a.P1.Y+a.P2.Y)*0.5
	l1dx, l1dy := a.P1.X-mpx, a.P1.Y-mpy
	l2dx, l2dy := b.X-mpx, b.Y-mpy
	dot := l1dx*l2dx + l1dy*l2dy
	return dot * dot / ((l1dx*l1dx + l1dy*l1dy) * (l2dx*l2dx + l2dy*l2dy))
}

// Distance2DLinePoint returns the distance point b is from line a.
func Distance2DLinePoint(a *Line2D, b *Vector2D) float64 {
	// http://local.wasp.uwa.edu.au/~pbourke/geometry/pointline/
	ldx, ldy := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y
	u := (ldx*(b.X-a.P1.X) + ldy*(b.Y-a.P1.Y)) / (ldx*ldx + ldy*ldy)
	x, y := b.X-(a.P1.X+ldx*u), b.Y-(a.P1.Y+ldy*u)
	return math.Sqrt(x*x + y*y)
}

// Distance2DLinePointSquared returns the squared distance point b is from line
// a.
func Distance2DLinePointSquared(a *Line2D, b *Vector2D) float64 {
	// http://local.wasp.uwa.edu.au/~pbourke/geometry/pointline/
	ldx, ldy := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y
	u := (ldx*(b.X-a.P1.X) + ldy*(b.Y-a.P1.Y)) / (ldx*ldx + ldy*ldy)
	x, y := b.X-(a.P1.X+ldx*u), b.Y-(a.P1.Y+ldy*u)
	return x*x + y*y
}

// Distance2DLineSegmentPoint returns the distance between line segment a and
// point b.
func Distance2DLineSegmentPoint(a *Line2D, b *Vector2D) float64 {
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

// Distance2DLineSegmentPointSquared returns the squared distance between line
// segment a and point b.
func Distance2DLineSegmentPointSquared(a *Line2D, b *Vector2D) float64 {
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

// Distance2DPointPoint returns the distance between points a and b.
func Distance2DPointPoint(a, b *Vector2D) float64 {
	dx, dy := b.X-a.X, b.Y-a.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// Distance2DPointPointSquared returns the squared distance between points a
// and b.
func Distance2DPointPointSquared(a, b *Vector2D) float64 {
	dx, dy := b.X-a.X, b.Y-a.Y
	return dx*dx + dy*dy
}

// Distance2DVectorVectorAngular returns the angle between a and b.
func Distance2DVectorVectorAngular(a, b *Vector2D) float64 {
	return math.Acos((a.X*b.X + a.Y*b.Y) /
		math.Sqrt((a.X*a.X+a.Y*a.Y)*(b.X*b.X+b.Y*b.Y)))
}

// Distance2DVectorVectorAngularCosSquared returns the cos of the squared angle
// between a and b.
func Distance2DVectorVectorAngularCosSquared(a, b *Vector2D) float64 {
	dot := (a.X*b.X + a.Y*b.Y)
	return dot * dot / ((a.X*a.X + a.Y*a.Y) * (b.X*b.X + b.Y*b.Y))
}
