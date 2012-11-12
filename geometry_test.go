package geometry

import (
	"math"
	"testing"
)

func (v *Vector2D) toPositiveInf() *Vector2D {
	if math.IsInf(v.X, -1) {
		v.X = math.Inf(1)
	}
	if math.IsInf(v.Y, -1) {
		v.Y = math.Inf(1)
	}
	return v
}

func nanFuzzyEqual(a, b float64) bool {
	return (math.IsNaN(a) && math.IsNaN(b)) || FuzzyEqual(a, b)
}

func (a *Vector2D) nanEqual(b *Vector2D) bool {
	return (math.IsNaN(a.X) && math.IsNaN(a.Y) && math.IsNaN(b.X) && math.IsNaN(b.Y)) || a.Equal(b)
}

func (a *Vector2D) nanFuzzyEqual(b *Vector2D) bool {
	return (math.IsNaN(a.X) && math.IsNaN(a.Y) && math.IsNaN(b.X) && math.IsNaN(b.Y)) || a.FuzzyEqual(b)
}

type fuzzyEqualTestData struct {
	a    float64 // any value
	far  float64 // value far enough from a that the two are not equal
	near float64 // value near enough to a that the two are equal
}

var fuzzyEqualTestValues = []fuzzyEqualTestData{
	{0, 1e-12, 1e-13},
	{1e-12, -1e-12, -1e-13},
	{1 + 1e-12, 1, 1 - 1e-13},
	{1e-12, 0, -1e-13},
	{1e-12, -1e-12, -1e-13},
	{1e11, 1e11 + 1e-1, 1e11 + 1e-2},
	{2, 1 + 2e-11, 2 + 1e-12},
}

func testFuzzyEqual(a, far, near float64, t *testing.T) {
	if a == far {
		t.Fatalf("geometry.FuzzyEqual: far input is exactly equal")
	}
	if a == near {
		t.Fatalf("geometry.FuzzyEqual: near input is exactly equal")
	}
	if FuzzyEqual(a, far) {
		t.Errorf("geometry.FuzzyEqual: %g == %g (far)", a, far)
	}
	if FuzzyEqual(far, a) {
		t.Errorf("geometry.FuzzyEqual: %g == %g (far)", a, far)
	}
	if !FuzzyEqual(a, near) {
		t.Errorf("geometry.FuzzyEqual: %g != %g (near)", a, near)
	}
	if !FuzzyEqual(near, a) {
		t.Errorf("geometry.FuzzyEqual: %f != %f (near)", a, near)
	}
}

func TestFuzzyEqual(t *testing.T) {
	for _, v := range fuzzyEqualTestValues {
		testFuzzyEqual(v.a, v.far, v.near, t)
		testFuzzyEqual(-v.a, -v.far, -v.near, t)
	}
}

func Benchmark_FuzzyEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FuzzyEqual(-1+1e-13, -1)
	}
}
