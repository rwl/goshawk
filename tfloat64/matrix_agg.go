package tfloat64

import (
	"math"
	"github.com/rwl/goshawk/common"
	"runtime"
)

func (m *Matrix) Aggregate(aggr Float64Float64Func, f Float64Func) float64 {
	if m.Size() == 0 {
		return math.NaN()
	}
	var a float64
	n := runtime.GOMAXPROCS(-1)
	if n > 1 && m.Rows()*m.Columns() > common.MatrixThreshold {
		n = common.Min(n, m.Rows())
		ch := make(chan float64, n)
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
				b := f(m.GetQuick(0, 0))
				d := 1 // First cell already done.
				for r := idx0; r < idx1; r++ {
					for c := d; c < m.Columns(); c++ {
						b = aggr(b, f(m.GetQuick(r, c)))
					}
					d = 0
				}
				ch <- b
			}()
		}
		a = <-ch
		for j := 1; j < n; j++ {
			a = aggr(a, <-ch)
		}
	} else {
		a = f(m.GetQuick(0, 0))
		d := 1 // First cell already done.
		for r := 0; r < m.Rows(); r++ {
			for c := d; c < m.Columns(); c++ {
				a = aggr(a, f(m.GetQuick(r, c)))
			}
			d = 0
		}
	}
	return a
}

func (m *Matrix) AggregateProcedure(aggr Float64Float64Func, f Float64Func, cond Float64Procedure) float64 {
	if m.Size() == 0 {
		return math.NaN()
	}
	var a float64
	n := runtime.GOMAXPROCS(-1)
	if n > 1 && m.Rows()*m.Columns() > common.MatrixThreshold {
		n = common.Min(n, m.Rows())
		ch := make(chan float64, n)
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
				elem := m.GetQuick(idx0, 0)
				b := 0
				if cond(elem) {
					b = aggr(b, f(elem))
				}
				d := 1
				for r := idx0; r < idx1; r++ {
					for c := d; c < m.Columns(); c++ {
						elem = m.GetQuick(r, c)
						if cond(elem) {
							b = aggr(b, f(elem))
						}
					}
					d = 0;
				}
				ch <- b
			}()
		}
		a = <-ch
		for j := 1; j < n; j++ {
			a = aggr(a, <-ch)
		}
	} else {
		a := 0.0
		elem := m.GetQuick(0, 0)
		if cond(elem) {
			a = aggr(a, f(elem))
		}
		d := 1 // First cell already done.
		for r := 0; r < m.Rows(); r++ {
			for c := d; c < m.Columns(); c++ {
				elem = m.GetQuick(r, c)
				if cond(elem) {
					a = aggr(a, f(elem))
				}
			}
			d = 0
		}
	}
	return a
}

func (m *Matrix) AggregateProcedureSelection(aggr Float64Float64Func, f Float64Func, rowList, columnList []int) float64 {
	if m.Size() == 0 {
		return math.NaN()
	}
	size := len(rowList)
	var a float64
	n := runtime.GOMAXPROCS(-1)
	if n > 1 && size > common.MatrixThreshold {
		n = common.Min(n, size)
		ch := make(chan float64, n)
		k := size / n
		var idx0, idx1 int
		for j := 0; j < n; j++ {
			idx0 = j * k
			if j == n - 1 {
				idx1 = size
			} else {
				idx1 = idx0 + k
			}
			go func() {
				b := f(m.GetQuick(rowList[idx0], columnList[idx0]))
				for i := idx0 + 1; i < idx1; i++ {
					elem := m.GetQuick(rowList[i], columnList[i])
					b = aggr(b, f(elem))
				}
				ch <- b
			}()
		}
		a = <-ch
		for j := 1; j < n; j++ {
			a = aggr(a, <-ch)
		}
	} else {
		a = f(m.GetQuick(rowList[0], columnList[0]))
		for i := 1; i < len(rowList); i++ {
			elem := m.GetQuick(rowList[i], columnList[i])
			a = aggr(a, f(elem))
		}
	}
	return a
}

func (m *Matrix) AggregateMatrix(other Mat, aggr Float64Float64Func, f Float64Float64Func) (float64, error) {
	err := m.checkShape(other)
	if err != nil {
		return math.NaN(), err
	}
	if m.Size() == 0 {
		return math.NaN(), nil
	}
	a := 0.0
	n := runtime.GOMAXPROCS(-1)
	if n > 1 && m.Rows()*m.Columns() > common.MatrixThreshold {
		n = common.Min(n, m.Rows())
		ch := make(chan float64, n)
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
				a := f(m.GetQuick(idx0, 0), other.GetQuick(idx0, 0))
				d := 1
				for r := idx0; r < idx1; r++ {
					for c := d; c < m.Columns(); c++ {
						a = aggr(a, f(m.GetQuick(r, c), other.GetQuick(r, c)))
					}
					d = 0
				}
				ch <- d
			}()
		}
		a = <-ch
		for j := 1; j < n; j++ {
			a = aggr(a, <-ch)
		}
	} else {
		a := f(m.GetQuick(0, 0), other.GetQuick(0, 0))
		d := 1 // First cell already done.
		for r := 0; r < m.Rows(); r++ {
			for c := d; c < m.Columns(); c++ {
				a = aggr(a, f(m.GetQuick(r, c), other.GetQuick(r, c)))
			}
			d = 0
		}
	}
	return a, nil
}
