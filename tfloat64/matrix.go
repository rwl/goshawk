
package tfloat64

import (
	"fmt"
	"math"
)

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

func (m *Matrix) Get(row int, column int) (float64, error) {
	if column < 0 || column >= m.Columns() || row < 0 || row >= m.Rows() {
		return math.NaN(), fmt.Errorf("row:%d, column:%d", row, column)
	}
	return m.GetQuick(row, column), nil
}

func (m *Matrix) Set(row, column int, value float64) error {
	if column < 0 || column >= m.Columns() || row < 0 || row >= m.Rows() {
		return fmt.Errorf("row:%d, column:%d", row, column)
	}
	m.SetQuick(row, column, value)
	return nil
}

func (m *Matrix) Cardinality() int {
	cardinality := 0
	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < m.Columns(); c++ {
			if m.GetQuick(r, c) != 0 {
				cardinality += 1
			}
		}
	}
	return cardinality
}

func (m *Matrix) Equals(value float64) bool {
	return prop.MatrixEqualsValue(m, value)
}

func (m *Matrix) EqualsMatrix(other MatrixData) bool {
	return prop.MatrixEqualsMatrix(m, other)
}

func (m *Matrix) ToArray() [][]float64 {
	values := make([][]float64, m.Rows())
	for r := 0; r < m.Rows(); r++ {
		values[r] = make([]float64, m.Columns())
		currentRow := values[r]
		for c := 0; c < m.Columns(); c++ {
			currentRow[c] = m.GetQuick(r, c)
		}
	}
	return values
}

func (m *Matrix) ForEachNonZero(function IntIntFloat64Func) *Matrix {
	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < m.Columns(); c++ {
			value := m.GetQuick(r, c)
			if value != 0 {
				a := function(r, c, value)
				if a != value {
					m.SetQuick(r, c, a)
				}
			}
		}
	}
	return m
}

func (m *Matrix) MaxLocation() (float64, int, int) {
	rowLocation := 0
	columnLocation := 0
	maxValue := m.GetQuick(0, 0)
	var elem float64
	d := 1
	for r := 0; r < m.Rows(); r++ {
		for c := d; c < m.Columns(); c++ {
			elem = m.GetQuick(r, c)
			if maxValue < elem {
				maxValue = elem
				rowLocation = r
				columnLocation = c
			}
		}
		d = 0
	}
	return maxValue, rowLocation, columnLocation
}

func (m *Matrix) MinLocation() (float64, int, int) {
	rowLocation := 0
	columnLocation := 0
	minValue := m.GetQuick(0, 0)
	var elem float64
	d := 1
	for r := 0; r < m.Rows(); r++ {
		for c := d; c < m.Columns(); c++ {
			elem = m.GetQuick(r, c)
			if minValue > elem {
				minValue = elem
				rowLocation = r
				columnLocation = c
			}
		}
		d = 0
	}
	return minValue, rowLocation, columnLocation
}

func (m *Matrix) NegativeValues(rowList, columnList *[]int, valueList *[]float64) {
	*rowList = make([]int, 0)
	*columnList = make([]int, 0)
	*valueList = make([]float64, 0)
	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < m.Columns(); c++ {
			value := m.GetQuick(r, c)
			if value < 0 {
				*rowList = append(*rowList, r)
				*columnList = append(*columnList, c)
				*valueList = append(*valueList, value)
			}
		}
	}
}

func (m *Matrix) NonZeros(rowList, columnList *[]int, valueList *[]float64) {
	*rowList = make([]int, 0)
	*columnList = make([]int, 0)
	*valueList = make([]float64, 0)
	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < m.Columns(); c++ {
			value := m.GetQuick(r, c)
			if value != 0 {
				*rowList = append(*rowList, r)
				*columnList = append(*columnList, c)
				*valueList = append(*valueList, value)
			}
		}
	}
}

func (m *Matrix) PositiveValues(rowList, columnList *[]int, valueList *[]float64) {
	*rowList = make([]int, 0)
	*columnList = make([]int, 0)
	*valueList = make([]float64, 0)
	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < m.Columns(); c++ {
			value := m.GetQuick(r, c)
			if value > 0 {
				*rowList = append(*rowList, r)
				*columnList = append(*columnList, c)
				*valueList = append(*valueList, value)
			}
		}
	}
}

func (m *Matrix) Normalize() *Matrix {
	min, _, _ := m.MinLocation()
	if min < 0 {
		m.AssignFunc(Subtract(min))
	}
	max, _, _ := m.MaxLocation()
	if max == 0 {
		m.Assign(1.0 / float64(m.Size()))
	} else {
		sumScaleFactor := m.ZSum()
		sumScaleFactor = 1.0 / sumScaleFactor
		m.AssignFunc(Multiply(sumScaleFactor))
	}
	return m
}

func (m *Matrix) ZSum() float64 {
	if m.Size() == 0 {
		return 0.0
	}
	return m.Aggregate(Plus, Identity)
}
