package geometry

import (
	"math"
)

// Distance2DLinePointAngular returns the angle the line segment a would have
// to rotate about its midpoint to pass through point b.
func Distance2DLinePointAngular(a *Line2D, b *Vector2D) float64 {
	mpx, mpy := a.P.X+0.5*a.V.X, a.P.Y+0.5*a.V.Y
	adx, ady := a.P.X-mpx, a.P.Y-mpy
	ldx, ldy := b.X-mpx, b.Y-mpy
	return math.Abs(math.Acos((adx*ldx+ady*ldy)/
		math.Sqrt((adx*adx+ady*ady)*(ldx*ldx+ldy*ldy))) - math.Pi/2)
}

// Distance2DLinePointAngularCosSquared returns the cos of the squared angle
// the line segment a would have to rotate about its midpoint to pass through
// point b.
func Distance2DLinePointAngularCosSquared(a *Line2D, b *Vector2D) float64 {
	mpx, mpy := a.P.X+0.5*a.V.X, a.P.Y+0.5*a.V.Y
	adx, ady := a.P.X-mpx, a.P.Y-mpy
	ldx, ldy := b.X-mpx, b.Y-mpy
	dot := adx*ldx + ady*ldy
	return dot * dot / ((adx*adx + ady*ady) * (ldx*ldx + ldy*ldy))
}

// Distance2DLinePoint returns the distance point b is from line a.
func Distance2DLinePoint(a *Line2D, b *Vector2D) float64 {
	// http://local.wasp.uwa.edu.au/~pbourke/geometry/pointline/
	u := (a.V.X*(b.X-a.P.X) + a.V.Y*(b.Y-a.P.Y)) / (a.V.X*a.V.X + a.V.Y*a.V.Y)
	x, y := b.X-(a.P.X+a.V.X*u), b.Y-(a.P.Y+a.V.Y*u)
	return math.Sqrt(x*x + y*y)
}

// Distance2DLinePointSquared returns the squared distance point b is from line
// a.
func Distance2DLinePointSquared(a *Line2D, b *Vector2D) float64 {
	// http://local.wasp.uwa.edu.au/~pbourke/geometry/pointline/
	u := (a.V.X*(b.X-a.P.X) + a.V.Y*(b.Y-a.P.Y)) / (a.V.X*a.V.X + a.V.Y*a.V.Y)
	x, y := b.X-(a.P.X+a.V.X*u), b.Y-(a.P.Y+a.V.Y*u)
	return x*x + y*y
}

// Distance2DLineSegmentPoint returns the distance between line segment a and
// point b.
func Distance2DLineSegmentPoint(a *Line2D, b *Vector2D) float64 {
	// http://softsurfer.com/Archive/algorithm_0102/algorithm_0102.htm
	c1 := a.V.X*(b.X-a.P.X) + a.V.Y*(b.Y-a.P.Y)
	if c1 <= 0 {
		x, y := b.X-a.P.X, b.Y-a.P.Y
		return math.Sqrt(x*x + y*y)
	}
	c2 := a.V.X*a.V.X + a.V.Y*a.V.Y
	if c2 <= c1 {
		x, y := b.X-(a.P.X+a.V.X), b.Y-(a.P.Y+a.V.Y)
		return math.Sqrt(x*x + y*y)
	}
	c1 /= c2
	x, y := b.X-(a.P.X+a.V.X*c1), b.Y-(a.P.Y+a.V.Y*c1)
	return math.Sqrt(x*x + y*y)
}

// Distance2DLineSegmentPointSquared returns the squared distance between line
// segment a and point b.
func Distance2DLineSegmentPointSquared(a *Line2D, b *Vector2D) float64 {
	// http://softsurfer.com/Archive/algorithm_0102/algorithm_0102.htm
	c1 := a.V.X*(b.X-a.P.X) + a.V.Y*(b.Y-a.P.Y)
	if c1 <= 0 {
		x, y := b.X-a.P.X, b.Y-a.P.Y
		return x*x + y*y
	}
	c2 := a.V.X*a.V.X + a.V.Y*a.V.Y
	if c2 <= c1 {
		x, y := b.X-(a.P.X+a.V.X), b.Y-(a.P.Y+a.V.Y)
		return x*x + y*y
	}
	c1 /= c2
	x, y := b.X-(a.P.X+a.V.X*c1), b.Y-(a.P.Y+a.V.Y*c1)
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
