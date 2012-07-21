package geometry

import (
	"math"
	"testing"
)

func TestVector2DCore(t *testing.T) {
	v := Vector2D{0, 0}
	v.Add(&Vector2D{2, 1})
	if !v.Equal(&Vector2D{2, 1}) {
		t.Error("Vector2D.Add")
	}
	v.Subtract(&Vector2D{1, 0})
	if !v.Equal(&Vector2D{1, 1}) {
		t.Error("Vector2D.Subtract")
	}
	v.Multiply(&Vector2D{0.5, 0.5})
	if !v.Equal(&Vector2D{0.5, 0.5}) {
		t.Error("Vector2D.Multiply")
	}
	v.Divide(&Vector2D{2, 2})
	if !v.Equal(&Vector2D{0.25, 0.25}) {
		t.Error("Vector2D.Divide")
	}
}

func Benchmark_Vector2D_Add(b *testing.B) {
	v1, v2 := &Vector2D{1, 1}, &Vector2D{2, 1}
	for i := 0; i < b.N; i++ {
		v1.Add(v2)
	}
}

func Benchmark_Vector2D_Subtract(b *testing.B) {
	v1, v2 := &Vector2D{1, 1}, &Vector2D{2, 1}
	for i := 0; i < b.N; i++ {
		v1.Subtract(v2)
	}
}

func Benchmark_Vector2D_Multiply(b *testing.B) {
	v1, v2 := &Vector2D{1, 1}, &Vector2D{2, 1}
	for i := 0; i < b.N; i++ {
		v1.Multiply(v2)
	}
}

func Benchmark_Vector2D_Divide(b *testing.B) {
	v1, v2 := &Vector2D{1, 1}, &Vector2D{2, 1}
	for i := 0; i < b.N; i++ {
		v1.Divide(v2)
	}
}

func TestVector2DScale(t *testing.T) {
	v := Vector2D{2, 2}
	v.Scale(0.5)
	if !v.Equal(&Vector2D{1, 1}) {
		t.Error("Vector2D.Scale")
	}
}

func Benchmark_Vector2D_Scale(b *testing.B) {
	v := &Vector2D{1, 1}
	for i := 0; i < b.N; i++ {
		v.Scale(1.5)
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

func Benchmark_Vector2D_Length(b *testing.B) {
	v := Vector2D{1, 2}
	for i := 0; i < b.N; i++ {
		v.Length()
	}
}

func Benchmark_Vector2D_LengthSquared(b *testing.B) {
	v := Vector2D{1, 2}
	for i := 0; i < b.N; i++ {
		v.LengthSquared()
	}
}

func TestScalarProjectionOnto2D(t *testing.T) {
	v1 := &Vector2D{2, 3}
	v2 := &Vector2D{2, 1}
	if v1.ScalarProjection(v2) != 7.0/5.0 {
		t.Error("Vector2D.ScalarProjection")
	}
}

func Benchmark_Vector2D_ScalarProjection(b *testing.B) {
	v1 := &Vector2D{2, 3}
	v2 := &Vector2D{2, 1}
	for i := 0; i < b.N; i++ {
		v1.ScalarProjection(v2)
	}
}

func TestVectorProjectionOnto2D(t *testing.T) {
	v1 := &Vector2D{2, 3}
	v2 := &Vector2D{2, 1}
	v1.ProjectedOnto(v2)
	if !v1.Equal(&Vector2D{2.8, 1.4}) {
		t.Error("Vector2D.ProjectedOnto")
	}
}

func Benchmark_Vector2D_ProjectedOnto(b *testing.B) {
	v1 := &Vector2D{2, 3}
	v2 := &Vector2D{2, 1}
	for i := 0; i < b.N; i++ {
		v1.ProjectedOnto(v2)
	}
}

func TestVector2DFuzzyEqual(t *testing.T) {
	v1 := &Vector2D{1.0, 1.0}
	v2 := &Vector2D{1.0, 1.0}
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

func Benchmark_Vector2D_Equal_1(b *testing.B) {
	v1 := &Vector2D{2, 3}
	v2 := &Vector2D{2, 1}
	for i := 0; i < b.N; i++ {
		v1.Equal(v2)
	}
}

func Benchmark_Vector2D_Equal_2(b *testing.B) {
	v1 := &Vector2D{2, 3}
	v2 := &Vector2D{4, 1}
	for i := 0; i < b.N; i++ {
		v1.Equal(v2)
	}
}

func Benchmark_Vector2D_FuzzyEqual(b *testing.B) {
	v1 := &Vector2D{2, 3}
	v2 := &Vector2D{4, 1}
	for i := 0; i < b.N; i++ {
		v1.FuzzyEqual(v2)
	}
}

func TestVector2DNormalize(t *testing.T) {
	v := Vector2D{15, 20}
	if v.Normalize(); !v.Equal(&Vector2D{15.0 / 25.0, 20.0 / 25.0}) {
		t.Error("Vector2D.Normalize")
	}
	v = Vector2D{0, 0}
	if v.Normalize(); !v.Equal(&Vector2D{0, 0}) {
		t.Error("Vector2D.Normalize")
	}
}

func Benchmark_Vector2D_Normalize(b *testing.B) {
	v := &Vector2D{2, 3}
	for i := 0; i < b.N; i++ {
		v.Normalize()
	}
}

func TestDotProduct2D(t *testing.T) {
	if DotProduct2D(&Vector2D{2, 4}, &Vector2D{1, 5}) != 22 {
		t.Error("DotProduct2D")
	}
}

func Benchmark_Vector2D_DotProduct(b *testing.B) {
	v1 := &Vector2D{2, 3}
	v2 := &Vector2D{4, 1}
	for i := 0; i < b.N; i++ {
		DotProduct2D(v1, v2)
	}
}

func TestAngleBetween2D(t *testing.T) {
	v1 := &Vector2D{1, 0}
	v2 := &Vector2D{1, 1}
	if !FuzzyEqual(v1.AngleBetween(v2), math.Pi/4) {
		t.Error("Vector2D.AngleBetween")
	}
	if !FuzzyEqual(v1.CosAngleBetween(v2), math.Cos(math.Pi/4)) {
		t.Error("Vector2D.AngleBetween")
	}
}

func Benchmark_Vector2D_AngleBetween(b *testing.B) {
	v1 := &Vector2D{2, 3}
	v2 := &Vector2D{4, 1}
	for i := 0; i < b.N; i++ {
		v1.AngleBetween(v2)
	}
}

func Benchmark_Vector2D_CosAngleBetween(b *testing.B) {
	v1 := &Vector2D{2, 3}
	v2 := &Vector2D{4, 1}
	for i := 0; i < b.N; i++ {
		v1.CosAngleBetween(v2)
	}
}
