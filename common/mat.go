package common

import "fmt"

type Mat interface {
	Base

	Rows() int
	Columns() int

	RowStride() int
	ColumnStride() int

	RowZero() int
	ColumnZero() int

	Index(int, int) int

	Size() int // Returns the number of cells.

	VFlip()
	VPart(index, width int) error
	VStrides(stride int) error
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

// Returns a short string representation of the receiver's shape.
func (m *CoreMat) StringShort() string {
	return MatrixShape(m)
}

// Returns the number of cells which is Rows()*Columns().
func (m *CoreMat) Size() int {
	return m.rows*m.columns
}

func (m *CoreMat) CheckShape(other Mat) error {
	if m.rows != other.Rows() || m.columns != other.Columns() {
		return fmt.Errorf("row sizes do not match: %d!=%d", m.rows, other.Rows())
	}
	if m.columns != other.Columns() {
		return fmt.Errorf("column sizes do not match: %d!=%d", m.columns, other.Columns())
	}
	return nil
}

func (m *CoreMat) CheckColumn(column int) error {
	if column < 0 || column >= m.columns {
		return fmt.Errorf("Attempted to access %s at column=%d", m.StringShort(), column)
	}
	return nil
}

func (m *CoreMat) CheckBox(row, column, height, width int) error {
	if column < 0 || width < 0 || column + width > m.columns || row < 0 || height < 0 || row + height > m.rows {
		return fmt.Errorf("%s, column:%d, row:%d, width:%d, height:%d", m.StringShort(), column, row, width, height)
	}
	return nil
}

func (m *CoreMat) CheckRow(row int) error {
	if row < 0 || row >= m.rows {
		return fmt.Errorf("Attempted to access %s at row=%d", m.StringShort(), row)
	}
	return nil
}

func (m *CoreMat) VColumnFlip() {
	if m.columns > 0 {
		m.columnZero += (m.columns - 1) * m.columnStride
		m.columnStride = -columnStride
		m.isView = true
	}
}

func (m *CoreMat) VDice() {
	var tmp int
	// swap
	tmp = m.rows
	m.rows = m.columns
	m.columns = tmp
	tmp = m.rowStride
	m.rowStride = m.columnStride
	m.columnStride = tmp
	tmp = m.rowZero
	m.rowZero = m.columnZero
	m.columnZero = tmp

	// flips stay unaffected

	m.isView = true
}

func (m *CoreMat) VPart(row, column, height, width int) error {
	err := m.CheckBox(row, column, height, width)
	if err != nil {
		return err
	}
	m.rowZero += m.rowStride * row
	m.columnZero += m.columnStride * column
	m.rows = height
	m.columns = width
	m.isView = true
	return nil
}

func (m *CoreMat) VRowFlip() {
	if m.rows > 0 {
		m.rowZero += (m.rows - 1) * m.rowStride
		m.rowStride = -m.rowStride
		m.isView = true
	}
}

func (m *CoreMat) VStrides(rowStride, columnStride int) error {
	if m.rowStride <= 0 || m.columnStride <= 0 {
		return fmt.Errorf("illegal strides: %d, %d", rowStride, columnStride)
	}
	m.rowStride *= rowStride
	m.columnStride *= columnStride
	if m.rows != 0 {
		m.rows = (m.rows - 1) / rowStride + 1
	}
	if m.columns != 0 {
		m.columns = (m.columns - 1) / columnStride + 1
	}
	m.isView = true
	return nil
}
