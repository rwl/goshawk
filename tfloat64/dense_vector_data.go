
package tfloat64

import (
	l4g "code.google.com/p/log4go"
	"bitbucket.org/rwl/colt"
)

func NewVector(size int) *Vector {
	return &Vector{
		DenseVectorData{
			colt.NewCoreVectorData(false, size, 0, 1),
			make([]float64, size),
		},
	}
}

type DenseVectorData struct {
	colt.CoreVectorData
	elements []float64 // The elements of this matrix.
}

func (v DenseVectorData) GetQuick(index int) float64 {
	return v.elements[v.Zero() + index*v.Stride()]
}

func (v DenseVectorData) SetQuick(index int, value float64) {
	v.elements[v.Zero() + index*v.Stride()] = value
}

func (v DenseVectorData) Elements() interface{} {
	return v.elements
}

func (v DenseVectorData) Like(size int) colt.VectorData {
	return &DenseVectorData{
		colt.NewCoreVectorData(false, size, 0, 1),
		make([]float64, size),
	}
}

func (v DenseVectorData) LikeMatrix(rows, columns int) colt.MatrixData {
	return nil/*DenseMatrixData{
		CoreMatrixData{
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

func (v DenseVectorData) ViewSelectionLike(offsets []int) colt.VectorData {
	return nil/*SelectedDenseVectorData{
		CoreVectorData: CoreVectorData{
			isView: false,
			size: size,
			zero: 0,
			stride: 1,
		},
		elements: v.elements,
		offsets: offsets,
		offset: 0,
	}*/
}

func (v DenseVectorData) View() colt.VectorData {
	return &DenseVectorData{
		colt.NewCoreVectorData(v.IsView(), v.Size(), v.Zero(), v.Stride()),
		v.elements,
	}
}

func (v DenseVectorData) ReshapeMatrix(rows, columns int) (colt.MatrixData, error) {
	if rows * columns != v.Size() {
		return nil, l4g.Error("rows*columns != size")
	}
	/*M, _ := NewMatrix(rows, columns)
	elementsOther := M.Elements().([]float64)
	zeroOther := M.index(0, 0)
	rowStrideOther := M.RowStride()
	columnStrideOther := M.ColumnStride()

	var idxOther int
	idx := v.zero
	for c := 0; c < columns; c++ {
		idxOther = zeroOther + c * columnStrideOther
		for r := 0; r < rows; r++ {
			elementsOther[idxOther] = v.elements[idx]
			idxOther += rowStrideOther
			idx += v.stride
		}
	}
	return M, nil*/
	return nil, nil
}

func (v DenseVectorData) ReshapeCube(slices, rows, columns int) (colt.CubeData, error) {
	if slices * rows * columns != v.Size() {
		return nil, l4g.Error("slices*rows*columns != size")
	}
	/*M := NewCube(slices, rows, columns)
	elementsOther := M.Elements().([]float64)
	zeroOther := M.index(0, 0, 0)
	sliceStrideOther := M.sliceStride()
	rowStrideOther := M.rowStride()
	columnStrideOther := M.columnStride()

	var idxOther int
	idx := v.zero
	for s := 0; s < slices; s++ {
		for c := 0; c < columns; c++ {
			idxOther = zeroOther + s * sliceStrideOther + c * columnStrideOther
			for r := 0; r < rows; r++ {
				elementsOther[idxOther] = v.elements[idx]
				idxOther += rowStrideOther
				idx += v.stride
			}
		}
	}
	return M, nil*/
	return nil, nil
}
