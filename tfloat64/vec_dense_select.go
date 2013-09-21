package tfloat64

import (
	common "github.com/rwl/goshawk"
	"fmt"
)

// Selection view on dense 1-d matrices holding float64 elements.
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
type SelectedDenseVec struct {
	*DenseVec
	offsets []int // The offsets of visible indexes of this matrix.
	offset  int   // The offset.
}

func (v *SelectedDenseVec) GetQuick(index int) float64 {
	return v.elements[v.offset + v.offsets[v.Zero() + index*v.Stride()]]
}

func (v *SelectedDenseVec) SetQuick(index int, value float64) {
	v.elements[v.offset + v.offsets[v.Zero() + index*v.Stride()]] = value
}

func (v *SelectedDenseVec) Index(rank int) int {
	return v.offset + v.offsets[v.Zero() + rank*v.Stride()]
}

func (v *SelectedDenseVec) ViewVec() Vec {
	return &SelectedDenseVec{
		&DenseVec{
			common.NewCoreVec(false, v.Size(), 0, 1),
			v.elements,
		},
		v.offsets, v.offset,
	}
}

func (v *SelectedDenseVec) ReshapeMatrix(rows, columns int) (*Matrix, error) {
	if rows*columns != v.Size() {
		return nil, fmt.Errorf("rows*columns != size")
	}
	M := NewMatrix(rows, columns)
	elementsOther := M.Elements().([]float64)
	zeroOther := M.Index(0, 0)

	idx := 0
	for c := 0; c < columns; c++ {
		idxOther := zeroOther + c*M.ColumnStride()
		for r := 0; r < rows; r++ {
			elementsOther[idxOther] = v.GetQuick(idx)
			idxOther += M.RowStride()
			idx++
		}
	}
	return M, nil
}

func (v *SelectedDenseVec) ReshapeCube(slices, rows, columns int) (*Cube, error) {
	if slices*rows*columns != v.Size() {
		return nil, fmt.Errorf("slices*rows*columns != size")
	}
	M := NewCube(slices, rows, columns)
	elementsOther := M.Elements().([]float64)
	zeroOther := M.Index(0, 0, 0)

	idx := 0;
	for s := 0; s < slices; s++ {
		for c := 0; c < columns; c++ {
			idxOther := zeroOther + s*M.SliceStride() + c*M.ColumnStride()
			for r := 0; r < rows; r++ {
				elementsOther[idxOther] = v.GetQuick(idx)
				idxOther += M.RowStride()
				idx++
			}
		}
	}
	return M, nil
}
