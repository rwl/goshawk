
package tfloat64

import (
	"github.com/rwl/goshawk/common"
	"fmt"
	"math"
)

// Either a view wrapping another matrix or a matrix whose views are wrappers.
type WrapperMatrix struct {
	Matrix
	mat WrapperMat
}

func newWrapperMatrix(content Matrix) *WrapperMatrix {
	wm := &WrapperMat{
		common.NewCoreMat(false, content.Rows(), content.Columns(), content.Columns(), 1, 0, 0),
		content,
	}
	return &WrapperMatrix{
		Matrix{wm},
		wm,
	}
}

func (wm *WrapperMatrix) AssignVector(values []float64) (*WrapperMatrix, error) {
	dm, ok := wm.content.(DiagonalMatrix)
	if ok {
		if len(values) != dm.dlength {
			return wm, fmt.Errorf("Must have same length: length=%d dlength=%d", len(values), dm.dlength)
		}
		for i, v := range values {
			dm.elements[i] = values[i]
		}
	} else {
		_, err := wm.Matrix.AssignVector(values)
		if err != nil {
			return wm, err
		}
	}
	return wm, nil
}

func (wm *WrapperMatrix) AssignMatrixFunc(y Mat, f Float64Float64Func) (*Matrix, error) {
	err := wm.checkShape(y)
	if err != nil {
		return wm, err
	}
	wy, ok := y.(WrapperMatrix)
	if ok {
		var rowList []int
		var columnList []int
		var valueList []float64
		wy.NonZeros(&rowList, &columnList, &valueList)
		wm.Assign(wy, f, rowList, columnList)
	} else {
		_, err := wm.Matrix.Assign(y, f)
		if err != nil {
			return wm, err
		}
	}
	return wm, nil
}

func (wm *WrapperMatrix) Equals(value float64) bool {
	dm, ok := wm.content.(DiagonalMatrix)
	if ok {
		epsilon := prop.Tolerance()
		elements := dm.Elements().([]float64)
		for r := 0; r < len(elements); r++ {
			x := elements[r]
			diff := math.Abs(value - x)
			if (diff != diff) && ((value != value && x != x) || value == x) {
				diff = 0
			}
			if !(diff <= epsilon) {
				return false
			}
		}
		return true
	} else {
		return wm.Matrix.Equals(value)
	}
}

func (wm *WrapperMatrix) EqualsMatrix(other Mat) bool {
	A, ok := wm.content.(DiagonalMatrix)
	B, ok2 := other.(DiagonalMatrix)
	if ok && ok2 {
		epsilon := prop.Tolerance()
		if wm == other {
			return true
		}
		if !(wm != nil && other != nil) {
			return false
		}
		if A.Columns() != B.Columns() || A.Rows() != B.Rows() || A.DiagonalIndex() != B.DiagonalIndex() || A.DiagonalLength() != B.DiagonalLength() {
			return false
		}
		AElements := A.Elements()
		BElements := B.Elements()
		for r := 0; r < len(AElements); r++ {
			x := AElements[r]
			value := BElements[r]
			diff := math.Abs(value - x)
			if (diff != diff) && ((value != value && x != x) || value == x) {
				diff = 0
			}
			if !(diff <= epsilon) {
				return false
			}
		}
		return true
	} else {
		return wm.Matrix.EqualsMatrix(other)
	}
}

func (wm *WrapperMatrix) Vectorize() Vector {
	v := MakeVector(wm.Size())
	idx := 0
	for c := 0; c < wm.Columns(); c++ {
		for r := 0; r < wm.Rows(); r++ {
			v.SetQuick(idx, w.GetQuick(r, c))
			idx++
		}
	}
	return v
}
