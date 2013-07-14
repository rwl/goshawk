
package colt

import (
	l4g "code.google.com/p/log4go"
)

// Interface for all vector backends.
type VectorData interface {
	// Returns the matrix cell value at coordinate "index".
	//
	// Provided with invalid parameters this method may cause a panic or
	// return invalid values without causing an error. You should only
	// use this method when you are absolutely sure that the coordinate
	// is within bounds.
	// Precondition (unchecked): index < 0 || index >= Size().
	GetQuick(int) float64

	// Sets the matrix cell at coordinate "index" to the specified value.
	//
	// Provided with invalid parameters this method may cause a panic or
	// access illegal indexes without causing an error. You should only use
	// this method when you are absolutely sure that the coordinate is
	// within bounds.
	// Precondition (unchecked): index < 0 || index >= Size().
	SetQuick(int, float64)

	IsView() bool

	// Returns the number of cells.
	Size() int

	Zero() int
	Stride() int

//	Like() VectorData
	Like(int) VectorData
	LikeMatrix(int, int) MatrixData
	ReshapeMatrix(int, int) (MatrixData, error)
	ReshapeCube(int, int, int) (CubeData, error)
	ViewSelectionLike(offsets []int) VectorData
	View() VectorData

	Elements() interface{}

	vFlip()
	vPart(index, width int) error
	vStrides(stride int) error
}

// Vector data common to all vector backends.
type CoreVectorData struct {
	isView bool // Whether the receiver is a view or not.
	size   int  // The number of cells this matrix (view) has.
	zero   int  // The index of the first element.
	stride int  // The number of indexes between any two elements.
}

func NewCoreVectorData(isView bool, size, zero, stride int) CoreVectorData {
	return CoreVectorData{isView, size, zero, stride}
}

// Returns whether the receiver is a view or not.
func (v CoreVectorData) IsView() bool {
	return v.isView
}

// Returns the number of cells this vector (view) has.
func (v CoreVectorData) Size() int {
	return v.size
}

// Returns the number of indexes between any two elements.
func (v CoreVectorData) Stride() int {
	return v.stride
}

// Returns the index of the first element.
func (v CoreVectorData) Zero() int {
	return v.zero
}

func (v CoreVectorData) checkRange(index, width int) error {
	if index < 0 || index + width > v.Size() {
		return l4g.Error("index: %d, width: %d, size=%d", index, width, v.Size())
	}
	return nil
}

// Self modifying version of viewFlip(). What used to be index 0 is
// now index Size()-1, ..., what used to be index Size()-1
// is now index 0.
func (v CoreVectorData) vFlip() {
	if v.size > 0 {
		v.zero += (v.size - 1)*v.stride
		v.stride = -v.stride
		v.isView = true
	}
}

// Self modifying version of ViewPart().
func (v CoreVectorData) vPart(index, width int) error {
	err := v.checkRange(index, width)
	if err != nil {
		return err
	}
	v.zero += v.stride*index
	v.size = width
	v.isView = true
	return nil
}

// Self modifying version of ViewStrides().
func (v CoreVectorData) vStrides(stride int) error {
	if stride <= 0 {
		return l4g.Error("illegal stride: %s", stride)
	}
	v.stride *= stride
	if v.size != 0 {
		v.size = (v.size - 1)/stride + 1
	}
	v.isView = true
	return nil
}
