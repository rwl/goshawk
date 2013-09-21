package tfloat64

import "math"

func (m *Matrix) Aggregate(aggr Float64Float64Func, f Float64Func) float64 {
	if m.Size() == 0 {
		return math.NaN()
	}
	a := f(m.GetQuick(0, 0))
	d := 1 // First cell already done.
	for r := 0; r < m.Rows(); r++ {
		for c := d; c < m.Columns(); c++ {
			a = aggr(a, f(m.GetQuick(r, c)))
		}
		d = 0
	}
	return a
}

func (m *Matrix) AggregateProcedure(aggr Float64Float64Func, f Float64Func, cond Float64Procedure) float64 {
	if m.Size() == 0 {
		return math.NaN()
	}
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
	return a
}

func (m *Matrix) AggregateProcedureSelection(aggr Float64Float64Func, f Float64Func, rowList, columnList []int) float64 {
	if m.Size() == 0 {
		return math.NaN()
	}
	var elem float64
	a := f(m.GetQuick(rowList[0], columnList[0]))
	for i := 1; i < len(rowList); i++ {
		elem = m.GetQuick(rowList[i], columnList[i])
		a = aggr(a, f(elem))
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
	a := f(m.GetQuick(0, 0), other.GetQuick(0, 0))
	d := 1 // First cell already done.
	for r := 0; r < m.Rows(); r++ {
		for c := d; c < m.Columns(); c++ {
			a = aggr(a, f(m.GetQuick(r, c), other.GetQuick(r, c)))
		}
		d = 0
	}
	return a, nil
}
