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

// Returns the result of vector addition.
func (v1 Vector3D) Plus(v2 Vector3D) Vector3D {
	return Vector3D{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

// Add a vector.
func (v1 *Vector3D) Add(v2 Vector3D) {
	v1.X += v2.X
	v1.Y += v2.Y
	v1.Z += v2.Z
}

// Returns the result of vector subtraction.
func (v1 Vector3D) Minus(v2 Vector3D) Vector3D {
	return Vector3D{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

// Subtract a vector.
func (v1 *Vector3D) Subtract(v2 Vector3D) {
	v1.X -= v2.X
	v1.Y -= v2.Y
	v1.Z -= v2.Z
}

// Returns the element by element multiplication.
func (v1 Vector3D) Times(v2 Vector3D) Vector3D {
	return Vector3D{v1.X * v2.X, v1.Y * v2.Y, v1.Z * v2.Z}
}

// Multiplies the vector element by element.
func (v1 *Vector3D) Multiply(v2 Vector3D) {
	v1.X *= v2.X
	v1.Y *= v2.X
	v1.Z *= v2.Z
}

// Returns the element by element division.
func (v1 Vector3D) Divided(v2 Vector3D) Vector3D {
	return Vector3D{v1.X / v2.X, v1.Y / v2.Y, v1.Z / v2.Z}
}

// Divides the vector element by element.
func (v1 *Vector3D) Divide(v2 Vector3D) {
	v1.X /= v2.X
	v1.Y /= v2.Y
	v1.Z /= v2.Z
}

// Returns the scaled vector.
func (v Vector3D) Scaled(s float64) Vector3D {
	return Vector3D{v.X * s, v.Y * s, v.Z * s}
}

// Scales the vector.
func (v *Vector3D) Scale(s float64) {
	v.X *= s
	v.Y *= s
	v.Z *= s
}
