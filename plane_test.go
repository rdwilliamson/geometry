package geometry

import (
	"testing"
)

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
