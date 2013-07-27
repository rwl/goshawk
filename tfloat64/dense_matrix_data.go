
package tfloat64

import "bitbucket.org/rwl/colt"

func NewMatrix(rows, columns int) *Matrix {
	return &Matrix{
		DenseMatrixData{
			colt.NewCoreMatrixData(false, rows, columns, columns, 1, 0, 0),
			make([]float64, rows*columns),
		},
	}
}

type DenseMatrixData struct {
	colt.CoreMatrixData
	elements []float64 // The elements of this matrix.
}

func (m DenseMatrixData) GetQuick(row, column int) float64 {
	return m.elements[m.Index(row, column)]
}

func (m DenseMatrixData) SetQuick(row, column int, value float64) {
	m.elements[m.Index(row, column)] = value
}

func (m DenseMatrixData) Elements() interface{} {
	return m.elements
}
