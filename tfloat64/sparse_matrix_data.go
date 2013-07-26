
package tfloat64

import "bitbucket.org/rwl/colt"

func NewSparseMatrix(rows, columns int) *Matrix {
	return &Matrix{
		SparseMatrixData{
			colt.NewCoreMatrixData(false, rows, columns, columns, 1, 0, 0),
			make(map[int]float64),
		},
	}
}

type SparseMatrixData struct {
	colt.CoreMatrixData
	elements map[int]float64 // The elements of this matrix. TODO: use int64?
}

func (m SparseMatrixData) GetQuick(row, column int) float64 {
	return m.elements[m.RowZero() + row * m.RowStride() + m.ColumnZero() + column * m.ColumnStride()]
}

func (m SparseMatrixData) SetQuick(row, column int, value float64) {
	index := m.RowZero() + row * m.RowStride() + m.ColumnZero() + column * m.ColumnStride()
	if value == 0 {
		delete(m.elements, index)
	} else {
		m.elements[index] = value
	}
}

func (m SparseMatrixData) Elements() interface{} {
	return m.elements
}

func (m SparseMatrixData) Index(row, column int) int {
	return m.RowZero() + row*m.RowStride() + m.ColumnZero() + column*m.ColumnStride()
}
