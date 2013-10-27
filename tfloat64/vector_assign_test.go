package tfloat64

import (
	"testing"
	"math/rand"
	"math"
)

type assignVector interface {
	Vec
	Assign(float64) *Vector
}

func testAssign(t *testing.T, A assignVector) {
	value := rand.Float64()
	A.Assign(value)
	for i := 0; i < A.Size(); i++ {
		result := A.GetQuick(i)
		if math.Abs(value - result) > tol {
			t.Errorf("expected:%g actual:%g", value, result)
		}
	}
}

type assignArrayVector interface {
	Vec
	AssignArray([]float64) (*Vector, error)
}

func testAssignArray(t *testing.T, A assignArrayVector) {
	expected := make([]float64, A.Size())
	for i := 0; i < A.Size(); i++ {
		expected[i] = rand.Float64()
	}
	A.AssignArray(expected)
	for i := 0; i < A.Size(); i++ {
		result := A.GetQuick(i)
		if math.Abs(expected[i] - result) > tol {
			t.Errorf("expected:%g actual:%g", expected[i], result)
		}
	}
}

type assignFuncVector interface {
	Vec
	AssignFunc(Float64Func) *Vector
	Copy() *Vector
}

func testAssignFunc(t *testing.T, A assignFuncVector) {
	Acopy := A.Copy()
	A.AssignFunc(math.Acos)
	for i := 0; i < A.Size(); i++ {
		expected := math.Acos(Acopy.GetQuick(i))
		result := A.GetQuick(i)
		if math.Abs(expected - result) > tol {
			t.Errorf("expected:%g actual:%g", expected, result)
		}
	}
}

type assignVectorVector interface {
	Vec
	AssignVector(Vec) (*Vector, error)
}

func testAssignVector(t *testing.T, A assignVectorVector, B Vec) {
	A.AssignVector(B)
	if A.Size() != B.Size() {
		t.Errorf("sizes must be equal: %d!=%d", A.Size(), B.Size())
	}
	for i := 0; i < A.Size(); i++ {
		expected := B.GetQuick(i)
		result := A.GetQuick(i)
		if math.Abs(expected - result) > tol {
			t.Errorf("expected:%g actual:%g", expected, result)
		}
	}
}

type assignVectorFuncVector interface {
	Vec
	AssignVectorFunc(Vec, Float64Float64Func) (*Vector, error)
	Copy() *Vector
}

func testAssignVectorFunc(t *testing.T, A assignVectorFuncVector, B *Vector) {
	Acopy := A.Copy()
	A.AssignVectorFunc(B, Div)
	for i := 0; i < A.Size(); i++ {
		expected := Acopy.GetQuick(i)/B.GetQuick(i)
		result := A.GetQuick(i)
		if math.Abs(expected - result) > tol {
			t.Errorf("expected:%g actual:%g", expected, result)
		}
	}
}

type assignProcedureVector interface {
Vec
	AssignProcedure(Float64Procedure, float64) *Vector
	Copy() *Vector
}

func testAssignProcedure(t *testing.T, A assignProcedureVector) {
	procedure := func(element float64) bool {
		if math.Abs(element) > 0.1 {
			return true
		} else {
			return false
		}
	}
	Acopy := A.Copy()
	A.AssignProcedure(procedure, -1.0)
	for i := 0; i < A.Size(); i++ {
		var expected, result float64
		if math.Abs(Acopy.GetQuick(i)) > 0.1 {
			expected = -1.0
			result = A.GetQuick(i)
		} else {
			expected = Acopy.GetQuick(i)
			result = A.GetQuick(i)
		}
		if math.Abs(expected - result) > tol {
			t.Errorf("expected:%g actual:%g", expected, result)
		}
	}
}

type assignProcedureFuncVector interface {
	Vec
	AssignProcedureFunc(Float64Procedure, Float64Func) *Vector
	Copy() *Vector
}

func testAssignProcedureFunc(t *testing.T, A assignProcedureFuncVector) {
	procedure := func(element float64) bool {
		if math.Abs(element) > 0.1 {
			return true
		} else {
			return false
		}
	}
	Acopy := A.Copy()
	A.AssignProcedureFunc(procedure, math.Tan)
	for i := 0; i < A.Size(); i++ {
		var expected, result float64
		if math.Abs(Acopy.GetQuick(i)) > 0.1 {
			expected = math.Tan(Acopy.GetQuick(i))
			result = A.GetQuick(i)
		} else {
			expected = Acopy.GetQuick(i)
			result = A.GetQuick(i)
		}
		if math.Abs(expected - result) > tol {
			t.Errorf("expected:%g actual:%g", expected, result)
		}
	}
}
