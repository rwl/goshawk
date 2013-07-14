
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

type indexedAggregator interface {
	Size() int
	GetQuick(int) float64
	AggregateIndexed(Float64Float64Func, Float64Func, []int) float64
}

func TestDenseAggregateIndexed(t *testing.T) {
	A, _ := makeDenseVectors()
	testAggregateIndexed(t, A)
}

func TestSparseAggregateIndexed(t *testing.T) {
	A, _ := makeSparseVectors()
	testAggregateIndexed(t, A)
}

func testAggregateIndexed(t *testing.T, A indexedAggregator) {
	indexList := make([]int, A.Size())
	for i := 0; i < A.Size(); i++ {
		indexList[i] = i
	}
	expected := 0.0
	for i := 0; i < A.Size(); i++ {
		elem := A.GetQuick(i)
		expected += elem * elem
	}
	result := A.AggregateIndexed(Plus, Square, indexList)
	if math.Abs(expected - result) > tol {
		t.Errorf("expected:%g actual:%g", expected, result)
	}
}

type vectorAggregator interface {
	Size() int
	GetQuick(int) float64
	AggregateVector(VectorData, Float64Float64Func, Float64Float64Func) (float64, error)
}

func TestDenseAggregateVector(t *testing.T) {
	A, B := makeDenseVectors()
	testAggregateVector(t, A, B)
}

func TestSparseAggregateVector(t *testing.T) {
	A, B := makeSparseVectors()
	testAggregateVector(t, A, B)
}

func testAggregateVector(t *testing.T, A vectorAggregator, B VectorData) {
	expected := 0.0
	for i := 0; i < A.Size(); i++ {
		elemA := A.GetQuick(i)
		elemB := B.GetQuick(i)
		expected += elemA * elemB
	}
	result, err := A.AggregateVector(B, Plus, Mult)
	if err != nil {
		t.Error(err)
	}
	if math.Abs(expected - result) > tol {
		t.Errorf("expected:%g actual:%g", expected, result)
	}
}

type assigner interface {
	Size() int
	GetQuick(int) float64
	Assign(float64) *Vector
}

func TestDenseAssign(t *testing.T) {
	A, _ := makeDenseVectors()
	testAssign(t, A)
}

func TestSparseAssign(t *testing.T) {
	A, _ := makeSparseVectors()
	testAssign(t, A)
}

func testAssign(t *testing.T, A assigner) {
	value := rand.Float64()
	A.Assign(value)
	for i := 0; i < A.Size(); i++ {
		result := A.GetQuick(i)
		if math.Abs(value - result) > tol {
			t.Errorf("expected:%g actual:%g", value, result)
		}
	}
}

type arrayAssigner interface {
	Size() int
	GetQuick(int) float64
	AssignArray([]float64) (*Vector, error)
}

func TestDenseAssignArray(t *testing.T) {
	A, _ := makeDenseVectors()
	testAssignArray(t, A)
}

func TestSparseAssignArray(t *testing.T) {
	A, _ := makeSparseVectors()
	testAssignArray(t, A)
}

func testAssignArray(t *testing.T, A arrayAssigner) {
	expected := make([]float64, A.Size())
	for i := 0; i < A.Size(); i++ {
		expected[i] = rand.Float64()
	}
	A.AssignArray(expected)
	for i := 0; i < A.Size(); i++ {
		result := A.GetQuick(i)
		if math.Abs(expected[i] - result) > tol {
			t.Errorf("expected:%g actual:%g", expected[i], result)
		}
	}
}

type funcAssigner interface {
	Size() int
	GetQuick(int) float64
	AssignFunc(Float64Func) *Vector
	Copy() *Vector
}

func TestDenseAssignFunc(t *testing.T) {
	A, _ := makeDenseVectors()
	testAssignFunc(t, A)
}

func TestSparseAssignFunc(t *testing.T) {
	A, _ := makeSparseVectors()
	testAssignFunc(t, A)
}

func testAssignFunc(t *testing.T, A funcAssigner) {
	Acopy := A.Copy()
	A.AssignFunc(math.Acos)
	for i := 0; i < A.Size(); i++ {
		expected := math.Acos(Acopy.GetQuick(i))
		result := A.GetQuick(i)
		if math.Abs(expected - result) > tol {
			t.Errorf("expected:%g actual:%g", expected, result)
		}
	}
}

type vectorAssigner interface {
	Size() int
	GetQuick(int) float64
	AssignVector(VectorData) (*Vector, error)
}

func TestDenseAssignVector(t *testing.T) {
	A, B := makeDenseVectors()
	testAssignVector(t, A, B)
}

func TestSparseAssignVector(t *testing.T) {
	A, B := makeSparseVectors()
	testAssignVector(t, A, B)
}

func testAssignVector(t *testing.T, A vectorAssigner, B VectorData) {
	A.AssignVector(B)
	if A.Size() != B.Size() {
		t.Errorf("sizes must be equal: %d!=%d", A.Size(), B.Size())
	}
	for i := 0; i < A.Size(); i++ {
		expected := B.GetQuick(i)
		result := A.GetQuick(i)
		if math.Abs(expected - result) > tol {
			t.Errorf("expected:%g actual:%g", expected, result)
		}
	}
}

type vectorFuncAssigner interface {
	Size() int
	GetQuick(int) float64
	AssignVectorFunc(VectorData, Float64Float64Func) (*Vector, error)
	Copy() *Vector
}

func TestDenseAssignVectorFunc(t *testing.T) {
	A, B := makeDenseVectors()
	testAssignVectorFunc(t, A, B)
}

func TestSparseAssignVectorFunc(t *testing.T) {
	A, B := makeSparseVectors()
	testAssignVectorFunc(t, A, B)
}

func testAssignVectorFunc(t *testing.T, A vectorFuncAssigner, B *Vector) {
	Acopy := A.Copy()
	A.AssignVectorFunc(B, Div)
	for i := 0; i < A.Size(); i++ {
		expected := Acopy.GetQuick(i) / B.GetQuick(i)
		result := A.GetQuick(i)
		if math.Abs(expected - result) > tol {
			t.Errorf("expected:%g actual:%g", expected, result)
		}
	}
}
