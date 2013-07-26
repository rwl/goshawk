
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
	return m.elements[m.RowZero() + row*m.RowStride() + m.ColumnZero() + column*m.ColumnStride()]
}

func (m DenseMatrixData) SetQuick(row, column int, value float64) {
	m.elements[m.RowZero() + row*m.RowStride() + m.ColumnZero() + column*m.ColumnStride()] = value
}

func (m DenseMatrixData) Elements() interface{} {
	return m.elements
}

func (m DenseMatrixData) Index(row, column int) int {
	return m.RowZero() + row*m.RowStride() + m.ColumnZero() + column*m.ColumnStride()
}
