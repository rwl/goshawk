package tfloat64

import (
	"testing"
	"math/rand"
	"math"
)

type assignMatrix interface {
	Mat
	Assign(value float64) *Matrix
}

func TestDenseMatrixAssign(t *testing.T) {
	A := makeDenseMatrix()
	testMatrixAssign(t, A)
}

func TestSparseMatrixAssign(t *testing.T) {
	A := makeSparseMatrix()
	testMatrixAssign(t, A)
}

func testMatrixAssign(t *testing.T, A assignMatrix) {
	value := rand.Float64()
	A.Assign(value)
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Columns(); c++ {
			result := A.GetQuick(r, c)
			if math.Abs(value - result) > tol {
				t.Errorf("expected:%g actual:%g", value, result)
			}
		}
	}
}

type assignArrayMatrix interface {
	Mat
	AssignArray(values [][]float64) (*Matrix, error)
}

func TestDenseMatrixAssignArray(t *testing.T) {
	A := makeDenseMatrix()
	testMatrixAssignArray(t, A)
}

func TestSparseMatrixAssignArray(t *testing.T) {
	A := makeSparseMatrix()
	testMatrixAssignArray(t, A)
}

func testMatrixAssignArray(t *testing.T, A assignArrayMatrix) {
	expected := make([][]float64, A.Rows())
	for r := 0; r < A.Rows(); r++ {
		expected[r] = make([]float64, A.Columns())
		for c := 0; c < A.Columns(); c++ {
			expected[r][c] = rand.Float64()
		}
	}
	A.AssignArray(expected)
	for r := 0; r < A.Rows(); r++ {
		if len(expected[r]) != A.Columns() {
			t.Errorf("expected:%g actual:%g", len(expected[r]), A.Columns())
		}
		for c := 0; c < A.Columns(); c++ {
			if math.Abs(expected[r][c] - A.GetQuick(r, c)) > tol {
				t.Errorf("expected:%g actual:%g", expected[r][c], A.GetQuick(r, c))
			}
		}
	}
}

type assignFuncMatrix interface {
	Mat
	AssignFunc(f Float64Func) *Matrix
	Copy() *Matrix
}

func TestDenseMatrixAssignFunc(t *testing.T) {
	A := makeDenseMatrix()
	testMatrixAssignFunc(t, A)
}

func TestSparseMatrixAssignFunc(t *testing.T) {
	A := makeSparseMatrix()
	testMatrixAssignFunc(t, A)
}

func testMatrixAssignFunc(t *testing.T, A assignFuncMatrix) {
	Acopy := A.Copy()
	A.AssignFunc(math.Acos)
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Columns(); c++ {
			expected := math.Acos(Acopy.GetQuick(r, c))
			if math.Abs(expected - A.GetQuick(r, c)) > tol {
				t.Errorf("expected:%g actual:%g", expected, A.GetQuick(r, c))
			}
		}
	}
}

type assignMatrixMatrix interface {
	Mat
	AssignMatrix(other Mat) (*Matrix, error)
}

func TestDenseMatrixAssignMatrix(t *testing.T) {
	A := makeDenseMatrix()
	B := makeDenseMatrix()
	testMatrixAssignMatrix(t, A, B)
}

func TestSparseMatrixAssignMatrix(t *testing.T) {
	A := makeSparseMatrix()
	B := makeSparseMatrix()
	testMatrixAssignMatrix(t, A, B)
}

func testMatrixAssignMatrix(t *testing.T, A, B assignMatrixMatrix) {
	A.AssignMatrix(B)
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Columns(); c++ {
			expected := B.GetQuick(r, c)
			actual := A.GetQuick(r, c)
			if math.Abs(expected - actual) > tol {
				t.Errorf("expected:%g actual:%g", expected, actual)
			}
		}
	}
}

type assignMatrixFuncMatrix interface {
	Mat
	AssignMatrixFunc(y Mat, f Float64Float64Func) (*Matrix, error)
	Copy() *Matrix
}

func TestDenseMatrixAssignMatrixFunc(t *testing.T) {
	A := makeDenseMatrix()
	B := makeDenseMatrix()
	testMatrixAssignMatrixFunc(t, A, B)
}

func TestSparseMatrixAssignMatrixFunc(t *testing.T) {
	A := makeSparseMatrix()
	B := makeSparseMatrix()
	testMatrixAssignMatrixFunc(t, A, B)
}

func testMatrixAssignMatrixFunc(t *testing.T, A, B assignMatrixFuncMatrix) {
	Acopy := A.Copy()
	A.AssignMatrixFunc(B, Plus)
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Columns(); c++ {
			expected := Acopy.GetQuick(r, c) + B.GetQuick(r, c)
			actual := A.GetQuick(r, c)
			if math.Abs(expected - actual) > tol {
				t.Errorf("expected:%g actual:%g", expected, actual)
			}
		}
	}
}

type assignMatrixFuncSelection interface {
	Mat
	AssignMatrixFuncSelection(y Mat, f Float64Float64Func, rowList, columnList []int) (*Matrix, error)
	Copy() *Matrix
}

func TestDenseMatrixAssignMatrixFuncSelection(t *testing.T) {
	A := makeDenseMatrix()
	B := makeDenseMatrix()
	testMatrixAssignMatrixFuncSelection(t, A, B)
}

func TestSparseMatrixAssignMatrixFuncSelection(t *testing.T) {
	A := makeSparseMatrix()
	B := makeSparseMatrix()
	testMatrixAssignMatrixFuncSelection(t, A, B)
}

func testMatrixAssignMatrixFuncSelection(t *testing.T, A, B assignMatrixFuncSelection) {
	var rowList []int
	var columnList []int
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Columns(); c++ {
			rowList = append(rowList, r)
			columnList = append(columnList, c)
		}
	}
	Acopy := A.Copy()
	A.AssignMatrixFuncSelection(B, Div, rowList, columnList)
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Columns(); c++ {
			expected := Acopy.GetQuick(r, c)/B.GetQuick(r, c)
			actual := A.GetQuick(r, c)
			if math.Abs(expected - actual) > tol {
				t.Errorf("expected:%g actual:%g", expected, actual)
			}
		}
	}
}

type assignProcedureMatrix interface {
	Mat
	AssignProcedure(cond Float64Procedure, value float64) *Matrix
	Copy() *Matrix
}

func TestDenseMatrixAssignProcedure(t *testing.T) {
	A := makeDenseMatrix()
	testMatrixAssignProcedure(t, A)
}

func TestSparseMatrixAssignProcedure(t *testing.T) {
	A := makeSparseMatrix()
	testMatrixAssignProcedure(t, A)
}

func testMatrixAssignProcedure(t *testing.T, A assignProcedureMatrix) {
	procedure := func(element float64) bool {
		if math.Abs(element) > 0.1 {
			return true
		} else {
			return false
		}
	}
	Acopy := A.Copy()
	A.AssignProcedure(procedure, -1.0)
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Columns(); c++ {
			var expected float64
			if math.Abs(Acopy.GetQuick(r, c)) > 0.1 {
				expected = -1.0
			} else {
				expected = Acopy.GetQuick(r, c)
			}
			actual := A.GetQuick(r, c)
			if math.Abs(expected - actual) > tol {
				t.Errorf("expected:%g actual:%g", expected, actual)
			}
		}
	}
}

type assignProcedureFuncMatrix interface {
	Mat
	AssignProcedureFunc(cond Float64Procedure, f Float64Func) *Matrix
	Copy() *Matrix
}

func TestDenseMatrixAssignProcedureFunc(t *testing.T) {
	A := makeDenseMatrix()
	testMatrixAssignProcedureFunc(t, A)
}

func TestSparseMatrixAssignProcedureFunc(t *testing.T) {
	A := makeSparseMatrix()
	testMatrixAssignProcedureFunc(t, A)
}

func testMatrixAssignProcedureFunc(t *testing.T, A assignProcedureFuncMatrix) {
	procedure := func(element float64) bool {
		if math.Abs(element) > 0.1 {
			return true
		} else {
			return false
		}
	}
	Acopy := A.Copy()
	A.AssignProcedureFunc(procedure, math.Tan)
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Columns(); c++ {
			var expected float64
			if math.Abs(Acopy.GetQuick(r, c)) > 0.1 {
				expected = math.Tan(Acopy.GetQuick(r, c))
			} else {
				expected = Acopy.GetQuick(r, c)
			}
			actual := A.GetQuick(r, c)
			if math.Abs(expected - actual) > tol {
				t.Errorf("expected:%g actual:%g", expected, actual)
			}
		}
	}
}
