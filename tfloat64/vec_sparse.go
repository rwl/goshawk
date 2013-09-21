package tfloat64

import (
	common "github.com/rwl/goshawk"
	"fmt"
)

func NewSparseVector(size int) *Vector {
	return &Vector{
		&SparseVec{
			common.NewCoreVec(false, size, 0, 1),
			make(map[int]float64),
		},
	}
}

type SparseVec struct {
	*common.CoreVec
	elements map[int]float64 // The elements of the matrix.
}

func (sv *SparseVec) GetQuick(index int) float64 {
	return sv.elements[sv.Index(index)]
}

func (sv *SparseVec) SetQuick(index int, value float64) {
	i := sv.Index(index)
	if value == 0.0 {
		delete(sv.elements, i)
	} else {
		sv.elements[i] = value
	}
}

func (sv *SparseVec) Elements() interface{} {
	return sv.elements
}

func (sv *SparseVec) Like(size int) Vec {
	return &SparseVec{
		common.NewCoreVec(false, size, 0, 1),
		make(map[int]float64),
	}
}

func (sv *SparseVec) LikeMatrix(rows, columns int) Mat {
	return nil
}

func (sv *SparseVec) ViewSelectionLike(offsets []int) Vec {
	return &SelectedSparseVec{
		&SparseVec{
			common.NewCoreVec(false, len(offsets), 0, 1),
			sv.elements,
		},
		offsets, 0,
	}
}

func (sv *SparseVec) ViewVec() Vec {
	return &SparseVec{
		common.NewCoreVec(sv.IsView(), sv.Size(), sv.Zero(), sv.Stride()),
		sv.elements,
	}
}

func (sv *SparseVec) ReshapeMatrix(rows, columns int) (*Matrix, error) {
	if rows*columns != sv.Size() {
		return nil, fmt.Errorf("rows*columns != size")
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

func (sv *SparseVec) ReshapeCube(slices, rows, columns int) (*Cube, error) {
	if slices*rows*columns != sv.Size() {
		return nil, fmt.Errorf("slices*rows*columns != size")
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
