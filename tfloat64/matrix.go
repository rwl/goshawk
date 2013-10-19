package tfloat64

import (
	"fmt"
	"math"
	"errors"
)

type Matrix struct {
	Mat
}

// Returns a string representation using default formatting.
func (m *Matrix) String() string {
	return fmtr.MatrixToString(m)
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

func (m *Matrix) Copy() *Matrix {
	copy := &Matrix{m.Like(m.Rows(), m.Columns())}
	copy.AssignMatrix(m)
	return copy
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

func (m *Matrix) EqualsMatrix(other Mat) bool {
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
		m.Assign(1.0/float64(m.Size()))
	} else {
		sumScaleFactor := m.ZSum()
		sumScaleFactor = 1.0/sumScaleFactor
		m.AssignFunc(Multiply(sumScaleFactor))
	}
	return m
}

func (m *Matrix) ViewColumn(column int) (*Matrix, error) {
	err := m.CheckColumn(column)
	if err != nil {
		return nil , err
	}
	viewSize := m.Rows()
	viewZero := m.Index(0, column)
	viewStride := m.RowStride()
	return m.Like1D(viewSize, viewZero, viewStride)
}

func (m *Matrix) ViewColumnFlip() *Matrix {
	v := m.View()
	v.VColumnFlip()
	return v
}

func (m *Matrix) ViewDice() *Matrix {
	v := m.View()
	v.VDice()
	return v
}

func (m *Matrix) ViewPart(row, column, height, width int) (*Matrix, error) {
	v := m.View()
	err := v.VPart(row, column, height, width)
	if err != nil {
		return m, err
	}
	return v, nil
}

func (m *Matrix) ViewRow(row int) (*Vector, error) {
	err := m.CheckRow(row)
	if err != nil {
		return nil, err
	}
	viewSize := m.Columns()
	viewZero := m.Index(row, 0)
	viewStride := m.ColumnStride()
	r := m.Like1D(viewSize, viewZero, viewStride)
	return r, nil
}

func (m *Matrix) ViewRowFlip() (*Matrix) {
	v := m.View()
	v.vRowFlip()
	return v
}

func (m *Matrix) ViewSelectionProcedure(condition VectorProcedure) *Matrix {
	matches := make([]int, 0)
	for i := 0; i < m.Rows(); i++ {
		if condition(m.ViewRow(i)) {
			matches = append(matches, i)
		}
	}
	return m.ViewSelection(matches, nil) // take all columns
}

func (m *Matrix) ViewSelection(rowIndexes, columnIndexes []int) (*Matrix, error) {
	// check for "all"
	if rowIndexes == nil {
		rowIndexes = make([]int, m.Rows())
		for i := 0; i < m.Rows(); i++ {
			rowIndexes[i] = i
		}
	}
	if columnIndexes == nil {
		columnIndexes = make([]int, m.Columns())
		for i := 0; i < m.Columns(); i++ {
			columnIndexes[i] = i
		}
	}

	err := m.CheckRowIndexes(rowIndexes)
	if err != nil {
		return nil, err
	}
	err = m.CheckColumnIndexes(columnIndexes)
	if err != nil {
		return nil, err
	}
	rowOffsets := make([]int, len(rowIndexes))
	columnOffsets := make([]int, len(columnIndexes))
	for i := 0; i < len(rowIndexes); i++ {
		rowOffsets[i] = _rowOffset(_rowRank(rowIndexes[i]))
	}
	for i := 0; i < len(columnIndexes); i++ {
		columnOffsets[i] = _columnOffset(_columnRank(columnIndexes[i]))
	}
	return m.ViewSelectionLike(rowOffsets, columnOffsets)
}

/*func (m *Matrix) ViewSorted(column int) *Matrix {
	return mergeSort.sort(m, column)
}*/

func (m *Matrix) ViewStrides(rowStride, columnStride int) (*Matrix, error) {
	v := m.View()
	err := v.VStrides(rowStride, columnStride)
	if err != nil {
		return m, err
	}
	return v, nil
}

func (m *Matrix) ZMult(y, z *Vector) (*Vector, error) {
	return m.ZMultConst(y, z, 1, 0, false)
}

func (m *Matrix) ZMultConst(y, z *Vector, alpha, beta float64, transposeA bool) (*Vector, error) {
	if transposeA {
		return m.ViewDice().ZMultConst(y, z, alpha, beta, false)
	}
	var zz *Vector
	if z == nil {
		zz = y.LikeVector(m.Rows())
	} else {
		zz = z
	}
	if m.Columns() != y.Size() || m.Rows() > zz.Size() {
		return zz, fmt.Errorf("Incompatible args: %s, %s, %s", m.StringShort(), y.StringShort(), zz.StringShort())
	}

	for r := 0; r < m.Rows(); r++ {
		s := 0.0
		for c := 0; c < m.Columns(); c++ {
			s += m.GetQuick(r, c) * y.GetQuick(c)
		}
		zz.SetQuick(r, alpha * s + beta * zz.getQuick(r))
	}
	return zz, nil
}

func (m *Matrix) ZMultMatrix(B, C *Matrix) (*Matrix, error) {
	return m.ZMultMatrixConst(B, C, 1, 0, false, false)
}

func (m *Matrix) ZMultMatrixConst(B, C *Matrix, alpha, beta float64, transposeA, transposeB bool) (*Matrix, error) {
	if transposeA {
		return m.ViewDice().ZMultmatrixConst(B, C, alpha, beta, false, transposeB)
	}
	if transposeB {
		return m.ZMultMatrixConst(B.ViewDice(), C, alpha, beta, transposeA, false)
	}

	m := m.Rows()
	n := m.Columns()
	p := B.Columns()
	var CC *Matrix
	if C == nil {
		CC = m.Like(m, p)
	} else {
		CC = C
	}
	if B.Rows != n {
		return CC, fmt.Errorf("Matrix2D inner dimensions must agree: %s, %s", m.StringShort(), B.StringShort())
	}
	if CC.Rows() != m || CC.Columns() != p {
		return CC, fmt.Errorf("Incompatibe result matrix: %s, %s, %s", m.StringShort(), B.StringShort(), CC.toStringShort())
	}
	if m == CC || B == CC {
		return CC, errors.New("Matrices must not be identical")
	}

	for a := 0; a < p; a++ {
		for b := 0; b < m; b++ {
			s := 0.0
			for c := 0; c < n; c++ {
				s += m.GetQuick(b, c) * B.GetQuick(c, a)
			}
			CC.SetQuick(b, a, alpha * s + beta * CC.GetQuick(b, a))
		}
	}
	return CC, nil
}

func (m *Matrix) ZSum() float64 {
	if m.Size() == 0 {
		return 0.0
	}
	return m.Aggregate(Plus, Identity)
}
