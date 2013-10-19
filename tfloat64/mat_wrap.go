
package tfloat64

import "github.com/rwl/goshawk/common"


type WrapperMat struct {
	*common.CoreMat
	content Mat // The elements of the matrix.
}

func (m *WrapperMat) GetQuick(row, column int) float64 {
	return m.content.GetQuick(row, column)
}

func (m *WrapperMat) SetQuick(row, column int, value float64) {
	m.content.SetQuick(row, column)
}

func (m *WrapperMat) Elements() interface{} {
	return m.content.Elements()
}

func (m *WrapperMat) Like(rows, columns int) Mat {
	return m.content.Like(rows, columns)
}

func (m *WrapperMat) LikeVector(size int) Vec {
	return m.content.LikeVector(size)
}
