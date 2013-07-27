
package tfloat64

import l4g "code.google.com/p/log4go"

var prop = &Property{1e-9}

type Vector struct {
	VectorData
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

// Returns a string representation using default formatting.
func (v *Vector) String() string {
	return ""//NewFormatter().VectorToString(v)
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

// Sets the matrix cell at coordinate index to the specified value.
func (v *Vector) Set(index int, value float64) error {
	if index < 0 || index >= v.Size() {
		return l4g.Error("Attempted to access %s at index=%d",
			v.StringShort(), index)
	}
	v.SetQuick(index, value)
	return nil
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

// Constructs and returns a new view equal to the receiver. The view is a
// shallow clone.
//
// Note that the view is not a deep copy. The returned matrix is
// backed by this matrix, so changes in the returned matrix are reflected in
// this matrix, and vice-versa.
//
// Use Copy() to construct an independent deep copy rather than a
// new view.
func (v *Vector) ViewVector() *Vector {
	return &Vector{v.ViewVectorData()}
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

// Fills the coordinates and values of cells having non-zero values into the
// specified lists. Fills into the lists, starting at index 0. After this
// call returns the specified lists all have a new size, the number of
// non-zero values.
//
// In general, fill order is unspecified. This implementation fills
// like: for (index = 0..size()-1) do... .
// Example:
//
//		0, 0, 8, 0, 7
//		-->
//		indexList  = (2,4)
//		valueList  = (8,7)
//
//		In other words, get(2)==8, get(4)==7.
func (v *Vector) NonZeros(indexList *[]int, valueList *[]float64) {
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
		if value != 0 {
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
		if value != 0 {
			if fillIndexList {
				*indexList = append(*indexList, i)
			}
			if fillValueList {
				*valueList = append(*valueList, value)
			}
		}
		value = v.GetQuick(i + 1)
		if value != 0 {
			if fillIndexList {
				*indexList = append(*indexList, i + 1)
			}
			if fillValueList {
				*valueList = append(*valueList, value)
			}
		}
	}
}

// Fills the coordinates and values of the first <tt>maxCardinality</tt>
// cells having non-zero values into the specified lists. Fills into the
// lists, starting at index 0. After this call returns the specified lists
// all have a new size, the number of non-zero values.
func (v *Vector) NonZerosCardinality(indexList *[]int, valueList *[]float64, maxCardinality int) {
	fillIndexList := indexList != nil
	fillValueList := valueList != nil
	if fillIndexList {
		*indexList = make([]int, 0)
	}
	if fillValueList {
		*valueList = make([]float64, 0)
	}
	currentSize := 0
	for i := 0; i < v.Size(); i++ {
		value := v.GetQuick(i)
		if value != 0 {
			if fillIndexList {
				*indexList = append(*indexList, i)
			}
			if fillValueList {
				*valueList = append(*valueList, value)
			}
			currentSize += 1
		}
		if currentSize >= maxCardinality {
			break
		}
	}
}

// Fills the coordinates and values of cells having positive values into the
// specified lists. Fills into the lists, starting at index 0. After this
// call returns the specified lists all have a new size, the number of
// non-zero values.
func (v *Vector) PositiveValues(indexList *[]int, valueList *[]float64) {
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
	if (rem == 1) {
		value = v.GetQuick(0)
		if value > 0 {
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
		if value > 0 {
			if fillIndexList {
				*indexList = append(*indexList, i)
			}
			if fillValueList {
				*valueList = append(*valueList, value)
			}
		}
		value = v.GetQuick(i + 1)
		if value > 0 {
			if fillIndexList {
				*indexList = append(*indexList, i + 1)
			}
			if fillValueList {
				*valueList = append(*valueList, value)
			}
		}
	}
}

// Normalizes this matrix, i.e. makes the sum of all elements equal to 1.0
// If the matrix contains negative elements then all the values are shifted
// to ensure non-negativity.
func (v *Vector) Normalize() {
	min, _ := v.MinLocation()
	if min < 0 {
		v.AssignFunc(Subtract(min))
	}
	max, _ := v.MaxLocation()
	if max == 0 {
		v.Assign(1.0 / float64(v.Size()))
	} else {
		sumScaleFactor := v.ZSum()
		sumScaleFactor = 1.0 / sumScaleFactor
		v.AssignFunc(Multiply(sumScaleFactor))
	}
}

// Swaps each element v[i] with other[i].
func (v *Vector) Swap(other VectorData) error {
	err := v.checkSize(other)
	if err != nil {
		return err
	}
	for i := 0; i < v.Size(); i++ {
		tmp := v.GetQuick(i)
		v.SetQuick(i, other.GetQuick(i))
		other.SetQuick(i, tmp)
	}
	return nil
}

// Constructs and returns a 1-dimensional array containing the cell values.
// The values are copied. So subsequent changes in <tt>values</tt> are not
// reflected in the matrix, and vice-versa. The returned array
// values has the form:
//     for i:=0; i < Size; i++ {values[i] = Get(i)}
func (v *Vector) ToArray() []float64 {
	values := make([]float64, v.Size())
	v.FillArray(values)
	return values
}

// Fills the cell values into the specified 1-dimensional array.
// The values are copied. So subsequent changes in <tt>values</tt> are not
// reflected in the matrix, and vice-versa. The returned array
// values has the form:
//     for i:=0; i < Size; i++ {values[i] = Get(i)}
func (v *Vector) FillArray(values []float64) error {
	if len(values) < v.Size() {
		return l4g.Error("values too small")
	}
	for i := 0; i < v.Size(); i++ {
		values[i] = v.GetQuick(i)
	}
	return nil
}


// Returns the sum of all cells; Sum( x[i] ).
func (v *Vector) ZSum() float64 {
	if v.Size() == 0 {
		return 0
	}
	return v.Aggregate(Plus, Identity)
}
