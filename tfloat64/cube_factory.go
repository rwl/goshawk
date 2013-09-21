
package tfloat64

import common "github.com/rwl/goshawk"

func NewCube(slices, rows, columns int) *Cube {
	return &Cube{
		&DenseCub{
			common.NewCoreCub(false, slices, rows, columns, rows*columns, columns, 1, 0, 0, 0),
			make([]float64, slices*rows*columns),
		},
	}
}

func NewSparseCube(slices, rows, columns int) *Cube {
	return &Cube{
		&SparseCub{
			common.NewCoreCub(false, slices, rows, columns, rows*columns, columns, 1, 0, 0, 0),
			make(map[int]float64),
		},
	}
}
