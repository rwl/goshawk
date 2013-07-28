
package tfloat64

import (
	"bitbucket.org/rwl/colt"
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
type SelectedDenseVectorData struct {
	*DenseVectorData
	offsets []int // The offsets of visible indexes of this matrix.
	offset  int   // The offset.
}

func (v *SelectedDenseVectorData) GetQuick(index int) float64 {
	return v.elements[v.offset + v.offsets[v.Zero() + index*v.Stride()]]
}

func (v *SelectedDenseVectorData) SetQuick(index int, value float64) {
	v.elements[v.offset + v.offsets[v.Zero() + index*v.Stride()]] = value
}

func (v *SelectedDenseVectorData) Index(rank int) int {
	return v.offset + v.offsets[v.Zero() + rank*v.Stride()]
}

func (v *SelectedDenseVectorData) ViewVectorData() VectorData {
	return &SelectedDenseVectorData{
		&DenseVectorData{
			colt.NewCoreVectorData(false, v.Size(), 0, 1),
			v.elements,
		},
		v.offsets, v.offset,
	}
}

func (v *SelectedDenseVectorData) ReshapeMatrix(rows, columns int) (*Matrix, error) {
	if rows * columns != v.Size() {
		return nil, fmt.Errorf("rows*columns != size")
	}
	M := NewMatrix(rows, columns)
	elementsOther := M.Elements().([]float64)
	zeroOther := M.Index(0, 0)

	idx := 0
	for c := 0; c < columns; c++ {
		idxOther := zeroOther + c * M.ColumnStride()
		for r := 0; r < rows; r++ {
			elementsOther[idxOther] = v.GetQuick(idx)
			idxOther += M.RowStride()
			idx++
		}
	}
	return M, nil
}

func (v *SelectedDenseVectorData) ReshapeCube(slices, rows, columns int) (*Cube, error) {
	if slices * rows * columns != v.Size() {
		return nil, fmt.Errorf("slices*rows*columns != size")
	}
	M := NewCube(slices, rows, columns)
	elementsOther := M.Elements().([]float64)
	zeroOther := M.Index(0, 0, 0)

	idx := 0;
	for s := 0; s < slices; s++ {
		for c := 0; c < columns; c++ {
			idxOther := zeroOther + s * M.SliceStride() + c * M.ColumnStride()
			for r := 0; r < rows; r++ {
				elementsOther[idxOther] = v.GetQuick(idx)
				idxOther += M.RowStride()
				idx++
			}
		}
	}
	return M, nil
}
