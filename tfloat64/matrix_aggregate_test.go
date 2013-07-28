
package tfloat64

import (
	"testing"
	"math"
)

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

type aggregateProcedureMatrix interface {
	MatrixData
	AggregateProcedure(aggr Float64Float64Func, f Float64Func, cond Float64Procedure) float64
}

func TestDenseMatrixAggregateProcedure(t *testing.T) {
	A := makeDenseMatrix()
	testMatrixAggregateProcedure(t, A)
}

func TestSparseMatrixAggregateProcedure(t *testing.T) {
	A := makeSparseMatrix()
	testMatrixAggregateProcedure(t, A)
}

func testMatrixAggregateProcedure(t *testing.T, A aggregateProcedureMatrix) {
	procedure := func(element float64) bool {
		if math.Abs(element) > 0.2 {
			return true
		} else {
			return false
		}
	}
	expected := 0.0
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Columns(); c++ {
			elem := A.GetQuick(r, c)
			if math.Abs(elem) > 0.2 {
				expected += elem * elem
			}
		}
	}

	result := A.AggregateProcedure(Plus, Square, procedure)
	if math.Abs(expected - result) > tol {
		t.Errorf("expected:%g actual:%g", expected, result)
	}
}

type aggregateProcedureSelectionMatrix interface {
MatrixData
	AggregateProcedureSelection(aggr Float64Float64Func, f Float64Func, rowList, columnList []int) float64
}

func TestDenseMatrixAggregateProcedureSelection(t *testing.T) {
	A := makeDenseMatrix()
	testMatrixAggregateProcedureSelection(t, A)
}

func TestSparseMatrixAggregateProcedureSelection(t *testing.T) {
	A := makeSparseMatrix()
	testMatrixAggregateProcedureSelection(t, A)
}

func testMatrixAggregateProcedureSelection(t *testing.T, A aggregateProcedureSelectionMatrix) {
	var rowList []int
	var columnList []int
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Columns(); c++ {
			rowList = append(rowList, r)
			columnList = append(columnList, c)
		}
	}
	expected := 0.0
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Columns(); c++ {
			elem := A.GetQuick(r, c)
			expected += elem * elem
		}
	}
	result := A.AggregateProcedureSelection(Plus, Square, rowList, columnList)
	if math.Abs(expected - result) > tol {
		t.Errorf("expected:%g actual:%g", expected, result)
	}
}

type aggregateMatrixMatrix interface {
MatrixData
	AggregateMatrix(other MatrixData, aggr Float64Float64Func, f Float64Float64Func) (float64, error)
}

func TestDenseMatrixAggregateMatrix(t *testing.T) {
	A := makeDenseMatrix()
	B := makeDenseMatrix()
	testMatrixAggregateMatrix(t, A, B)
}

func TestSparseMatrixAggregateMatrix(t *testing.T) {
	A := makeSparseMatrix()
	B := makeSparseMatrix()
	testMatrixAggregateMatrix(t, A, B)
}

func testMatrixAggregateMatrix(t *testing.T, A, B aggregateMatrixMatrix) {
	expected := 0.0
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Columns(); c++ {
			elemA := A.GetQuick(r, c)
			elemB := B.GetQuick(r, c)
			expected += elemA * elemB
		}
	}
	result, _ := A.AggregateMatrix(B, Plus, Mult)
	if math.Abs(expected - result) > tol {
		t.Errorf("expected:%g actual:%g", expected, result)
	}
}
