package tfloat64

import (
	"testing"
	"math"
)

type aggregateVector interface {
Vec
	Aggregate(Float64Float64Func, Float64Func) float64
}

func TestDenseAggregate(t *testing.T) {
	A := makeDenseVector()
	testAggregate(t, A)
}

func TestSparseAggregate(t *testing.T) {
	A := makeSparseVector()
	testAggregate(t, A)
}

func testAggregate(t *testing.T, A aggregateVector) {
	expected := 0.0
	for i := 0; i < A.Size(); i++ {
		elem := A.GetQuick(i)
		expected += elem*elem
	}

	result := A.Aggregate(Plus, Square)
	if math.Abs(expected - result) > tol {
		t.Errorf("expected:%g actual:%g", expected, result)
	}
}

type aggregateIndexedVector interface {
Vec
	AggregateIndexed(Float64Float64Func, Float64Func, []int) float64
}

func TestDenseAggregateIndexed(t *testing.T) {
	A := makeDenseVector()
	testAggregateIndexed(t, A)
}

func TestSparseAggregateIndexed(t *testing.T) {
	A := makeSparseVector()
	testAggregateIndexed(t, A)
}

func testAggregateIndexed(t *testing.T, A aggregateIndexedVector) {
	indexList := make([]int, A.Size())
	for i := 0; i < A.Size(); i++ {
		indexList[i] = i
	}
	expected := 0.0
	for i := 0; i < A.Size(); i++ {
		elem := A.GetQuick(i)
		expected += elem*elem
	}
	result := A.AggregateIndexed(Plus, Square, indexList)
	if math.Abs(expected - result) > tol {
		t.Errorf("expected:%g actual:%g", expected, result)
	}
}

type aggregatorVectorVector interface {
Vec
	AggregateVector(Vec, Float64Float64Func, Float64Float64Func) (float64, error)
}

func TestDenseAggregateVector(t *testing.T) {
	A := makeDenseVector()
	B := makeDenseVector()
	testAggregateVector(t, A, B)
}

func TestSparseAggregateVector(t *testing.T) {
	A := makeSparseVector()
	B := makeSparseVector()
	testAggregateVector(t, A, B)
}

func testAggregateVector(t *testing.T, A aggregatorVectorVector, B Vec) {
	expected := 0.0
	for i := 0; i < A.Size(); i++ {
		elemA := A.GetQuick(i)
		elemB := B.GetQuick(i)
		expected += elemA*elemB
	}
	result, err := A.AggregateVector(B, Plus, Mult)
	if err != nil {
		t.Error(err)
	}
	if math.Abs(expected - result) > tol {
		t.Errorf("expected:%g actual:%g", expected, result)
	}
}
