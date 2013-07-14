
package colt

type MatrixData interface {
	GetQuick(int, int) float64
	SetQuick(int, int, float64)
	Rows() int
	Columns() int
}
