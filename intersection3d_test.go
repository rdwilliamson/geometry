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

func testLine3DFromLineBetween(d intersection3DLineLineData, t *testing.T) {
	var l Line3D
	if Intersection3DLineLine(&d.l1, &d.l2, &l); !l.SegmentFuzzyEqual(&d.lb) {
		t.Error("Line3D.FromLineBetween", d.l1, d.l2, "want", d.lb, "got", l)
	}
}

func TestLine3DFromLineBetween(t *testing.T) {
	for _, v := range intersection3DLineLineValues {
		testLine3DFromLineBetween(v, t)
		v.l1.P1, v.l1.P2 = v.l1.P2, v.l1.P1
		testLine3DFromLineBetween(v, t)
		v.l1, v.l2 = v.l2, v.l1
		testLine3DFromLineBetween(v, t)
		v.l1.P1, v.l1.P2 = v.l1.P2, v.l1.P1
		testLine3DFromLineBetween(v, t)
		v.l1, v.l2 = v.l2, v.l1
		testLine3DFromLineBetween(v, t)
	}
}

func Benchmark_Line3D_FromLineBetween(b *testing.B) {
	l1 := &Line3D{Vector3D{1, 2, 1}, Vector3D{5, 4, 5}}
	l2 := &Line3D{Vector3D{5, 6, 1}, Vector3D{1, 4, 5}}
	r := &Line3D{}
	for i := 0; i < b.N; i++ {
		Intersection3DLineLine(l1, l2, r)
	}
}
