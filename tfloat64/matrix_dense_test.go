
package tfloat64

import (
	"testing"
	"math/rand"
)

const (
	nrows    = 13
	ncolumns = 17
)

func makeDenseMatrix() (*Matrix) {
	A := NewMatrix(nrows, ncolumns)
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Columns(); c++ {
			A.SetQuick(r, c, rand.Float64())
		}
	}
	return A
}

// Aggregate tests.

func TestDenseMatrixAggregate(t *testing.T) {
	A := makeDenseMatrix()
	testMatrixAggregate(t, A)
}

func TestDenseMatrixAggregateProcedure(t *testing.T) {
	A := makeDenseMatrix()
	testMatrixAggregateProcedure(t, A)
}

func TestDenseMatrixAggregateProcedureSelection(t *testing.T) {
	A := makeDenseMatrix()
	testMatrixAggregateProcedureSelection(t, A)
}

func TestDenseMatrixAggregateMatrix(t *testing.T) {
	A := makeDenseMatrix()
	B := makeDenseMatrix()
	testMatrixAggregateMatrix(t, A, B)
}

// Assign tests.

func TestDenseMatrixAssign(t *testing.T) {
	A := makeDenseMatrix()
	testMatrixAssign(t, A)
}

func TestDenseMatrixAssignArray(t *testing.T) {
	A := makeDenseMatrix()
	testMatrixAssignArray(t, A)
}

func TestDenseMatrixAssignFunc(t *testing.T) {
	A := makeDenseMatrix()
	testMatrixAssignFunc(t, A)
}

func TestDenseMatrixAssignMatrix(t *testing.T) {
	A := makeDenseMatrix()
	B := makeDenseMatrix()
	testMatrixAssignMatrix(t, A, B)
}

func TestDenseMatrixAssignMatrixFunc(t *testing.T) {
	A := makeDenseMatrix()
	B := makeDenseMatrix()
	testMatrixAssignMatrixFunc(t, A, B)
}

func TestDenseMatrixAssignMatrixFuncSelection(t *testing.T) {
	A := makeDenseMatrix()
	B := makeDenseMatrix()
	testMatrixAssignMatrixFuncSelection(t, A, B)
}

func TestDenseMatrixAssignProcedure(t *testing.T) {
	A := makeDenseMatrix()
	testMatrixAssignProcedure(t, A)
}

func TestDenseMatrixAssignProcedureFunc(t *testing.T) {
	A := makeDenseMatrix()
	testMatrixAssignProcedureFunc(t, A)
}
