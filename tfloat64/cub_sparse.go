package tfloat64

import "github.com/rwl/goshawk/common"

type SparseCub struct {
	*common.CoreCub
	elements map[int]float64 // The elements of this matrix.
}

func (m *SparseCub) GetQuick(slice, row, column int) float64 {
	return m.elements[m.Index(slice, row, column)]
}

func (m *SparseCub) SetQuick(slice, row, column int, value float64) {
	index := m.SliceZero() + slice*m.SliceStride() + m.RowZero() + row*m.RowStride() + m.ColumnZero() + column*m.ColumnStride()
	if value == 0 {
		delete(m.elements, index)
	} else {
		m.elements[index] = value
	}
}

func (m *SparseCub) Elements() interface{} {
	return m.elements
}
