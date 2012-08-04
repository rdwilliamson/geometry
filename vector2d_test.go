package geometry

import (
	"math"
	"testing"
)

func TestNewVector2D(t *testing.T) {
	if !NewVector2D(1, 2).Equal(&Vector2D{1, 2}) {
		t.Error("NewVector2D")
	}
}

func Benchmark_Vector2D_New(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewVector2D(0, 0)
	}
}

func TestVector2DAdd(t *testing.T) {
	r, v1, v2 := &Vector2D{}, &Vector2D{1, 2}, &Vector2D{3, 4}
	if !r.Add(v1, v2).Equal(&Vector2D{4, 6}) {
		t.Error("Vector2D.Add")
	}
}

func Benchmark_Vector2D_Add(b *testing.B) {
	r, v1, v2 := &Vector2D{}, &Vector2D{1, 2}, &Vector2D{3, 4}
	for i := 0; i < b.N; i++ {
		r.Add(v1, v2)
	}
}

func TestVector2DAngularDifference(t *testing.T) {
	v1, v2 := &Vector2D{1, 0}, &Vector2D{0, 1}
	if v1.AngularDifference(v2) != math.Pi/2 {
		t.Error("Vector2D.AngularDifference")
	}
}

func Benchmark_Vector2D_AngularDifference(b *testing.B) {
	v1, v2 := &Vector2D{1, 2}, &Vector2D{3, 4}
	for i := 0; i < b.N; i++ {
		v1.AngularDifference(v2)
	}
}

func TestVector2DAngularCosSquaredDifference(t *testing.T) {
	v1, v2 := &Vector2D{1, 0}, &Vector2D{0, 1}
	if FuzzyEqual(v1.AngularCosSquaredDifference(v2), math.Cos(math.Pi/2)) {
		t.Error("Vector2D.AngularCosSquaredDifference")
	}
}

func Benchmark_Vector2D_AngularCosSquaredDifference(b *testing.B) {
	v1, v2 := &Vector2D{1, 2}, &Vector2D{3, 4}
	for i := 0; i < b.N; i++ {
		v1.AngularCosSquaredDifference(v2)
	}
}

func TestVector2DDirectionEqual(t *testing.T) {
	v1, v2 := &Vector2D{1, 1}, &Vector2D{2, 2}
	if !v1.DirectionEqual(v2) {
		t.Error("Vector2D.DirectionEqual")
	}
	if !v2.DirectionEqual(v1) {
		t.Error("Vector2D.DirectionEqual")
	}
}

func Benchmark_Vector2D_DirectionEqual(b *testing.B) {
	v1, v2 := &Vector2D{1, 1}, &Vector2D{2, 2}
	for i := 0; i < b.N; i++ {
		v1.DirectionEqual(v2)
	}
}

func TestVector2DDirectionFuzzyEqual(t *testing.T) {
	v1, v2 := &Vector2D{1, 1}, &Vector2D{2, 2.000000000001}
	if !v1.DirectionFuzzyEqual(v2) {
		t.Error("Vector2D.DirectionFuzzyEqual")
	}
	if !v2.DirectionFuzzyEqual(v1) {
		t.Error("Vector2D.DirectionFuzzyEqual")
	}
	v2.Y = 2.00000000001
	if v1.DirectionFuzzyEqual(v2) {
		t.Error("Vector2D.DirectionFuzzyEqual")
	}
	if v2.DirectionFuzzyEqual(v1) {
		t.Error("Vector2D.DirectionFuzzyEqual")
	}
}

func Benchmark_Vector2D_DirectionFuzzyEqual(b *testing.B) {
	v1, v2 := &Vector2D{1, 1}, &Vector2D{2, 2}
	for i := 0; i < b.N; i++ {
		v1.DirectionFuzzyEqual(v2)
	}
}

func TestVector2DDistance(t *testing.T) {
	v1, v2 := &Vector2D{}, &Vector2D{1, 0}
	if v1.Distance(v2) != 1 {
		t.Error("Vector2D.Distance")
	}
}

func Benchmark_Vector2D_Distance(b *testing.B) {
	v1, v2 := &Vector2D{}, &Vector2D{1, 0}
	for i := 0; i < b.N; i++ {
		v1.Distance(v2)
	}
}

func TestVector2DDistanceSquared(t *testing.T) {
	v1, v2 := &Vector2D{}, &Vector2D{1, 0}
	if v1.DistanceSquared(v2) != 1 {
		t.Error("Vector2D.DistanceSquared")
	}
}

func Benchmark_Vector2D_DistanceSquared(b *testing.B) {
	v1, v2 := &Vector2D{}, &Vector2D{1, 0}
	for i := 0; i < b.N; i++ {
		v1.DistanceSquared(v2)
	}
}

func TestVector2DDotProduct(t *testing.T) {
	v1, v2 := &Vector2D{1, 2}, &Vector2D{3, 4}
	if v1.DotProduct(v2) != 11 {
		t.Error("Vector2D.DotProduct")
	}
}

func Benchmark_Vector2D_DotProduct(b *testing.B) {
	v1, v2 := &Vector2D{1, 2}, &Vector2D{3, 4}
	for i := 0; i < b.N; i++ {
		v1.DotProduct(v2)
	}
}

func TestVector2DMagnitude(t *testing.T) {
	v := &Vector2D{3, 4}
	if v.Magnitude() != 5 {
		t.Error("Vector2D.Magnitude")
	}
}

func Benchmark_Vector2D_Magnitude(b *testing.B) {
	v := &Vector2D{3, 4}
	for i := 0; i < b.N; i++ {
		v.Magnitude()
	}
}

func TestVector2DMagnitudeSquared(t *testing.T) {
	v := &Vector2D{3, 4}
	if v.MagnitudeSquared() != 25 {
		t.Error("Vector2D.MagnitudeSquared")
	}
}

func Benchmark_Vector2D_MagnitudeSquared(b *testing.B) {
	v := &Vector2D{3, 4}
	for i := 0; i < b.N; i++ {
		v.MagnitudeSquared()
	}
}

func TestVector2DEqual(t *testing.T) {
	v1, v2 := &Vector2D{}, &Vector2D{1, 0}
	if v1.Equal(v2) {
		t.Error("Vector2D.Equal")
	}
	v2.X = 0
	if !v1.Equal(v2) {
		t.Error("Vector2D.Equal")
	}
}

func Benchmark_Vector2D_Equal(b *testing.B) {
	v1, v2 := &Vector2D{}, &Vector2D{1, 0}
	for i := 0; i < b.N; i++ {
		v1.Equal(v2)
	}
}

func TestVector2DFuzzyEqual(t *testing.T) {
	v1, v2 := &Vector2D{}, &Vector2D{0.000000000001, 0}
	if v1.FuzzyEqual(v2) {
		t.Error("Vector2D.FuzzyEqual")
	}
	v2.X = 0.0000000000001
	if !v1.FuzzyEqual(v2) {
		t.Error("Vector2D.FuzzyEqual")
	}
}

func Benchmark_Vector2D_FuzzyEqual(b *testing.B) {
	v1, v2 := &Vector2D{}, &Vector2D{1, 0}
	for i := 0; i < b.N; i++ {
		v1.FuzzyEqual(v2)
	}
}

func TestVector2DSet(t *testing.T) {
	v1, v2 := &Vector2D{}, &Vector2D{1, 0}
	if !v1.Set(v2).Equal(v2) {
		t.Error("Vector2D.Set")
	}
}

func Benchmark_Vector2D_Set(b *testing.B) {
	v1, v2 := &Vector2D{}, &Vector2D{1, 0}
	for i := 0; i < b.N; i++ {
		v1.Set(v2)
	}
}

func TestVector2DSubtract(t *testing.T) {
	r, v1, v2 := &Vector2D{}, &Vector2D{1, 2}, &Vector2D{3, 4}
	if !r.Subtract(v1, v2).Equal(&Vector2D{-2, -2}) {
		t.Error("Vector2D.Subtract")
	}
}

func Benchmark_Vector2D_Subtract(b *testing.B) {
	r, v1, v2 := &Vector2D{}, &Vector2D{1, 2}, &Vector2D{3, 4}
	for i := 0; i < b.N; i++ {
		r.Subtract(v1, v2)
	}
}
