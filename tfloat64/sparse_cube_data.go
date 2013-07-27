
package tfloat64

import "bitbucket.org/rwl/colt"

func NewSparseCube(slices, rows, columns int) *Cube {
	return &Cube{
		SparseCubeData{
			colt.NewCoreCubeData(false, slices, rows, columns, rows*columns, columns, 1, 0, 0, 0),
			make(map[int]float64),
		},
	}
}

type SparseCubeData struct {
	colt.CoreCubeData
	elements map[int]float64 // The elements of this matrix.
}

func (m SparseCubeData) GetQuick(slice, row, column int) float64 {
	return m.elements[m.Index(slice, row, column)]
}

func (m SparseCubeData) SetQuick(slice, row, column int, value float64) {
	index := m.SliceZero() + slice * m.SliceStride() + m.RowZero() + row * m.RowStride() + m.ColumnZero() + column * m.ColumnStride()
	if value == 0 {
		delete(m.elements, index)
	} else {
		m.elements[index] = value
	}
}

func (m SparseCubeData) Elements() interface{} {
	return m.elements
}
