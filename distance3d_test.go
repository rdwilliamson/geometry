package geometry

import (
	"math"
	"testing"
)

func TestDistance3DLinePointAngular(t *testing.T) {
	l := &Line3D{Vector3D{0, 0, 0}, Vector3D{1, 0, 0}}
	p := &Vector3D{1, 0.5, 0}
	if Distance3DLinePointAngular(l, p) != math.Pi/4 {
		t.Error("Distance3D.LinePointAngular")
	}
}

func Benchmark_Distance3D_LinePointAngular(b *testing.B) {
	l := &Line3D{Vector3D{0, 0, 0}, Vector3D{1, 0, 0}}
	p := &Vector3D{1, 0.5, 0}
	for i := 0; i < b.N; i++ {
		Distance3DLinePointAngular(l, p)
	}
}

func TestDistance3DLinePointAngularCosSquared(t *testing.T) {
	l := &Line3D{Vector3D{0, 0, 0}, Vector3D{1, 0, 0}}
	p := &Vector3D{1, 0.5, 0}
	if Distance3DLinePointAngularCosSquared(l, p) != 0.5 {
		t.Error("Distance3D.LinePointAngularCosSquared")
	}
}

func Benchmark_Distance3D_LinePointAngularCosSquared(b *testing.B) {
	l := &Line3D{Vector3D{0, 0, 0}, Vector3D{1, 0, 0}}
	p := &Vector3D{1, 0.5, 0}
	for i := 0; i < b.N; i++ {
		Distance3DLinePointAngularCosSquared(l, p)
	}
}

func TestDistance3DLinePoint(t *testing.T) {
	l := &Line3D{Vector3D{0, 0, 0}, Vector3D{1, 0, 0}}
	p := &Vector3D{0, 1, 0}
	if Distance3DLinePoint(l, p) != 1 {
		t.Error("Distance3D.LinePoint")
	}
}

func Benchmark_Distance3D_LinePoint(b *testing.B) {
	l := &Line3D{Vector3D{0, 0, 0}, Vector3D{1, 0, 0}}
	p := &Vector3D{0, 1, 0}
	for i := 0; i < b.N; i++ {
		Distance3DLinePoint(l, p)
	}
}

func TestDistance3DLinePointSquared(t *testing.T) {
	l := &Line3D{Vector3D{0, 0, 0}, Vector3D{1, 0, 0}}
	p := &Vector3D{0, 1, 0}
	if Distance3DLinePointSquared(l, p) != 1 {
		t.Error("Distance3D.LinePointSquared")
	}
}

func Benchmark_Distance3D_LinePointSquared(b *testing.B) {
	l := &Line3D{Vector3D{0, 0, 0}, Vector3D{1, 0, 0}}
	p := &Vector3D{0, 1, 0}
	for i := 0; i < b.N; i++ {
		Distance3DLinePointSquared(l, p)
	}
}

type distance3DLineSegmentPointData struct {
	l Line3D
	p Vector3D
	d float64
}

var distance3DLineSegmentPointValues = []distance3DLineSegmentPointData{
	{Line3D{Vector3D{}, Vector3D{1, 0, 0}}, Vector3D{0, 2, 0}, 2},
	{Line3D{Vector3D{}, Vector3D{1, 0, 0}}, Vector3D{0, 0, 2}, 2},
	{Line3D{Vector3D{}, Vector3D{1, 0, 0}}, Vector3D{3, 0, 0}, 2},
	{Line3D{Vector3D{}, Vector3D{1, 0, 0}}, Vector3D{-2, 0, 0}, 2},
	{Line3D{Vector3D{}, Vector3D{1, 0, 0}}, Vector3D{0.5, 0, 0}, 0},
}

func testDistance3DLineSegmentPoint(d distance3DLineSegmentPointData, t *testing.T) {
	if got := Distance3DLineSegmentPoint(&d.l, &d.p); got != d.d {
		t.Error("Distance3D.LineSegmentPointDistance", d.l, d.p, "want", d.d,
			"got", got)
	}
}

func TestDistance3DLineSegmentPoint(t *testing.T) {
	for _, v := range distance3DLineSegmentPointValues {
		testDistance3DLineSegmentPoint(v, t)
	}
}

func Benchmark_Distance3D_LineSegmentPoint(b *testing.B) {
	l := &Line3D{Vector3D{}, Vector3D{1, 1, 1}}
	p := &Vector3D{1, 2, 3}
	for i := 0; i < b.N; i++ {
		Distance3DLineSegmentPoint(l, p)
	}
}

func testDistance3DLineSegmentPointSquared(d distance3DLineSegmentPointData, t *testing.T) {
	if got := Distance3DLineSegmentPointSquared(&d.l, &d.p); got != d.d*d.d {
		t.Error("Distance3D.LineSegmentPointDistanceSquared", d.l, d.p, "want",
			d.d*d.d, "got", got)
	}
}

func TestDistance3DLineSegmentPointSquared(t *testing.T) {
	for _, v := range distance3DLineSegmentPointValues {
		testDistance3DLineSegmentPointSquared(v, t)
	}
}

func Benchmark_Distance3D_LineSegmentPointSquared(b *testing.B) {
	l := &Line3D{Vector3D{}, Vector3D{1, 1, 1}}
	p := &Vector3D{1, 2, 3}
	for i := 0; i < b.N; i++ {
		Distance3DLineSegmentPointSquared(l, p)
	}
}

func TestDistance3DPlaneNormalizedPoint(t *testing.T) {
	pl, pt := &Plane{2, -2, 5, 8}, &Vector3D{4, -4, 3}
	pl.Normalize(pl)
	if d := Distance3DPlaneNormalizedPoint(pl, pt); !FuzzyEqual(d, 39/math.Sqrt(33)) {
		t.Error("Distance3D.PlaneNormalizedPoint", *pl, *pt, "want",
			39/math.Sqrt(33), "got", d)
	}
}

func Benchmark_Distance3D_PlaneNormalizedPoint(b *testing.B) {
	pl := &Plane{1, 2, 3, 4}
	pl.Normalize(pl)
	pt := &Vector3D{5, 6, 7}
	for i := 0; i < b.N; i++ {
		Distance3DPlaneNormalizedPoint(pl, pt)
	}
}

func TestDistance3DPlanePoint(t *testing.T) {
	pl, pt := &Plane{2, -2, 5, 8}, &Vector3D{4, -4, 3}
	if d := Distance3DPlanePoint(pl, pt); d != 39/math.Sqrt(33) {
		t.Error("Distance3D.PlanePoint", *pl, *pt, "want", 39/math.Sqrt(33),
			"got", d)
	}
}

func Benchmark_Distance3D_PlanePoint(b *testing.B) {
	pl := &Plane{1, 2, 3, 4}
	pt := &Vector3D{5, 6, 7}
	for i := 0; i < b.N; i++ {
		Distance3DPlanePoint(pl, pt)
	}
}

func TestDistance3DPlanePointSquared(t *testing.T) {
	pl, pt := &Plane{2, -2, 5, 8}, &Vector3D{4, -4, 3}
	if d := Distance3DPlanePointSquared(pl, pt); d != 39.0*39.0/33.0 {
		t.Error("Distance3D.PlanePointSquared", *pl, *pt, "want",
			39.0*39.0/33.0, "got", d)
	}
}

func Benchmark_Distance3D_PlanePointSquared(b *testing.B) {
	pl := &Plane{1, 2, 3, 4}
	pt := &Vector3D{5, 6, 7}
	for i := 0; i < b.N; i++ {
		Distance3DPlanePointSquared(pl, pt)
	}
}

func TestDistance3DPointPoint(t *testing.T) {
	p1, p2 := &Vector3D{}, &Vector3D{1, 0, 0}
	if Distance3DPointPoint(p1, p2) != 1 {
		t.Error("Distance3D.Distance3DPointPoint")
	}
}

func Benchmark_Vector3D_Distance(b *testing.B) {
	p1, p2 := &Vector3D{}, &Vector3D{1, 2, 3}
	for i := 0; i < b.N; i++ {
		Distance3DPointPoint(p1, p2)
	}
}

func TestDistance3DPointPointSquared(t *testing.T) {
	p1, p2 := &Vector3D{}, &Vector3D{1, 0, 0}
	if Distance3DPointPointSquared(p1, p2) != 1 {
		t.Error("Distance3D.Distance3DPointPointSquared")
	}
}

func Benchmark_Vector3D_DistanceSquared(b *testing.B) {
	p1, p2 := &Vector3D{}, &Vector3D{1, 2, 3}
	for i := 0; i < b.N; i++ {
		Distance3DPointPointSquared(p1, p2)
	}
}

func TestDistance3DVectorVectorAngular(t *testing.T) {
	v1, v2 := &Vector3D{1, 0, 0}, &Vector3D{0, 0, 1}
	if Distance3DVectorVectorAngular(v1, v2) != math.Pi/2 {
		t.Error("Distance3D.VectorVectorAngular")
	}
}

func Benchmark_Distance3D_VectorVectorAngular(b *testing.B) {
	v1, v2 := &Vector3D{1, 2, 5}, &Vector3D{3, 4, 6}
	for i := 0; i < b.N; i++ {
		Distance3DVectorVectorAngular(v1, v2)
	}
}

func TestDistance3DVectorVectorAngularCosSquared(t *testing.T) {
	v1, v2 := &Vector3D{1, 0, 0}, &Vector3D{0, 0, 1}
	if FuzzyEqual(Distance3DVectorVectorAngularCosSquared(v1, v2), math.Sqrt2/2) {
		t.Error("Vector3D.AngularDifferenceCosSquared")
	}
}

func Benchmark_Distance3D_VectorVectorAngularCosSquared(b *testing.B) {
	v1, v2 := &Vector3D{1, 2, 5}, &Vector3D{3, 4, 6}
	for i := 0; i < b.N; i++ {
		Distance3DVectorVectorAngularCosSquared(v1, v2)
	}
}
