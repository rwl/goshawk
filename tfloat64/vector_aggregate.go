
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

// Applies a function to all cells with a given indexes and aggregates the
// results.
func (v *Vector) AggregateIndexed(aggr Float64Float64Func, f Float64Func, indexList []int) float64 {
	if v.Size() == 0 {
		return math.NaN()
	}
	size := len(indexList)
	var elem float64
	a := f(v.GetQuick(indexList[0]))
	for i := 1; i < size; i++ {
		elem = v.GetQuick(indexList[i])
		a = aggr(a, f(elem))
	}
	return a
}

// Applies a function to each corresponding cell of two matrices and
// aggregates the results. Returns a value v such that
// v==a(size()) where
// a(i) == aggr( a(i-1), f(get(i), other.get(i)) ) and terminators
// are a(1) == f(get(0), other.get(0)), a(0)==NaN.
//
// Example:
//
// 	 x = 0 1 2 3
// 	 y = 0 1 2 3
//
// 	 // Sum( x[i]*y[i] )
// 	 x.aggregate(y, Plus, Mult)
// 	 --> 14
//
// 	 // Sum( (x[i]+y[i])^2 )
// 	 x.aggregate(y, Plus, Chain(Square, Plus))
// 	 --> 56
func (v *Vector) AggregateVector(other VectorData, aggr, f Float64Float64Func) (float64, error) {
	err := v.checkSize(other)
	if err != nil {
		return math.NaN(), err
	}
	if v.Size() == 0 {
		return math.NaN(), nil
	}
	a := f(v.GetQuick(0), other.GetQuick(0))
	for i := 1; i < v.Size(); i++ {
		a = aggr(a, f(v.GetQuick(i), other.GetQuick(i)))
	}
	return a, nil
}
