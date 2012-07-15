package geometry

// 3D Vector.
type Vector3D struct {
	X, Y, Z float64
}

// Returns the dot product of two vectors.
func DotProduct3D(v1, v2 Vector3D) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

// Returns the cross product of two vectors.
func CrossProduct3D(v1, v2 Vector3D) Vector3D {
	return Vector3D{v1.Y*v2.Z - v1.Z*v2.Y, v1.Z*v2.X - v1.X*v2.Z, v1.X*v2.Y - v1.Y*v2.X}
}

// Returns true if the two vectors are the same.
func (v1 Vector3D) Equal(v2 Vector3D) bool {
	return v1.X == v2.X && v1.Y == v2.Y && v1.Z == v2.Z
}

// Returns true if the two vectors are close.
func (v1 Vector3D) FuzzyEqual(v2 Vector3D) bool {
	return FuzzyEqual(v1.X, v2.X) && FuzzyEqual(v1.Y, v2.Y) && FuzzyEqual(v1.Z, v2.Z)
}
