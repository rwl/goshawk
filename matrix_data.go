
package colt

type BaseMatrixData interface {
	IsView() bool
	Rows() int
	Columns() int
	RowStride() int
	ColumnStride() int
	RowZero() int
	ColumnZero() int

	Elements() interface{}

	Index(row, column int) int
}

type CoreMatrixData struct {
	isView              bool
	rows, columns       int // The number of columns and rows this matrix (view) has.
	rowStride           int // The number of elements between two rows, i.e. index(i+1,j,k) - index(i,j,k).
	columnStride        int // The number of elements between two columns, i.e. index(i,j+1,k) - index(i,j,k).
	rowZero, columnZero int // The index of the first element.
}

func NewCoreMatrixData(isView bool, rows, columns, rowStride, columnStride, rowZero, columnZero int) CoreMatrixData {
	return CoreMatrixData{isView, rows, columns, rowStride, columnStride, rowZero, columnZero}
}

// Returns whether the receiver is a view or not.
func (v CoreMatrixData) IsView() bool {
	return v.isView
}

// Returns the number of rows this matrix (view) has.
func (m CoreMatrixData) Rows() int {
	return m.rows
}

// Returns the number of columns this matrix (view) has.
func (m CoreMatrixData) Columns() int {
	return m.columns
}

// Returns the number of elements between two rows, i.e. index(i+1,j,k) - index(i,j,k).
func (m CoreMatrixData) RowStride() int {
	return m.rowStride
}

// The number of elements between two columns, i.e. index(i,j+1,k) - index(i,j,k).
func (m CoreMatrixData) ColumnStride() int {
	return m.columnStride
}

// Returns the index of the first element.
func (m CoreMatrixData) RowZero() int {
	return m.rowZero
}

// Returns the index of the first element.
func (m CoreMatrixData) ColumnZero() int {
	return m.columnZero
}

func (m CoreMatrixData) Index(row, column int) int {
	return m.RowZero() + row*m.RowStride() + m.ColumnZero() + column*m.ColumnStride()
}
