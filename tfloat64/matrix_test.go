
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

type zMultMatrix interface {
	Mat
}

func testMatrixZMult(t *testing.T, A zMultMatrix) {
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

type zMultMatrixMatrix interface {
	Mat
	ZMultMatrix(*Matrix, *Matrix, float64, float64, bool, bool)
}

func testMatrixZMultMatrix(t *testing.T, A zMultMatrixMatrix) {
	alpha := 3.0
	beta := 5.0
	C := RandomMatrix(A.Rows(), A.Rows())
	expected := C.ToArray()
	C = A.ZMultMatrix(Bt, C, alpha, beta, false, false)
	for j := 0; j < A.Rows(); j++ {
		for i := 0; i < A.Rows(); i++ {
			s := 0.0
			for k := 0; k < A.Columns(); k++ {
				s += A.GetQuick(i, k) * Bt.GetQuick(k, j)
			}
			expected[i][j] = s * alpha + expected[i][j] * beta
		}
	}
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Rows(); c++ {
			actual := C.GetQuick(r, c)
			if math.Abs(expected[r][c] - actual) > tol {
				t.Errorf("expected:%g actual:%g", expected[r][c], actual)
			}
		}
	}

	//---
	C = nil
	C = A.ZMultMatrix(Bt, C, alpha, beta, false, false)
	expected = make([][]float64, A.Rows(), A.Rows())
	for j := 0; j < A.Rows(); j++ {
		for i := 0; i < A.Rows(); i++ {
			s := 0.0
			for k := 0; k < A.Columns(); k++ {
				s += A.GetQuick(i, k) * Bt.GetQuick(k, j)
			}
			expected[i][j] = s * alpha
		}
	}
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Rows(); c++ {
			actual := C.GetQuick(r, c)
			if math.Abs(expected[r][c] - actual) > tol {
				t.Errorf("expected:%g actual:%g", expected[r][c], actual)
			}
		}
	}

	//transposeA
	C = RandomMatrix(A.Columns(), A.Columns())
	expected = C.ToArray()
	C = A.ZMultMatrix(B, C, alpha, beta, true, false)
	for j := 0; j < A.Columns(); j++ {
		for i := 0; i < A.Columns(); i++ {
			s := 0.0
			for k := 0; k < A.Rows(); k++ {
				s += A.GetQuick(k, i) * B.GetQuick(k, j)
			}
			expected[i][j] = s * alpha + expected[i][j] * beta
		}
	}
	for r := 0; r < A.Columns(); r++ {
		for c := 0; c < A.Columns(); c++ {
			actual := C.GetQuick(r, c)
			if math.Abs(expected[r][c] - actual) > tol {
				t.Errorf("expected:%g actual:%g", expected[r][c], actual)
			}
		}
	}
	//---
	C = nil
	C = A.ZMult(B, C, alpha, beta, true, false)
	expected = make([][]float64, A.Columns(), A.Columns())
	for j := 0; j < A.Columns(); j++ {
		for i := 0; i < A.Columns(); i++ {
			s := 0.0
			for k := 0; k < A.Rows(); k++ {
				s += A.GetQuick(k, i) * B.GetQuick(k, j)
			}
			expected[i][j] = s * alpha
		}
	}
	for r := 0; r < A.Columns(); r++ {
		for c := 0; c < A.Columns(); c++ {
			actual := C.GetQuick(r, c)
			if math.Abs(expected[r][c] - actual) > tol {
				t.Errorf("expected:%g actual:%g", expected[r][c], actual)
			}
		}
	}

	//transposeB
	C = RandomMatrix(A.Rows(), A.Rows())
	expected = C.ToArray()
	C = A.ZMultMatrix(B, C, alpha, beta, false, true)
	for j := 0; j < A.Rows(); j++ {
		for i := 0; i < A.Rows(); i++ {
			s := 0.0
			for k := 0; k < A.Columns(); k++ {
				s += A.GetQuick(i, k) * B.GetQuick(j, k)
			}
			expected[i][j] = s * alpha + expected[i][j] * beta
		}
	}
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Rows(); c++ {
			actual := C.GetQuick(r, c)
			if math.Abs(expected[r][c] - actual) > tol {
				t.Errorf("expected:%g actual:%g", expected[r][c], actual)
			}
		}
	}
	//---
	C = nil
	C = A.ZMultMatrix(B, C, alpha, beta, false, true)
	expected = make([][]float64, A.Rows(), A.Rows())
	for j := 0; j < A.Rows(); j++ {
		for i := 0; i < A.Rows(); i++ {
			s := 0.0
			for k := 0; k < A.Columns(); k++ {
				s += A.GetQuick(i, k) * B.GetQuick(j, k)
			}
			expected[i][j] = s * alpha
		}
	}
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Rows(); c++ {
			actual := C.GetQuick(r, c)
			if math.Abs(expected[r][c] - actual) > tol {
				t.Errorf("expected:%g actual:%g", expected[r][c], actual)
			}
		}
	}
	//transposeA and transposeB
	C = RandomMatrix(A.Columns(), A.Columns())
	expected = C.ToArray()
	C = A.ZMultMatrix(Bt, C, alpha, beta, true, true)
	for j := 0; j < A.Columns(); j++ {
		for i := 0; i < A.Columns(); i++ {
			s := 0.0
			for k := 0; k < A.Rows(); k++ {
				s += A.GetQuick(k, i) * Bt.GetQuick(j, k)
			}
			expected[i][j] = s * alpha + expected[i][j] * beta
		}
	}
	for r := 0; r < A.Columns(); r++ {
		for c := 0; c < A.Columns(); c++ {
			actual := C.GetQuick(r, c)
			if math.Abs(expected[r][c] - actual) > tol {
				t.Errorf("expected:%g actual:%g", expected[r][c], actual)
			}
		}
	}
	//---
	C = nil
	C = A.ZMultMatrix(Bt, C, alpha, beta, true, true)
	expected = make([][]float64, A.Columns(), A.Columns())
	for j := 0; j < A.Columns(); j++ {
		for i := 0; i < A.Columns(); i++ {
			s := 0.0
			for k := 0; k < A.Rows(); k++ {
				s += A.GetQuick(k, i) * Bt.GetQuick(j, k)
			}
			expected[i][j] = s * alpha
		}
	}
	for r := 0; r < A.Columns(); r++ {
		for c := 0; c < A.Columns(); c++ {
			actual := C.GetQuick(r, c)
			if math.Abs(expected[r][c] - actual) > tol {
				t.Errorf("expected:%g actual:%g", expected[r][c], actual)
			}
		}
	}
}

type zSumMatrix interface {
	Mat
	ZSum() float64
}

func testMatrixZSum(t *testing.T, A zSumMatrix) {
	sum := A.ZSum()
	expected := 0
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Columns(); c++ {
			expected += A.GetQuick(r, c)
		}
	}
	if math.Abs(expected - sum) > tol {
		t.Errorf("expected:%g actual:%g", expected, sum)
	}
}
