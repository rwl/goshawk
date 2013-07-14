
package colt

type CubeData interface {
	GetQuick(slice, row, col int) float64
	Slices() int
	Rows() int
	Columns() int
}
