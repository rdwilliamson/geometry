package geometry

import (
	"math"
	"testing"
)

// Treats converts -inf to +inf to simplify test code.
func (a *Vector2D) infEqual(b *Vector2D) bool {
	if math.IsInf(a.X, -1) {
		a.X = math.Inf(1)
	}
	if math.IsInf(a.Y, -1) {
		a.Y = math.Inf(1)
	}
	if math.IsInf(b.X, -1) {
		b.X = math.Inf(1)
	}
	if math.IsInf(b.Y, -1) {
		b.Y = math.Inf(1)
	}
	return a.Equal(b)
}

func TestNewVector2D(t *testing.T) {
	if !NewVector2D(1, 2).Equal(&Vector2D{1, 2}) {
		t.Error("NewVector2D")
	}
}

func Benchmark_Vector2D_New(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewVector2D(0, 0)
	}
}

func TestVector2DAdd(t *testing.T) {
	r, v1, v2 := &Vector2D{}, &Vector2D{1, 2}, &Vector2D{3, 4}
	if !r.Add(v1, v2).Equal(&Vector2D{4, 6}) {
		t.Error("Vector2D.Add")
	}
}

func Benchmark_Vector2D_Add(b *testing.B) {
	r, v1, v2 := &Vector2D{}, &Vector2D{1, 2}, &Vector2D{3, 4}
	for i := 0; i < b.N; i++ {
		r.Add(v1, v2)
	}
}

func TestVector2DAngularDifference(t *testing.T) {
	v1, v2 := &Vector2D{1, 0}, &Vector2D{0, 1}
	if v1.AngularDifference(v2) != math.Pi/2 {
		t.Error("Vector2D.AngularDifference")
	}
}

func Benchmark_Vector2D_AngularDifference(b *testing.B) {
	v1, v2 := &Vector2D{1, 2}, &Vector2D{3, 4}
	for i := 0; i < b.N; i++ {
		v1.AngularDifference(v2)
	}
}

func TestVector2DAngularDifferenceCosSquared(t *testing.T) {
	v1, v2 := &Vector2D{1, 0}, &Vector2D{0, 1}
	if FuzzyEqual(v1.AngularDifferenceCosSquared(v2), math.Sqrt2/2) {
		t.Error("Vector2D.AngularDifferenceCosSquared")
	}
}

func Benchmark_Vector2D_AngularDifferenceCosSquared(b *testing.B) {
	v1, v2 := &Vector2D{1, 2}, &Vector2D{3, 4}
	for i := 0; i < b.N; i++ {
		v1.AngularDifferenceCosSquared(v2)
	}
}

func TestVector2DCopy(t *testing.T) {
	v1, v2 := &Vector2D{}, &Vector2D{1, 0}
	if !v1.Copy(v2).Equal(v2) {
		t.Error("Vector2D.Copy")
	}
}

func Benchmark_Vector2D_Copy(b *testing.B) {
	v1, v2 := &Vector2D{}, &Vector2D{1, 0}
	for i := 0; i < b.N; i++ {
		v1.Copy(v2)
	}
}

func TestVector2DDirectionEqual(t *testing.T) {
	v1, v2 := &Vector2D{1, 1}, &Vector2D{2, 2}
	if !v1.DirectionEqual(v2) {
		t.Error("Vector2D.DirectionEqual")
	}
	if !v2.DirectionEqual(v1) {
		t.Error("Vector2D.DirectionEqual")
	}
	*v2 = Vector2D{-2, -2}
	if v1.DirectionEqual(v2) {
		t.Error("Vector2D.DirectionEqual")
	}
	if v2.DirectionEqual(v1) {
		t.Error("Vector2D.DirectionEqual")
	}
}

func Benchmark_Vector2D_DirectionEqual(b *testing.B) {
	v1, v2 := &Vector2D{1, 1}, &Vector2D{2, 2}
	for i := 0; i < b.N; i++ {
		v1.DirectionEqual(v2)
	}
}

func TestVector2DFromLine(t *testing.T) {
	l, v := &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}, &Vector2D{}
	if !v.FromLine(l).Equal(&Vector2D{2, 2}) {
		t.Error("Vector2D.FromLine")
	}
}

func Benchmark_Vector2D_FromLine(b *testing.B) {
	l, v := &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}, &Vector2D{}
	for i := 0; i < b.N; i++ {
		v.FromLine(l)
	}
}

type vector2DFromLineIntersectionData struct {
	l1, l2 Line2D
	p      Vector2D
}

var vector2DFromLineIntersectionValues = []vector2DFromLineIntersectionData{
	{Line2D{Vector2D{0, 0}, Vector2D{1, 1}}, Line2D{Vector2D{0, 1}, Vector2D{1, 0}}, Vector2D{0.5, 0.5}},
	{Line2D{Vector2D{0, 0}, Vector2D{1, 1}}, Line2D{Vector2D{1, 0}, Vector2D{2, 1}}, Vector2D{math.Inf(1), math.Inf(1)}},
	{Line2D{Vector2D{0, 0}, Vector2D{-1, -1}}, Line2D{Vector2D{0, 1}, Vector2D{1, 0}}, Vector2D{0.5, 0.5}},
}

func testVector2DFromLineIntersection(d vector2DFromLineIntersectionData, t *testing.T) {
	var p Vector2D
	if !p.FromLineIntersection(&d.l1, &d.l2).infEqual(&d.p) {
		t.Error("Vector2D.FromLineIntersection", d.l1, d.l2, "want", d.p, "got", p)
	}
	if !p.FromLineIntersection(&d.l2, &d.l1).infEqual(&d.p) {
		t.Error("Vector2D.FromLineIntersection", d.l2, d.l1, "want", d.p, "got", p)
	}
	d.l1.P1, d.l1.P2 = d.l1.P2, d.l1.P1
	if !p.FromLineIntersection(&d.l1, &d.l2).infEqual(&d.p) {
		t.Error("Vector2D.FromLineIntersection", d.l1, d.l2, "want", d.p, "got", p)
	}
	if !p.FromLineIntersection(&d.l2, &d.l1).infEqual(&d.p) {
		t.Error("Vector2D.FromLineIntersection", d.l2, d.l1, "want", d.p, "got", p)
	}
	d.l2.P1, d.l2.P2 = d.l2.P2, d.l2.P1
	if !p.FromLineIntersection(&d.l1, &d.l2).infEqual(&d.p) {
		t.Error("Vector2D.FromLineIntersection", d.l1, d.l2, "want", d.p, "got", p)
	}
	if !p.FromLineIntersection(&d.l2, &d.l1).infEqual(&d.p) {
		t.Error("Vector2D.FromLineIntersection", d.l2, d.l1, "want", d.p, "got", p)
	}
}

func TestVector2DFromLineIntersection(t *testing.T) {
	for _, v := range vector2DFromLineIntersectionValues {
		testVector2DFromLineIntersection(v, t)
	}
}

func Benchmark_Vector2D_FromLineIntersection(b *testing.B) {
	l1 := &Line2D{Vector2D{0, 0}, Vector2D{1, 1}}
	l2 := &Line2D{Vector2D{0, 1}, Vector2D{1, 0}}
	p := &Vector2D{}
	for i := 0; i < b.N; i++ {
		p.FromLineIntersection(l1, l2)
	}
}

func TestVector2DFromLineNormal(t *testing.T) {
	l, v := &Line2D{Vector2D{1, 1}, Vector2D{3, 1}}, &Vector2D{}
	if !v.FromLineNormal(l).Equal(&Vector2D{0, -2}) {
		t.Error("Vector2D.FromLineNormal", v)
	}
}

func Benchmark_Vector2D_FromLineNormal(b *testing.B) {
	l, v := &Line2D{Vector2D{1, 1}, Vector2D{3, 1}}, &Vector2D{}
	for i := 0; i < b.N; i++ {
		v.FromLineNormal(l)
	}
}

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

type vector2DDirectionFuzzyEqualData struct {
	v1, v2 Vector2D
	equal  bool
}

var vector2DDirectionFuzzyEqualValues = []vector2DDirectionFuzzyEqualData{
	{Vector2D{1, 1}, Vector2D{2, 2 + 2e-12}, false},
	{Vector2D{1, 1}, Vector2D{2 + 2e-12, 2}, false},
	{Vector2D{1, 1}, Vector2D{2, 2 + 2e-13}, true},
	{Vector2D{1, 1}, Vector2D{2 + 2e-13, 2}, true},
}

func testVector2DDirectionFuzzyEqual(d vector2DDirectionFuzzyEqualData, t *testing.T) {
	if d.v1.DirectionFuzzyEqual(&d.v2) != d.equal {
		t.Error("Vector2D.DirectionFuzzyEqual:", d.v1, d.v2, "want", d.equal, "got", !d.equal)
	}
	if d.v2.DirectionFuzzyEqual(&d.v1) != d.equal {
		t.Error("Vector2D.DirectionFuzzyEqual:", d.v2, d.v1, "want", d.equal, "got", !d.equal)
	}
}

func TestVector2DDirectionFuzzyEqual(t *testing.T) {
	for _, v := range vector2DDirectionFuzzyEqualValues {
		testVector2DDirectionFuzzyEqual(v, t)
	}
}

func Benchmark_Vector2D_DirectionFuzzyEqual(b *testing.B) {
	v1, v2 := &Vector2D{1, 1}, &Vector2D{2, 2}
	for i := 0; i < b.N; i++ {
		v1.DirectionFuzzyEqual(v2)
	}
}

func TestVector2DDistance(t *testing.T) {
	v1, v2 := &Vector2D{}, &Vector2D{1, 0}
	if v1.Distance(v2) != 1 {
		t.Error("Vector2D.Distance")
	}
}

func Benchmark_Vector2D_Distance(b *testing.B) {
	v1, v2 := &Vector2D{}, &Vector2D{1, 0}
	for i := 0; i < b.N; i++ {
		v1.Distance(v2)
	}
}

func TestVector2DDistanceSquared(t *testing.T) {
	v1, v2 := &Vector2D{}, &Vector2D{1, 0}
	if v1.DistanceSquared(v2) != 1 {
		t.Error("Vector2D.DistanceSquared")
	}
}

func Benchmark_Vector2D_DistanceSquared(b *testing.B) {
	v1, v2 := &Vector2D{}, &Vector2D{1, 0}
	for i := 0; i < b.N; i++ {
		v1.DistanceSquared(v2)
	}
}

func TestVector2DDotProduct(t *testing.T) {
	v1, v2 := &Vector2D{1, 2}, &Vector2D{3, 4}
	if v1.DotProduct(v2) != 11 {
		t.Error("Vector2D.DotProduct")
	}
}

func Benchmark_Vector2D_DotProduct(b *testing.B) {
	v1, v2 := &Vector2D{1, 2}, &Vector2D{3, 4}
	for i := 0; i < b.N; i++ {
		v1.DotProduct(v2)
	}
}

func TestVector2DEqual(t *testing.T) {
	v1, v2 := &Vector2D{}, &Vector2D{1, 0}
	if v1.Equal(v2) {
		t.Error("Vector2D.Equal")
	}
	v2.X = 0
	if !v1.Equal(v2) {
		t.Error("Vector2D.Equal")
	}
}

func Benchmark_Vector2D_Equal(b *testing.B) {
	v1, v2 := &Vector2D{}, &Vector2D{1, 0}
	for i := 0; i < b.N; i++ {
		v1.Equal(v2)
	}
}

type fuzzyEqualData struct {
	v1, v2 Vector2D
	equal  bool
}

var fuzzyEqualValues = []fuzzyEqualData{
	{Vector2D{}, Vector2D{0, 1e-12}, false},
	{Vector2D{}, Vector2D{0, 1e-13}, true},
}

func testVector2DFuzzyEqual(d fuzzyEqualData, t *testing.T) {
	if d.v1.FuzzyEqual(&d.v2) != d.equal {
		t.Error("Vector2D.FuzzyEqual", d.v1, d.v2, d.equal)
	}
}

func TestVector2DFuzzyEqual(t *testing.T) {
	for _, v := range fuzzyEqualValues {
		testVector2DFuzzyEqual(v, t)
	}
}

func Benchmark_Vector2D_FuzzyEqual(b *testing.B) {
	v1, v2 := &Vector2D{}, &Vector2D{1, 0}
	for i := 0; i < b.N; i++ {
		v1.FuzzyEqual(v2)
	}
}

func TestVector2DMagnitude(t *testing.T) {
	v := &Vector2D{3, 4}
	if v.Magnitude() != 5 {
		t.Error("Vector2D.Magnitude")
	}
}

func Benchmark_Vector2D_Magnitude(b *testing.B) {
	v := &Vector2D{3, 4}
	for i := 0; i < b.N; i++ {
		v.Magnitude()
	}
}

func TestVector2DMagnitudeSquared(t *testing.T) {
	v := &Vector2D{3, 4}
	if v.MagnitudeSquared() != 25 {
		t.Error("Vector2D.MagnitudeSquared")
	}
}

func Benchmark_Vector2D_MagnitudeSquared(b *testing.B) {
	v := &Vector2D{3, 4}
	for i := 0; i < b.N; i++ {
		v.MagnitudeSquared()
	}
}

func TestVector2DNormalize(t *testing.T) {
	v1, v2 := &Vector2D{3, 4}, &Vector2D{3.0 / 5.0, 4.0 / 5.0}
	if !v1.Normalize(v1).FuzzyEqual(v2) {
		t.Error("Vector2D.Normalize")
	}
}

func Benchmark_Vector2D_Normalize(b *testing.B) {
	v1, v2 := &Vector2D{3, 4}, &Vector2D{}
	for i := 0; i < b.N; i++ {
		v1.Normalize(v2)
	}
}

func TestVector2DSubtract(t *testing.T) {
	r, v1, v2 := &Vector2D{}, &Vector2D{1, 2}, &Vector2D{3, 4}
	if !r.Subtract(v1, v2).Equal(&Vector2D{-2, -2}) {
		t.Error("Vector2D.Subtract")
	}
}

func Benchmark_Vector2D_Subtract(b *testing.B) {
	r, v1, v2 := &Vector2D{}, &Vector2D{1, 2}, &Vector2D{3, 4}
	for i := 0; i < b.N; i++ {
		r.Subtract(v1, v2)
	}
}
