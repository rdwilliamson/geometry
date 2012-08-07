package geometry

import "testing"

type fuzzyTestData struct {
	a    float64 // any value
	far  float64 // value far enough from a that the two are not equal
	near float64 // value near enough to a that the two are equal
}

var fuzzyTestValues = []fuzzyTestData{
	{0, 1e-12, 1e-13},
	{1e-12, -1e-12, -1e-13},
	{1 + 1e-12, 1, 1 - 1e-13},
	{1e-12, 0, -1e-13},
	{1e-12, -1e-12, -1e-13},
	{1e11, 1e11 + 1e-1, 1e11 + 1e-2},
}

func testFuzzyEqual(a, far, near float64, t *testing.T) {
	if a == far {
		t.Fatalf("geometry.FuzzyEqual: far input is exactly equal")
	}
	if a == near {
		t.Fatalf("geometry.FuzzyEqual: near input is exactly equal")
	}
	if FuzzyEqual(a, far) {
		t.Errorf("geometry.FuzzyEqual: %f == %f", a, far)
	}
	if FuzzyEqual(far, a) {
		t.Errorf("geometry.FuzzyEqual: %f == %f", a, far)
	}
	if !FuzzyEqual(a, near) {
		t.Errorf("geometry.FuzzyEqual: %f != %f", a, near)
	}
	if !FuzzyEqual(near, a) {
		t.Errorf("geometry.FuzzyEqual: %f != %f", a, near)
	}
}

func TestFuzzyEqual(t *testing.T) {
	for _, v := range fuzzyTestValues {
		testFuzzyEqual(v.a, v.far, v.near, t)
		testFuzzyEqual(-v.a, -v.far, -v.near, t)
	}
}

func Benchmark_FuzzyEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FuzzyEqual(-1+1e-13, -1)
	}
}
