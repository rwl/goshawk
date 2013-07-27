
package colt

type BaseCubeData interface {
	BaseData
	Slices() int
	Rows() int
	Columns() int

	SliceStride() int
	RowStride() int
	ColumnStride() int

	SliceZero() int
	RowZero() int
	ColumnZero() int

	Index(int, int, int) int
}

type CoreCubeData struct {
	CoreData

	// The number of slices this cube (view) has.
	slices int

	// The number of rows this cube (view) has.
	rows int

	// The number of columns this cube (view) has.
	columns int

	// The number of elements between two slices, i.e. index(k+1,i,j) - index(k,i,j).
	sliceStride int

	// The number of elements between two rows, i.e. index(k,i+1,j) - index(k,i,j).
	rowStride int

	// The number of elements between two columns, i.e. index(k,i,j+1) - index(k,i,j).
	columnStride int

	// The index of the first element.
	sliceZero, rowZero, columnZero int
}

func NewCoreCubeData(isView bool, slices, rows, columns, sliceStride, rowStride, columnStride, sliceZero, rowZero, columnZero int) CoreCubeData {
	return CoreCubeData{CoreData{isView}, slices, rows, columns, sliceStride, rowStride, columnStride, sliceZero, rowZero, columnZero}
}

// Returns the number of slices this cube (view) has.
func (m CoreCubeData) Slices() int {
	return m.slices
}

// Returns the number of rows this cube (view) has.
func (m CoreCubeData) Rows() int {
	return m.rows
}

// Returns the number of columns this cube (view) has.
func (m CoreCubeData) Columns() int {
	return m.columns
}

// Returns the number of elements between two slices, i.e. index(k+1,i,j) - index(k,i,j).
func (m CoreCubeData) SliceStride() int {
	return m.sliceStride
}

// Returns the number of elements between two rows, i.e. index(i+1,j,k) - index(i,j,k).
func (m CoreCubeData) RowStride() int {
	return m.rowStride
}

// The number of elements between two columns, i.e. index(i,j+1,k) - index(i,j,k).
func (m CoreCubeData) ColumnStride() int {
	return m.columnStride
}

// Returns the index of the first element.
func (m CoreCubeData) SliceZero() int {
	return m.sliceZero
}

// Returns the index of the first element.
func (m CoreCubeData) RowZero() int {
	return m.rowZero
}

// Returns the index of the first element.
func (m CoreCubeData) ColumnZero() int {
	return m.columnZero
}

// Returns the position of the given coordinate within the (virtual or non-virtual) internal 1-dimensional array.
func (m CoreCubeData) Index(slice, row, column int) int {
	return m.SliceZero() + slice * m.SliceStride() + m.RowZero() + row * m.RowStride() + m.ColumnZero() + column * m.ColumnStride()
}
