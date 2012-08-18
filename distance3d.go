package geometry

import (
	"math"
)

// Distance3DLinePointAngular returns the angle the line segment a would have
// to rotate about its midpoint to pass through point b.
func Distance3DLinePointAngular(a *Line3D, b *Vector3D) float64 {
	mpx, mpy, mpz := (a.P1.X+a.P2.X)*0.5, (a.P1.Y+a.P2.Y)*0.5,
		(a.P1.Z+a.P2.Z)*0.5
	adx, ady, adz := a.P1.X-mpx, a.P1.Y-mpy, a.P1.Z-mpz
	ldx, ldy, ldz := b.X-mpx, b.Y-mpy, b.Z-mpz
	return math.Abs(math.Acos((adx*ldx+ady*ldy+adz*ldz)/
		math.Sqrt((adx*adx+ady*ady+adz*adz)*(ldx*ldx+ldy*ldy+ldz*ldz))) -
		math.Pi/2)
}

// Distance3DLinePointAngularCosSquared returns the cos of the squared angle
// the line segment a would have to rotate about its midpoint to pass through
// point b.
func Distance3DLinePointAngularCosSquared(a *Line3D, b *Vector3D) float64 {
	mpx, mpy, mpz := (a.P1.X+a.P2.X)*0.5, (a.P1.Y+a.P2.Y)*0.5,
		(a.P1.Z+a.P2.Z)*0.5
	adx, ady, adz := a.P1.X-mpx, a.P1.Y-mpy, a.P1.Z-mpz
	ldx, ldy, ldz := b.X-mpx, b.Y-mpy, b.Z-mpz
	dot := adx*ldx + ady*ldy + adz*ldz
	return dot * dot / ((adx*adx + ady*ady + adz*adz) *
		(ldx*ldx + ldy*ldy + ldz*ldz))
}

// Distance3DLinePoint returns the distance point b is from line a.
func Distance3DLinePoint(a *Line3D, b *Vector3D) float64 {
	// http://local.wasp.uwa.edu.au/~pbourke/geometry/pointline/
	ldx, ldy, ldz := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y, a.P2.Z-a.P1.Z
	u := (ldx*(b.X-a.P1.X) + ldy*(b.Y-a.P1.Y) + ldz*(b.Z-a.P1.Z)) /
		(ldx*ldx + ldy*ldy + ldz*ldz)
	x, y, z := b.X-(a.P1.X+ldx*u), b.Y-(a.P1.Y+ldy*u), b.Z-(a.P1.Z+ldz*u)
	return math.Sqrt(x*x + y*y + z*z)
}

// Distance3DLinePointSquared returns the squared distance point b is from
// line a.
func Distance3DLinePointSquared(a *Line3D, b *Vector3D) float64 {
	// http://local.wasp.uwa.edu.au/~pbourke/geometry/pointline/
	ldx, ldy, ldz := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y, a.P2.Z-a.P1.Z
	u := (ldx*(b.X-a.P1.X) + ldy*(b.Y-a.P1.Y) + ldz*(b.Z-a.P1.Z)) /
		(ldx*ldx + ldy*ldy + ldz*ldz)
	x, y, z := b.X-(a.P1.X+ldx*u), b.Y-(a.P1.Y+ldy*u), b.Z-(a.P1.Z+ldz*u)
	return x*x + y*y + z*z
}

// Distance3DLineSegmentPoint returns the distance between line segment a and
// point b.
func Distance3DLineSegmentPoint(a *Line3D, b *Vector3D) float64 {
	adx, ady, adz := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y, a.P2.Z-a.P1.Z
	b1dx, b1dy, b1dz := b.X-a.P1.X, b.Y-a.P1.Y, b.Z-a.P1.Z
	u := (adx*b1dx + ady*b1dy + adz*b1dz) / (adx*adx + ady*ady + adz*adz)
	var x, y, z float64
	if u < 0 {
		x, y, z = b1dx, b1dy, b1dz
	} else if u > 1 {
		x, y, z = b.X-a.P2.X, b.Y-a.P2.Y, b.Z-a.P2.Z
	} else {
		x, y, z = b.X-(a.P1.X+adx*u), b.Y-(a.P1.Y+ady*u), b.Z-(a.P1.Z+adz*u)
	}
	return math.Sqrt(x*x + y*y + z*z)
}

// Distance3DLineSegmentPointSquared returns the squared distance between line
// segment a and point b.
func Distance3DLineSegmentPointSquared(a *Line3D, b *Vector3D) float64 {
	adx, ady, adz := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y, a.P2.Z-a.P1.Z
	b1dx, b1dy, b1dz := b.X-a.P1.X, b.Y-a.P1.Y, b.Z-a.P1.Z
	u := (adx*b1dx + ady*b1dy + adz*b1dz) / (adx*adx + ady*ady + adz*adz)
	var x, y, z float64
	if u < 0 {
		x, y, z = b1dx, b1dy, b1dz
	} else if u > 1 {
		x, y, z = b.X-a.P2.X, b.Y-a.P2.Y, b.Z-a.P2.Z
	} else {
		x, y, z = b.X-(a.P1.X+adx*u), b.Y-(a.P1.Y+ady*u), b.Z-(a.P1.Z+adz*u)
	}
	return x*x + y*y + z*z
}

// Distance3DPlaneNormalizedPoint returns the distance (may be negative) point
// b is from plane a assuming it is normalized.
func Distance3DPlaneNormalizedPoint(a *Plane, b *Vector3D) float64 {
	return a.A*b.X + a.B*b.Y + a.C*b.Z + a.D
}

// Distance3DPlanePoint returns the distance (may be negative) point b is from
// plane a.
func Distance3DPlanePoint(a *Plane, b *Vector3D) float64 {
	return (a.A*b.X + a.B*b.Y + a.C*b.Z + a.D) /
		math.Sqrt(a.A*a.A+a.B*a.B+a.C*a.C)
}

// Distance3DPlanePointSquared returns the squared distance point b is from
// plane a.
func Distance3DPlanePointSquared(a *Plane, b *Vector3D) float64 {
	n := a.A*b.X + a.B*b.Y + a.C*b.Z + a.D
	return (n * n) / (a.A*a.A + a.B*a.B + a.C*a.C)
}

// Distance3DPointPoint returns the distance between points a and b.
func Distance3DPointPoint(a, b *Vector3D) float64 {
	dx, dy, dz := b.X-a.X, b.Y-a.Y, b.Z-a.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

// Distance3DPointPointSquared returns the squared distance between points a
// and b.
func Distance3DPointPointSquared(a, b *Vector3D) float64 {
	dx, dy, dz := b.X-a.X, b.Y-a.Y, b.Z-a.Z
	return dx*dx + dy*dy + dz*dz
}

// Distance3DVectorVectorAngular returns the angle between a and b.
func Distance3DVectorVectorAngular(a, b *Vector3D) float64 {
	return math.Acos((a.X*b.X + a.Y*b.Y + a.Z*b.Z) /
		math.Sqrt((a.X*a.X+a.Y*a.Y+a.Z*a.Z)*(b.X*b.X+b.Y*b.Y+b.Z*b.Z)))
}

// Distance3DVectorVectorAngularCosSquared returns the cos of the squared angle
// between a and b.
func Distance3DVectorVectorAngularCosSquared(a, b *Vector3D) float64 {
	dot := (a.X*b.X + a.Y*b.Y + a.Z*b.Z)
	return dot * dot /
		((a.X*a.X + a.Y*a.Y + a.Z*a.Z) * (b.X*b.X + b.Y*b.Y + b.Z*b.Z))
}
