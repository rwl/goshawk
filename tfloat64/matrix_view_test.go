
package tfloat64

import (
	"testing"
	"math"
)

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
