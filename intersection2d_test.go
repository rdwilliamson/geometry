package geometry

import (
	"testing"
)

type intersection2DLineLineData struct {
	l1, l2 Line2D
	p      Vector2D
	n      int
}

var intersection2DLineLineValues = []intersection2DLineLineData{
	{Line2D{Vector2D{0, 0}, Vector2D{1, 1}},
		Line2D{Vector2D{0, 1}, Vector2D{1, 0}},
		Vector2D{0.5, 0.5}, 1},
	{Line2D{Vector2D{0, 0}, Vector2D{-1, -1}},
		Line2D{Vector2D{0, 1}, Vector2D{1, 0}},
		Vector2D{0.5, 0.5}, 1},
	{Line2D{Vector2D{0, 0}, Vector2D{1, 1}},
		Line2D{Vector2D{1, 0}, Vector2D{2, 1}},
		Vector2D{}, 0},
	{Line2D{Vector2D{0, 0}, Vector2D{1, 1}},
		Line2D{Vector2D{-1, -1}, Vector2D{0, 0}},
		Vector2D{}, -1},
}

func testVector2DFromLineIntersection(d intersection2DLineLineData, t *testing.T) {
	var p Vector2D
	if n := Intersection2DLineLine(&d.l1, &d.l2, &p); n != d.n ||
		(d.n == 1 && !p.Equal(&d.p)) {
		if d.n == 1 {
			t.Error("Vector2D.FromLineIntersection", d.l1, d.l2, "want", d.n,
				d.p, "got", n, p)
		} else {
			t.Error("Vector2D.FromLineIntersection", d.l1, d.l2, "want", d.n,
				"got", n)
		}
	}
	if n := Intersection2DLineLine(&d.l2, &d.l1, &p); n != d.n ||
		(d.n == 1 && !p.Equal(&d.p)) {
		if d.n == 1 {
			t.Error("Vector2D.FromLineIntersection", d.l2, d.l1, "want", d.n,
				d.p, "got", n, p)
		} else {
			t.Error("Vector2D.FromLineIntersection", d.l2, d.l1, "want", d.n,
				"got", n)
		}
	}
}

func TestIntersection2DLineLine(t *testing.T) {
	// for _, v := range intersection2DLineLineValues {
	// 	testVector2DFromLineIntersection(v, t)
	// 	v.l1.P1, v.l1.P2 = v.l1.P2, v.l1.P1
	// 	testVector2DFromLineIntersection(v, t)
	// 	v.l2.P1, v.l2.P2 = v.l2.P2, v.l2.P1
	// 	testVector2DFromLineIntersection(v, t)
	// }
}

func Benchmark_Intersection2D_LineLine_Coincident(b *testing.B) {
	l1 := &Line2D{Vector2D{0, 0}, Vector2D{1, 1}}
	l2 := &Line2D{Vector2D{-1, -1}, Vector2D{0, 0}}
	p := &Vector2D{}
	for i := 0; i < b.N; i++ {
		Intersection2DLineLine(l1, l2, p)
	}
}

func Benchmark_Intersection2D_LineLine_Parallel(b *testing.B) {
	l1 := &Line2D{Vector2D{0, 0}, Vector2D{1, 1}}
	l2 := &Line2D{Vector2D{1, 0}, Vector2D{2, 1}}
	p := &Vector2D{}
	for i := 0; i < b.N; i++ {
		Intersection2DLineLine(l1, l2, p)
	}
}

func Benchmark_Intersection2D_LineLine_Intersect(b *testing.B) {
	l1 := &Line2D{Vector2D{0, 0}, Vector2D{1, 1}}
	l2 := &Line2D{Vector2D{0, 1}, Vector2D{1, 0}}
	p := &Vector2D{}
	for i := 0; i < b.N; i++ {
		Intersection2DLineLine(l1, l2, p)
	}
}

func Benchmark_Intersection2DFuzzy_LineLine_Coincident(b *testing.B) {
	l1 := &Line2D{Vector2D{0, 0}, Vector2D{1, 1}}
	l2 := &Line2D{Vector2D{-1, -1}, Vector2D{0, 0}}
	p := &Vector2D{}
	for i := 0; i < b.N; i++ {
		Intersection2DFuzzyLineLine(l1, l2, p)
	}
}

func Benchmark_Intersection2DFuzzy_LineLine_Parallel(b *testing.B) {
	l1 := &Line2D{Vector2D{0, 0}, Vector2D{1, 1}}
	l2 := &Line2D{Vector2D{1, 0}, Vector2D{2, 1}}
	p := &Vector2D{}
	for i := 0; i < b.N; i++ {
		Intersection2DFuzzyLineLine(l1, l2, p)
	}
}

func Benchmark_Intersection2DFuzzy_LineLine_Intersect(b *testing.B) {
	l1 := &Line2D{Vector2D{0, 0}, Vector2D{1, 1}}
	l2 := &Line2D{Vector2D{0, 1}, Vector2D{1, 0}}
	p := &Vector2D{}
	for i := 0; i < b.N; i++ {
		Intersection2DFuzzyLineLine(l1, l2, p)
	}
}

/*
type vector2DFromLineSegmentIntersection struct {
	l1, l2 Line2D
	p      Vector2D
	onBoth bool
}

var vector2DFromLineSegmentIntersectionValue = []vector2DFromLineSegmentIntersection{
	{Line2D{Vector2D{0, 0}, Vector2D{1, 1}}, Line2D{Vector2D{0, 1}, Vector2D{1, 0}}, Vector2D{0.5, 0.5}, true},
	{Line2D{Vector2D{0, 0}, Vector2D{1, 1}}, Line2D{Vector2D{1, 0}, Vector2D{2, 1}}, Vector2D{math.Inf(1), math.Inf(1)}, false},
	{Line2D{Vector2D{0, 0}, Vector2D{-1, -1}}, Line2D{Vector2D{0, 1}, Vector2D{1, 0}}, Vector2D{0.5, 0.5}, false},
}

func testVector2DFromLineSegmentIntersection(d vector2DFromLineSegmentIntersection, t *testing.T) {
	p := &Vector2D{}
	if b := p.FromLineSegmentIntersection(&d.l1, &d.l2); b != d.onBoth || !d.p.infEqual(p) {
		t.Error("Vector2D.FromLineSegmentIntersection", d.l1, d.l2, "want", d.p, "got", *p, "want", d.onBoth, "got", b)
	}
	if b := p.FromLineSegmentIntersection(&d.l2, &d.l1); b != d.onBoth || !d.p.infEqual(p) {
		t.Error("Vector2D.FromLineSegmentIntersection", d.l2, d.l1, "want", d.p, "got", *p, "want", d.onBoth, "got", b)
	}
	d.l1.P1, d.l1.P2 = d.l1.P2, d.l1.P1
	if b := p.FromLineSegmentIntersection(&d.l1, &d.l2); b != d.onBoth || !d.p.infEqual(p) {
		t.Error("Vector2D.FromLineSegmentIntersection", d.l1, d.l2, "want", d.p, "got", *p, "want", d.onBoth, "got", b)
	}
	if b := p.FromLineSegmentIntersection(&d.l2, &d.l1); b != d.onBoth || !d.p.infEqual(p) {
		t.Error("Vector2D.FromLineSegmentIntersection", d.l2, d.l1, "want", d.p, "got", *p, "want", d.onBoth, "got", b)
	}
	d.l2.P1, d.l2.P2 = d.l2.P2, d.l2.P1
	if b := p.FromLineSegmentIntersection(&d.l1, &d.l2); b != d.onBoth || !d.p.infEqual(p) {
		t.Error("Vector2D.FromLineSegmentIntersection", d.l1, d.l2, "want", d.p, "got", *p, "want", d.onBoth, "got", b)
	}
	if b := p.FromLineSegmentIntersection(&d.l2, &d.l1); b != d.onBoth || !d.p.infEqual(p) {
		t.Error("Vector2D.FromLineSegmentIntersection", d.l2, d.l1, "want", d.p, "got", *p, "want", d.onBoth, "got", b)
	}
}

func TestLine2DSegmentIntersection(t *testing.T) {
	for _, v := range vector2DFromLineSegmentIntersectionValue {
		testVector2DFromLineSegmentIntersection(v, t)
	}
}

func Benchmark_Line2D_SegmentIntersection_Set(b *testing.B) {
	l1 := &Line2D{Vector2D{-1, -1}, Vector2D{0, 0}}
	l2 := &Line2D{Vector2D{0, 1}, Vector2D{1, 0}}
	p := &Vector2D{}
	for i := 0; i < b.N; i++ {
		p.FromLineSegmentIntersection(l1, l2)
	}
}
*/
