
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

type aggregateVector interface {
	VectorData
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

func testAggregate(t *testing.T, A aggregateVector) {
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

type aggregateIndexedVector interface {
	VectorData
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

func testAggregateIndexed(t *testing.T, A aggregateIndexedVector) {
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

type aggregatorVectorVector interface {
	VectorData
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

func testAggregateVector(t *testing.T, A aggregatorVectorVector, B VectorData) {
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

type assignVector interface {
	VectorData
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

func testAssign(t *testing.T, A assignVector) {
	value := rand.Float64()
	A.Assign(value)
	for i := 0; i < A.Size(); i++ {
		result := A.GetQuick(i)
		if math.Abs(value - result) > tol {
			t.Errorf("expected:%g actual:%g", value, result)
		}
	}
}

type assignArrayVector interface {
	VectorData
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

func testAssignArray(t *testing.T, A assignArrayVector) {
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

type assignFuncVector interface {
	VectorData
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

func testAssignFunc(t *testing.T, A assignFuncVector) {
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

type assignVectorVector interface {
	VectorData
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

func testAssignVector(t *testing.T, A assignVectorVector, B VectorData) {
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

type assignVectorFuncVector interface {
	VectorData
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

func testAssignVectorFunc(t *testing.T, A assignVectorFuncVector, B *Vector) {
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

type assignProcedureVector interface {
	VectorData
	AssignProcedure(Float64Procedure, float64) *Vector
	Copy() *Vector
}

func TestDenseAssignProcedure(t *testing.T) {
	A, _ := makeDenseVectors()
	testAssignProcedure(t, A)
}

func TestSparseAssignProcedure(t *testing.T) {
	A, _ := makeSparseVectors()
	testAssignProcedure(t, A)
}

func testAssignProcedure(t *testing.T, A assignProcedureVector) {
	procedure := func(element float64) bool {
		if math.Abs(element) > 0.1 {
			return true
		} else {
			return false
		}
	}
	Acopy := A.Copy()
	A.AssignProcedure(procedure, -1.0)
	for i := 0; i < A.Size(); i++ {
		var expected, result float64
		if math.Abs(Acopy.GetQuick(i)) > 0.1 {
			expected = -1.0
			result = A.GetQuick(i)
		} else {
			expected = Acopy.GetQuick(i)
			result = A.GetQuick(i)
		}
		if math.Abs(expected - result) > tol {
			t.Errorf("expected:%g actual:%g", expected, result)
		}
	}
}

type assignProcedureFuncVector interface {
	VectorData
	AssignProcedureFunc(Float64Procedure, Float64Func) *Vector
	Copy() *Vector
}

func TestDenseAssignProcedureFunc(t *testing.T) {
	A, _ := makeDenseVectors()
	testAssignProcedureFunc(t, A)
}

func TestSparseAssignProcedureFunc(t *testing.T) {
	A, _ := makeSparseVectors()
	testAssignProcedureFunc(t, A)
}

func testAssignProcedureFunc(t *testing.T, A assignProcedureFuncVector) {
	procedure := func(element float64) bool {
		if math.Abs(element) > 0.1 {
			return true
		} else {
			return false
		}
	}
	Acopy := A.Copy()
	A.AssignProcedureFunc(procedure, math.Tan)
	for i := 0; i < A.Size(); i++ {
		var expected, result float64
		if math.Abs(Acopy.GetQuick(i)) > 0.1 {
			expected = math.Tan(Acopy.GetQuick(i))
			result = A.GetQuick(i)
		} else {
			expected = Acopy.GetQuick(i)
			result = A.GetQuick(i)
		}
		if math.Abs(expected - result) > tol {
			t.Errorf("expected:%g actual:%g", expected, result)
		}
	}
}

type cardinalityVector interface {
	VectorData
	Cardinality() int
}

func TestDenseCardinality(t *testing.T) {
	A, _ := makeDenseVectors()
	testCardinality(t, A)
}

func TestSparseCardinality(t *testing.T) {
	A, _ := makeSparseVectors()
	testCardinality(t, A)
}

func testCardinality(t *testing.T, A cardinalityVector) {
	card := A.Cardinality()
	if A.Size() != card {
		t.Errorf("expected:%g actual:%g", A.Size(), card)
	}
}

type equalsVector interface {
	Assign(float64) *Vector
	Equals(float64) bool
}

func TestDenseEquals(t *testing.T) {
	A, _ := makeDenseVectors()
	testEquals(t, A)
}

func TestSparseEquals(t *testing.T) {
	A, _ := makeSparseVectors()
	testEquals(t, A)
}

func testEquals(t *testing.T, A equalsVector) {
	value := 1.0
	A.Assign(value)
	if !A.Equals(value) {
		t.Errorf("expected:%g", value)
	}
	if A.Equals(2) {
		t.Fail()
	}
}

type equalsVectorVector interface {
	VectorData
	EqualsVector(VectorData) bool
}

func TestDenseEqualsVector(t *testing.T) {
	A, B := makeDenseVectors()
	testEqualsVector(t, A, B)
}

func TestSparseEqualsVector(t *testing.T) {
	A, B := makeSparseVectors()
	testEqualsVector(t, A, B)
}

func testEqualsVector(t *testing.T, A, B equalsVectorVector) {
	if !A.EqualsVector(A) {
		t.Fail()
	}
	if A.EqualsVector(B) {
		t.Fail()
	}
}

type maxLocationVector interface {
	VectorData
	Assign(float64) *Vector
	MaxLocation() (float64, int)
}

func TestDenseMaxLocation(t *testing.T) {
	A, _ := makeDenseVectors()
	testMaxLocation(t, A)
}

func TestSparseMaxLocation(t *testing.T) {
	A, _ := makeSparseVectors()
	testMaxLocation(t, A)
}

func testMaxLocation(t *testing.T, A maxLocationVector) {
	A.Assign(0)
	value := 0.7
	A.SetQuick(A.Size() / 3, value)
	A.SetQuick(A.Size() / 2, 0.1)
	max, loc := A.MaxLocation()
	if math.Abs(value - max) > tol {
		t.Errorf("expected:%g actual:%g", value, max)
	}
	if A.Size() / 3 != loc {
		t.Errorf("expected:%d actual:%d", A.Size() / 3, loc)
	}
}

type minLocationVector interface {
	VectorData
	Assign(float64) *Vector
	MinLocation() (float64, int)
}

func TestDenseMinLocation(t *testing.T) {
	A, _ := makeDenseVectors()
	testMinLocation(t, A)
}

func TestSparseMinLocation(t *testing.T) {
	A, _ := makeSparseVectors()
	testMinLocation(t, A)
}

func testMinLocation(t *testing.T, A minLocationVector) {
	A.Assign(0)
	value := -0.7
	A.SetQuick(A.Size() / 3, value)
	A.SetQuick(A.Size() / 2, -0.1)
	min, loc := A.MinLocation()
	if math.Abs(value - min) > tol {
		t.Errorf("expected:%g actual:%g", value, min)
	}
	if A.Size() / 3 != loc {
		t.Errorf("expected:%d actual:%d", A.Size() / 3, loc)
	}
}

type negativeValuesVector interface {
	VectorData
	Assign(float64) *Vector
	NegativeValues(*[]int, *[]float64)
}

func TestDenseNegativeValues(t *testing.T) {
	A, _ := makeDenseVectors()
	testGetNegativeValues(t, A)
}

func TestSparseNegativeValues(t *testing.T) {
	A, _ := makeSparseVectors()
	testGetNegativeValues(t, A)
}

func testGetNegativeValues(t *testing.T, A negativeValuesVector) {
	A.Assign(0)
	A.SetQuick(A.Size() / 3, -0.7)
	A.SetQuick(A.Size() / 2, -0.1)
	var indexList []int
	var valueList []float64
	A.NegativeValues(&indexList, &valueList)
	if len(indexList) != 2 {
		t.Errorf("expected:%d actual:%d", 2, len(indexList))
	}
	if len(valueList) != 2 {
		t.Errorf("expected:%d actual:%d", 2, len(valueList))
	}
	if !ContainsInt(indexList, A.Size() / 3) {
		t.Errorf("missing:%d", A.Size() / 3)
	}
	if !ContainsInt(indexList, A.Size() / 2) {
		t.Errorf("missing:%d", A.Size() / 2)
	}
	if !ContainsFloat(valueList, -0.7, tol) {
		t.Errorf("missing:%g", -0.7)
	}
	if !ContainsFloat(valueList, -0.1, tol) {
		t.Errorf("missing:%g", -0.1)
	}
}

type nonZerosVector interface {
	VectorData
	Assign(float64) *Vector
	NonZeros(*[]int, *[]float64)
}

func TestDenseNonZeros(t *testing.T) {
	A, _ := makeDenseVectors()
	testNonZeros(t, A)
}

func TestSparseNonZeros(t *testing.T) {
	A, _ := makeSparseVectors()
	testNonZeros(t, A)
}

func testNonZeros(t *testing.T, A nonZerosVector) {
	A.Assign(0)
	A.SetQuick(A.Size() / 3, 0.7)
	A.SetQuick(A.Size() / 2, 0.1)
	var indexList []int
	var valueList []float64
	A.NonZeros(&indexList, &valueList)
	if len(indexList) != 2 {
		t.Errorf("expected:%d actual:%d", 2, len(indexList))
	}
	if len(valueList) != 2 {
		t.Errorf("expected:%d actual:%d", 2, len(valueList))
	}
	if !ContainsInt(indexList, A.Size() / 3) {
		t.Errorf("missing:%d", A.Size() / 3)
	}
	if !ContainsInt(indexList, A.Size() / 2) {
		t.Errorf("missing:%d", A.Size() / 2)
	}
	if !ContainsFloat(valueList, 0.7, tol) {
		t.Errorf("missing:%g", 0.7)
	}
	if !ContainsFloat(valueList, 0.1, tol) {
		t.Errorf("missing:%g", 0.1)
	}
}

type positiveValuesVector interface {
	VectorData
	Assign(float64) *Vector
	PositiveValues(*[]int, *[]float64)
}

func TestDensePositiveValues(t *testing.T) {
	A, _ := makeDenseVectors()
	testPositiveValues(t, A)
}

func TestSparsePositiveValues(t *testing.T) {
	A, _ := makeSparseVectors()
	testPositiveValues(t, A)
}

func testPositiveValues(t *testing.T, A positiveValuesVector) {
	A.Assign(0)
	A.SetQuick(A.Size() / 3, 0.7)
	A.SetQuick(A.Size() / 2, 0.1)
	var indexList []int
	var valueList []float64
	A.PositiveValues(&indexList, &valueList)
	if len(indexList) != 2 {
		t.Errorf("expected:%d actual:%d", 2, len(indexList))
	}
	if len(valueList) != 2 {
		t.Errorf("expected:%d actual:%d", 2, len(valueList))
	}
	if !ContainsInt(indexList, A.Size() / 3) {
		t.Errorf("missing:%d", A.Size() / 3)
	}
	if !ContainsInt(indexList, A.Size() / 2) {
		t.Errorf("missing:%d", A.Size() / 2)
	}
	if !ContainsFloat(valueList, 0.7, tol) {
		t.Errorf("missing:%g", 0.7)
	}
	if !ContainsFloat(valueList, 0.1, tol) {
		t.Errorf("missing:%g", 0.1)
	}
}

type toArrayVector interface {
	VectorData
	ToArray() []float64
}

func TestDenseToArray(t *testing.T) {
	A, _ := makeDenseVectors()
	testToArray(t, A)
}

func TestSparseToArray(t *testing.T) {
	A, _ := makeSparseVectors()
	testToArray(t, A)
}

func testToArray(t *testing.T, A toArrayVector) {
	array := A.ToArray()
	if A.Size() != len(array) {
		t.Errorf("expected:%d actual:%d", A.Size(), len(array))
	}
	for i := 0; i < A.Size(); i++ {
		expected := array[i]
		result := A.GetQuick(i)
		if math.Abs(expected - result) > tol {
			t.Errorf("expected:%g actual:%g", expected, result)
		}
	}
}

type fillArrayVector interface {
	VectorData
	FillArray([]float64) error
}

func TestDenseFillArray(t *testing.T) {
	A, _ := makeDenseVectors()
	testFillArray(t, A)
}

func TestSparseFillArray(t *testing.T) {
	A, _ := makeSparseVectors()
	testFillArray(t, A)
}

func testFillArray(t *testing.T, A fillArrayVector) {
	array := make([]float64, A.Size())
	err := A.FillArray(array)
	if err != nil {
		t.Fail()
	}
	for i := 0; i < A.Size(); i++ {
		expected := A.GetQuick(i)
		result := array[i]
		if math.Abs(expected - result) > tol {
			t.Errorf("expected:%g actual:%g", expected, result)
		}
	}
}

type reshapeMatrixVector interface {
	VectorData
}

func TestDenseReshapeMatrix(t *testing.T) {
	A, _ := makeDenseVectors()
	testReshapeMatrix(t, A)
}

func TestSparseReshapeMatrix(t *testing.T) {
	A, _ := makeSparseVectors()
	testReshapeMatrix(t, A)
}

func testReshapeMatrix(t *testing.T, A reshapeMatrixVector) {
	rows := 10
	columns := 17
	B, err := A.ReshapeMatrix(rows, columns)
	if err != nil {
		t.Fail()
	}
	idx := 0
	for c := 0; c < columns; c++ {
		for r := 0; r < rows; r++ {
			if math.Abs(A.GetQuick(idx) - B.GetQuick(r, c)) > tol {
				t.Errorf("idx:%d r:%d c:%d expected:%g actual:%g",
					idx, r, c, A.GetQuick(idx), B.GetQuick(r, c))
			}
			idx++
		}
	}
}

type reshapeCubeVector interface {
	VectorData
}

func TestDenseReshapeCube(t *testing.T) {
	A, _ := makeDenseVectors()
	testReshapeCube(t, A)
}

func TestSparseReshapeCube(t *testing.T) {
	A, _ := makeSparseVectors()
	testReshapeCube(t, A)
}

func testReshapeCube(t *testing.T, A reshapeCubeVector) {
	slices := 2
	rows := 5
	columns := 17
	B, err := A.ReshapeCube(slices, rows, columns)
	if err != nil {
		t.Fail()
	}
	idx := 0
	for s := 0; s < slices; s++ {
		for c := 0; c < columns; c++ {
			for r := 0; r < rows; r++ {
				if math.Abs(A.GetQuick(idx) - B.GetQuick(s, r, c)) > tol {
					t.Errorf("idx:%d s:%d r:%d c:%d expected:%g actual:%g",
						idx, s, r, c, A.GetQuick(idx), B.GetQuick(s, r, c))
				}
				idx++
			}
		}
	}
}

type swapVector interface {
	VectorData
	Swap(VectorData) error
	Copy() *Vector
}

func TestDenseSwap(t *testing.T) {
	A, B := makeDenseVectors()
	testSwap(t, A, B)
}

func TestSparseSwap(t *testing.T) {
	A, B := makeSparseVectors()
	testSwap(t, A, B)
}

func testSwap(t *testing.T, A, B swapVector) {
	Acopy := A.Copy()
	Bcopy := B.Copy()
	A.Swap(B)
	for i := 0; i < A.Size(); i++ {
		expected := Bcopy.GetQuick(i)
		result := A.GetQuick(i)
		if math.Abs(expected - result) > tol {
			t.Errorf("expected:%g actual:%g", expected, result)
		}

		expected = Acopy.GetQuick(i)
		result = B.GetQuick(i)
		if math.Abs(expected - result) > tol {
			t.Errorf("expected:%g actual:%g", expected, result)
		}
	}
}

type viewFlipVector interface {
	VectorData
	ViewFlip() *Vector
}

func TestDenseViewFlip(t *testing.T) {
	A, _ := makeDenseVectors()
	testViewFlip(t, A)
}

func TestSparseViewFlip(t *testing.T) {
	A, _ := makeSparseVectors()
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
			t.Errorf("zero:%d stride:%d", b.Zero(), b.Stride())
			t.Errorf("expected:%g actual:%g", expected, result)
		}
	}
}
