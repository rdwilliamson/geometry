package geometry

import (
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

func TestVector2DCopy(t *testing.T) {
	v1, v2 := &Vector2D{}, &Vector2D{1, 0}
	if !v1.Copy(v2).Equal(v2) {
		t.Error("Vector2D.Copy")
	}
}

func Benchmark_Vector2D_Copy(b *testing.B) {
	v1, v2 := &Vector2D{}, &Vector2D{1, 0}
	for i := 0; i < b.N; i++ {
		v1.Copy(v2)
	}
}

type vector2DDirectionEqualData struct {
	v1, v2 Vector2D
	equal  bool
}

var vector2DDirectionEqualValues = []vector2DDirectionEqualData{
	{Vector2D{1, 1}, Vector2D{2, 2}, true},
	{Vector2D{1, 1}, Vector2D{-2, -2}, false},
	{Vector2D{0, 1}, Vector2D{0, 2}, true},
	{Vector2D{0, 1}, Vector2D{0, -2}, false},
	{Vector2D{1, 0}, Vector2D{2, 0}, true},
	{Vector2D{1, 0}, Vector2D{-2, 0}, false},
}

func testVector2DDirectionEqual(d vector2DDirectionEqualData, t *testing.T) {
	if d.v1.DirectionEqual(&d.v2) != d.equal {
		t.Error("Vector2D.DirectionEqual", d.v1, d.v2, "want", d.equal)
	}
	if d.v2.DirectionEqual(&d.v1) != d.equal {
		t.Error("Vector2D.DirectionEqual", d.v2, d.v1, "want", d.equal)
	}
}

func TestVector2DDirectionEqual(t *testing.T) {
	for _, v := range vector2DDirectionEqualValues {
		testVector2DDirectionEqual(v, t)
	}
}

func Benchmark_Vector2D_DirectionEqual(b *testing.B) {
	v1, v2 := &Vector2D{1, 1}, &Vector2D{2, 2}
	for i := 0; i < b.N; i++ {
		v1.DirectionEqual(v2)
	}
}

type vector2DDirectionFuzzyEqualData struct {
	v1, v2 Vector2D
	equal  bool
}

var vector2DDirectionFuzzyEqualValues = []vector2DDirectionFuzzyEqualData{
	{Vector2D{1, 1}, Vector2D{2, 2 + 2e-12}, false},
	{Vector2D{1, 1}, Vector2D{2 + 2e-12, 2}, false},
	{Vector2D{1, 1}, Vector2D{2, 2 + 2e-13}, true},
	{Vector2D{1, 1}, Vector2D{2 + 2e-13, 2}, true},
	{Vector2D{0, 1}, Vector2D{0, 2}, true},
	{Vector2D{0, 1}, Vector2D{0, -2}, false},
	{Vector2D{1, 0}, Vector2D{2, 0}, true},
	{Vector2D{1, 0}, Vector2D{-2, 0}, false},
}

func testVector2DDirectionFuzzyEqual(d vector2DDirectionFuzzyEqualData, t *testing.T) {
	if d.v1.DirectionFuzzyEqual(&d.v2) != d.equal {
		t.Error("Vector2D.DirectionFuzzyEqual:", d.v1, d.v2, "want", d.equal, "got", !d.equal)
	}
	if d.v2.DirectionFuzzyEqual(&d.v1) != d.equal {
		t.Error("Vector2D.DirectionFuzzyEqual:", d.v2, d.v1, "want", d.equal, "got", !d.equal)
	}
}

func TestVector2DDirectionFuzzyEqual(t *testing.T) {
	for _, v := range vector2DDirectionFuzzyEqualValues {
		testVector2DDirectionFuzzyEqual(v, t)
	}
}

func Benchmark_Vector2D_DirectionFuzzyEqual(b *testing.B) {
	v1, v2 := &Vector2D{1, 1}, &Vector2D{2, 2}
	for i := 0; i < b.N; i++ {
		v1.DirectionFuzzyEqual(v2)
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

type fuzzyEqualData struct {
	v1, v2 Vector2D
	equal  bool
}

var fuzzyEqualValues = []fuzzyEqualData{
	{Vector2D{}, Vector2D{0, 1e-12}, false},
	{Vector2D{}, Vector2D{0, 1e-13}, true},
}

func testVector2DFuzzyEqual(d fuzzyEqualData, t *testing.T) {
	if d.v1.FuzzyEqual(&d.v2) != d.equal {
		t.Error("Vector2D.FuzzyEqual", d.v1, d.v2, d.equal)
	}
}

func TestVector2DFuzzyEqual(t *testing.T) {
	for _, v := range fuzzyEqualValues {
		testVector2DFuzzyEqual(v, t)
	}
}

func Benchmark_Vector2D_FuzzyEqual(b *testing.B) {
	v1, v2 := &Vector2D{}, &Vector2D{1, 0}
	for i := 0; i < b.N; i++ {
		v1.FuzzyEqual(v2)
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

func TestVector2DNormalize(t *testing.T) {
	v1, v2 := &Vector2D{3, 4}, &Vector2D{3.0 / 5.0, 4.0 / 5.0}
	if !v1.Normalize(v1).FuzzyEqual(v2) {
		t.Error("Vector2D.Normalize")
	}
}

func Benchmark_Vector2D_Normalize(b *testing.B) {
	v1, v2 := &Vector2D{3, 4}, &Vector2D{}
	for i := 0; i < b.N; i++ {
		v1.Normalize(v2)
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
