package geometry

import (
	"math"
	"testing"
)

func TestDistance2DLinePointAngular(t *testing.T) {
	l := &Line2D{Vector2D{}, Vector2D{1, 1}}
	p := &Vector2D{1, 0.5}
	if Distance2DLinePointAngular(l, p) != math.Pi/4 {
		t.Error("Distance2D.LinePointAngular")
	}
	l.P1, l.P2 = l.P2, l.P1
	if !FuzzyEqual(Distance2DLinePointAngular(l, p), math.Pi/4) {
		t.Error("Distance2D.LinePointAngular")
	}
}

func Benchmark_Distance2D_LinePointAngular(b *testing.B) {
	l, p := &Line2D{Vector2D{1, 1}, Vector2D{}}, &Vector2D{1, 0.5}
	for i := 0; i < b.N; i++ {
		Distance2DLinePointAngular(l, p)
	}
}

func TestDistance2DLinePointAngularCosSquared(t *testing.T) {
	l := &Line2D{Vector2D{}, Vector2D{1, 1}}
	p := &Vector2D{1, 0.5}
	if Distance2DLinePointAngularCosSquared(l, p) != 0.5 {
		t.Error("Distance2D.LinePointAngularCosSquared")
	}
	l.P1, l.P2 = l.P2, l.P1
	if Distance2DLinePointAngularCosSquared(l, p) != 0.5 {
		t.Error("Distance2D.LinePointAngularCosSquared")
	}
}

func Benchmark_Distance2D_LinePointAngularCosSquared(b *testing.B) {
	l, p := &Line2D{Vector2D{1, 1}, Vector2D{}}, &Vector2D{1, 0.5}
	for i := 0; i < b.N; i++ {
		Distance2DLinePointAngularCosSquared(l, p)
	}
}

func TestDistance2DLinePoint(t *testing.T) {
	l, p := &Line2D{Vector2D{0, 0}, Vector2D{1, 0}}, &Vector2D{0, 1}
	if Distance2DLinePoint(l, p) != 1 {
		t.Error("Distance2D.LinePoint")
	}
}

func Benchmark_Distance2D_LinePoint(b *testing.B) {
	l, p := &Line2D{Vector2D{0, 0}, Vector2D{1, 0}}, &Vector2D{0, 1}
	for i := 0; i < b.N; i++ {
		Distance2DLinePoint(l, p)
	}
}

func TestDistance2DLinePointSquared(t *testing.T) {
	l, p := &Line2D{Vector2D{0, 0}, Vector2D{1, 0}}, &Vector2D{0, 1}
	if Distance2DLinePointSquared(l, p) != 1 {
		t.Error("Distance2D.LinePointSquared")
	}
}

func Benchmark_Distance2D_PointDistanceSquared(b *testing.B) {
	l, p := &Line2D{Vector2D{0, 0}, Vector2D{1, 0}}, &Vector2D{0, 1}
	for i := 0; i < b.N; i++ {
		Distance2DLinePointSquared(l, p)
	}
}

func TestDistance2DLineSegmentPoint(t *testing.T) {
	l := &Line2D{Vector2D{0, 1}, Vector2D{1, 1}}
	p := &Vector2D{-1, 0}
	if Distance2DLineSegmentPoint(l, p) != math.Sqrt2 {
		t.Error("Distance2D.LineSegmentPoint")
	}
	p.X = 0.5
	if Distance2DLineSegmentPoint(l, p) != 1 {
		t.Error("Distance2D.LineSegmentPoint")
	}
	p.X = 2
	if Distance2DLineSegmentPoint(l, p) != math.Sqrt2 {
		t.Error("Distance2D.LineSegmentPoint")
	}
}

func Benchmark_Distance2D_LineSegmentPoint(b *testing.B) {
	l := &Line2D{Vector2D{0, 1}, Vector2D{1, 1}}
	p := &Vector2D{0.5, 0}
	for i := 0; i < b.N; i++ {
		Distance2DLineSegmentPoint(l, p)
	}
}

func TestDistance2DLineSegmentPointSquared(t *testing.T) {
	l := &Line2D{Vector2D{0, 1}, Vector2D{1, 1}}
	p := &Vector2D{-1, 0}
	if Distance2DLineSegmentPointSquared(l, p) != 2 {
		t.Error("Distance2DLineSegmentPointSquared")
	}
	p.X = 0.5
	if Distance2DLineSegmentPointSquared(l, p) != 1 {
		t.Error("Distance2DLineSegmentPointSquared")
	}
	p.X = 2
	if Distance2DLineSegmentPointSquared(l, p) != 2 {
		t.Error("Distance2DLineSegmentPointSquared")
	}
}

func Benchmark_Distance2D_LineSegmentPointSquared(b *testing.B) {
	l := &Line2D{Vector2D{0, 1}, Vector2D{1, 1}}
	p := &Vector2D{0.5, 0}
	for i := 0; i < b.N; i++ {
		Distance2DLineSegmentPointSquared(l, p)
	}
}

func TestDistance2DPointPoint(t *testing.T) {
	p1, p2 := &Vector2D{}, &Vector2D{1, 0}
	if Distance2DPointPoint(p1, p2) != 1 {
		t.Error("Distance2D.PointPoint")
	}
}

func Benchmark_Distance2D_PointPoint(b *testing.B) {
	p1, p2 := &Vector2D{}, &Vector2D{1, 0}
	for i := 0; i < b.N; i++ {
		Distance2DPointPoint(p1, p2)
	}
}

func TestDistance2DPointPointSquared(t *testing.T) {
	p1, p2 := &Vector2D{}, &Vector2D{1, 0}
	if Distance2DPointPointSquared(p1, p2) != 1 {
		t.Error("Distance2D.PointPointSquared")
	}
}

func Benchmark_Distance2DPointPointSquared(b *testing.B) {
	p1, p2 := &Vector2D{}, &Vector2D{1, 0}
	for i := 0; i < b.N; i++ {
		Distance2DPointPointSquared(p1, p2)
	}
}

func TestDistance2DVectorVectorAngular(t *testing.T) {
	v1, v2 := &Vector2D{1, 0}, &Vector2D{0, 1}
	if Distance2DVectorVectorAngular(v1, v2) != math.Pi/2 {
		t.Error("Distance2D.VectorVectorAngular")
	}
}

func Benchmark_Distance2D_VectorVectorAngular(b *testing.B) {
	v1, v2 := &Vector2D{1, 2}, &Vector2D{3, 4}
	for i := 0; i < b.N; i++ {
		Distance2DVectorVectorAngular(v1, v2)
	}
}

func TestDistance2DVectorVectorAngularCosSquared(t *testing.T) {
	v1, v2 := &Vector2D{1, 0}, &Vector2D{0, 1}
	if FuzzyEqual(Distance2DVectorVectorAngularCosSquared(v1, v2),
		math.Sqrt2/2) {
		t.Error("Distance2D.VectorVectorAngularCosSquared")
	}
}

func Benchmark_Distance2D_VectorVectorAngularCosSquared(b *testing.B) {
	v1, v2 := &Vector2D{1, 2}, &Vector2D{3, 4}
	for i := 0; i < b.N; i++ {
		Distance2DVectorVectorAngularCosSquared(v1, v2)
	}
}
