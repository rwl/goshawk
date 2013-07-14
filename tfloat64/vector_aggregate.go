
package tfloat64

import "math"

// Applies a function to each cell and aggregates the results. Returns a
// value v such that v==a(size()) where
// a(i) == aggr( a(i-1), f(get(i)) ) and terminators are
// a(1) == f(get(0)), a(0)==NaN.
//
// Example:
//
// 	 matrix = 0 1 2 3
//
// 	 // Sum( x[i]*x[i] )
// 	 matrix.Aggregate(Plus, Square)
// 	 --> 14
func (v *Vector) Aggregate(aggr Float64Float64Func, f Float64Func) float64 {
	if v.Size() == 0 {
		return math.NaN()
	}
	a := f(v.GetQuick(0))
	for i := 1; i < v.Size(); i++ {
		a = aggr(a, f(v.GetQuick(i)))
	}
	return a
}
