package tfloat64

import (
	"fmt"
	"sort"
	"github.com/rwl/goshawk/common"
	"runtime"
)

var (
	prop = &Property{1e-9}
	fmtr = NewFormatter()
)

type Vector struct {
	Vec
}

// Returns a string representation using default formatting.
func (v *Vector) String() string {
	return fmtr.VectorToString(v)
}

// Returns the matrix cell value at coordinate "index".
func (v *Vector) Get(index int) (float64, error) {
	if index < 0 || index >= v.Size() {
		return 0.0, fmt.Errorf("Attempted to access %s at index=%d",
			v.StringShort(), index)
	}
	return v.GetQuick(index), nil
}

// Sets the matrix cell at coordinate index to the specified value.
func (v *Vector) Set(index int, value float64) error {
	if index < 0 || index >= v.Size() {
		return fmt.Errorf("Attempted to access %s at index=%d",
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
	return &Vector{v.ViewVec()}
}

// Returns the number of cells having non-zero values; ignores tolerance.
func (v *Vector) Cardinality() int {
	var cardinality int
	n := runtime.GOMAXPROCS(-1)
	if n > 1 && v.Size() > common.VectorThreshold {
		n = common.Min(n, v.Size())
		c := make(chan int, n)
		k := v.Size() / n
		var idx0, idx1 int
		for j := 0; j < n; j++ {
			idx0 = j * k
			if j == n - 1 {
				idx1 = v.Size()
			} else {
				idx1 = idx0 + k
			}
			go func() {
				card := 0
				for i := idx0; i < idx1; i++ {
					if v.GetQuick(i) != 0.0 {
						card += 1
					}
				}
				c <- card
			}()
		}
		cardinality = <-c
		for j := 1; j < n; j++ {
			cardinality += <-c
		}
	} else {
		cardinality = 0
		for i := 0; i < v.Size(); i++ {
			if v.GetQuick(i) != 0.0 {
				cardinality += 1
			}
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
func (v *Vector) EqualsVector(other Vec) bool {
	return prop.VectorEqualsVector(v, other)
}

// Return the maximum value of this matrix together with its location.
func (v *Vector) MaxLocation() (float64, int) {
	var location int
	var maxValue float64
	n := runtime.GOMAXPROCS(-1)
	if n > 1 && v.Size() > common.VectorThreshold {
		n = common.Min(n, v.Size())
		c := make(chan vectorElement, n)
		k := v.Size() / n
		var idx0, idx1 int
		for j := 0; j < n; j++ {
			idx0 = j * k
			if j == n - 1 {
				idx1 = v.Size()
			} else {
				idx1 = idx0 + k
			}
			go func() {
				loc := idx0
				max := v.GetQuick(loc)
				for i := idx0 + 1; i < idx1; i++ {
					elem := v.GetQuick(i)
					if max < elem {
						max = elem
						loc = i
					}
				}
				c <- vectorElement{max, loc}
			}()
		}
		vl0 := <-c
		maxValue = vl0.value
		location = vl0.location
		for j := 1; j < n; j++ {
			vl := <-c
			if maxValue < vl.value {
				maxValue = vl.value
				location = vl.location
			}
		}
	} else {
		location = 0
		maxValue = v.GetQuick(location)
		for i := 1; i < v.Size(); i++ {
			elem := v.GetQuick(i)
			if maxValue < elem {
				maxValue = elem
				location = i
			}
		}
	}
	return maxValue, location
}

// Return the minimum value of this matrix together with its location.
func (v *Vector) MinLocation() (float64, int) {
	var location int
	var minValue float64
	n := runtime.GOMAXPROCS(-1)
	if n > 1 && v.Size() > common.VectorThreshold {
		n = common.Min(n, v.Size())
		c := make(chan vectorElement, n)
		k := v.Size() / n
		var idx0, idx1 int
		for j := 0; j < n; j++ {
			idx0 = j * k
			if j == n - 1 {
				idx1 = v.Size()
			} else {
				idx1 = idx0 + k
			}
			go func() {
				loc := idx0
				min := v.GetQuick(loc)
				for i := idx0 + 1; i < idx1; i++ {
					elem := v.GetQuick(i)
					if min > elem {
						min = elem
						loc = i
					}
				}
				c <- vectorElement{min, loc}
			}()
		}
		vl0 := <-c
		minValue = vl0.value
		location = vl0.location
		for j := 1; j < n; j++ {
			vl := <-c
			if minValue > vl.value {
				minValue = vl.value
				location = vl.location
			}
		}
	} else {
		location = 0
		minValue = v.GetQuick(location)
		for i := 1; i < v.Size(); i++ {
			elem := v.GetQuick(i)
			if minValue > elem {
				minValue = elem
				location = i
			}
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
	rem := v.Size()%2
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
	rem := v.Size()%2
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
func (v *Vector) NonZerosCardinality(indexList *[]int, valueList *[]float64,
maxCardinality int) {
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
	rem := v.Size()%2
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
		v.Assign(1.0/float64(v.Size()))
	} else {
		sumScaleFactor := v.ZSum()
		sumScaleFactor = 1.0/sumScaleFactor
		v.AssignFunc(Multiply(sumScaleFactor))
	}
}

// Swaps each element v[i] with other[i].
func (v *Vector) Swap(other Vec) error {
	err := v.checkSize(other)
	if err != nil {
		return err
	}
	n := runtime.GOMAXPROCS(-1)
	if n > 1 && v.Size() > common.VectorThreshold {
		n = common.Min(n, v.Size())
		done := make(chan bool, n)
		k := v.Size() / n
		var idx0, idx1 int
		for j := 0; j < n; j++ {
			idx0 = j * k
			if j == n - 1 {
				idx1 = v.Size()
			} else {
				idx1 = idx0 + k
			}
			go func() {
				for i := idx0; i < idx1; i++ {
					tmp := v.GetQuick(i)
					v.SetQuick(i, other.GetQuick(i))
					other.SetQuick(i, tmp)
				}
				done <- true
			}()
		}
		for j := 0; j < n; j++ {
			<-done
		}
	} else {
		for i := 0; i < v.Size(); i++ {
			tmp := v.GetQuick(i)
			v.SetQuick(i, other.GetQuick(i))
			other.SetQuick(i, tmp)
		}
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
		return fmt.Errorf("values too small")
	}
	n := runtime.GOMAXPROCS(-1)
	if n > 1 && v.Size() > common.VectorThreshold {
		n = common.Min(n, v.Size())
		done := make(chan bool, n)
		k := v.Size() / n
		var idx0, idx1 int
		for j := 0; j < n; j++ {
			idx0 = j * k
			if j == n - 1 {
				idx1 = v.Size()
			} else {
				idx1 = idx0 + k
			}
			go func() {
				for i := idx0; i < idx1; i++ {
					values[i] = v.GetQuick(i)
				}
				done <- true
			}()
		}
		for j := 0; j < n; j++ {
			<-done
		}
	} else {
		for i := 0; i < v.Size(); i++ {
			values[i] = v.GetQuick(i)
		}
	}
	return nil
}

// Returns the dot product of two vectors x and y, which is
// Sum(x[i]*y[i]). Where x == this. Operates on cells at
// indexes 0 .. math.Min(size(), y.size()).
func (v *Vector) ZDotProduct(y Vec) float64 {
	return v.ZDotProductRange(y, 0, v.Size())
}

// Returns the dot product of two vectors x and y, which is
// Sum(x[i]*y[i]). Where x == this. Operates on cells at
// indexes from .. Min(Size(), y.Size(),from+length)-1.
func (v *Vector) ZDotProductRange(y Vec, from, length int) float64 {
	if from < 0 || length <= 0 {
		return 0
	}

	tail := from + length
	if v.Size() < tail {
		tail = v.Size()
	}
	if y.Size() < tail {
		tail = y.Size()
	}
	length = tail - from

	var sum float64
	n := runtime.GOMAXPROCS(-1)
	if n > 1 && v.Size() > common.VectorThreshold {
		n = common.Min(n, length)
		c := make(chan float64, n)
		k := length / n
		var idx0, idx1 int
		for j := 0; j < n; j++ {
			idx0 = j * k
			if j == n - 1 {
				idx1 = length
			} else {
				idx1 = idx0 + k
			}
			go func() {
				s := 0.0
				for i := idx0; i < idx1; i++ {
					idx := k + from
					s += v.GetQuick(idx)*y.GetQuick(idx)
				}
				c <- s
			}()
		}
		sum = <-c
		for j := 1; j < n; j++ {
			sum += <-c
		}
	} else {
		sum = 0.0
		i := tail - 1
		for k := 0; k < length; i-- {
			sum += v.GetQuick(i)*y.GetQuick(i)
			k++
		}
	}
	return sum
}

// Returns the dot product of two vectors x and y, which is
// Sum(x[i]*y[i]). Where x == this.
func (v *Vector) ZDotProductSelection(y Vec, from, length int,
nonZeroIndexes []int) float64 {
	// determine minimum length
	if from < 0 || length <= 0 {
		return 0
	}

	tail := from + length
	if v.Size() < tail {
		tail = v.Size()
	}
	if y.Size() < tail {
		tail = y.Size()
	}
	length = tail - from
	if length <= 0 {
		return 0
	}
	indexesCopy := make([]int, len(nonZeroIndexes))
	copy(indexesCopy, nonZeroIndexes)
	sort.Ints(indexesCopy)
	index := 0
	s := len(indexesCopy)
	// skip to start
	for (index < s) && nonZeroIndexes[index] < from {
		index++
	}
	// now the sparse dot product
	i := nonZeroIndexes[index]
	sum := 0.0
	length--
	for length >= 0 && index < s && i < tail {
		sum += v.GetQuick(i)*y.GetQuick(i)
		index++
		i = nonZeroIndexes[index]
		length--
	}
	return sum
}

// Returns the sum of all cells; Sum( x[i] ).
func (v *Vector) ZSum() float64 {
	if v.Size() == 0 {
		return 0
	}
	return v.Aggregate(Plus, Identity)
}

type vectorElement struct {
	value float64
	location int
}
