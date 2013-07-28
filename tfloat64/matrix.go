
package tfloat64

import "fmt"

type Matrix struct {
	MatrixData
}

func (m *Matrix) checkShape(other MatrixData) error {
	if m.Rows() != other.Rows() || m.Columns() != other.Columns() {
		return fmt.Errorf("row sizes do not match: %d!=%d", m.Rows(), other.Rows())
	}
	if m.Columns() != other.Columns() {
		return fmt.Errorf("column sizes do not match: %d!=%d", m.Columns(), other.Columns())
	}
	return nil
}

func (m *Matrix) checkColumn(column int) error {
	if column < 0 || column >= m.Columns() {
		return fmt.Errorf("Attempted to access %s at column=%d", m.StringShort(), column)
	}
	return nil
}

func (m *Matrix) checkBox(row, column, height, width int) error {
	if column < 0 || width < 0 || column + width > m.Columns() || row < 0 || height < 0 || row + height > m.Rows() {
		return fmt.Errorf("%s, column:%d, row:%d, width:%d, height:%d", m.StringShort(), column, row, width, height)
	}
	return nil
}

func (m *Matrix) checkRow(row int) error {
	if row < 0 || row >= m.Rows() {
		return fmt.Errorf("Attempted to access %s at row=%d", m.StringShort(), row)
	}
	return nil
}

// Returns a string representation using default formatting.
func (m *Matrix) String() string {
	return fmtr.MatrixToString(m)
}

// Returns a short string representation of the receiver's shape.
func (m *Matrix) StringShort() string {
	return fmtr.MatrixShape(m)
}

// Returns the number of cells which is Rows()*Columns().
func (m *Matrix) Size() int {
	return m.Rows() * m.Columns()
}
