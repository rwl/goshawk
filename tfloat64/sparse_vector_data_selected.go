
package tfloat64

import "bitbucket.org/rwl/colt"
import l4g "code.google.com/p/log4go"

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
type SelectedSparseVectorData struct {
	*SparseVectorData
	offsets []int // The offsets of visible indexes of this matrix.
	offset  int   // The offset.
}

func (v *SelectedSparseVectorData) GetQuick(index int) float64 {
	return v.elements[v.Index(index)]
}

func (v *SelectedSparseVectorData) SetQuick(index int, value float64) {
	v.elements[v.Index(index)] = value
}

func (v *SelectedSparseVectorData) Index(rank int) int {
	return v.offset + v.offsets[v.Zero() + rank * v.Stride()]
}

func (v *SelectedSparseVectorData) ViewVectorData() VectorData {
	return &SelectedSparseVectorData{
		&SparseVectorData{
			colt.NewCoreVectorData(false, v.Size(), 0, 1),
			v.elements,
		},
		v.offsets, v.offset,
	}
}

func (v *SelectedSparseVectorData) ReshapeMatrix(rows, columns int) (*Matrix, error) {
	if rows * columns != v.Size() {
		return nil, l4g.Error("rows*columns != size")
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

func (v *SelectedSparseVectorData) ReshapeCube(slices, rows, columns int) (*Cube, error) {
	if slices * rows * columns != v.Size() {
		return nil, l4g.Error("slices*rows*columns != size")
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
