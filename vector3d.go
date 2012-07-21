package geometry

import (
	"math"
)

// 3D Vector.
type Vector3D struct {
	X, Y, Z float64
}

// Returns a copy of the vector.
func (v *Vector3D) Copy() Vector3D {
	return *v
}

// Add a vector.
func (v1 *Vector3D) Add(v2 *Vector3D) {
	v1.X += v2.X
	v1.Y += v2.Y
	v1.Z += v2.Z
}

// Subtract a vector.
func (v1 *Vector3D) Subtract(v2 *Vector3D) {
	v1.X -= v2.X
	v1.Y -= v2.Y
	v1.Z -= v2.Z
}

// Multiplies the vector element by element.
func (v1 *Vector3D) Multiply(v2 *Vector3D) {
	v1.X *= v2.X
	v1.Y *= v2.X
	v1.Z *= v2.Z
}

// Divides the vector element by element.
func (v1 *Vector3D) Divide(v2 *Vector3D) {
	v1.X /= v2.X
	v1.Y /= v2.Y
	v1.Z /= v2.Z
}

// Scales the vector.
func (v *Vector3D) Scale(s float64) {
	v.X *= s
	v.Y *= s
	v.Z *= s
}

// Returns the length of the vector.
func (v *Vector3D) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Returns the squared length of the vector.
func (v *Vector3D) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Returns the dot product of two vectors.
func DotProduct3D(v1, v2 *Vector3D) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

// Returns the cross product of two vectors.
func CrossProduct3D(v1, v2 *Vector3D) Vector3D {
	return Vector3D{v1.Y*v2.Z - v1.Z*v2.Y, v1.Z*v2.X - v1.X*v2.Z, v1.X*v2.Y - v1.Y*v2.X}
}

// Returns true if the two vectors are the same.
func (v1 *Vector3D) Equal(v2 *Vector3D) bool {
	return v1.X == v2.X && v1.Y == v2.Y && v1.Z == v2.Z
}

// Returns true if the two vectors are close.
func (v1 *Vector3D) FuzzyEqual(v2 *Vector3D) bool {
	dx, dy, dz := v2.X-v1.X, v2.Y-v1.Y, v2.Z-v1.Z
	return dx*dx+dy*dy+dz*dz < 0.000000000001*0.000000000001
}

// Normalizes the vector. The zero vector remains unchanged.
func (v *Vector3D) Normalize() {
	if v.X == 0 && v.Y == 0 && v.Z == 0 {
		return
	}
	l := 1 / math.Sqrt(v.X*v.X+v.Y*v.Y+v.Z*v.Z)
	v.X *= l
	v.Y *= l
	v.Z *= l
}

// Returns the scalar projection.
func (v1 *Vector3D) ScalarProjection(v2 *Vector3D) float64 {
	return (v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z) / math.Sqrt(v2.X*v2.X+v2.Y*v2.Y+v2.Z*v2.Z)
}

// Returns the vector projection.
func (v1 *Vector3D) ProjectedOnto(v2 *Vector3D) Vector3D {
	s := (v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z) / (v2.X*v2.X + v2.Y*v2.Y + v2.Z*v2.Z)
	return Vector3D{v2.X * s, v2.Y * s, v2.Z * s}
}

// Returns the angle between two vectors.
func (v1 *Vector3D) AngleBetween(v2 *Vector3D) float64 {
	dot := v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
	v1d := (v1.X*v1.X + v1.Y*v1.Y + v1.Z*v1.Z)
	v2d := (v2.X*v2.X + v2.Y*v2.Y + v2.Z*v2.Z)
	return math.Acos(dot / math.Sqrt(v1d*v2d))
}

// Returns the cos of the angle between two vectors.
func (v1 *Vector3D) CosAngleBetween(v2 *Vector3D) float64 {
	dot := v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
	v1d := (v1.X*v1.X + v1.Y*v1.Y + v1.Z*v1.Z)
	v2d := (v2.X*v2.X + v2.Y*v2.Y + v2.Z*v2.Z)
	return dot / math.Sqrt(v1d*v2d)
}
