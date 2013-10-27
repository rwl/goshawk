
package tfloat64

import (
	"testing"
	"math/rand"
)

func makeSparseMatrix() (*Matrix) {
	A := NewSparseMatrix(nrows, ncolumns)
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Columns(); c++ {
			A.SetQuick(r, c, rand.Float64())
		}
	}
	return A
}

// Aggregate tests.

func TestSparseMatrixAggregate(t *testing.T) {
	A := makeSparseMatrix()
	testMatrixAggregate(t, A)
}

func TestSparseMatrixAggregateProcedure(t *testing.T) {
	A := makeSparseMatrix()
	testMatrixAggregateProcedure(t, A)
}

func TestSparseMatrixAggregateProcedureSelection(t *testing.T) {
	A := makeSparseMatrix()
	testMatrixAggregateProcedureSelection(t, A)
}

func TestSparseMatrixAggregateMatrix(t *testing.T) {
	A := makeSparseMatrix()
	B := makeSparseMatrix()
	testMatrixAggregateMatrix(t, A, B)
}

// Assign tests.

func TestSparseMatrixAssign(t *testing.T) {
	A := makeSparseMatrix()
	testMatrixAssign(t, A)
}

func TestSparseMatrixAssignArray(t *testing.T) {
	A := makeSparseMatrix()
	testMatrixAssignArray(t, A)
}

func TestSparseMatrixAssignFunc(t *testing.T) {
	A := makeSparseMatrix()
	testMatrixAssignFunc(t, A)
}

func TestSparseMatrixAssignMatrix(t *testing.T) {
	A := makeSparseMatrix()
	B := makeSparseMatrix()
	testMatrixAssignMatrix(t, A, B)
}

func TestSparseMatrixAssignMatrixFunc(t *testing.T) {
	A := makeSparseMatrix()
	B := makeSparseMatrix()
	testMatrixAssignMatrixFunc(t, A, B)
}

func TestSparseMatrixAssignMatrixFuncSelection(t *testing.T) {
	A := makeSparseMatrix()
	B := makeSparseMatrix()
	testMatrixAssignMatrixFuncSelection(t, A, B)
}

func TestSparseMatrixAssignProcedure(t *testing.T) {
	A := makeSparseMatrix()
	testMatrixAssignProcedure(t, A)
}

func TestSparseMatrixAssignProcedureFunc(t *testing.T) {
	A := makeSparseMatrix()
	testMatrixAssignProcedureFunc(t, A)
}
