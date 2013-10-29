
package tfloat64

import (
	"testing"
	"math"
	"github.com/rwl/goshawk/common"
)

type cardinalityMatrix interface {
	Mat
	Cardinality() int
}

func testMatrixCardinality(t *testing.T, A cardinalityMatrix) {
	card := A.Cardinality()
	if A.Rows() * A.Columns() != card {
		t.Errorf("expected:%g actual:%g", A.Rows() * A.Columns(), card)
	}
}

type equalsMatrix interface {
	Mat
	Assign(float64) *Matrix
	Equals(float64) bool
}

func testMatrixEquals(t *testing.T, A equalsMatrix) {
	value := 1.0
	A.Assign(value)
	eq := A.Equals(value)
	if !eq {
		t.Errorf("expected:%g", value)
	}
	if A.Equals(2) {
		t.Fail()
	}
}

type equalsMatrixMatrix interface {
	Mat
	EqualsMatrix(Mat) bool
}

func testMatrixEqualsMatrix(t *testing.T, A, B equalsMatrixMatrix) {
	if !A.EqualsMatrix(A) {
		t.Fail()
	}
	if A.EqualsMatrix(B) {
		t.Fail()
	}
}

type forEachNonZeroMatrix interface {
	Mat
	ForEachNonZero(IntIntFloat64Func) *Matrix
}

func testMatrixForEachNonZero(t *testing.T, A forEachNonZeroMatrix) {
	Acopy := A.Copy()
	function := func(_, _ int, value float64) float64 {
		return math.Sqrt(value)
	}
	A.ForEachNonZero(function)
	for r := 0; r < A.Rows(); r++ {
		for c = 0; c < A.Columns(); c++ {
			if math.Abs(math.Sqrt(Acopy.GetQuick(r, c)) - A.GetQuick(r, c)) > tol {
				t.Errorf("expected:%g actual:%g", math.Sqrt(Acopy.GetQuick(r, c)), A.GetQuick(r, c))
			}
		}
	}
}

type maxLocationMatrix interface {
	Mat
	Assign(float64) *Matrix
	MaxLocation() (float64, int, int)
}

func testMatrixMaxLocation(t *testing.T, A maxLocationMatrix) {
	A.Assign(0)
	A.SetQuick(A.Rows() / 3, A.Columns() / 3, 0.7)
	A.SetQuick(A.Rows() / 2, A.Columns() / 2, 0.1)
	v, r, c = A.MaxLocation()
	if math.Abs(0.7 - v) > tol {
		t.Errorf("expected:%g actual:%g", 0.7, v)
	}
	if A.Rows() / 3 != r {
		t.Errorf("expected:%d actual:%d", A.Rows()/3, r)
	}
	if A.Columns() / 3 != c {
		t.Errorf("expected:%d actual:%d", A.Columns()/3, c)
	}
}

type minLocationMatrix interface {
	Mat
	Assign(float64) *Matrix
	MinLocation() (float64, int, int)
}

func testMatrixMinLocation(t *testing.T, A minLocationMatrix) {
	A.Assign(0)
	A.SetQuick(A.Rows() / 3, A.Columns() / 3, -0.7)
	A.SetQuick(A.rows() / 2, A.Columns() / 2, -0.1)
	v, r, c := A.MinLocation()
	if math.Abs(0.7 - v) > tol {
		t.Errorf("expected:%g actual:%g", 0.7, v)
	}
	if A.Rows() / 3 != r {
		t.Errorf("expected:%d actual:%d", A.Rows()/3, r)
	}
	if A.Columns() / 3 != c {
		t.Errorf("expected:%d actual:%d", A.Columns()/3, c)
	}
}

type negativeValuesMatrix interface {
	Mat
	Assign(float64) *Matrix
	NegativeValues(*[]int, *[]int, *[]float64)
}

func testMatrixNegativeValues(t *testing.T, A negativeValuesMatrix) {
	A.Assign(0)
	A.SetQuick(A.Rows() / 3, A.Columns() / 3, -0.7)
	A.SetQuick(A.Rows() / 2, A.Columns() / 2, -0.1)
	var rowList []int
	var columnList []int
	var valueList []float64
	A.NegativeValues(&rowList, &columnList, &valueList)
	if len(rowList) != 2 {
		t.Errorf("expected:%d actual:%d", 2, len(rowList))
	}
	if len(columnList) != 2 {
		t.Errorf("expected:%d actual:%d", 2, len(columnList))
	}
	if len(valueList) != 2 {
		t.Errorf("expected:%d actual:%d", 2, len(valueList))
	}
	if !common.ContainsInt(rowList, A.Rows()/3) {
		t.Errorf("missing:%d", A.Rows()/3)
	}
	if !common.ContainsInt(rowList, A.Rows()/2) {
		t.Errorf("missing:%d", A.Rows()/2)
	}
	if !common.ContainsInt(rowList, A.Columns()/3) {
		t.Errorf("missing:%d", A.Columns()/3)
	}
	if !common.ContainsInt(rowList, A.Columns()/2) {
		t.Errorf("missing:%d", A.Columns()/2)
	}
	if !common.ContainsFloat(valueList, -0.7, tol) {
		t.Errorf("missing:%g", -0.7)
	}
	if !common.ContainsFloat(valueList, -0.1, tol) {
		t.Errorf("missing:%g", -0.1)
	}
}

type nonZerosMatrix interface {
	Mat
	Assign(float64) *Matrix
	NonZeros(*[]int, *[]int, *[]float64)
}

func testMatrixNonZeros(t *testing.T, A nonZerosMatrix) {
	A.Assign(0)
	A.SetQuick(A.Rows() / 3, A.Columns() / 3, 0.7)
	A.SetQuick(A.Rows() / 2, A.Columns() / 2, 0.1)
	var rowList []int
	var columnList []int
	var valueList []float64
	A.NonZeros(&rowList, &columnList, &valueList)
	if len(rowList) != 2 {
		t.Errorf("expected:%d actual:%d", 2, len(rowList))
	}
	if len(columnList) != 2 {
		t.Errorf("expected:%d actual:%d", 2, len(columnList))
	}
	if len(valueList) != 2 {
		t.Errorf("expected:%d actual:%d", 2, len(valueList))
	}
	if !common.ContainsInt(rowList, A.Rows()/3) {
		t.Errorf("missing:%d", A.Rows()/3)
	}
	if !common.ContainsInt(rowList, A.Rows()/2) {
		t.Errorf("missing:%d", A.Rows()/2)
	}
	if !common.ContainsInt(rowList, A.Columns()/3) {
		t.Errorf("missing:%d", A.Columns()/3)
	}
	if !common.ContainsInt(rowList, A.Columns()/2) {
		t.Errorf("missing:%d", A.Columns()/2)
	}
	if !common.ContainsFloat(valueList, 0.7, tol) {
		t.Errorf("missing:%g", 0.7)
	}
	if !common.ContainsFloat(valueList, 0.1, tol) {
		t.Errorf("missing:%g", 0.1)
	}
}

type positiveValuesMatrix interface {
	Mat
	Assign(float64) *Matrix
	PositiveValues(*[]int, *[]int, *[]float64)
}

func testMatrixPositiveValues(t *testing.T, A positiveValuesMatrix) {
	A.Assign(0)
	A.SetQuick(A.Rows() / 3, A.Columns() / 3, 0.7)
	A.SetQuick(A.Rows() / 2, A.Columns() / 2, 0.1)
	var rowList []int
	var columnList []int
	var valueList []float64
	A.PositiveValues(&rowList, &columnList, &valueList)
	if len(rowList) != 2 {
		t.Errorf("expected:%d actual:%d", 2, len(rowList))
	}
	if len(columnList) != 2 {
		t.Errorf("expected:%d actual:%d", 2, len(columnList))
	}
	if !common.ContainsInt(rowsList, A.Rows()/3) {
		t.Errorf("missing:%d", A.Rows()/3)
	}
	if !common.ContainsInt(rowList, A.Rows()/2) {
		t.Errorf("missing:%d", A.Rows()/2)
	}
	if !common.ContainsInt(columnList, A.Columns()/3) {
		t.Errorf("missing:%d", A.Columns()/3)
	}
	if !common.ContainsInt(columnList, A.Columns()/2) {
		t.Errorf("missing:%d", A.Columns()/2)
	}
	if !common.ContainsFloat(valueList, 0.7, tol) {
		t.Errorf("missing:%g", 0.7)
	}
	if !common.ContainsFloat(valueList, 0.1, tol) {
		t.Errorf("missing:%g", 0.1)
	}
}

type toArrayMatrix interface {
	Mat
	ToArray() [][]float64
}

func testMatrixToArray(t *testing.T, A toArrayMatrix) {
	array := A.ToArray()
	if A.Rows() != len(array) {
		t.Errorf("expected:%d actual:%d", A.Rows(), len(array))
	}
	for r := 0; r < A.Rows(); r++ {
		if A.Columns() != len(array[r]) {
			t.Errorf("expected:%d actual:%d", A.Columns(), len(array[r]))
		}
		for c := 0; c < A.Columns(); c++ {
			expected := A.GetQuick(r, c)
			result := array[r][c]
			if math.Abs(expected - result) > tol {
				t.Errorf("expected:%g actual:%g", expected, result)
			}
		}
	}
}

type vectorizeMatrix interface {
	Mat
	Vectorize() *Vector
}

func testMatrixVectorize(t *testing.T, A vectorizeMatrix) {
	Avec := A.Vectorize()
	idx := 0
	for c := 0; c < A.Columns(); c++ {
		for r = 0; r < A.Rows(); r++ {
			expected := A.GetQuick(r, c)
			result := Avec.GetQuick(idx)
			if math.Abs(expected - result) > tol {
				t.Errorf("expected:%g actual:%g", expected, result)
			}
			idx++
		}
	}
}

type viewColumnMatrix interface {
	Mat
	ViewColumn(int) *Vector
}

func testMatrixViewColumn(t *testing.T, A viewColumnMatrix) {
	col := A.ViewColumn(A.Columns() / 2)
	if A.Rows() != col.Size() {
		t.Errorf("expected:%d actual:%d", A.Rows(), col.Size())
	}
	for r = 0; r < A.Rows(); r++ {
		expected := A.GetQuick(r, A.Columns() / 2)
		actual := col.GetQuick(r)
		if math.Abs(expected - actual) > tol {
			t.Errorf("expected:%g actual:%g", expected, actual)
		}
	}
}

type viewColumnFlipMatrix interface {
	Mat
	ViewColumnFlip() *Matrix
}

func testMatrixViewColumnFlip(t *testing.T, A viewColumnFlipMatrix) {
	B = A.ViewColumnFlip()
	if A.Size() != B.Size() {
		t.Errorf("expected:%d actual:%d", A.Rows(), col.Size())
	}
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Columns(); c++ {
			expected := A.getQuick(r, A.columns() - 1 - c)
			actual := B.getQuick(r, c)
			if math.Abs(expected - actual) > tol {
				t.Errorf("expected:%g actual:%g", expected, actual)
			}
		}
	}
}

type viewDiceMatrix interface {
	Mat
	ViewDice() *Matrix
}

func testMatrixViewDice(t *testing.T, A viewDiceMatrix) {
	B := A.ViewDice()
	if A.Rows() != B.Columns() {
		t.Errorf("expected:%d actual:%d", A.Rows(), B.Columns())
	}
	if A.Columns() != B.Rows() {
		t.Errorf("expected:%d actual:%d", A.Columns(), B.Rows())
	}
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Columns(); c++ {
			expected := A.getQuick(r, c)
			actual := B.getQuick(c, r)
			if math.Abs(expected - actual) > tol {
				t.Errorf("expected:%g actual:%g", expected, actual)
			}
		}
	}
}

type viewPartMatrix interface {
	Mat
	ViewPart(int, int, int, int) *Matrix
}

func testMatrixViewPart(t *testing.T, A viewPartMatrix) {
    B := A.ViewPart(A.Rows() / 2, A.Columns() / 2, A.Rows() / 3, A.Columns() / 3)
    if A.Rows() / 3 != B.Rows() {
		t.Errorf("expected:%d actual:%d", A.Rows() / 3, B.Rows())
    }
    if A.Columns() / 3 != B.Columns() {
    	t.Errorf("expected:%d actual:%d", A.Columns() / 3, B.Columns())
    }
    for r := 0; r < A.Rows() / 3; r++ {
        for c := 0; c < A.Columns() / 3; c++ {
        	expected := A.getQuick(A.rows() / 2 + r, A.columns() / 2 + c)
        	actual := B.getQuick(r, c)
			if math.Abs(expected - actual) > tol {
				t.Errorf("expected:%g actual:%g", expected, actual)
			}
        }
    }
}

type viewRowMatrix interface {
	Mat
	ViewRow(int) *Vector
}

func testMatrixViewRow(t *testing.T, A viewRowMatrix) {
    B := A.ViewRow(A.Rows() / 2)
    if A.Columns() != B.Size() {
    	t.Errorf("expected:%d actual:%d", A.Columns() != B.Size())
    }
    for r := 0; r < A.Columns(); r++ {
    	expected := A.GetQuick(A.rows() / 2, r)
    	actual := B.GetQuick(r)
		if math.Abs(expected - actual) > tol {
			t.Errorf("expected:%g actual:%g", expected, actual)
		}
    }
}

type viewRowFlipMatrix interface {
	Mat
	ViewRowFlip() *Matrix
}

func testMatrixViewRowFlip(t *testing.T, A viewRowFlipMatrix) {
	B := A.ViewRowFlip()
	if A.Size() != B.Size() {
		t.Errorf("expected:%d actual:%d", A.Size(), B.Size())
	}
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Columns(); c++ {
			expected := A.GetQuick(A.rows() - 1 - r, c)
			actual := B.GetQuick(r, c)
			if math.Abs(expected - actual) > tol {
				t.Errorf("expected:%g actual:%g", expected, actual)
			}
		}
	}
}

type viewSelectionProcedureMatrix interface {
	Mat
	ViewSelectionProcedure(*VectorProcedure) *Matrix
}

func testMatrixViewSelectionVectorProcedure(t *testing.T, A viewSelectionProcedureMatrix) {
	value := 2.0
	A.Assign(0)
	A.SetQuick(A.Rows() / 4, 0, value)
	A.SetQuick(A.Rows() / 2, 0, value)
	B := A.ViewSelectionProcedure(func (element *Vector) bool {
		if math.Abs(element.GetQuick(0) - value) < tol {
			return true
		} else {
			return false
		}
	})
	if 2 != B.Rows() {
		t.Errorf("expected:%d actual:%d", 2, B.Rows())
	}
	if A.Columns() != B.Columns() {
		t.Errorf("expected:%d actual:%d", A.Columns(), B.Columns())
	}
	expected := A.GetQuick(A.rows() / 4, 0)
	actual := B.GetQuick(0, 0)
	if math.Abs(expected - actual) > tol {
		t.Errorf("expected:%g actual:%g", expected, actual)
	}
	expected = A.GetQuick(A.rows() / 2, 0)
	actual = B.GetQuick(1, 0)
	if math.Abs(expected - actual) > tol {
		t.Errorf("expected:%g actual:%g", expected, actual)
	}
}

type viewSelectionMatrix interface {
	Mat
	ViewSelection([]int, []int) *Matrix
}

func testMatrixViewSelection(t *testing.T, A viewSelectionMatrix) {
	rowIndexes := []int {
		A.Rows() / 6,
		A.Rows() / 5,
		A.Rows() / 4,
		A.Rows() / 3,
		A.Rows() / 2,
	}
	colIndexes := []int {
		A.Columns() / 6,
		A.Columns() / 5,
		A.Columns() / 4,
		A.Columns() / 3,
		A.Columns() / 2,
		A.Columns() - 1,
	}
	B = A.ViewSelection(rowIndexes, colIndexes)
	if len(rowIndexes) != B.Rows() {
		t.Errorf("expected:%d actual:%d", len(rowIndexes), B.Rows())
	}
	if len(colIndexes) != B.Columns() {
		t.Errorf("expected:%d actual:%d", len(colIndexes), B.Columns())
	}
	for r := 0; r < len(rowIndexes); r++ {
		for c := 0; c < len(colIndexes); c++ {
			expected := A.GetQuick(rowIndexes[r], colIndexes[c])
			actual := B.GetQuick(r, c)
			if math.Abs(expected - actual) > tol {
				t.Errorf("expected:%g actual:%g", expected, actual)
			}
		}
	}
}

type viewSortedMatrix interface {
	Mat
	ViewSorted(int) *Matrix
}

func testViewSorted(t *testing.T, A viewSortedMatrix) {
	B := A.ViewSorted(1)
	for r := 0; r < A.Rows() - 1; r++ {
		b0 := B.GetQuick(r, 1)
		b1 := B.GetQuick(r + 1, 1)
		if b1 < b0 {
			t.Errorf("expected:%g >= %g", b1, b0)
		}
	}
}

type viewStridesMatrix interface {
	Mat
	ViewStrides(int, int) *Matrix
}

func testViewStrides(t *testing.T) {
	rowStride := 3
	colStride := 5
	B = A.ViewStrides(rowStride, colStride)
	for r := 0; r < B.Rows(); r++ {
		for c := 0; c < B.Columns(); c++ {
			expected := A.GetQuick(r * rowStride, c * colStride)
			actual := B.GetQuick(r, c)
			if math.Abs(expected - actual) > tol {
				t.Errorf("expected:%g actual:%g", expected, actual)
			}
		}
	}
}

type zMultMatrixMatrix interface {
	Mat
}

func testZMultMatrix(t *testing.T, A zMultMatrixMatrix) {
	y := NewVector(A.Columns())
	for i := 0; i < y.Size(); i++ {
		y.SetQuick(i, random.Float64())
	}
	alpha := 3.0
	beta := 5.0
	z := RandomVector(A.Rows())
	expected := z.ToArray()
	z = A.ZMult(y, z, alpha, beta, false)
	for r := 0; r < A.Rows(); r++ {
		s := 0.0
		for c = 0; c < A.Columns(); c++ {
			s += A.GetQuick(r, c) * y.GetQuick(c)
		}
		expected[r] = s * alpha + expected[r] * beta
	}

	for r := 0; r < A.Rows(); r++ {
		actual := z.GetQuick(r)
		if math.Abs(expected[r] - actual) > tol {
			t.Errorf("expected:%g actual:%g", expected[r], actual)
		}
	}
	//---
	z = nil
	z = A.ZMult(y, z, alpha, beta, false)
	expected = make([]float64, A.Rows())
	for r := 0; r < A.Rows(); r++ {
		s := 0.0
		for c := 0; c < A.Columns(); c++ {
			s += A.GetQuick(r, c) * y.GetQuick(c)
		}
		expected[r] = s * alpha
	}
	for r = 0; r < A.Rows(); r++ {
		actual := z.GetQuick(r)
		if math.Abs(expected[r] - actual) > tol {
			t.Errorf("expected:%g actual:%g", expected[r], actual)
		}
	}

	//transpose
	y = NewVector(A.Rows())
	for i := 0; i < y.Size(); i++ {
		y.SetQuick(i, random.Float64())
	}
	z = RandomVector(A.Columns())
	expected = z.ToArray()
	z = A.ZMult(y, z, alpha, beta, true)
	for r := 0; r < A.Columns(); r++ {
		s := 0.0
		for c := 0; c < A.Rows(); c++ {
			s += A.GetQuick(c, r) * y.GetQuick(c)
		}
		expected[r] = s * alpha + expected[r] * beta
	}
	for r = 0; r < A.Columns(); r++ {
		actual := z.GetQuick(r)
		if math.Abs(expected[r] - actual) > tol {
			t.Errorf("expected:%g actual:%g", expected[r], actual)
		}
	}
	//---
	z = nil
	z = A.ZMult(y, z, alpha, beta, true)
	expected = make([]float64, A.Columns())
	for r := 0; r < A.Columns(); r++ {
		s := 0.0
		for c := 0; c < A.Rows(); c++ {
			s += A.GetQuick(c, r) * y.GetQuick(c)
		}
		expected[r] = s * alpha
	}
	for r := 0; r < A.Columns(); r++ {
		actual := z.GetQuick(r)
		if math.Abs(expected[r] - actual) > tol {
			t.Errorf("expected:%g actual:%g", expected[r], actual)
		}
	}
}
