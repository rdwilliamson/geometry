package geometry

import (
	"math"
	"testing"
)

func TestScalarProjectionOnto2D(t *testing.T) {
	v1 := Vector2D{2, 3}
	v2 := Vector2D{2, 1}
	if v1.ScalarProjectionOnto(v2) != 7.0/5.0 {
		t.Error("Vector2D.ScalarProjectionOnto")
	}
}

func TestVectorProjectionOnto2D(t *testing.T) {
	v1 := Vector2D{2, 3}
	v2 := Vector2D{2, 1}
	if !v1.VectorProjectionOnto(v2).Equal(Vector2D{2.8, 1.4}) {
		t.Error("Vector2D.VectorProjectionOnto")
	}
}

func TestVector2DCore(t *testing.T) {
	v := Vector2D{0, 0}
	v = v.Plus(Vector2D{1, 2})
	if !v.Equal(Vector2D{1, 2}) {
		t.Error("Vector2D.Plus")
	}
	v.Add(Vector2D{2, 1})
	if !v.Equal(Vector2D{3, 3}) {
		t.Error("Vector2D.Add")
	}
	v = v.Minus(Vector2D{1, 2})
	if !v.Equal(Vector2D{2, 1}) {
		t.Error("Vector2D.Minus")
	}
	v.Subtract(Vector2D{1, 0})
	if !v.Equal(Vector2D{1, 1}) {
		t.Error("Vector2D.Subtract")
	}
	v = v.Times(Vector2D{2, 2})
	if !v.Equal(Vector2D{2, 2}) {
		t.Error("Vector2D.Times")
	}
	v.Multiply(Vector2D{0.5, 0.5})
	if !v.Equal(Vector2D{1, 1}) {
		t.Error("Vector2D.Multiply")
	}
	v = v.Divided(Vector2D{0.5, 0.5})
	if !v.Equal(Vector2D{2, 2}) {
		t.Error("Vector2D.Divided")
	}
	v.Divide(Vector2D{2, 2})
	if !v.Equal(Vector2D{1, 1}) {
		t.Error("Vector2D.Divide")
	}
}

func TestVector2DScale(t *testing.T) {
	v := Vector2D{1, 1}
	v = v.Scaled(2)
	if !v.Equal(Vector2D{2, 2}) {
		t.Error("Vector2D.Scaled")
	}
	v.Scale(0.5)
	if !v.Equal(Vector2D{1, 1}) {
		t.Error("Vector2D.Scale")
	}
}

func TestVector2DAngle(t *testing.T) {
	v := Vector2D{1, 1}
	if v.Angle() != math.Pi/4 {
		t.Error("Vector2D.Angle")
	}
	v = Vector2D{-1, 1}
	if v.Angle() != 3*math.Pi/4 {
		t.Error("Vector2D.Angle")
	}
	v = Vector2D{-1, -1}
	if v.Angle() != -3*math.Pi/4 {
		t.Error("Vector2D.Angle")
	}
	v = Vector2D{1, -1}
	if v.Angle() != -math.Pi/4 {
		t.Error("Vector2D.Angle")
	}
	v = Vector2D{1, 0}
	if v.Angle() != 0 {
		t.Error("Vector2D.Angle")
	}
	v = Vector2D{0, 1}
	if v.Angle() != math.Pi/2 {
		t.Error("Vector2D.Angle")
	}
	v = Vector2D{-1, 0}
	if v.Angle() != math.Pi {
		t.Error("Vector2D.Angle")
	}
	v = Vector2D{0, -1}
	if v.Angle() != -math.Pi/2 {
		t.Error("Vector2D.Angle")
	}
}

func TestVector2DFuzzyEqual(t *testing.T) {
	v1 := Vector2D{1.0, 1.0}
	v2 := v1
	v2.X += 0.0000000000001
	if v1.Equal(v2) {
		t.Error("Vector2D.Equal")
	}
	if !v1.FuzzyEqual(v2) {
		t.Error("Vector2D.FuzzyEqual")
	}
	v2.Y += 0.000000000001
	if v1.Equal(v2) {
		t.Error("Vector2D.Equal")
	}
	if v1.FuzzyEqual(v2) {
		t.Error("Vector2D.FuzzyEqual")
	}
}

func TestVector2DLength(t *testing.T) {
	v := Vector2D{3, 4}
	if v.Length() != 5 {
		t.Error("Vector2D.Length")
	}
	if v.LengthSquared() != 25 {
		t.Error("Vector2D.LengthSquared")
	}
}

func TestVector2DNormalize(t *testing.T) {
	v := Vector2D{15, 20}
	v = v.Normalized()
	if !v.Equal(Vector2D{15.0 / 25.0, 20.0 / 25.0}) {
		t.Error("Vector2D.Normalized", v)
	}
	v = Vector2D{15, 20}
	v.Normalize()
	if !v.Equal(Vector2D{15.0 / 25.0, 20.0 / 25.0}) {
		t.Error("Vector2D.Normalize")
	}
	v = Vector2D{0, 0}
	if !v.Normalized().Equal(Vector2D{0, 0}) {
		t.Error("Vector2D.Normalized")
	}
	v.Normalize()
	if !v.Equal(Vector2D{0, 0}) {
		t.Error("Vector2D.Normalize")
	}
}

func TestDotProduct2D(t *testing.T) {
	if DotProduct2D(Vector2D{2, 4}, Vector2D{1, 5}) != 22 {
		t.Error("DotProduct2D")
	}
}
