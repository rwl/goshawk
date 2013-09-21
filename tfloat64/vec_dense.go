package tfloat64

import (
	common "github.com/rwl/goshawk"
	"fmt"
)

func NewVector(size int) *Vector {
	return &Vector{
		&DenseVec{
			common.NewCoreVec(false, size, 0, 1),
			make([]float64, size),
		},
	}
}

type DenseVec struct {
	*common.CoreVec
	elements []float64 // The elements of this matrix.
}

func (v *DenseVec) GetQuick(index int) float64 {
	return v.elements[v.Index(index)]
}

func (v *DenseVec) SetQuick(index int, value float64) {
	v.elements[v.Index(index)] = value
}

func (v *DenseVec) Elements() interface{} {
	return v.elements
}

func (v *DenseVec) Like(size int) Vec {
	return &DenseVec{
		common.NewCoreVec(false, size, 0, 1),
		make([]float64, size),
	}
}

func (v *DenseVec) LikeMatrix(rows, columns int) Mat {
	return nil/*DenseMat{
		CoreMat{
			isView: false,
			columns: columns,
			rows: rows,
			rowStride: 1,
			columnStride: 1,
			rowZero: 0,
			columnZero: 0,
		},
		elements: make([]float64, rows*columns),
	}*/
}

func (v *DenseVec) ViewSelectionLike(offsets []int) Vec {
	return &SelectedDenseVec{
		&DenseVec{
			common.NewCoreVec(false, len(offsets), 0, 1),
			v.elements,
		},
		offsets, 0,
	}
}

func (v *DenseVec) ViewVec() Vec {
	return &DenseVec{
		common.NewCoreVec(v.IsView(), v.Size(), v.Zero(), v.Stride()),
		v.elements,
	}
}

func (v *DenseVec) ReshapeMatrix(rows, columns int) (*Matrix, error) {
	if rows*columns != v.Size() {
		return nil, fmt.Errorf("rows*columns != size")
	}
	M := NewMatrix(rows, columns)
	elementsOther := M.Elements().([]float64)
	zeroOther := M.Index(0, 0)

	var idxOther int
	idx := v.Zero()
	for c := 0; c < columns; c++ {
		idxOther = zeroOther + c*M.ColumnStride()
		for r := 0; r < rows; r++ {
			elementsOther[idxOther] = v.elements[idx]
			idxOther += M.RowStride()
			idx += v.Stride()
		}
	}
	return M, nil
}

func (v *DenseVec) ReshapeCube(slices, rows, columns int) (*Cube, error) {
	if slices*rows*columns != v.Size() {
		return nil, fmt.Errorf("slices*rows*columns != size")
	}
	M := NewCube(slices, rows, columns)
	elementsOther := M.Elements().([]float64)
	zeroOther := M.Index(0, 0, 0)

	var idxOther int
	idx := v.Zero()
	for s := 0; s < slices; s++ {
		for c := 0; c < columns; c++ {
			idxOther = zeroOther + s*M.SliceStride() + c*M.ColumnStride()
			for r := 0; r < rows; r++ {
				elementsOther[idxOther] = v.elements[idx]
				idxOther += M.RowStride()
				idx += v.Stride()
			}
		}
	}
	return M, nil
}
