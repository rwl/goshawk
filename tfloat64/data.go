
package tfloat64

// Interface for all vector backends.
type VectorData interface {
	// Returns the matrix cell value at coordinate "index".
	//
	// Provided with invalid parameters this method may cause a panic or
	// return invalid values without causing an error. You should only
	// use this method when you are absolutely sure that the coordinate
	// is within bounds.
	// Precondition (unchecked): index < 0 || index >= Size().
	GetQuick(int) float64

	// Sets the matrix cell at coordinate "index" to the specified value.
	//
	// Provided with invalid parameters this method may cause a panic or
	// access illegal indexes without causing an error. You should only use
	// this method when you are absolutely sure that the coordinate is
	// within bounds.
	// Precondition (unchecked): index < 0 || index >= Size().
	SetQuick(int, float64)

	IsView() bool

	// Returns the number of cells.
	Size() int

	Zero() int
	Stride() int

	//	Like() VectorData
	Like(int) VectorData
	LikeMatrix(int, int) MatrixData
	ReshapeMatrix(int, int) (MatrixData, error)
	ReshapeCube(int, int, int) (CubeData, error)
	ViewSelectionLike(offsets []int) VectorData
	View() VectorData

	Elements() interface{}

	VectorFlip()
	VectorPart(index, width int) error
	VectorStrides(stride int) error
}

type MatrixData interface {
	GetQuick(int, int) float64
	SetQuick(int, int, float64)
	IsView() bool
	Rows() int
	Columns() int
	RowStride() int
	ColumnStride() int
	RowZero() int
	ColumnZero() int

	Elements() interface{}

	index(row, column int) int
}

type CubeData interface {
	GetQuick(slice, row, col int) float64
	Slices() int
	Rows() int
	Columns() int
}
