
package tfloat64

import "bitbucket.org/rwl/colt"
import l4g "code.google.com/p/log4go"

func NewSparseVector(size int) *Vector {
	return &Vector{
		&SparseVectorData{
			colt.NewCoreVectorData(false, size, 0, 1),
			make(map[int]float64),
		},
	}
}

type SparseVectorData struct {
	*colt.CoreVectorData
	elements map[int]float64 // The elements of the matrix.
}

func (sv *SparseVectorData) GetQuick(index int) float64 {
	return sv.elements[sv.Index(index)]
}

func (sv *SparseVectorData) SetQuick(index int, value float64) {
	i := sv.Index(index)
	if value == 0.0 {
		delete(sv.elements, i)
	} else {
		sv.elements[i] = value
	}
}

func (sv *SparseVectorData) Elements() interface{} {
	return sv.elements
}

func (sv *SparseVectorData) Like(size int) VectorData {
	return &SparseVectorData{
		colt.NewCoreVectorData(false, size, 0, 1),
		make(map[int]float64),
	}
}

func (sv *SparseVectorData) LikeMatrix(rows, columns int) MatrixData {
	return nil
}

func (sv *SparseVectorData) ViewSelectionLike(offsets []int) VectorData {
	return &SelectedSparseVectorData{
		&SparseVectorData{
			colt.NewCoreVectorData(false, len(offsets), 0, 1),
			sv.elements,
		},
		offsets, 0,
	}
}

func (sv *SparseVectorData) ViewVectorData() VectorData {
	return &SparseVectorData{
		colt.NewCoreVectorData(sv.IsView(), sv.Size(), sv.Zero(), sv.Stride()),
		sv.elements,
	}
}

func (sv *SparseVectorData) ReshapeMatrix(rows, columns int) (*Matrix, error) {
	if rows * columns != sv.Size() {
		return nil, l4g.Error("rows*columns != size")
	}
	M := NewSparseMatrix(rows, columns)
	idx := 0
	for c := 0; c < columns; c++ {
		for r := 0; r < rows; r++ {
			elem := sv.GetQuick(idx)
			idx++
			if elem != 0 {
				M.SetQuick(r, c, elem)
			}
		}
	}
	return M, nil
}

func (sv *SparseVectorData) ReshapeCube(slices, rows, columns int) (*Cube, error) {
	if slices * rows * columns != sv.Size() {
		return nil, l4g.Error("slices*rows*columns != size")
	}
	M := NewSparseCube(slices, rows, columns)
	idx := 0
	for s := 0; s < slices; s++ {
		for c := 0; c < columns; c++ {
			for r := 0; r < rows; r++ {
				elem := sv.GetQuick(idx)
				idx++
				if elem != 0 {
					M.SetQuick(s, r, c, elem)
				}
			}
		}
	}
	return M, nil
}
