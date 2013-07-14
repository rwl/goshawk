
package tfloat64

import "bitbucket.org/rwl/colt"

func NewSparseVector(size int) *Vector {
	return &Vector{
		SparseVectorData{
			colt.NewCoreVectorData(false, size, 0, 1),
			make(map[int]float64),
		},
	}
}

type SparseVectorData struct {
	colt.CoreVectorData
	elements map[int]float64 // The elements of the matrix.
}

func (sv SparseVectorData) GetQuick(index int) float64 {
	return sv.elements[sv.Zero() + index * sv.Stride()]
}

func (sv SparseVectorData) SetQuick(index int, value float64) {
	i := sv.Zero() + index * sv.Stride()
	if value == 0.0 {
		delete(sv.elements, i)
	} else {
		sv.elements[i] = value
	}
}

func (sv SparseVectorData) Elements() interface{} {
	return sv.elements
}

func (sv SparseVectorData) Like(size int) colt.VectorData {
	return &SparseVectorData{
		colt.NewCoreVectorData(false, size, 0, 1),
		make(map[int]float64),
	}
}

func (sv SparseVectorData) LikeMatrix(rows, columns int) colt.MatrixData {
	return nil
}

func (sv SparseVectorData) ViewSelectionLike(offsets []int) colt.VectorData {
	return nil
}

func (sv SparseVectorData) View() colt.VectorData {
	return &SparseVectorData{
		colt.NewCoreVectorData(sv.IsView(), sv.Size(), sv.Zero(), sv.Stride()),
		sv.elements,
	}
}

func (sv SparseVectorData) ReshapeMatrix(rows, columns int) (colt.MatrixData, error) {
	return nil, nil
}

func (sv SparseVectorData) ReshapeCube(slices, rows, columns int) (colt.CubeData, error) {
	return nil, nil
}
