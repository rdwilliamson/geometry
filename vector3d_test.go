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
