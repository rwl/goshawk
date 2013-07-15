
package tfloat64

import l4g "code.google.com/p/log4go"

var prop = &Property{1e-9}

type Vector struct {
	VectorData
}

func newVector(impl VectorData) Vector {
	return Vector{impl}
}

func (v *Vector) checkSize(other VectorData) error {
	if v.Size() != other.Size() {
//		formatter := NewFormatter()
//		return l4g.Error("Incompatible sizes: %s and %s",
//			formatter.VectorShape(v), formatter.VectorShape(other))
		return l4g.Error("Incompatible sizes: %d and %d",
			v.Size(), other.Size())
	}
	return nil
}

// Returns a short string representation of the receiver's shape.
func (v *Vector) StringShort() string {
	return ""//NewFormatter().VectorShape(v)
}

// Returns the matrix cell value at coordinate "index".
func (v *Vector) Get(index int) (float64, error) {
	if index < 0 || index >= v.Size() {
		return 0.0, l4g.Error("Attempted to access %s at index=%d",
			v.StringShort(), index)
	}
	return v.GetQuick(index), nil
}

// Constructs and returns a deep copy of the receiver.
//
// Note that the returned matrix is an independent deep copy. The
// returned matrix is not backed by this matrix, so changes in the returned
// matrix are not reflected in this matrix, and vice-versa.
func (v *Vector) Copy() *Vector {
	copy := &Vector{v.Like(v.Size())}
	copy.AssignVector(v)
	return copy
}

// Returns the number of cells having non-zero values; ignores tolerance.
func (v *Vector) Cardinality() int {
	cardinality := 0
	for i := 0; i < v.Size(); i++ {
		if v.GetQuick(i) != 0.0 {
			cardinality += 1
		}
	}
	return cardinality
}

// Returns whether all cells are equal to the given value.
func (v *Vector) Equals(value float64) bool {
	return prop.VectorEqualsValue(v, value)
}

// Compares this vector against the specified vector. The result is
// true if and only if the argument is not nil
// and is at least a vector that has the same
// size as the receiver and has exactly the same values at the same
// indexes.
func (v *Vector) EqualsVector(other VectorData) bool {
	return prop.VectorEqualsVector(v, other)
}

// Return the maximum value of this matrix together with its location.
func (v *Vector) MaxLocation() (float64, int) {
	location := 0
	maxValue := v.GetQuick(location)
	var elem float64
	for i := 1; i < v.Size(); i++ {
		elem = v.GetQuick(i)
		if maxValue < elem {
			maxValue = elem
			location = i
		}
	}
	return maxValue, location
}

// Return the minimum value of this matrix together with its location.
func (v *Vector) MinLocation() (float64, int) {
	location := 0
	minValue := v.GetQuick(location)
	var elem float64
	for i := 1; i < v.Size(); i++ {
		elem = v.GetQuick(i)
		if minValue > elem {
			minValue = elem
			location = i
		}
	}
	return minValue, location
}


// Fills the coordinates and values of cells having negative values into the
// specified lists. Fills into the lists, starting at index 0. After this
// call returns the specified lists all have a new size, the number of
// non-zero values.
func (v *Vector) NegativeValues(indexList *[]int, valueList *[]float64) {
	fillIndexList := indexList != nil
	fillValueList := valueList != nil
	if fillIndexList {
		*indexList = make([]int, 0)
	}
	if fillValueList {
		*valueList = make([]float64, 0)
	}
	rem := v.Size() % 2
	var value float64
	if rem == 1 {
		value = v.GetQuick(0)
		if value < 0 {
			if fillIndexList {
				*indexList = append(*indexList, 0)
			}
			if fillValueList {
				*valueList = append(*valueList, value)
			}
		}
	}

	for i := rem; i < v.Size(); i += 2 {
		value = v.GetQuick(i)
		if value < 0 {
			if fillIndexList {
				*indexList = append(*indexList, i)
			}
			if fillValueList {
				*valueList = append(*valueList, value)
			}
		}
		value = v.GetQuick(i + 1)
		if value < 0 {
			if fillIndexList {
				*indexList = append(*indexList, i + 1)
			}
			if fillValueList {
				*valueList = append(*valueList, value)
			}
		}
	}
}
