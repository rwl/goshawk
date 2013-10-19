package tfloat64

import (
	"fmt"
	"github.com/rwl/goshawk/common"
	"runtime"
)

func (m *Matrix) AssignFunc(f Float64Func) *Matrix {
	n := runtime.GOMAXPROCS(-1)
	if n > 1 && m.Rows()*m.Columns() > common.MatrixThreshold {
		n = common.Min(n, m.Rows())
		done := make(chan bool, n)
		k := m.Rows() / n
		var idx0, idx1 int
		for j := 0; j < n; j++ {
			idx0 = j * k
			if j == n - 1 {
				idx1 = m.Rows()
			} else {
				idx1 = idx0 + k
			}
			go func() {
				for r := idx0; r < idx1; r++ {
					for c := 0; c < m.Columns(); c++ {
						m.SetQuick(r, c, f(m.GetQuick(r, c)))
					}
				}
				done <- true
			}()
		}
		for j := 0; j < n; j++ {
			<-done
		}
	} else {
		for r := 0; r < m.Rows(); r++ {
			for c := 0; c < m.Columns(); c++ {
				m.SetQuick(r, c, f(m.GetQuick(r, c)))
			}
		}
	}
	return m
}

func (m *Matrix) AssignProcedureFunc(cond Float64Procedure, f Float64Func) *Matrix {
	var elem float64
	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < m.Columns(); c++ {
			elem = m.GetQuick(r, c)
			if cond(elem) {
				m.SetQuick(r, c, f(elem))
			}
		}
	}
	return m
}

func (m *Matrix) AssignProcedure(cond Float64Procedure, value float64) *Matrix {
	var elem float64
	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < m.Columns(); c++ {
			elem = m.GetQuick(r, c)
			if cond(elem) {
				m.SetQuick(r, c, value)
			}
		}
	}
	return m
}

func (m *Matrix) Assign(value float64) *Matrix {
	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < m.Columns(); c++ {
			m.SetQuick(r, c, value)
		}
	}
	return m
}

func (m *Matrix) AssignVector(values []float64) (*Matrix, error) {
	if len(values) != m.Size() {
		return m, fmt.Errorf("Must have same length: length=%d rows()*columns()=%d", len(values), m.Rows()*m.Columns())
	}
	idx := 0
	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < m.Columns(); c++ {
			m.SetQuick(r, c, values[idx])
			idx++
		}
	}
	return m, nil
}

func (m *Matrix) AssignArray(values [][]float64) (*Matrix, error) {
	if len(values) != m.Rows() {
		return m, fmt.Errorf("Must have same number of rows: rows=%d rows()=%d", len(values), m.Rows())
	}
	for r := 0; r < m.Rows(); r++ {
		currentRow := values[r]
		if len(currentRow) != m.Columns() {
			return m, fmt.Errorf("Must have same number of columns in every row: columns=%d columns()=%d", len(currentRow), m.Columns())
		}
		for c := 0; c < m.Columns(); c++ {
			m.SetQuick(r, c, currentRow[c])
		}
	}
	return m, nil
}

func (m *Matrix) AssignMatrix(other Mat) (*Matrix, error) {
	if other == m {
		return m, nil
	}
	err := m.checkShape(other)
	if err != nil {
		return m, err
	}
	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < m.Columns(); c++ {
			m.SetQuick(r, c, other.GetQuick(r, c))
		}
	}
	return m, nil
}

func (m *Matrix) AssignMatrixFunc(y Mat, f Float64Float64Func) (*Matrix, error) {
	err := m.checkShape(y)
	if err != nil {
		return m, err
	}
	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < m.Columns(); c++ {
			m.SetQuick(r, c, f(m.GetQuick(r, c), y.GetQuick(r, c)))
		}
	}
	return m, nil
}

func (m *Matrix) AssignMatrixFuncSelection(y Mat, f Float64Float64Func, rowList, columnList []int) (*Matrix, error) {
	err := m.checkShape(y)
	if err != nil {
		return m, err
	}
	for i := 0; i < m.Size(); i++ {
		m.SetQuick(rowList[i], columnList[i], f(m.GetQuick(rowList[i], columnList[i]), y.GetQuick(rowList[i], columnList[i])))
	}
	return m, nil
}
