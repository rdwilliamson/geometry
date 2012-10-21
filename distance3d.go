package geometry

import (
	"math"
)

// Distance3DLinePointAngular returns the angle the line segment a would have
// to rotate about its midpoint to pass through point b.
func Distance3DLinePointAngular(a *Line3D, b *Vector3D) float64 {
	mpx, mpy, mpz := a.P.X+a.V.X*0.5, a.P.Y+a.V.Y*0.5, a.P.Z+a.V.Z*0.5
	adx, ady, adz := a.P.X-mpx, a.P.Y-mpy, a.P.Z-mpz
	ldx, ldy, ldz := b.X-mpx, b.Y-mpy, b.Z-mpz
	return math.Abs(math.Acos((adx*ldx+ady*ldy+adz*ldz)/
		math.Sqrt((adx*adx+ady*ady+adz*adz)*(ldx*ldx+ldy*ldy+ldz*ldz))) - math.Pi/2)
}

// Distance3DLinePointAngularCosSquared returns the cos of the squared angle
// the line segment a would have to rotate about its midpoint to pass through
// point b.
func Distance3DLinePointAngularCosSquared(a *Line3D, b *Vector3D) float64 {
	mpx, mpy, mpz := a.P.X+a.V.X*0.5, a.P.Y+a.V.Y*0.5, a.P.Z+a.V.Z*0.5
	adx, ady, adz := a.P.X-mpx, a.P.Y-mpy, a.P.Z-mpz
	ldx, ldy, ldz := b.X-mpx, b.Y-mpy, b.Z-mpz
	dot := adx*ldx + ady*ldy + adz*ldz
	return dot * dot / ((adx*adx + ady*ady + adz*adz) * (ldx*ldx + ldy*ldy + ldz*ldz))
}

// Distance3DLinePoint returns the distance point b is from line a.
func Distance3DLinePoint(a *Line3D, b *Vector3D) float64 {
	// http://local.wasp.uwa.edu.au/~pbourke/geometry/pointline/
	u := (a.V.X*(b.X-a.P.X) + a.V.Y*(b.Y-a.P.Y) + a.V.Z*(b.Z-a.P.Z)) / (a.V.X*a.V.X + a.V.Y*a.V.Y + a.V.Z*a.V.Z)
	x, y, z := b.X-(a.P.X+a.V.X*u), b.Y-(a.P.Y+a.V.Y*u), b.Z-(a.P.Z+a.V.Z*u)
	return math.Sqrt(x*x + y*y + z*z)
}

// Distance3DLinePointSquared returns the squared distance point b is from
// line a.
func Distance3DLinePointSquared(a *Line3D, b *Vector3D) float64 {
	// http://local.wasp.uwa.edu.au/~pbourke/geometry/pointline/
	u := (a.V.X*(b.X-a.P.X) + a.V.Y*(b.Y-a.P.Y) + a.V.Z*(b.Z-a.P.Z)) / (a.V.X*a.V.X + a.V.Y*a.V.Y + a.V.Z*a.V.Z)
	x, y, z := b.X-(a.P.X+a.V.X*u), b.Y-(a.P.Y+a.V.Y*u), b.Z-(a.P.Z+a.V.Z*u)
	return x*x + y*y + z*z
}

// Distance3DLineSegmentPoint returns the distance between line segment a and
// point b.
func Distance3DLineSegmentPoint(a *Line3D, b *Vector3D) float64 {
	b1dx, b1dy, b1dz := b.X-a.P.X, b.Y-a.P.Y, b.Z-a.P.Z
	u := (a.V.X*b1dx + a.V.Y*b1dy + a.V.Z*b1dz) / (a.V.X*a.V.X + a.V.Y*a.V.Y + a.V.Z*a.V.Z)
	var x, y, z float64
	if u < 0 {
		x, y, z = b1dx, b1dy, b1dz
	} else if u > 1 {
		x, y, z = b.X-(a.P.X+a.V.X), b.Y-(a.P.Y+a.V.Y), b.Z-(a.P.Z+a.V.Z)
	} else {
		x, y, z = b.X-(a.P.X+a.V.X*u), b.Y-(a.P.Y+a.V.Y*u), b.Z-(a.P.Z+a.V.Z*u)
	}
	return math.Sqrt(x*x + y*y + z*z)
}

// Distance3DLineSegmentPointSquared returns the squared distance between line
// segment a and point b.
func Distance3DLineSegmentPointSquared(a *Line3D, b *Vector3D) float64 {
	b1dx, b1dy, b1dz := b.X-a.P.X, b.Y-a.P.Y, b.Z-a.P.Z
	u := (a.V.X*b1dx + a.V.Y*b1dy + a.V.Z*b1dz) / (a.V.X*a.V.X + a.V.Y*a.V.Y + a.V.Z*a.V.Z)
	var x, y, z float64
	if u < 0 {
		x, y, z = b1dx, b1dy, b1dz
	} else if u > 1 {
		x, y, z = b.X-(a.P.X+a.V.X), b.Y-(a.P.Y+a.V.Y), b.Z-(a.P.Z+a.V.Z)
	} else {
		x, y, z = b.X-(a.P.X+a.V.X*u), b.Y-(a.P.Y+a.V.Y*u), b.Z-(a.P.Z+a.V.Z*u)
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
