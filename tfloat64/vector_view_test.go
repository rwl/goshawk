
package tfloat64

import (
	"testing"
	"math"
)

type viewFlipVector interface {
	VectorData
	ViewFlip() *Vector
}

func TestDenseViewFlip(t *testing.T) {
	A := makeDenseVector()
	testViewFlip(t, A)
}

func TestSparseViewFlip(t *testing.T) {
	A := makeSparseVector()
	testViewFlip(t, A)
}

func testViewFlip(t *testing.T, A viewFlipVector) {
	b := A.ViewFlip()
	if A.Size() != b.Size() {
		t.Errorf("expected:%d actual:%d", A.Size(), b.Size())
	}
	for i := 0; i < A.Size(); i++ {
		expected := A.GetQuick(i)
		result := b.GetQuick(A.Size() - 1 - i)
		if math.Abs(expected - result) > tol {
			t.Errorf("expected:%g actual:%g", expected, result)
		}
	}
}

type viewPartVector interface {
	VectorData
	ViewPart(int, int) *Vector
}

func TestDenseViewPart(t *testing.T) {
	A := makeDenseVector()
	testViewPart(t, A)
}

func TestSparseViewPart(t *testing.T) {
	A := makeSparseVector()
	testViewPart(t, A)
}

func testViewPart(t *testing.T, A viewPartVector) {
	b := A.ViewPart(15, 11)
	for i := 0; i < 11; i++ {
		expected := A.GetQuick(15 + i)
		result := b.GetQuick(i)
		if math.Abs(expected - result) > tol {
			t.Errorf("expected:%g actual:%g", expected, result)
		}
	}
}

type viewProcedureVector interface {
	VectorData
	ViewProcedure(Float64Procedure) *Vector
}

func TestDenseViewProcedure(t *testing.T) {
	A := makeDenseVector()
	testViewProcedure(t, A)
}

func TestSparseViewProcedure(t *testing.T) {
	A := makeSparseVector()
	testViewProcedure(t, A)
}

func testViewProcedure(t *testing.T, A viewProcedureVector) {
	b := A.ViewProcedure(func(element float64) bool {
		return math.Remainder(element, 2) == 0
	})
	for i := 0; i < b.Size(); i++ {
		el := b.GetQuick(i)
		if math.Remainder(el, 2) != 0 {
			t.Fail()
		}
	}
}

type viewVector interface {
VectorData
	View([]int) (*Vector, error)
}

func TestDenseView(t *testing.T) {
	A := makeDenseVector()
	testView(t, A)
}

func TestSparseView(t *testing.T) {
	A := makeSparseVector()
	testView(t, A)
}

func testView(t *testing.T, A viewVector) {
	indexes := []int { 5, 11, 22, 37, 101 }
	b, _ := A.View(indexes)
	for i := 0; i < len(indexes); i++ {
		expected := A.GetQuick(indexes[i])
		result := b.GetQuick(i)
		if math.Abs(expected - result) > tol {
			t.Errorf("expected:%g actual:%g", expected, result)
		}
	}
}

/*type viewSortedVector interface { TODO: implement sort
	VectorData
	ViewSorted() *Vector
}

func TestDenseViewSorted(t *testing.T) {
	A := makeDenseVector()
	testViewSorted(t, A)
}

func TestSparseViewSorted(t *testing.T) {
	A := makeSparseVector()
	testViewSorted(t, A)
}

func testViewSorted(t *testing.T, A viewSortedVector) {
	b := A.ViewSorted()
	for i := 0; i < A.Size() - 1; i++ {
		if b.GetQuick(i + 1) < b.GetQuick(i) {
			t.Errorf("%g < %g", b.GetQuick(i + 1), b.GetQuick(i))
		}
	}
}*/

type viewStridesVector interface {
VectorData
	ViewStrides(int) *Vector
}

func TestDenseViewStrides(t *testing.T) {
	A := makeDenseVector()
	testViewStrides(t, A)
}

func TestSparseViewStrides(t *testing.T) {
	A := makeSparseVector()
	testViewStrides(t, A)
}

func testViewStrides(t *testing.T, A viewStridesVector) {
	stride := 3
	b := A.ViewStrides(stride)
	for i := 0; i < b.Size(); i++ {
		expected := A.GetQuick(i * stride)
		result := b.GetQuick(i)
		if math.Abs(expected - result) > tol {
			t.Errorf("expected:%g actual:%g", expected, result)
		}
	}
}
