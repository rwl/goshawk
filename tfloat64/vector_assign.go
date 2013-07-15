
package tfloat64

import (
	l4g "code.google.com/p/log4go"
)

// Assigns the result of a function to each cell; x[i] = f(x[i]).
func (v *Vector) AssignFunc(f Float64Func) *Vector {
	for i := 0; i < v.Size(); i++ {
		v.SetQuick(i, f(v.GetQuick(i)))
	}
	return v
}

// Sets all cells to the state specified by "value".
func (v *Vector) Assign(value float64) *Vector {
	for i := 0; i < v.Size(); i++ {
		v.SetQuick(i, value)
	}
	return v
}

// Sets all cells to the state specified by "values". "values"
// is required to have the same number of cells as the receiver.
//
// The values are copied. So subsequent changes in "values" are not
// reflected in the matrix, and vice-versa.
func (v *Vector) AssignArray(values []float64) (*Vector, error) {
	if len(values) != v.Size() {
		return v, l4g.Error("Must have same number of cells: length=%d size()=%d",
			len(values), v.Size())
	}
	for i, val := range values {
		v.SetQuick(i, val)
	}
	return v, nil
}

// Replaces all cell values of the receiver with the values of another
// matrix. Both matrices must have the same size. If both matrices share
// the same cells (as is the case if they are views derived from the same
// matrix) and intersect in an ambiguous way, then replaces as if
// using an intermediate auxiliary deep copy of "other".
func (v *Vector) AssignVector(other VectorData) (*Vector, error) {
	if v.Size() != other.Size() {
		return v, l4g.Error("Incompatible sizes: %d and %d",
			v.Size(), other.Size())
//		return v, l4g.Error("Incompatible sizes: %s and %s",
//			v.StringShort(), NewFormatter().VectorShape(other))
	}
	for i := 0; i < v.Size(); i++ {
		v.SetQuick(i, other.GetQuick(i))
	}
	return v, nil
}

// Assigns to each cell the result of a function taking as first argument
// the current cell's value of this matrix, and as second argument the
// current cell's value of "y".
func (v *Vector) AssignVectorFunc(y VectorData, f Float64Float64Func) (*Vector, error) {
	if y.Size() != v.Size() {
		return v, l4g.Error("Incompatible sizes: %d and %d",
				y.Size(), v.Size())
//		return v, l4g.Error("Incompatible sizes: %s and %s",
//			v.StringShort(), NewFormatter().VectorShape(y))
	}
	// the general case x[i] = f(x[i],y[i])
	for i := 0; i < v.Size(); i++ {
		v.SetQuick(i, f(v.GetQuick(i), y.GetQuick(i)))
	}
	return v, nil
}

// Assigns the result of a function to all cells that satisfy a condition.
func (v *Vector) AssignProcedureFunc(cond Float64Procedure, f Float64Func) *Vector {
	var elem float64
	for i := 0; i < v.Size(); i++ {
		elem = v.GetQuick(i)
		if cond(elem) {
			v.SetQuick(i, f(elem))
		}
	}
	return v
}

// Assigns a value to all cells that satisfy a condition.
func (v *Vector) AssignProcedure(cond Float64Procedure, value float64) *Vector {
	var elem float64
	for i := 0; i < v.Size(); i++ {
		elem = v.GetQuick(i)
		if cond(elem) {
			v.SetQuick(i, value)
		}
	}
	return v
}
