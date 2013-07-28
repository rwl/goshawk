
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
