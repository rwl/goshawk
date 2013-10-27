package tfloat64

import (
	"testing"
	"math"
)

type aggregateMatrix interface {
	Mat
	Aggregate(aggr Float64Float64Func, f Float64Func) float64
}

func testMatrixAggregate(t *testing.T, A aggregateMatrix) {
	expected := 0.0
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Columns(); c++ {
			elem := A.GetQuick(r, c)
			expected += elem*elem
		}
	}
	result := A.Aggregate(Plus, Square)
	if math.Abs(expected - result) > tol {
		t.Errorf("expected:%g actual:%g", expected, result)
	}
}

type aggregateProcedureMatrix interface {
	Mat
	AggregateProcedure(aggr Float64Float64Func, f Float64Func, cond Float64Procedure) float64
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
				expected += elem*elem
			}
		}
	}

	result := A.AggregateProcedure(Plus, Square, procedure)
	if math.Abs(expected - result) > tol {
		t.Errorf("expected:%g actual:%g", expected, result)
	}
}

type aggregateProcedureSelectionMatrix interface {
	Mat
	AggregateProcedureSelection(aggr Float64Float64Func, f Float64Func, rowList, columnList []int) float64
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
			expected += elem*elem
		}
	}
	result := A.AggregateProcedureSelection(Plus, Square, rowList, columnList)
	if math.Abs(expected - result) > tol {
		t.Errorf("expected:%g actual:%g", expected, result)
	}
}

type aggregateMatrixMatrix interface {
	Mat
	AggregateMatrix(other Mat, aggr Float64Float64Func, f Float64Float64Func) (float64, error)
}

func testMatrixAggregateMatrix(t *testing.T, A, B aggregateMatrixMatrix) {
	expected := 0.0
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Columns(); c++ {
			elemA := A.GetQuick(r, c)
			elemB := B.GetQuick(r, c)
			expected += elemA*elemB
		}
	}
	result, _ := A.AggregateMatrix(B, Plus, Mult)
	if math.Abs(expected - result) > tol {
		t.Errorf("expected:%g actual:%g", expected, result)
	}
}
