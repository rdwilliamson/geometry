package geometry

import (
	"math"
	"testing"
)

func TestVector3DCore(t *testing.T) {
	v := Vector3D{0, 0, 0}
	v.Add(&Vector3D{3, 2, 1})
	if !v.Equal(&Vector3D{3, 2, 1}) {
		t.Error("Vector3D.Add")
	}
	v.Subtract(&Vector3D{2, 1, 0})
	if !v.Equal(&Vector3D{1, 1, 1}) {
		t.Error("Vector3D.Subtract")
	}
	v.Multiply(&Vector3D{0.5, 0.5, 0.5})
	if !v.Equal(&Vector3D{0.5, 0.5, 0.5}) {
		t.Error("Vector3D.Multiply")
	}
	v.Divide(&Vector3D{2, 2, 2})
	if !v.Equal(&Vector3D{0.25, 0.25, 0.25}) {
		t.Error("Vector3D.Divide")
	}
}

func TestDotProduct3D(t *testing.T) {
	if DotProduct3D(&Vector3D{1, 2, 3}, &Vector3D{4, 5, 6}) != 32 {
		t.Error("DotProduct3D")
	}
}

func TestCrossProduct3D(t *testing.T) {
	if v := CrossProduct3D(&Vector3D{1, 2, 3}, &Vector3D{4, 5, 6}); !v.Equal(&Vector3D{-3, 6, -3}) {
		t.Error("CrossProduct3D")
	}
}

func TestVector3DFuzzyEqual(t *testing.T) {
	v1 := &Vector3D{1.0, 1.0, 1.0}
	v2 := &Vector3D{1.0, 1.0, 1.0}
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

func TestVector3DScale(t *testing.T) {
	v := &Vector3D{1, 1, 1}
	v.Scale(0.5)
	if !v.Equal(&Vector3D{0.5, 0.5, 0.5}) {
		t.Error("Vector3D.Scale")
	}
}

func TestVector3DLength(t *testing.T) {
	v := Vector3D{-6, 3, -1}
	if v.Length() != math.Sqrt(46) {
		t.Error("Vector3D.Length")
	}
	if v.LengthSquared() != 46 {
		t.Error("Vector3D.LengthSquared")
	}
}

func TestVector3DNormalize(t *testing.T) {
	v := Vector3D{-6, 3, -1}
	v.Normalize()
	if !v.Equal(&Vector3D{-6 / math.Sqrt(46), 3 / math.Sqrt(46), -1 / math.Sqrt(46)}) {
		t.Error("Vector3D.Normalize")
	}
	v = Vector3D{0, 0, 0}
	v.Normalize()
	if !v.Equal(&Vector3D{0, 0, 0}) {
		t.Error("Vector3D.Normalize")
	}
}

func TestScalarProjectionOnto3D(t *testing.T) {
	v1 := &Vector3D{-3, 2, -4}
	v2 := &Vector3D{2, -5, 1}
	if v1.ScalarProjectionOnto(v2) != -20/math.Sqrt(30) {
		t.Error("Vector3D.ScalarProjectionOnto")
	}
}

func TestVectorProjectionOnto3D(t *testing.T) {
	v1 := &Vector3D{-3, 2, -4}
	v2 := &Vector3D{2, -5, 1}
	if v1.ProjectedOnto(v2); v1.FuzzyEqual(&Vector3D{-4.0 / 3.0, 10.0 / 3.0, -2.0 / 3.0}) {
		t.Error("Vector3D.VectorProjectionOnto")
	}
}

func TestAngleBetween3D(t *testing.T) {
	v1 := &Vector3D{1, 0, 0}
	v2 := &Vector3D{1, 1, 0}
	if !FuzzyEqual(v1.AngleBetween(v2), math.Pi/4) {
		t.Error("Vector3D.AngleBetween")
	}
	if !FuzzyEqual(v2.AngleBetween(v1), math.Pi/4) {
		t.Error("Vector3D.AngleBetween")
	}
}
