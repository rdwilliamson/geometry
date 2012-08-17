package geometry

import (
	"math"
	"testing"
)

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

func Benchmark_Distance3DPlanePointSquared(b *testing.B) {
	pl := &Plane{1, 2, 3, 4}
	pt := &Vector3D{5, 6, 7}
	for i := 0; i < b.N; i++ {
		Distance3DPlanePointSquared(pl, pt)
	}
}
