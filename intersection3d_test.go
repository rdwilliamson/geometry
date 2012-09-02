package geometry

import (
	"testing"
)

type intersection3DLineLineData struct {
	l1, l2, lb Line3D
}

var intersection3DLineLineValues = []intersection3DLineLineData{
	{Line3D{Vector3D{1, 2, 1}, Vector3D{3, 3, 3}},
		Line3D{Vector3D{1, 2, 1}, Vector3D{1, 2, 3}},
		Line3D{Vector3D{1, 2, 1}, Vector3D{1, 2, 1}}},
	{Line3D{Vector3D{1, 2, 1}, Vector3D{5, 4, 5}},
		Line3D{Vector3D{5, 6, 1}, Vector3D{1, 4, 5}},
		Line3D{Vector3D{3.4, 3.2, 3.4}, Vector3D{2.6, 4.8, 3.4}}},
	{Line3D{Vector3D{5, 4, 5}, Vector3D{3, 3, 3}},
		Line3D{Vector3D{1, 2, 5}, Vector3D{1, 2, 3}},
		Line3D{Vector3D{1, 2, 1}, Vector3D{1, 2, 1}}},
	{Line3D{Vector3D{1, 2, 1}, Vector3D{-3, 0, -3}},
		Line3D{Vector3D{5, 6, 1}, Vector3D{9, 8, -3}},
		Line3D{Vector3D{3.4, 3.2, 3.4}, Vector3D{2.6, 4.8, 3.4}}},
}

func testIntersection3DLineLine(d intersection3DLineLineData, t *testing.T) {
	var l Line3D
	if Intersection3DLineLine(&d.l1, &d.l2, &l); !l.SegmentFuzzyEqual(&d.lb) {
		t.Error("Intersection3D.LineLine", d.l1, d.l2, "want", d.lb, "got", l)
	}
}

func TestIntersection3DLineLine(t *testing.T) {
	for _, v := range intersection3DLineLineValues {
		testIntersection3DLineLine(v, t)
		v.l1.P1, v.l1.P2 = v.l1.P2, v.l1.P1
		testIntersection3DLineLine(v, t)
		v.l1, v.l2 = v.l2, v.l1
		testIntersection3DLineLine(v, t)
		v.l1.P1, v.l1.P2 = v.l1.P2, v.l1.P1
		testIntersection3DLineLine(v, t)
		v.l1, v.l2 = v.l2, v.l1
		testIntersection3DLineLine(v, t)
	}
}

func Benchmark_Intersection3D_LineLine(b *testing.B) {
	l1 := &Line3D{Vector3D{1, 2, 1}, Vector3D{5, 4, 5}}
	l2 := &Line3D{Vector3D{5, 6, 1}, Vector3D{1, 4, 5}}
	r := &Line3D{}
	for i := 0; i < b.N; i++ {
		Intersection3DLineLine(l1, l2, r)
	}
}

type intersection3DLineSegmentLineSegmentData struct {
	l1, l2, lb Line3D
}

var intersection3DLineSegmentLineSegmentValues = []intersection3DLineSegmentLineSegmentData{
	{Line3D{Vector3D{-1, 1, -1}, Vector3D{3, 3, 3}},
		Line3D{Vector3D{1, 2, -1}, Vector3D{1, 2, 3}},
		Line3D{Vector3D{1, 2, 1}, Vector3D{1, 2, 1}}},
	{Line3D{Vector3D{-1, 1, -1}, Vector3D{-3, 0, -3}},
		Line3D{Vector3D{1, 2, -1}, Vector3D{1, 2, 3}},
		Line3D{Vector3D{-1, 1, -1}, Vector3D{1, 2, 1}}},
	{Line3D{Vector3D{-1, 1, -1}, Vector3D{3, 3, 3}},
		Line3D{Vector3D{1, 2, -1}, Vector3D{1, 2, -3}},
		Line3D{Vector3D{1, 2, 1}, Vector3D{1, 2, -1}}},
	{Line3D{Vector3D{-1, 1, -1}, Vector3D{-3, 0, -3}},
		Line3D{Vector3D{1, 2, -1}, Vector3D{1, 2, -3}},
		Line3D{Vector3D{-1, 1, -1}, Vector3D{1, 2, -1}}},
	{Line3D{Vector3D{1, 2, 1}, Vector3D{5, 4, 5}},
		Line3D{Vector3D{5, 6, 1}, Vector3D{1, 4, 5}},
		Line3D{Vector3D{3.4, 3.2, 3.4}, Vector3D{2.6, 4.8, 3.4}}},
}

func testIntersection3DLineSegmentLineSegment(d intersection3DLineSegmentLineSegmentData, t *testing.T) {
	var l Line3D
	if Intersection3DLineSegmentLineSegment(&d.l1, &d.l2, &l); !l.SegmentFuzzyEqual(&d.lb) {
		t.Error("Intersection3D.LineSegmentLineSegment", d.l1, d.l2, "want",
			d.lb, "got", l)
	}
}

func TestIntersection3DLineSegmentLineSegment(t *testing.T) {
	for _, v := range intersection3DLineSegmentLineSegmentValues {
		testIntersection3DLineSegmentLineSegment(v, t)
		v.l1.P1, v.l1.P2 = v.l1.P2, v.l1.P1
		testIntersection3DLineSegmentLineSegment(v, t)
		v.l1, v.l2 = v.l2, v.l1
		testIntersection3DLineSegmentLineSegment(v, t)
		v.l1.P1, v.l1.P2 = v.l1.P2, v.l1.P1
		testIntersection3DLineSegmentLineSegment(v, t)
		v.l1, v.l2 = v.l2, v.l1
		testIntersection3DLineSegmentLineSegment(v, t)
	}
}

func Benchmark_Intersection3D_LineSegmentLineSegment(b *testing.B) {
	l1 := &Line3D{Vector3D{1, 2, 1}, Vector3D{5, 4, 5}}
	l2 := &Line3D{Vector3D{5, 6, 1}, Vector3D{1, 4, 5}}
	r := &Line3D{}
	for i := 0; i < b.N; i++ {
		Intersection3DLineSegmentLineSegment(l1, l2, r)
	}
}

func TestIntersection3DPlaneLine(t *testing.T) {
	p := &Plane{1, -2, 1, -7}
	l := &Line3D{Vector3D{5, -1, -1}, Vector3D{8, -2, 0}}
	want, got := &Vector3D{11.0 / 2.0, -7.0 / 6.0, -5.0 / 6.0}, &Vector3D{}
	if Intersection3DPlaneLine(p, l, got); !got.Equal(want) {
		t.Error("Intersection3D.PlaneLine", *p, *l, "want", want, "got", got)
	}
}

func Benchmark_Intersection3D_PlaneLine(b *testing.B) {
	p := &Plane{1, -2, 1, -7}
	l := &Line3D{Vector3D{5, -1, -1}, Vector3D{8, -2, 0}}
	r := &Vector3D{}
	for i := 0; i < b.N; i++ {
		Intersection3DPlaneLine(p, l, r)
	}
}

func TestIntersection3DPlanePlane(t *testing.T) {
	p1, p2 := &Plane{1, 1, 1, 1}, &Plane{1, 2, 3, 4}
	got, want := &Line3D{}, &Line3D{Vector3D{3, -5, 1}, Vector3D{4, -7, 2}}
	if Intersection3DPlanePlane(p1, p2, got); !got.FuzzyEqual(want) {
		t.Error("Intersection3D.PlanePlane", *p1, *p2, "want", want, "got", got)
	}
}

func Benchmark_Intersection3D_PlanePlane(b *testing.B) {
	p1, p2 := &Plane{1, 1, 1, 1}, &Plane{1, 2, 3, 4}
	l := &Line3D{}
	for i := 0; i < b.N; i++ {
		Intersection3DPlanePlane(p1, p2, l)
	}
}

func TestIntersection3DPlanePlanePlane(t *testing.T) {
	p1, p2, p3 := &Plane{1, -3, 3, 4}, &Plane{2, 3, -1, -15},
		&Plane{4, -3, -1, -19}
	got, want := &Vector3D{}, &Vector3D{5, 1, -2}
	if Intersection3DPlanePlanePlane(p1, p2, p3, got); !got.FuzzyEqual(want) {
		t.Error("Intersection3D.PlanePlanePlane", *p1, *p2, *p3, "want", want,
			"got", got)
	}
}

func Benchmark_Intersection3D_PlanePlanePlane(b *testing.B) {
	p1, p2, p3 := &Plane{1, -3, 3, 4}, &Plane{2, 3, -1, -15},
		&Plane{4, -3, -1, -19}
	pt := &Vector3D{}
	for i := 0; i < b.N; i++ {
		Intersection3DPlanePlanePlane(p1, p2, p3, pt)
	}
}

type intersection3DFuzzyPlaneLineData struct {
	pl Plane
	l  Line3D
	pt Vector3D
	n  int
}

var intersection3DFuzzyPlaneLineValues = []intersection3DFuzzyPlaneLineData{
	{Plane{1, -2, 1, -7}, Line3D{Vector3D{5, -1, -1}, Vector3D{8, -2, 0}},
		Vector3D{11.0 / 2.0, -7.0 / 6.0, -5.0 / 6.0}, 1},
	{Plane{1, 0, 0, 0}, Line3D{Vector3D{1, 0, 0}, Vector3D{1, 1, 0}},
		Vector3D{}, 0},
	{Plane{1, 0, 0, 0}, Line3D{Vector3D{0, 0, 0}, Vector3D{0, 1, 0}},
		Vector3D{}, -1},
}

func testIntersection3DFuzzyPlaneLine(d intersection3DFuzzyPlaneLineData, t *testing.T) {
	pt := &Vector3D{}
	got := Intersection3DFuzzyPlaneLine(&d.pl, &d.l, pt)
	if got != d.n {
		t.Error("Intersection3D.FuzzyPlaneLine", d.pl, d.l, "want", d.n, "got",
			got)
	} else if !d.pt.FuzzyEqual(pt) {
		t.Error("Intersection3D.FuzzyPlaneLine", d.pl, d.l, "want", d.pt,
			"got", *pt)
	}
	d.l.P1, d.l.P2 = d.l.P2, d.l.P1
	got = Intersection3DFuzzyPlaneLine(&d.pl, &d.l, pt)
	if got != d.n {
		t.Error("Intersection3D.FuzzyPlaneLine", d.pl, d.l, "want", d.n, "got",
			got)
	} else if !d.pt.FuzzyEqual(pt) {
		t.Error("Intersection3D.FuzzyPlaneLine", d.pl, d.l, "want", d.pt,
			"got", *pt)
	}
}

func TestIntersection3DFuzzyPlaneLine(t *testing.T) {
	for _, v := range intersection3DFuzzyPlaneLineValues {
		testIntersection3DFuzzyPlaneLine(v, t)
	}
}

func Benchmark_Intersection3D_FuzzyPlaneLine(b *testing.B) {
	p := &Plane{1, 2, 3, 4}
	l := &Line3D{Vector3D{1, 2, 3}, Vector3D{4, 5, 6}}
	pt := &Vector3D{}
	for i := 0; i < b.N; i++ {
		Intersection3DFuzzyPlaneLine(p, l, pt)
	}
}

type intersection3DFuzzyPlanePlaneData struct {
	p1, p2 Plane
	l      Line3D
	n      int
}

var intersection3DFuzzyPlanePlaneValue = []intersection3DFuzzyPlanePlaneData{
	{Plane{1, 0, 0, 1}, Plane{1, 0, 0, 1}, Line3D{Vector3D{}, Vector3D{}}, -1},
	{Plane{1, 0, 0, 1}, Plane{1, 0, 0, 2}, Line3D{Vector3D{}, Vector3D{}}, 0},
	{Plane{1, 1, 1, 1}, Plane{1, 2, 3, 4},
		Line3D{Vector3D{3, -5, 1}, Vector3D{4, -7, 2}}, 1},
	{Plane{1 + 1e-13, 0, 0, 1}, Plane{1, 0, 0, 1},
		Line3D{Vector3D{}, Vector3D{}}, -1},
}

func testIntersection3DFuzzyPlanePlane(d intersection3DFuzzyPlanePlaneData, t *testing.T) {
	var l Line3D
	if n := Intersection3DFuzzyPlanePlane(&d.p1, &d.p2, &l); n != d.n {
		t.Error("Intersection3D.FuzzyPlanePlane", d.p1, d.p2, "want", d.n,
			"got", n)
	}
}

func TestIntersection3DFuzzyPlanePlane(t *testing.T) {
	for _, v := range intersection3DFuzzyPlanePlaneValue {
		testIntersection3DFuzzyPlanePlane(v, t)
		v.p1, v.p2 = v.p2, v.p1
		testIntersection3DFuzzyPlanePlane(v, t)
	}
}

func Benchmark_Intersection3D_FuzzyPlanePlane(b *testing.B) {
	p1, p2 := &Plane{1, 2, 3, 4}, &Plane{5, 6, 7, 8}
	var l Line3D
	for i := 0; i < b.N; i++ {
		Intersection3DFuzzyPlanePlane(p1, p2, &l)
	}
}

func Benchmark_Intersection3D_FuzzyPlanePlane_Coincident(b *testing.B) {
	p1, p2 := &Plane{1, 0, 0, 1}, &Plane{1, 0, 0, 1}
	var l Line3D
	for i := 0; i < b.N; i++ {
		Intersection3DFuzzyPlanePlane(p1, p2, &l)
	}
}

type intersection3DFuzzyPlanePlanePlaneData struct {
	p1, p2, p3 Plane
	p          Vector3D
	n          int
}

var intersection3DFuzzyPlanePlanePlaneValues = []intersection3DFuzzyPlanePlanePlaneData{
	{Plane{1, -3, 3, 4}, Plane{2, 3, -1, -15}, Plane{4, -3, -1, -19},
		Vector3D{5, 1, -2}, 1},
	{Plane{1, 0, 0, 0}, Plane{1, 0, 0, 1}, Plane{1, 0, 0, 2}, Vector3D{}, 0},
	{Plane{1, 0, 0, 0}, Plane{1, 0, 0, 0}, Plane{1, 0, 0, 0}, Vector3D{}, -1},
	{Plane{1, 0, 0, 0}, Plane{0, 1, 0, 0}, Plane{1, 1, 0, 0}, Vector3D{}, -2},
	{Plane{1, 0, 0, 0}, Plane{1, 0, 0, 1}, Plane{1, 1, 0, 0}, Vector3D{}, -3},
	{Plane{-1, 0, 0, 0}, Plane{1, 0, 0, 1}, Plane{1, 1, 0, 0}, Vector3D{}, -3},
}

func testIntersection3DFuzzyPlanePlanePlane(d intersection3DFuzzyPlanePlanePlaneData, t *testing.T) {
	var p Vector3D
	n := Intersection3DFuzzyPlanePlanePlane(&d.p1, &d.p2, &d.p3, &p)
	if n != d.n {
		t.Error("Intersection3DFuzzyPlanePlanePlane", d.p1, d.p2, d.p3, "want",
			d.n, "got", n)
	} else if !p.FuzzyEqual(&d.p) {
		t.Error("Intersection3DFuzzyPlanePlanePlane", d.p1, d.p2, d.p3, "want",
			d.p, "got", p)
	}
}

func TestIntersection3DFuzzyPlanePlanePlane(t *testing.T) {
	for _, v := range intersection3DFuzzyPlanePlanePlaneValues {
		testIntersection3DFuzzyPlanePlanePlane(v, t)
		v.p2, v.p3 = v.p3, v.p2
		testIntersection3DFuzzyPlanePlanePlane(v, t)
		v.p1, v.p3 = v.p3, v.p1
		testIntersection3DFuzzyPlanePlanePlane(v, t)
		v.p2, v.p3 = v.p3, v.p2
		testIntersection3DFuzzyPlanePlanePlane(v, t)
		v.p1, v.p3 = v.p3, v.p1
		testIntersection3DFuzzyPlanePlanePlane(v, t)
		v.p2, v.p3 = v.p3, v.p2
		testIntersection3DFuzzyPlanePlanePlane(v, t)
	}
}

func Benchmark_Intersection3D_FuzzyPlanePlanePlane(b *testing.B) {
	p1, p2, p3 := &Plane{1, -3, 3, 4}, &Plane{2, 3, -1, -15},
		&Plane{4, -3, -1, -19}
	pt := &Vector3D{}
	for i := 0; i < b.N; i++ {
		Intersection3DFuzzyPlanePlanePlane(p1, p2, p3, pt)
	}
}
