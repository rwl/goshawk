
package tfloat64

import (
	"math/rand"
	"testing"
)

const test_size = 2*17*5

func makeDenseVector() (*Vector) {
	A := NewVector(test_size)
	for i := 0; i < test_size; i++ {
		A.SetQuick(i, rand.Float64())
	}
	return A
}

func TestDenseCardinality(t *testing.T) {
	A := makeDenseVector()
	testCardinality(t, A)
}

func TestDenseEquals(t *testing.T) {
	A := makeDenseVector()
	testEquals(t, A)
}

func TestDenseEqualsVector(t *testing.T) {
	A := makeDenseVector()
	B := makeDenseVector()
	testEqualsVector(t, A, B)
}

func TestDenseMaxLocation(t *testing.T) {
	A := makeDenseVector()
	testMaxLocation(t, A)
}

func TestDenseMinLocation(t *testing.T) {
	A := makeDenseVector()
	testMinLocation(t, A)
}

func TestDenseNegativeValues(t *testing.T) {
	A := makeDenseVector()
	testGetNegativeValues(t, A)
}

func TestDenseNonZeros(t *testing.T) {
	A := makeDenseVector()
	testNonZeros(t, A)
}

func TestDensePositiveValues(t *testing.T) {
	A := makeDenseVector()
	testPositiveValues(t, A)
}

func TestDenseToArray(t *testing.T) {
	A := makeDenseVector()
	testToArray(t, A)
}

func TestDenseFillArray(t *testing.T) {
	A := makeDenseVector()
	testFillArray(t, A)
}

func TestDenseReshapeMatrix(t *testing.T) {
	A := makeDenseVector()
	testReshapeMatrix(t, A)
}

func TestDenseReshapeCube(t *testing.T) {
	A := makeDenseVector()
	testReshapeCube(t, A)
}

func TestDenseSwap(t *testing.T) {
	A := makeDenseVector()
	B := makeDenseVector()
	testSwap(t, A, B)
}

func TestDenseZDotProduct(t *testing.T) {
	A := makeDenseVector()
	B := makeDenseVector()
	testZDotProduct(t, A, B)
}

func TestDenseZDotProductRange(t *testing.T) {
	A := makeDenseVector()
	B := makeDenseVector()
	testZDotProductRange(t, A, B)
}

func TestDenseZDotProductSelection(t *testing.T) {
	A := makeDenseVector()
	B := makeDenseVector()
	testZDotProductSelection(t, A, B)
}

func TestDenseZSum(t *testing.T) {
	A := makeDenseVector()
	testZSum(t, A)
}

// View tests.

func TestDenseViewFlip(t *testing.T) {
	A := makeDenseVector()
	testViewFlip(t, A)
}

func TestDenseViewPart(t *testing.T) {
	A := makeDenseVector()
	testViewPart(t, A)
}

func TestDenseViewProcedure(t *testing.T) {
	A := makeDenseVector()
	testViewProcedure(t, A)
}

func TestDenseView(t *testing.T) {
	A := makeDenseVector()
	testView(t, A)
}
/*
func TestDenseViewSorted(t *testing.T) {
	A := makeDenseVector()
	testViewSorted(t, A)
}
*/

func TestDenseViewStrides(t *testing.T) {
	A := makeDenseVector()
	testViewStrides(t, A)
}

// Aggregate tests.

func TestDenseAggregate(t *testing.T) {
	A := makeDenseVector()
	testAggregate(t, A)
}

func TestDenseAggregateIndexed(t *testing.T) {
	A := makeDenseVector()
	testAggregateIndexed(t, A)
}

func TestDenseAggregateVector(t *testing.T) {
	A := makeDenseVector()
	B := makeDenseVector()
	testAggregateVector(t, A, B)
}

// Assign tests.

func TestDenseAssign(t *testing.T) {
	A := makeDenseVector()
	testAssign(t, A)
}

func TestDenseAssignArray(t *testing.T) {
	A := makeDenseVector()
	testAssignArray(t, A)
}

func TestDenseAssignFunc(t *testing.T) {
	A := makeDenseVector()
	testAssignFunc(t, A)
}

func TestDenseAssignVector(t *testing.T) {
	A := makeDenseVector()
	B := makeDenseVector()
	testAssignVector(t, A, B)
}

func TestDenseAssignVectorFunc(t *testing.T) {
	A := makeDenseVector()
	B := makeDenseVector()
	testAssignVectorFunc(t, A, B)
}

func TestDenseAssignProcedure(t *testing.T) {
	A := makeDenseVector()
	testAssignProcedure(t, A)
}

func TestDenseAssignProcedureFunc(t *testing.T) {
	A := makeDenseVector()
	testAssignProcedureFunc(t, A)
}
