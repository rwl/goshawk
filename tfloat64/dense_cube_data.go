
package tfloat64

import "bitbucket.org/rwl/colt"

func NewCube(slices, rows, columns int) *Cube {
	return &Cube{
		&DenseCubeData{
			colt.NewCoreCubeData(false, slices, rows, columns, rows*columns, columns, 1, 0, 0, 0),
			make([]float64, slices*rows*columns),
		},
	}
}

type DenseCubeData struct {
	*colt.CoreCubeData
	elements []float64 // The elements of this matrix.
}

func (m *DenseCubeData) GetQuick(slice, row, column int) float64 {
	return m.elements[m.Index(slice, row, column)]
}

func (m *DenseCubeData) SetQuick(slice, row, column int, value float64) {
	m.elements[m.Index(slice, row, column)] = value
}

func (m *DenseCubeData) Elements() interface{} {
	return m.elements
}
