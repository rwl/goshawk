package tfloat64

import (
	"fmt"
	"log"
)

// Constructs and returns a new flip view. What used to be index
// 0 is now index size()-1, ..., what used to be index
// size()-1 is now index 0. The returned view is backed by
// this matrix, so changes in the returned view are reflected in this
// matrix, and vice-versa.
func (v *Vector) ViewFlip() *Vector {
	view := v.ViewVector()
	view.VectorFlip()
	return view
}

// Constructs and returns a new sub-range view that is a
// width sub matrix starting at index.
//
// Operations on the returned view can only be applied to the restricted
// range.
// Note that the view is really just a range restriction: The
// returned matrix is backed by this matrix, so changes in the returned
// matrix are reflected in this matrix, and vice-versa.
//
// The view contains the cells from index..index+width-1. and has
// view.size() == width. A view's legal coordinates are again zero
// based, as usual. In other words, legal coordinates of the view are
// 0 .. view.size()-1==width-1.
func (v *Vector) ViewPart(index, width int) *Vector {
	view := v.ViewVector()
	view.VectorPart(index, width)
	return view
}

// Constructs and returns a new "selection view" that is a vector
// holding the indicated cells. There holds
// view.Size == len(indexes) and
// view.Get(i) == this.Get(indexes[i]). Indexes can occur multiple
// times and can be in arbitrary order.
// Example:
//
// 	 this     = (0,0,8,0,7)
// 	 indexes  = (0,2,4,2)
// 	 -->
// 	 view     = (0,8,7,8)
//
// Note that modifying "indexes" after this call has returned has no
// effect on the view. The returned view is backed by this matrix, so
// changes in the returned view are reflected in this matrix, and
// vice-versa. To indicate that all cells shall be visible,
// simply set this parameter to nil.
func (v *Vector) View(indexes []int) (*Vector, error) {
	// check for "all"
	if indexes == nil {
		indexes = make([]int, v.Size())
		for i := 0; i < v.Size(); i++ {
			indexes[i] = i
		}
	} else {
		for _, index := range indexes {
			if index < 0 || index >= v.Size() {
				return nil, fmt.Errorf("Attempted to access %s at index=%d",
					v.StringShort(), index)
			}
		}
	}
	offsets := make([]int, len(indexes))
	for i, idx := range indexes {
		offsets[i] = v.Index(idx)
	}
	return &Vector{v.ViewSelectionLike(offsets)}, nil
}

// Constructs and returns a new selection view that is a matrix
// holding the cells matching the given condition. Applies the condition to
// each cell and takes only those cells where
// condition(get(i)) yields true.
//
// Example:
//
// 	 // extract and view all cells with even value
// 	 matrix = 0 1 2 3
// 	 matrix.ViewSelectionProcedure(
// 	    func(a float64) bool { return a % 2 == 0 }
// 	 )
// 	 -->
// 	 matrix == 0 2
func (v *Vector) ViewProcedure(condition Float64Procedure) *Vector {
	matches := make([]int, 0)
	for i := 0; i < v.Size(); i++ {
		if condition(v.GetQuick(i)) {
			matches = append(matches, i)
		}
	}
	view, _ := v.View(matches)
	return view
}

// Sorts the vector into ascending order, according to the natural
// ordering. This sort is guaranteed to be stable.
func (v *Vector) ViewSorted() *Vector {
	log.Fatal("sort not implemented")
	return v//mergeSort.sortVector(v) TODO: implement sort
}

// Constructs and returns a new stride view which is a sub matrix
// consisting of every i-th cell. More specifically, the view has size
// this.size()/stride holding cells this.get(i*stride) for
// all i = 0..size()/stride - 1.
func (v *Vector) ViewStrides(stride int) *Vector {
	view := v.ViewVector()
	view.VectorStrides(stride)
	return view
}
