package tfloat64

import "github.com/rwl/goshawk/common"

type DenseMat struct {
	*common.CoreMat
	elements []float64 // The elements of this matrix.
}

func (m *DenseMat) GetQuick(row, column int) float64 {
	return m.elements[m.Index(row, column)]
}

func (m *DenseMat) SetQuick(row, column int, value float64) {
	m.elements[m.Index(row, column)] = value
}

func (m *DenseMat) Elements() interface{} {
	return m.elements
}

func (m *DenseMat) Like(rows, columns int) Mat {
	return &DenseMat{
		common.NewCoreMat(false, rows, columns, columns, 1, 0, 0),
		make([]float64, rows*columns),
	}
}
