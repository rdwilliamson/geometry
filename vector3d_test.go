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

func Benchmark_Vector3D_Add(b *testing.B) {
	v1 := &Vector3D{0, 0, 0}
	v2 := &Vector3D{1, 2, 3}
	for i := 0; i < b.N; i++ {
		v1.Add(v2)
	}
}

func Benchmark_Vector3D_Subtract(b *testing.B) {
	v1 := &Vector3D{0, 0, 0}
	v2 := &Vector3D{1, 2, 3}
	for i := 0; i < b.N; i++ {
		v1.Subtract(v2)
	}
}

func Benchmark_Vector3D_Multiply(b *testing.B) {
	v1 := &Vector3D{3, 4, 5}
	v2 := &Vector3D{6, 2, 3}
	for i := 0; i < b.N; i++ {
		v1.Multiply(v2)
	}
}

func Benchmark_Vector3D_Divide(b *testing.B) {
	v1 := &Vector3D{3, 4, 5}
	v2 := &Vector3D{6, 2, 3}
	for i := 0; i < b.N; i++ {
		v1.Divide(v2)
	}
}

func TestDotProduct3D(t *testing.T) {
	v1 := &Vector3D{1, 2, 3}
	v2 := &Vector3D{4, 5, 6}
	if v1.DotProduct(v2) != 32 {
		t.Error("DotProduct3D")
	}
}

func Benchmark_Vector3D_DotProduct(b *testing.B) {
	v1 := &Vector3D{3, 4, 5}
	v2 := &Vector3D{6, 2, 3}
	for i := 0; i < b.N; i++ {
		v1.DotProduct(v2)
	}
}

func TestCrossProduct3D(t *testing.T) {
	v1 := &Vector3D{1, 2, 3}
	v2 := &Vector3D{4, 5, 6}
	if v := v1.CrossProduct(v2); !v.Equal(&Vector3D{-3, 6, -3}) {
		t.Error("CrossProduct3D")
	}
}

func Benchmark_Vector3D_CrossProduct(b *testing.B) {
	v1 := &Vector3D{3, 4, 5}
	v2 := &Vector3D{6, 2, 3}
	for i := 0; i < b.N; i++ {
		v1.CrossProduct(v2)
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

func Benchmark_Vector3D_Equal_1(b *testing.B) {
	v1 := &Vector3D{3, 4, 5}
	v2 := &Vector3D{1, 4, 5}
	for i := 0; i < b.N; i++ {
		v1.Equal(v2)
	}
}

func Benchmark_Vector3D_Equal_2(b *testing.B) {
	v1 := &Vector3D{3, 4, 5}
	v2 := &Vector3D{3, 2, 5}
	for i := 0; i < b.N; i++ {
		v1.Equal(v2)
	}
}

func Benchmark_Vector3D_Equal_3(b *testing.B) {
	v1 := &Vector3D{3, 4, 5}
	v2 := &Vector3D{3, 4, 5}
	for i := 0; i < b.N; i++ {
		v1.Equal(v2)
	}
}

func Benchmark_Vector3D_FuzzyEqual(b *testing.B) {
	v1 := &Vector3D{3, 4, 5}
	v2 := &Vector3D{3, 4, 5}
	for i := 0; i < b.N; i++ {
		v1.FuzzyEqual(v2)
	}
}

func TestVector3DScale(t *testing.T) {
	v := &Vector3D{1, 1, 1}
	v.Scale(0.5)
	if !v.Equal(&Vector3D{0.5, 0.5, 0.5}) {
		t.Error("Vector3D.Scale")
	}
}

func Benchmark_Vector3D_Scale(b *testing.B) {
	v := &Vector3D{3, 4, 5}
	for i := 0; i < b.N; i++ {
		v.Scale(1.5)
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

func Benchmark_Vector3D_Length(b *testing.B) {
	v := &Vector3D{3, 4, 5}
	for i := 0; i < b.N; i++ {
		v.Length()
	}
}

func Benchmark_Vector3D_LengthSquared(b *testing.B) {
	v := &Vector3D{3, 4, 5}
	for i := 0; i < b.N; i++ {
		v.LengthSquared()
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

func Benchmark_Vector3D_Normalize(b *testing.B) {
	v := &Vector3D{3, 4, 5}
	for i := 0; i < b.N; i++ {
		v.Normalize()
	}
}

func TestScalarProjection3D(t *testing.T) {
	v1 := &Vector3D{-3, 2, -4}
	v2 := &Vector3D{2, -5, 1}
	if v1.ScalarProjection(v2) != -20/math.Sqrt(30) {
		t.Error("Vector3D.ScalarProjection")
	}
}

func Benchmark_Vector3D_ScalarProjection(b *testing.B) {
	v1 := &Vector3D{-3, 2, -4}
	v2 := &Vector3D{2, -5, 1}
	for i := 0; i < b.N; i++ {
		v1.ScalarProjection(v2)
	}
}

func TestVectorProjectionOnto3D(t *testing.T) {
	v1 := &Vector3D{-3, 2, -4}
	v2 := &Vector3D{2, -5, 1}
	if v1.ProjectedOnto(v2); v1.FuzzyEqual(&Vector3D{-4.0 / 3.0, 10.0 / 3.0, -2.0 / 3.0}) {
		t.Error("Vector3D.ProjectedOnto")
	}
}

func Benchmark_Vector3D_ProjectedOnto(b *testing.B) {
	v1 := &Vector3D{-3, 2, -4}
	v2 := &Vector3D{2, -5, 1}
	for i := 0; i < b.N; i++ {
		v1.ProjectedOnto(v2)
	}
}

func TestAngleBetween3D(t *testing.T) {
	v1 := &Vector3D{1, 0, 0}
	v2 := &Vector3D{1, 1, 0}
	if !FuzzyEqual(v1.AngleBetween(v2), math.Pi/4) {
		t.Error("Vector3D.AngleBetween")
	}
	if !FuzzyEqual(v1.CosAngleBetween(v2), math.Cos(math.Pi/4)) {
		t.Error("Vector3D.CosAngleBetween")
	}
}

func Benchmark_Vector3D_AngleBetween(b *testing.B) {
	v1 := &Vector3D{-3, 2, -4}
	v2 := &Vector3D{2, -5, 1}
	for i := 0; i < b.N; i++ {
		v1.AngleBetween(v2)
	}
}

func Benchmark_Vector3D_CosAngleBetween(b *testing.B) {
	v1 := &Vector3D{-3, 2, -4}
	v2 := &Vector3D{2, -5, 1}
	for i := 0; i < b.N; i++ {
		v1.CosAngleBetween(v2)
	}
}
