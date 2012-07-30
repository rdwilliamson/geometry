package geometry

import "testing"

func TestLine3DFuzzyFuzzyEqual(t *testing.T) {
	l1 := &Line3D{Vector3D{1, 2, 3}, Vector3D{1, 1, 1}}
	l2 := &Line3D{Vector3D{1, 1, 1}, Vector3D{1, 2, 3}}
	l1.P1.X += 0.0000000000001
	if !l1.FuzzyEqual(l2) || !l1.FuzzyEqual(l1) {
		t.Error("Line3D.FuzzyEqual")
	}
	l2.P1.X += 0.000000000001
	if l1.FuzzyEqual(l2) || !l1.FuzzyEqual(l1) {
		t.Error("Line3D.FuzzyEqual")
	}
}

func Benchmark_Line3D_FuzzyEqual(b *testing.B) {
	l1 := &Line3D{Vector3D{1, 2, 3}, Vector3D{1, 1, 1}}
	l2 := &Line3D{Vector3D{1, 1, 1}, Vector3D{1, 2, 3}}
	for i := 0; i < b.N; i++ {
		l1.FuzzyEqual(l2)
	}
}

func TestLine3DPointDistance(t *testing.T) {
	l := Line3D{Vector3D{0, 0, 0}, Vector3D{1, 0, 0}}
	p := &Vector3D{0, 1, 0}
	if l.PointDistance(p) != 1 {
		t.Error("Line3D.PointDistance")
	}
}

func Benchmark_Line3D_PointDistance(b *testing.B) {
	l := Line3D{Vector3D{0, 0, 0}, Vector3D{1, 0, 0}}
	p := &Vector3D{0, 1, 0}
	for i := 0; i < b.N; i++ {
		l.PointDistance(p)
	}
}

func TestLine3DPointSquaredDistance(t *testing.T) {
	l := Line3D{Vector3D{0, 0, 0}, Vector3D{1, 0, 0}}
	p := &Vector3D{0, 1, 0}
	if l.PointSquaredDistance(p) != 1 {
		t.Error("Line3D.PointSquaredDistance")
	}
}

func Benchmark_Line3D_PointSquaredDistance(b *testing.B) {
	l := Line3D{Vector3D{0, 0, 0}, Vector3D{1, 0, 0}}
	p := &Vector3D{0, 1, 0}
	for i := 0; i < b.N; i++ {
		l.PointSquaredDistance(p)
	}
}

func TestLine3DLineBetween(t *testing.T) {
	l1 := &Line3D{Vector3D{1, 2, 1}, Vector3D{3, 3, 3}}
	l2 := &Line3D{Vector3D{1, 2, 1}, Vector3D{1, 2, 3}}
	lb := &Line3D{Vector3D{1, 2, 1}, Vector3D{1, 2, 1}}
	if l := l1.LineBetween(l2); !l.Equal(lb) {
		t.Error("Line3D.LineBetween", l)
	}
	l1 = &Line3D{Vector3D{1, 2, 1}, Vector3D{5, 4, 5}}
	l2 = &Line3D{Vector3D{5, 6, 1}, Vector3D{1, 4, 5}}
	lb = &Line3D{Vector3D{3.4, 3.2, 3.4}, Vector3D{2.6, 4.8, 3.4}}
	if l := l1.LineBetween(l2); !l.FuzzyEqual(lb) {
		t.Error("Line3D.LineBetween")
	}
}

func Benchmark_Line3D_LineBetween(b *testing.B) {
	l1 := &Line3D{Vector3D{1, 2, 1}, Vector3D{5, 4, 5}}
	l2 := &Line3D{Vector3D{5, 6, 1}, Vector3D{1, 4, 5}}
	for i := 0; i < b.N; i++ {
		l1.LineBetween(l2)
	}
}
