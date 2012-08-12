package geometry

import (
	"math"
)

// A Plane represents a plane by ax + by + cz + d = 0. A, B, and C represent
// the normal direction only (could have any magnitude). D divided by the
// magnitude of the normal represents the distance of the plane from the
// origin and the sign, if positive, signifies the origin is in the half space
// determined by the direction of the normal or the other half space if
// negative.
type Plane struct {
	A, B, C, D float64
}

// NewPLane returns a new Plane.
func NewPlane(a, b, c, d float64) *Plane {
	return &Plane{a, b, c, d}
}

// Equal returns true if the two planes are exactly equal or false otherwise.
func (a *Plane) Equal(b *Plane) bool {
	// check normal and distance from origin direction
	s := a.A / b.A
	if a.B != s*b.B || a.C != s*b.C || s*a.D*b.D < 0 {
		return false
	}
	// check distance (squared) from origin
	return b.D*b.D*(a.A*a.A+a.B*a.B+a.C*a.C) ==
		a.D*a.D*(b.A*b.A+b.B*b.B+b.C*b.C)
}

// FuzzyEqual returns true if the two planes are very close or false otherwise.
func (a *Plane) FuzzyEqual(b *Plane) bool {
	s := a.A / b.A
	if s*a.D*b.D < 0 || !FuzzyEqual(a.B, s*b.B) || !FuzzyEqual(a.C, s*b.C) {
		return false
	}
	return FuzzyEqual(b.D*b.D*(a.A*a.A+a.B*a.B+a.C*a.C),
		a.D*a.D*(b.A*b.A+b.B*b.B+b.C*b.C))
}

// Nomrmal sets z to x's normal then returns z.
func (x *Plane) Normal(z *Vector3D) *Vector3D {
	z.X = x.A
	z.Y = x.B
	z.Z = x.C
	return z
}

// Normalize sets z to the Hessian normal form of x where the normal is a unit
// vector and D is the distance from the origin, then returns z.
func (x *Plane) Normalize(z *Plane) *Plane {
	s := 1 / math.Sqrt(x.A*x.A+x.B*x.B+x.C*x.C)
	z.A = x.A * s
	z.B = x.B * s
	z.C = x.C * s
	z.D = x.D * s
	return z
}

// NormalizedEqual returns true if the two planes (assumed to be in Hessian
// normal form) are exactly equal or false otherwise.
func (a *Plane) NormalizedEqual(b *Plane) bool {
	if a.D*b.D < 0 {
		return -a.A == b.A && -a.B == b.B && -a.C == b.C && -a.D == b.D
	}
	return a.A == b.A && a.B == b.B && a.C == b.C && a.D == b.D
}

// PointDistance returns the distance (may be negative) point b is from plane a
// assuming it is normalized.
func (a *Plane) NormalizedPointDistance(b *Vector3D) float64 {
	return a.A*b.X + a.B*b.Y + a.C*b.Z + a.D
}

// PointDistance returns the distance (may be negative) point b is from plane a.
func (a *Plane) PointDistance(b *Vector3D) float64 {
	return (a.A*b.X + a.B*b.Y + a.C*b.Z + a.D) /
		math.Sqrt(a.A*a.A+a.B*a.B+a.C*a.C)
}

// PointDistance returns the squared distance point b is from plane a.
func (a *Plane) PointDistanceSquared(b *Vector3D) float64 {
	n := a.A*b.X + a.B*b.Y + a.C*b.Z + a.D
	return (n * n) / (a.A*a.A + a.B*a.B + a.C*a.C)
}

// Set sets z to x then returns z.
func (z *Plane) Set(x *Plane) *Plane {
	z.A = x.A
	z.B = x.B
	z.C = x.C
	z.D = x.D
	return z
}
