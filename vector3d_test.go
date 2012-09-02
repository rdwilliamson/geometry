package geometry

import (
	"math"
	"testing"
)

func TestNewVector3D(t *testing.T) {
	if !NewVector3D(1, 2, 3).Equal(&Vector3D{1, 2, 3}) {
		t.Error("NewVector3D")
	}
}

func Benchmark_Vector3D_New(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewVector3D(0, 0, 0)
	}
}

func TestVector3DAdd(t *testing.T) {
	r, v1, v2 := &Vector3D{}, &Vector3D{1, 2, 5}, &Vector3D{3, 4, 6}
	if !r.Add(v1, v2).Equal(&Vector3D{4, 6, 11}) {
		t.Error("Vector3D.Add")
	}
}

func Benchmark_Vector3D_Add(b *testing.B) {
	r, v1, v2 := &Vector3D{}, &Vector3D{1, 2, 5}, &Vector3D{3, 4, 6}
	for i := 0; i < b.N; i++ {
		r.Add(v1, v2)
	}
}

func TestVector3DCopy(t *testing.T) {
	v1, v2 := &Vector3D{}, &Vector3D{1, 2, 3}
	if !v1.Copy(v2).Equal(v2) {
		t.Error("Vector3D.Copy")
	}
}

func Benchmark_Vector3D_Copy(b *testing.B) {
	v1, v2 := &Vector3D{}, &Vector3D{1, 2, 3}
	for i := 0; i < b.N; i++ {
		v1.Copy(v2)
	}
}

func TestVector3DCrossProduct(t *testing.T) {
	r, v1, v2 := &Vector3D{}, &Vector3D{1, 2, 3}, &Vector3D{4, 5, 6}
	if !r.CrossProduct(v1, v2).Equal(&Vector3D{-3, 6, -3}) {
		t.Error("Vector3D.CrossProduct")
	}
}

func Benchmark_Vector3D_CrossProduct(b *testing.B) {
	r, v1, v2 := &Vector3D{}, &Vector3D{1, 2, 3}, &Vector3D{4, 5, 6}
	for i := 0; i < b.N; i++ {
		r.CrossProduct(v1, v2)
	}
}

type vector3DDirectionEqualData struct {
	v1, v2 Vector3D
	equal  bool
}

var vector3DDirectionEqualValues = []vector3DDirectionEqualData{
	{Vector3D{1, 1, 1}, Vector3D{2, 2, 2}, true},
	{Vector3D{1, 1, 1}, Vector3D{-2, -2, -2}, false},
	{Vector3D{1, 0, 0}, Vector3D{2, 0, 0}, true},
	{Vector3D{1, 0, 0}, Vector3D{-2, 0, 0}, false},
	{Vector3D{0, 1, 0}, Vector3D{0, 2, 0}, true},
	{Vector3D{0, 1, 0}, Vector3D{0, -2, 0}, false},
	{Vector3D{0, 0, 1}, Vector3D{0, 0, 2}, true},
	{Vector3D{0, 0, 1}, Vector3D{0, 0, -2}, false},
}

func testVector3DDirectionEqual(d vector3DDirectionEqualData, t *testing.T) {
	if d.v1.DirectionEqual(&d.v2) != d.equal {
		t.Error("Vector3D.DirectionEqual", d.v1, d.v2, "want", d.equal)
	}
	if d.v2.DirectionEqual(&d.v1) != d.equal {
		t.Error("Vector3D.DirectionEqual", d.v2, d.v1, "want", d.equal)
	}
}

func TestVector3DDirectionEqual(t *testing.T) {
	for _, v := range vector3DDirectionEqualValues {
		testVector3DDirectionEqual(v, t)
	}
}

func Benchmark_Vector3D_DirectionEqual(b *testing.B) {
	v1, v2 := &Vector3D{1, 1, 1}, &Vector3D{2, 2, 2}
	// v1, v2 := &Vector3D{0, 0, 1}, &Vector3D{0, 0, 2}
	for i := 0; i < b.N; i++ {
		v1.DirectionEqual(v2)
	}
}

type vector3DDirectionFuzzyEqualData struct {
	v1, v2 Vector3D
	equal  bool
}

var vector3DDirectionFuzzyEqualValues = []vector3DDirectionFuzzyEqualData{
	{Vector3D{1, 1, 1}, Vector3D{2, 2, 2 + 1e-11}, false},
	{Vector3D{1, 1, 1}, Vector3D{2, 2, 2 + 1e-12}, true},
	{Vector3D{1, 1, 1}, Vector3D{2, 2 + 1e-11, 2}, false},
	{Vector3D{1, 1, 1}, Vector3D{2, 2 + 1e-12, 2}, true},
	{Vector3D{1, 1, 1}, Vector3D{2 + 1e-11, 2, 2}, false},
	{Vector3D{1, 1, 1}, Vector3D{2 + 1e-12, 2, 2}, true},
	{Vector3D{1, 1, 1}, Vector3D{1 + 1e-12, 1, 1}, false},
	{Vector3D{1, 1, 1}, Vector3D{1 + 1e-13, 1, 1}, true},
	{Vector3D{1, 0, 0}, Vector3D{2, 0, 0}, true},
	{Vector3D{1, 0, 0}, Vector3D{-2, 0, 0}, false},
	{Vector3D{0, 1, 0}, Vector3D{0, 2, 0}, true},
	{Vector3D{0, 1, 0}, Vector3D{0, -2, 0}, false},
	{Vector3D{0, 0, 1}, Vector3D{0, 0, 2}, true},
	{Vector3D{0, 0, 1}, Vector3D{0, 0, -2}, false},
}

func testVector3DDirectionFuzzyEqual(d vector3DDirectionFuzzyEqualData, t *testing.T) {
	if d.v1.DirectionFuzzyEqual(&d.v2) != d.equal {
		t.Error("Vector3D.DirectionFuzzyEqual", d.v1, d.v2, "want", d.equal)
	}
	if d.v2.DirectionFuzzyEqual(&d.v1) != d.equal {
		t.Error("Vector3D.DirectionFuzzyEqual", d.v2, d.v1, "want", d.equal)
	}
}

func TestVector3DDirectionFuzzyEqual(t *testing.T) {
	for _, v := range vector3DDirectionFuzzyEqualValues {
		testVector3DDirectionFuzzyEqual(v, t)
	}
}

func Benchmark_Vector3D_DirectionFuzzyEqual(b *testing.B) {
	v1, v2 := &Vector3D{1, 1, 1}, &Vector3D{2, 2, 2}
	// v1, v2 := &Vector3D{0, 0, 1}, &Vector3D{0, 0, 2}
	for i := 0; i < b.N; i++ {
		v1.DirectionFuzzyEqual(v2)
	}
}

func TestVector3DDotProduct(t *testing.T) {
	v1, v2 := &Vector3D{1, 2, 5}, &Vector3D{3, 4, 6}
	if v1.DotProduct(v2) != 41 {
		t.Error("Vector3D.DotProduct")
	}
}

func Benchmark_Vector3D_DotProduct(b *testing.B) {
	v1, v2 := &Vector3D{1, 2, 5}, &Vector3D{3, 4, 6}
	for i := 0; i < b.N; i++ {
		v1.DotProduct(v2)
	}
}

func TestVector3DEqual(t *testing.T) {
	v1, v2 := &Vector3D{}, &Vector3D{1, 0, 0}
	if v1.Equal(v2) {
		t.Error("Vector3D.Equal")
	}
	v2.X = 0
	if !v1.Equal(v2) {
		t.Error("Vector3D.Equal")
	}
}

func Benchmark_Vector3D_Equal(b *testing.B) {
	v1, v2 := &Vector3D{}, &Vector3D{1, 0, 0}
	for i := 0; i < b.N; i++ {
		v1.Equal(v2)
	}
}

type vector3DFuzzyEqualData struct {
	v1, v2 Vector3D
	equal  bool
}

var vector3DFuzzyEqualValues = []vector3DFuzzyEqualData{
	{Vector3D{1, 1, 1}, Vector3D{1, 1, 1 + 1e-12}, false},
	{Vector3D{1, 1, 1}, Vector3D{1, 1, 1 + 1e-13}, true},
	{Vector3D{1, 1, 1}, Vector3D{1, 1 + 1e-12, 1}, false},
	{Vector3D{1, 1, 1}, Vector3D{1, 1 + 1e-13, 1}, true},
	{Vector3D{1, 1, 1}, Vector3D{1 + 1e-12, 1, 1}, false},
	{Vector3D{1, 1, 1}, Vector3D{1 + 1e-13, 1, 1}, true},
}

func testVector3DFuzzyEqual(d vector3DFuzzyEqualData, t *testing.T) {
	if d.v1.FuzzyEqual(&d.v2) != d.equal {
		t.Error("Vector3D.FuzzyEqual", d.v1, d.v2, d.equal)
	}
	if d.v2.FuzzyEqual(&d.v1) != d.equal {
		t.Error("Vector3D.FuzzyEqual", d.v2, d.v1, d.equal)
	}
}

func TestVector3DFuzzyEqual(t *testing.T) {
	for _, v := range vector3DFuzzyEqualValues {
		testVector3DFuzzyEqual(v, t)
	}
}

func Benchmark_Vector3D_FuzzyEqual(b *testing.B) {
	v1, v2 := &Vector3D{}, &Vector3D{1, 0, 0}
	for i := 0; i < b.N; i++ {
		v1.FuzzyEqual(v2)
	}
}

func TestVector3DMagnitude(t *testing.T) {
	v := &Vector3D{3, 4, 5}
	if v.Magnitude() != math.Sqrt(50) {
		t.Error("Vector3D.Magnitude")
	}
}

func Benchmark_Vector3D_Magnitude(b *testing.B) {
	v := &Vector3D{3, 4, 5}
	for i := 0; i < b.N; i++ {
		v.Magnitude()
	}
}

func TestVector3DMagnitudeSquared(t *testing.T) {
	v := &Vector3D{3, 4, 5}
	if v.MagnitudeSquared() != 50 {
		t.Error("Vector3D.MagnitudeSquared")
	}
}

func Benchmark_Vector3D_MagnitudeSquared(b *testing.B) {
	v := &Vector3D{3, 4, 5}
	for i := 0; i < b.N; i++ {
		v.MagnitudeSquared()
	}
}

func TestVector3DNormalize(t *testing.T) {
	s := 1 / math.Sqrt(29)
	v1, v2 := &Vector3D{2, 3, 4}, &Vector3D{s * 2, s * 3, s * 4}
	if !v1.Normalize(v1).FuzzyEqual(v2) {
		t.Error("Vector3D.Normalize")
	}
}

func Benchmark_Vector3D_Normalize(b *testing.B) {
	v1, v2 := &Vector3D{2, 3, 4}, &Vector3D{}
	for i := 0; i < b.N; i++ {
		v1.Normalize(v2)
	}
}

func TestVector3DSubtract(t *testing.T) {
	r, v1, v2 := &Vector3D{}, &Vector3D{1, 2, 5}, &Vector3D{3, 4, 6}
	if !r.Subtract(v1, v2).Equal(&Vector3D{-2, -2, -1}) {
		t.Error("Vector3D.Subtract")
	}
}

func Benchmark_Vector3D_Subtract(b *testing.B) {
	r, v1, v2 := &Vector3D{}, &Vector3D{1, 2, 5}, &Vector3D{3, 4, 6}
	for i := 0; i < b.N; i++ {
		r.Subtract(v1, v2)
	}
}
