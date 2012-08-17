package geometry

import (
	"math"
)

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
