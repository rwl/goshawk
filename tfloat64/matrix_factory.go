
package tfloat64

import common "github.com/rwl/goshawk"

// Returns a new dense matrix with the given number of rows and columns.
func NewMatrix(rows, columns int) *Matrix {
	return &Matrix{
		&DenseMat{
			common.NewCoreMat(false, rows, columns, columns, 1, 0, 0),
			make([]float64, rows*columns),
		},
	}
}

// Returns a new sparse matrix with the given number of rows and columns.
func NewSparseMatrix(rows, columns int) *Matrix {
	return &Matrix{
		&SparseMat{
			common.NewCoreMat(false, rows, columns, columns, 1, 0, 0),
			make(map[int]float64),
		},
	}
}
