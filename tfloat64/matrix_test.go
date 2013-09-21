package tfloat64

import (
	"math/rand"
)

const (
	nrows    = 13
	ncolumns = 17
)

func makeDenseMatrix() (*Matrix) {
	A := NewMatrix(nrows, ncolumns)
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Columns(); c++ {
			A.SetQuick(r, c, rand.Float64())
		}
	}
	return A
}

func makeSparseMatrix() (*Matrix) {
	A := NewSparseMatrix(nrows, ncolumns)
	for r := 0; r < A.Rows(); r++ {
		for c := 0; c < A.Columns(); c++ {
			A.SetQuick(r, c, rand.Float64())
		}
	}
	return A
}
