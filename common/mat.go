package common

type Mat interface {
Base

	Rows() int
	Columns() int

	RowStride() int
	ColumnStride() int

	RowZero() int
	ColumnZero() int

	Index(int, int) int
}

type CoreMat struct {
	*Core
	rows,  columns       int // The number of columns and rows this matrix (view) has.
	rowStride            int // The number of elements between two rows, i.e. index(i+1,j,k) - index(i,j,k).
	columnStride         int // The number of elements between two columns, i.e. index(i,j+1,k) - index(i,j,k).
	rowZero,  columnZero int // The index of the first element.
}

func NewCoreMat(isView bool, rows, columns, rowStride, columnStride, rowZero, columnZero int) *CoreMat {
	return &CoreMat{&Core{isView}, rows, columns, rowStride, columnStride, rowZero, columnZero}
}

// Returns the number of rows this matrix (view) has.
func (m *CoreMat) Rows() int {
	return m.rows
}

// Returns the number of columns this matrix (view) has.
func (m *CoreMat) Columns() int {
	return m.columns
}

// Returns the number of elements between two rows, i.e. index(i+1,j,k) - index(i,j,k).
func (m *CoreMat) RowStride() int {
	return m.rowStride
}

// The number of elements between two columns, i.e. index(i,j+1,k) - index(i,j,k).
func (m *CoreMat) ColumnStride() int {
	return m.columnStride
}

// Returns the index of the first element.
func (m *CoreMat) RowZero() int {
	return m.rowZero
}

// Returns the index of the first element.
func (m *CoreMat) ColumnZero() int {
	return m.columnZero
}

// Returns the position of the given coordinate within the (virtual or non-virtual) internal 1-dimensional array.
func (m *CoreMat) Index(row, column int) int {
	return m.RowZero() + row*m.RowStride() + m.ColumnZero() + column*m.ColumnStride()
}
