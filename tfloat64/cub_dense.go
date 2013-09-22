package tfloat64

import "github.com/rwl/goshawk/common"

type DenseCub struct {
	*common.CoreCub
	elements []float64 // The elements of this matrix.
}

func (m *DenseCub) GetQuick(slice, row, column int) float64 {
	return m.elements[m.Index(slice, row, column)]
}

func (m *DenseCub) SetQuick(slice, row, column int, value float64) {
	m.elements[m.Index(slice, row, column)] = value
}

func (m *DenseCub) Elements() interface{} {
	return m.elements
}
