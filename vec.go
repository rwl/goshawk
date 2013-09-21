package goshawk

import "fmt"

// Interface for all vector backends.
type Vec interface {
Base

	// Returns the number of cells.
	Size() int
	Zero() int
	Stride() int
	Index(int) int

	VectorFlip()
	VectorPart(index, width int) error
	VectorStrides(stride int) error
}

// Vector data common to all vector backends.
type CoreVec struct {
	*Core
	size   int // The number of cells this matrix (view) has.
	zero   int // The index of the first element.
	stride int // The number of indexes between any two elements.
}

func NewCoreVec(isView bool, size, zero, stride int) *CoreVec {
	return &CoreVec{&Core{isView}, size, zero, stride}
}

// Returns the number of cells this vector (view) has.
func (v *CoreVec) Size() int {
	return v.size
}

// Returns the number of indexes between any two elements.
func (v *CoreVec) Stride() int {
	return v.stride
}

// Returns the index of the first element.
func (v *CoreVec) Zero() int {
	return v.zero
}

// Returns the position of the given coordinate within the (virtual or non-virtual) internal 1-dimensional array.
func (v *CoreVec) Index(rank int) int {
	return v.zero + rank*v.stride
}

func (v *CoreVec) checkRange(index, width int) error {
	if index < 0 || index + width > v.Size() {
		return fmt.Errorf("index: %d, width: %d, size=%d", index, width, v.Size())
	}
	return nil
}

// Self modifying version of viewFlip(). What used to be index 0 is
// now index Size()-1, ..., what used to be index Size()-1
// is now index 0.
func (v *CoreVec) VectorFlip() {
	if v.size > 0 {
		v.zero += (v.size - 1)*v.stride
		v.stride = -v.stride
		v.isView = true
	}
}

// Self modifying version of ViewPart().
func (v *CoreVec) VectorPart(index, width int) error {
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
func (v *CoreVec) VectorStrides(stride int) error {
	if stride <= 0 {
		return fmt.Errorf("illegal stride: %s", stride)
	}
	v.stride *= stride
	if v.size != 0 {
		v.size = (v.size - 1)/stride + 1
	}
	v.isView = true
	return nil
}
