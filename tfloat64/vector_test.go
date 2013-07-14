
package tfloat64

import (
	"testing"
	"math/rand"
	"math"
)

const (
	test_size = 2 * 17 * 5
	tol = 1e-10
)

func makeDenseVectors() (*Vector, *Vector) {
	A := NewVector(test_size)
	B := NewVector(test_size)

	for i := 0; i < test_size; i++ {
		A.SetQuick(i, rand.Float64())
		B.SetQuick(i, rand.Float64())
	}
	return A, B
}

func makeSparseVectors() (*Vector, *Vector) {
	A := NewSparseVector(test_size)
	B := NewSparseVector(test_size)

	for i := 0; i < test_size; i++ {
		A.SetQuick(i, rand.Float64())
		B.SetQuick(i, rand.Float64())
	}
	return A, B
}

type aggregator interface {
	Size() int
	GetQuick(int) float64
	Aggregate(Float64Float64Func, Float64Func) float64
}

func TestDenseAggregate(t *testing.T) {
	A, _ := makeDenseVectors()
	testAggregate(t, A)
}

func TestSparseAggregate(t *testing.T) {
	A, _ := makeSparseVectors()
	testAggregate(t, A)
}

func testAggregate(t *testing.T, A aggregator) {
	expected := 0.0
	for i := 0; i < A.Size(); i++ {
		elem := A.GetQuick(i)
		expected += elem * elem
	}

	result := A.Aggregate(Plus, Square)
	if math.Abs(expected - result) > tol {
		t.Errorf("expected:%g actual:%g", expected, result)
	}
}
