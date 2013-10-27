package tfloat64

import (
	"testing"
	"math"
	"github.com/rwl/goshawk/common"
)

const tol = 1e-10

type cardinalityVector interface {
	Vec
	Cardinality() int
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
	Vec
	EqualsVector(Vec) bool
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
	Vec
	Assign(float64) *Vector
	MaxLocation() (float64, int)
}

func testMaxLocation(t *testing.T, A maxLocationVector) {
	A.Assign(0)
	value := 0.7
	A.SetQuick(A.Size()/3, value)
	A.SetQuick(A.Size()/2, 0.1)
	max, loc := A.MaxLocation()
	if math.Abs(value - max) > tol {
		t.Errorf("expected:%g actual:%g", value, max)
	}
	if A.Size()/3 != loc {
		t.Errorf("expected:%d actual:%d", A.Size()/3, loc)
	}
}

type minLocationVector interface {
	Vec
	Assign(float64) *Vector
	MinLocation() (float64, int)
}

func testMinLocation(t *testing.T, A minLocationVector) {
	A.Assign(0)
	value := -0.7
	A.SetQuick(A.Size()/3, value)
	A.SetQuick(A.Size()/2, -0.1)
	min, loc := A.MinLocation()
	if math.Abs(value - min) > tol {
		t.Errorf("expected:%g actual:%g", value, min)
	}
	if A.Size()/3 != loc {
		t.Errorf("expected:%d actual:%d", A.Size()/3, loc)
	}
}

type negativeValuesVector interface {
	Vec
	Assign(float64) *Vector
	NegativeValues(*[]int, *[]float64)
}

func testGetNegativeValues(t *testing.T, A negativeValuesVector) {
	A.Assign(0)
	A.SetQuick(A.Size()/3, -0.7)
	A.SetQuick(A.Size()/2, -0.1)
	var indexList []int
	var valueList []float64
	A.NegativeValues(&indexList, &valueList)
	if len(indexList) != 2 {
		t.Errorf("expected:%d actual:%d", 2, len(indexList))
	}
	if len(valueList) != 2 {
		t.Errorf("expected:%d actual:%d", 2, len(valueList))
	}
	if !common.ContainsInt(indexList, A.Size()/3) {
		t.Errorf("missing:%d", A.Size()/3)
	}
	if !common.ContainsInt(indexList, A.Size()/2) {
		t.Errorf("missing:%d", A.Size()/2)
	}
	if !common.ContainsFloat(valueList, -0.7, tol) {
		t.Errorf("missing:%g", -0.7)
	}
	if !common.ContainsFloat(valueList, -0.1, tol) {
		t.Errorf("missing:%g", -0.1)
	}
}

type nonZerosVector interface {
Vec
	Assign(float64) *Vector
	NonZeros(*[]int, *[]float64)
}

func testNonZeros(t *testing.T, A nonZerosVector) {
	A.Assign(0)
	A.SetQuick(A.Size()/3, 0.7)
	A.SetQuick(A.Size()/2, 0.1)
	var indexList []int
	var valueList []float64
	A.NonZeros(&indexList, &valueList)
	if len(indexList) != 2 {
		t.Errorf("expected:%d actual:%d", 2, len(indexList))
	}
	if len(valueList) != 2 {
		t.Errorf("expected:%d actual:%d", 2, len(valueList))
	}
	if !common.ContainsInt(indexList, A.Size()/3) {
		t.Errorf("missing:%d", A.Size()/3)
	}
	if !common.ContainsInt(indexList, A.Size()/2) {
		t.Errorf("missing:%d", A.Size()/2)
	}
	if !common.ContainsFloat(valueList, 0.7, tol) {
		t.Errorf("missing:%g", 0.7)
	}
	if !common.ContainsFloat(valueList, 0.1, tol) {
		t.Errorf("missing:%g", 0.1)
	}
}

type positiveValuesVector interface {
	Vec
	Assign(float64) *Vector
	PositiveValues(*[]int, *[]float64)
}

func testPositiveValues(t *testing.T, A positiveValuesVector) {
	A.Assign(0)
	A.SetQuick(A.Size()/3, 0.7)
	A.SetQuick(A.Size()/2, 0.1)
	var indexList []int
	var valueList []float64
	A.PositiveValues(&indexList, &valueList)
	if len(indexList) != 2 {
		t.Errorf("expected:%d actual:%d", 2, len(indexList))
	}
	if len(valueList) != 2 {
		t.Errorf("expected:%d actual:%d", 2, len(valueList))
	}
	if !common.ContainsInt(indexList, A.Size()/3) {
		t.Errorf("missing:%d", A.Size()/3)
	}
	if !common.ContainsInt(indexList, A.Size()/2) {
		t.Errorf("missing:%d", A.Size()/2)
	}
	if !common.ContainsFloat(valueList, 0.7, tol) {
		t.Errorf("missing:%g", 0.7)
	}
	if !common.ContainsFloat(valueList, 0.1, tol) {
		t.Errorf("missing:%g", 0.1)
	}
}

type toArrayVector interface {
	Vec
	ToArray() []float64
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
	Vec
	FillArray([]float64) error
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
	Vec
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
	Vec
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
	Vec
	Swap(Vec) error
	Copy() *Vector
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

type dotProductVector interface {
	Vec
	ZDotProduct(Vec) float64
}

func testZDotProduct(t *testing.T, A, B dotProductVector) {
	product := A.ZDotProduct(B)
	var expected float64 = 0
	for i := 0; i < A.Size(); i++ {
		expected += A.GetQuick(i)*B.GetQuick(i)
	}
	if math.Abs(expected - product) > tol {
		t.Errorf("expected:%g actual:%g", expected, product)
	}
}

type dotProductRangeVector interface {
	Vec
	ZDotProductRange(Vec, int, int) float64
}

func testZDotProductRange(t *testing.T, A, B dotProductRangeVector) {
	product := A.ZDotProductRange(B, 5, B.Size() - 10)
	var expected float64 = 0
	for i := 5; i < A.Size() - 5; i++ {
		expected += A.GetQuick(i)*B.GetQuick(i)
	}
	if math.Abs(expected - product) > tol {
		t.Errorf("expected:%g actual:%g", expected, product)
	}
}

type dotProductSelectionVector interface {
	Vec
	ZDotProductSelection(Vec, int, int, []int) float64
	NonZeros(*[]int, *[]float64)
}

func testZDotProductSelection(t *testing.T, A, B dotProductSelectionVector) {
	var indexList []int
	B.NonZeros(&indexList, nil)
	product := A.ZDotProductSelection(B, 5, B.Size() - 10, indexList)
	var expected float64 = 0
	for i := 5; i < A.Size() - 5; i++ {
		expected += A.GetQuick(i)*B.GetQuick(i)
	}
	if math.Abs(expected - product) > tol {
		t.Errorf("expected:%g actual:%g", expected, product)
	}
}

type sumVector interface {
	Vec
	ZSum() float64
}

func testZSum(t *testing.T, A sumVector) {
	sum := A.ZSum()
	var expected float64 = 0
	for i := 0; i < A.Size(); i++ {
		expected += A.GetQuick(i)
	}
	if math.Abs(expected - sum) > tol {
		t.Errorf("expected:%g actual:%g", expected, sum)
	}
}
