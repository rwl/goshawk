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

func TestDenseAssign(t *testing.T) {
	A := makeDenseVector()
	testAssign(t, A)
}

func TestSparseAssign(t *testing.T) {
	A := makeSparseVector()
	testAssign(t, A)
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

func TestDenseAssignArray(t *testing.T) {
	A := makeDenseVector()
	testAssignArray(t, A)
}

func TestSparseAssignArray(t *testing.T) {
	A := makeSparseVector()
	testAssignArray(t, A)
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

func TestDenseAssignFunc(t *testing.T) {
	A := makeDenseVector()
	testAssignFunc(t, A)
}

func TestSparseAssignFunc(t *testing.T) {
	A := makeSparseVector()
	testAssignFunc(t, A)
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

func TestDenseAssignVector(t *testing.T) {
	A := makeDenseVector()
	B := makeDenseVector()
	testAssignVector(t, A, B)
}

func TestSparseAssignVector(t *testing.T) {
	A := makeSparseVector()
	B := makeSparseVector()
	testAssignVector(t, A, B)
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

func TestDenseAssignVectorFunc(t *testing.T) {
	A := makeDenseVector()
	B := makeDenseVector()
	testAssignVectorFunc(t, A, B)
}

func TestSparseAssignVectorFunc(t *testing.T) {
	A := makeSparseVector()
	B := makeSparseVector()
	testAssignVectorFunc(t, A, B)
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

func TestDenseAssignProcedure(t *testing.T) {
	A := makeDenseVector()
	testAssignProcedure(t, A)
}

func TestSparseAssignProcedure(t *testing.T) {
	A := makeSparseVector()
	testAssignProcedure(t, A)
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

func TestDenseAssignProcedureFunc(t *testing.T) {
	A := makeDenseVector()
	testAssignProcedureFunc(t, A)
}

func TestSparseAssignProcedureFunc(t *testing.T) {
	A := makeSparseVector()
	testAssignProcedureFunc(t, A)
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
