package tfloat64

import (
	common "github.com/rwl/goshawk"
	"fmt"
)

// Selection view on sparse 1-d matrices holding float64 elements.
//
// Instances of this type are typically constructed via viewIndexes
// methods on some source vector. From a user point of view there is
// nothing special about this type; it presents the same functionality
// with the same signatures and semantics as its original vector while
// introducing no additional functionality.
//
// This class uses no delegation. Its instances point directly to the
// data. Cell addressing overhead is 1 additional array index access
// per get/set.
type SelectedSparseVec struct {
	*SparseVec
	offsets []int // The offsets of visible indexes of this matrix.
	offset  int   // The offset.
}

func (v *SelectedSparseVec) GetQuick(index int) float64 {
	return v.elements[v.Index(index)]
}

func (v *SelectedSparseVec) SetQuick(index int, value float64) {
	v.elements[v.Index(index)] = value
}

func (v *SelectedSparseVec) Index(rank int) int {
	return v.offset + v.offsets[v.Zero() + rank*v.Stride()]
}

func (v *SelectedSparseVec) ViewVec() Vec {
	return &SelectedSparseVec{
		&SparseVec{
			common.NewCoreVec(false, v.Size(), 0, 1),
			v.elements,
		},
		v.offsets, v.offset,
	}
}

func (v *SelectedSparseVec) ReshapeMatrix(rows, columns int) (*Matrix, error) {
	if rows*columns != v.Size() {
		return nil, fmt.Errorf("rows*columns != size")
	}
	M := NewSparseMatrix(rows, columns)
	idx := 0
	for c := 0; c < columns; c++ {
		for r := 0; r < rows; r++ {
			M.SetQuick(r, c, v.GetQuick(idx))
			idx++
		}
	}
	return M, nil
}

func (v *SelectedSparseVec) ReshapeCube(slices, rows, columns int) (*Cube, error) {
	if slices*rows*columns != v.Size() {
		return nil, fmt.Errorf("slices*rows*columns != size")
	}
	M := NewSparseCube(slices, rows, columns)
	idx := 0
	for s := 0; s < slices; s++ {
		for c := 0; c < columns; c++ {
			for r := 0; r < rows; r++ {
				M.SetQuick(s, r, c, v.GetQuick(idx))
				idx++
			}
		}
	}
	return M, nil
}
