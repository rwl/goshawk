package tfloat64

import "github.com/rwl/goshawk/common"

type SparseMat struct {
	*common.CoreMat
	elements map[int]float64 // The elements of this matrix.
}

func (m *SparseMat) GetQuick(row, column int) float64 {
	return m.elements[m.RowZero() + row*m.RowStride() + m.ColumnZero() + column*m.ColumnStride()]
}

func (m *SparseMat) SetQuick(row, column int, value float64) {
	index := m.RowZero() + row*m.RowStride() + m.ColumnZero() + column*m.ColumnStride()
	if value == 0 {
		delete(m.elements, index)
	} else {
		m.elements[index] = value
	}
}

func (m *SparseMat) Elements() interface{} {
	return m.elements
}

func (m *SparseMat) Like(rows, columns int) Mat {
	return &SparseMat{
		common.NewCoreMat(false, rows, columns, columns, 1, 0, 0),
		make(map[int]float64),
	}
}
