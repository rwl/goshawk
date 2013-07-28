
package tfloat64

import (
	"testing"
	"math/rand"
	"math"
)

type assignMatrix interface {
	MatrixData
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
	MatrixData
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
	MatrixData
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
	MatrixData
	AssignMatrix(other MatrixData) (*Matrix, error)
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
