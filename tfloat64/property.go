package tfloat64

import "math"

type Property struct {
	tolerance float64
}

// Returns whether all cells of the given matrix A are equal to the
// given value. The result is true if and only if
// A != nil and !(math.Abs(value - A[i]) > tolerance)
// holds for all coordinates.
func (p *Property) VectorEqualsValue(v Vec, value float64) bool {
	for i := 0; i < v.Size(); i++ {
		x := v.GetQuick(i)
		diff := math.Abs(value - x)
		if (diff != diff) && ((value != value && x != x) || value == x) {
			diff = 0
		}
		if diff > p.tolerance {
			return false
		}
	}
	return true
}

// Returns whether both given matrices A and B are equal.
// The result is true if A==B. Otherwise, the result is
// true if and only if both arguments are != nil, have
// the same size and !(math.Abs(A[i] - B[i]) > tolerance) holds
// for all indexes.
func (p *Property) VectorEqualsVector(A Vec, B Vec) bool {
	if A == B {
		return true
	}
	if !(A != nil && B != nil) {
		return false
	}
	size := A.Size()
	if size != B.Size() {
		return false
	}

	for i := 0; i < size; i++ {
		x := A.GetQuick(i)
		value := B.GetQuick(i)
		diff := math.Abs(value - x)
		if (diff != diff) && ((value != value && x != x) || value == x) {
			diff = 0
		}
		if diff > p.tolerance {
			return false
		}
	}
	return true
}

// Returns whether all cells of the given matrix A are equal to the
// given value. The result is true if and only if
// A != nil and
// ! (math.Abs(value - A[row, col]) > tolerance) holds for all
// coordinates.
func (p *Property) MatrixEqualsValue(A Mat, value float64) bool {
	if A == nil {
		return false
	}
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Columns(); c++ {
			x := A.GetQuick(r, c)
			diff := math.Abs(value - x)
			if (diff != diff) && ((value != value && x != x) || value == x) {
				diff = 0
			}
			if !(diff <= p.tolerance) {
				return false
			}
		}
	}
	return true
}

// Returns whether both given matrices A and B are equal.
// The result is true if A==B. Otherwise, the result is
// true if and only if both arguments are != nil, have
// the same number of columns and rows and
// ! (math.Abs(A[row,col] - B[row,col]) > tolerance) holds for
// all coordinates.
func (p *Property) MatrixEqualsMatrix(A, B Mat) bool {
	if A == B {
		return true
	}
	if !(A != nil && B != nil) {
		return false
	}
	if A.Columns() != B.Columns() || A.Rows() != B.Rows() {
		return false
	}
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Columns(); c++ {
			x := A.GetQuick(r, c)
			value := B.GetQuick(r, c)
			diff := math.Abs(value - x)
			if (diff != diff) && ((value != value && x != x) || value == x) {
				diff = 0
			}
			if !(diff <= p.tolerance) {
				return false
			}
		}
	}
	return true
}
