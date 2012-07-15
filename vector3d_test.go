package geometry

import (
	"testing"
)

func TestDotProduct3D(t *testing.T) {
	if DotProduct3D(Vector3D{1, 2, 3}, Vector3D{4, 5, 6}) != 32 {
		t.Error("DotProduct3D")
	}
}

func TestCrossProduct3D(t *testing.T) {
	if !CrossProduct3D(Vector3D{1, 2, 3}, Vector3D{4, 5, 6}).Equal(Vector3D{-3, 6, -3}) {
		t.Error("CrossProduct3D")
	}
}

func TestVector3DFuzzyEqual(t *testing.T) {
	v1 := Vector3D{1.0, 1.0, 1.0}
	v2 := v1
	v2.X += 0.0000000000001
	if v1.Equal(v2) {
		t.Error("Vector3D.Equal")
	}
	if !v1.FuzzyEqual(v2) {
		t.Error("Vector3D.FuzzyEqual")
	}
	v2.Y += 0.000000000001
	if v1.Equal(v2) {
		t.Error("Vector3D.Equal")
	}
	if v1.FuzzyEqual(v2) {
		t.Error("Vector3D.FuzzyEqual")
	}
}

func TestVector3DCore(t *testing.T) {
	v := Vector3D{0, 0, 0}
	v = v.Plus(Vector3D{1, 2, 3})
	if !v.Equal(Vector3D{1, 2, 3}) {
		t.Error("Vector3D.Plus")
	}
	v.Add(Vector3D{3, 2, 1})
	if !v.Equal(Vector3D{4, 4, 4}) {
		t.Error("Vector3D.Add")
	}
	v = v.Minus(Vector3D{1, 2, 3})
	if !v.Equal(Vector3D{3, 2, 1}) {
		t.Error("Vector3D.Minus")
	}
	v.Subtract(Vector3D{2, 1, 0})
	if !v.Equal(Vector3D{1, 1, 1}) {
		t.Error("Vector3D.Subtract")
	}
	v = v.Times(Vector3D{2, 2, 2})
	if !v.Equal(Vector3D{2, 2, 2}) {
		t.Error("Vector3D.Times")
	}
	v.Multiply(Vector3D{0.5, 0.5, 0.5})
	if !v.Equal(Vector3D{1, 1, 1}) {
		t.Error("Vector3D.Multiply")
	}
	v = v.Divided(Vector3D{0.5, 0.5, 0.5})
	if !v.Equal(Vector3D{2, 2, 2}) {
		t.Error("Vector3D.Divided")
	}
	v.Divide(Vector3D{2, 2, 2})
	if !v.Equal(Vector3D{1, 1, 1}) {
		t.Error("Vector3D.Divide")
	}
}

func TestVector3DScale(t *testing.T) {
	v := Vector3D{1, 1, 1}
	v = v.Scaled(2)
	if !v.Equal(Vector3D{2, 2, 2}) {
		t.Error("Vector3D.Scaled")
	}
	v.Scale(0.5)
	if !v.Equal(Vector3D{1, 1, 1}) {
		t.Error("Vector3D.Scale")
	}
}
