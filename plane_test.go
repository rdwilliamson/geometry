package geometry

import (
	"testing"
)

func TestNewPlane(t *testing.T) {
	p := NewPlane(1, 2, 3, 4)
	if !p.NormalizedEqual(&Plane{1, 2, 3, 4}) {
		t.Error("NewPlane")
	}
}

func Benchmark_NewPlane(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPlane(1, 2, 3, 4)
	}
}

type planeEqualData struct {
	p1, p2 Plane
	equal  bool
}

var planeEqualValues = []planeEqualData{
	{Plane{-1, -2, -3, -4}, Plane{1, 2, 3, 4}, true},
	{Plane{-1, -2, -3, -4}, Plane{1, -2, 3, 4}, false},
	{Plane{-1, -2, -3, -4}, Plane{1, 2, 3, -4}, false},
	{Plane{2, 4, 6, 8}, Plane{1, 2, 3, 4}, true},
}

func testPlaneEqual(d planeEqualData, t *testing.T) {
	if d.p1.Equal(&d.p2) != d.equal {
		t.Error("Plane.Equal", d.p1, d.p2, d.equal)
	}
	if d.p2.Equal(&d.p1) != d.equal {
		t.Error("Plane.Equal", d.p2, d.p1, d.equal)
	}
}

func TestPlaneEqual(t *testing.T) {
	for _, v := range planeEqualValues {
		testPlaneEqual(v, t)
	}
}

func Benchmark_Plane_Equal(b *testing.B) {
	p1, p2 := &Plane{1, 2, 3, 4}, &Plane{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		p1.Equal(p2)
	}
}

type planeFuzzyEqualData struct {
	p1, p2 Plane
	equal  bool
}

var planeFuzzyEqualValues = []planeFuzzyEqualData{
	{Plane{-1, -2, -3, -4}, Plane{1 + 1e-12, 2, 3, 4}, false},
	{Plane{-1, -2, -3, -4}, Plane{1 + 1e-13, 2, 3, 4}, true},
	{Plane{-1, -2, -3, -4}, Plane{1, 2 + 1e-11, 3, 4}, false},
	{Plane{-1, -2, -3, -4}, Plane{1, 2 + 1e-12, 3, 4}, true},
	{Plane{-1, -2, -3, -4}, Plane{1, 2, 3 + 1e-11, 4}, false},
	{Plane{-1, -2, -3, -4}, Plane{1, 2, 3 + 1e-12, 4}, true},
	{Plane{-1, -2, -3, -4}, Plane{1, 2, 3, 4 + 1e-11}, false},
	{Plane{-1, -2, -3, -4}, Plane{1, 2, 3, 4 + 1e-12}, true},
	{Plane{-1, -2, -3, -4}, Plane{1, -2, 3, 4}, false},
	{Plane{-1, -2, -3, -4}, Plane{1, 2, 3, -4}, false},
	{Plane{2, 4, 6, 8}, Plane{1, 2, 3, 4}, true},
}

func testPlaneFuzzyEqual(d planeFuzzyEqualData, t *testing.T) {
	if d.p1.FuzzyEqual(&d.p2) != d.equal {
		t.Error("Plane.FuzzyEqual", d.p1, d.p2, d.equal)
	}
	if d.p2.FuzzyEqual(&d.p1) != d.equal {
		t.Error("Plane.FuzzyEqual", d.p2, d.p1, d.equal)
	}
}

func TestPlaneFuzzyEqual(t *testing.T) {
	for _, v := range planeFuzzyEqualValues {
		testPlaneFuzzyEqual(v, t)
	}
}

func Benchmark_Plane_FuzzyEqual(b *testing.B) {
	p1, p2 := &Plane{1, 2, 3, 4}, &Plane{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		p1.FuzzyEqual(p2)
	}
}

func TestNormalizedEqual(t *testing.T) {
	p := Plane{1, 2, 3, 4}
	if !p.NormalizedEqual(&Plane{1, 2, 3, 4}) {
		t.Error("Plane.NormalizedEqual")
	}
}

func Benchmark_Plane_NormalizedEqual(b *testing.B) {
	p1, p2 := &Plane{1, 2, 3, 4}, &Plane{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		p1.NormalizedEqual(p2)
	}
}

func Benchmark_Plane_NormalizedPointDistance(b *testing.B) {
	pl := &Plane{1, 2, 3, 4}
	pl.Normalize(pl)
	pt := &Vector3D{5, 6, 7}
	for i := 0; i < b.N; i++ {
		pl.NormalizedPointDistance(pt)
	}
}

func Benchmark_Plane_PointDistance(b *testing.B) {
	pl := &Plane{1, 2, 3, 4}
	pt := &Vector3D{5, 6, 7}
	for i := 0; i < b.N; i++ {
		pl.PointDistance(pt)
	}
}

func Benchmark_Plane_PointDistanceSquared(b *testing.B) {
	pl := &Plane{1, 2, 3, 4}
	pt := &Vector3D{5, 6, 7}
	for i := 0; i < b.N; i++ {
		pl.PointDistanceSquared(pt)
	}
}
