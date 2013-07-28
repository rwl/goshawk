
package tfloat64

import (
	"math/rand"
	"testing"
	"math"
)

const (
	nrows = 13
	ncolumns = 17
)

func makeDenseMatrix() (*Matrix) {
	A := NewMatrix(nrows, ncolumns)
	makeMatrixRandom(A)
	return A
}

func makeSparseMatrix() (*Matrix) {
	A := NewSparseMatrix(nrows, ncolumns)
	makeMatrixRandom(A)
	return A
}

func makeMatrixRandom(A *Matrix) {
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Columns(); c++ {
			A.SetQuick(r, c, rand.Float64())
		}
	}
}

type aggregateMatrix interface {
	MatrixData
	Aggregate(aggr Float64Float64Func, f Float64Func) float64
}

func TestDenseMatrixAggregate(t *testing.T) {
	A := makeDenseMatrix()
	testMatrixAggregate(t, A)
}

func TestSparseMatrixAggregate(t *testing.T) {
	A := makeSparseMatrix()
	testMatrixAggregate(t, A)
}

func testMatrixAggregate(t *testing.T, A aggregateMatrix) {
	expected := 0.0
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Columns(); c++ {
			elem := A.GetQuick(r, c)
			expected += elem * elem
		}
	}
	result := A.Aggregate(Plus, Square)
	if math.Abs(expected - result) > tol {
		t.Errorf("expected:%g actual:%g", expected, result)
	}
}
