package geometry

import (
	"math"
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

func TestPlaneCopy(t *testing.T) {
	p1, p2 := &Plane{2, -2, 5, 8}, &Plane{}
	if !p2.Copy(p1).NormalizedEqual(p1) {
		t.Error("Plane.Copy")
	}
}

func Benchmark_Plane_Copy(b *testing.B) {
	p1, p2 := &Plane{2, -2, 5, 8}, &Plane{}
	for i := 0; i < b.N; i++ {
		p2.Copy(p1)
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
	{Plane{2, 4, 6, 8}, Plane{1, 2, 3, 4}, true},
	{Plane{0, 4, 6, 8}, Plane{0, 2, 3, 4}, true},
	{Plane{2, 0, 6, 8}, Plane{1, 0, 3, 4}, true},
	{Plane{2, 4, 0, 8}, Plane{1, 2, 0, 4}, true},
	{Plane{0, 0, 6, 8}, Plane{0, 0, 3, 4}, true},
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

func TestPlaneFromPoints(t *testing.T) {
	got, want := Plane{}, &Plane{28, 11, -26, 38}
	pt1, pt2, pt3 := &Vector3D{1, -6, 0}, &Vector3D{-4, 2, -2}, &Vector3D{-2, 4, 1}
	if !got.FromPoints(pt1, pt2, pt3).NormalizedEqual(want) {
		t.Error("Plane.FromPoints")
	}
}

func Benchmark_Plane_FromPoints(b *testing.B) {
	pl := Plane{}
	pt1, pt2, pt3 := &Vector3D{1, -6, 0}, &Vector3D{-4, 2, -2}, &Vector3D{-2, 4, 1}
	for i := 0; i < b.N; i++ {
		pl.FromPoints(pt1, pt2, pt3)
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
	{Plane{0, 4, 6, 8}, Plane{0, 2, 3, 4}, true},
	{Plane{2, 0, 6, 8}, Plane{1, 0, 3, 4}, true},
	{Plane{2, 4, 0, 8}, Plane{1, 2, 0, 4}, true},
	{Plane{0, 0, 6, 8}, Plane{0, 0, 3, 4}, true},
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

func TestPlaneNormal(t *testing.T) {
	p := &Plane{1, 2, 3, 4}
	v := &Vector3D{}
	if !p.Normal(v).Equal(&Vector3D{1, 2, 3}) {
		t.Error("Plane.Normal")
	}
}

func Benchmark_Plane_Normal(b *testing.B) {
	p := &Plane{1, 2, 3, 4}
	v := &Vector3D{}
	for i := 0; i < b.N; i++ {
		p.Normal(v)
	}
}

func TestPlaneNormalize(t *testing.T) {
	s := 1 / math.Sqrt(29)
	p1, p2 := &Plane{2, 3, 4, 5}, &Plane{s * 2, s * 3, s * 4, s * 5}
	if !p1.Normalize(p1).NormalizedEqual(p2) {
		t.Error("Plane.Normalize")
	}
}

func Benchmark_Plane_Normalize(b *testing.B) {
	p1, p2 := &Plane{2, 3, 4, -5}, &Plane{}
	for i := 0; i < b.N; i++ {
		p1.Normalize(p2)
	}
}

func TestNormalizedEqual(t *testing.T) {
	s := 1 / math.Sqrt(14)
	p, pe := &Plane{1, 2, 3, 4}, &Plane{s * 1, s * 2, s * 3, s * 4}
	p.Normalize(p)
	if !p.NormalizedEqual(pe) {
		t.Error("Plane.NormalizedEqual")
	}
	*p = Plane{-1, -2, -3, -4}
	p.Normalize(p)
	if !p.NormalizedEqual(pe) {
		t.Error("Plane.NormalizedEqual")
	}
}

func Benchmark_Plane_NormalizedEqual(b *testing.B) {
	p1, p2 := &Plane{1, 2, 3, 4}, &Plane{-1, -2, -3, -4}
	for i := 0; i < b.N; i++ {
		p1.NormalizedEqual(p2)
	}
}
