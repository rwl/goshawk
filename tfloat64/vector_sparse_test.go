
package tfloat64

import (
	"math/rand"
	"testing"
)

func makeSparseVector() (*Vector) {
	A := NewSparseVector(test_size)
	for i := 0; i < test_size; i++ {
		A.SetQuick(i, rand.Float64())
	}
	return A
}

func TestSparseCardinality(t *testing.T) {
	A := makeSparseVector()
	testCardinality(t, A)
}

func TestSparseEquals(t *testing.T) {
	A := makeSparseVector()
	testEquals(t, A)
}

func TestSparseEqualsVector(t *testing.T) {
	A := makeSparseVector()
	B := makeSparseVector()
	testEqualsVector(t, A, B)
}

func TestSparseMaxLocation(t *testing.T) {
	A := makeSparseVector()
	testMaxLocation(t, A)
}

func TestSparseMinLocation(t *testing.T) {
	A := makeSparseVector()
	testMinLocation(t, A)
}

func TestSparseNegativeValues(t *testing.T) {
	A := makeSparseVector()
	testGetNegativeValues(t, A)
}

func TestSparseNonZeros(t *testing.T) {
	A := makeSparseVector()
	testNonZeros(t, A)
}

func TestSparsePositiveValues(t *testing.T) {
	A := makeSparseVector()
	testPositiveValues(t, A)
}

func TestSparseToArray(t *testing.T) {
	A := makeSparseVector()
	testToArray(t, A)
}

func TestSparseFillArray(t *testing.T) {
	A := makeSparseVector()
	testFillArray(t, A)
}

func TestSparseReshapeMatrix(t *testing.T) {
	A := makeSparseVector()
	testReshapeMatrix(t, A)
}

func TestSparseReshapeCube(t *testing.T) {
	A := makeSparseVector()
	testReshapeCube(t, A)
}

func TestSparseSwap(t *testing.T) {
	A := makeSparseVector()
	B := makeSparseVector()
	testSwap(t, A, B)
}

func TestSparseZDotProduct(t *testing.T) {
	A := makeSparseVector()
	B := makeSparseVector()
	testZDotProduct(t, A, B)
}

func TestSparseZDotProductRange(t *testing.T) {
	A := makeSparseVector()
	B := makeSparseVector()
	testZDotProductRange(t, A, B)
}

func TestSparseZDotProductSelection(t *testing.T) {
	A := makeSparseVector()
	B := makeSparseVector()
	testZDotProductSelection(t, A, B)
}

func TestSparseZSum(t *testing.T) {
	A := makeSparseVector()
	testZSum(t, A)
}

// View tests.

func TestSparseViewFlip(t *testing.T) {
	A := makeSparseVector()
	testViewFlip(t, A)
}

func TestSparseViewPart(t *testing.T) {
	A := makeSparseVector()
	testViewPart(t, A)
}

func TestSparseViewProcedure(t *testing.T) {
	A := makeSparseVector()
	testViewProcedure(t, A)
}

func TestSparseView(t *testing.T) {
	A := makeSparseVector()
	testView(t, A)
}
/*
func TestSparseViewSorted(t *testing.T) {
	A := makeSparseVector()
	testViewSorted(t, A)
}
*/

func TestSparseViewStrides(t *testing.T) {
	A := makeSparseVector()
	testViewStrides(t, A)
}

// Aggregate tests.

func TestSparseAggregate(t *testing.T) {
	A := makeSparseVector()
	testAggregate(t, A)
}

func TestSparseAggregateIndexed(t *testing.T) {
	A := makeSparseVector()
	testAggregateIndexed(t, A)
}

func TestSparseAggregateVector(t *testing.T) {
	A := makeSparseVector()
	B := makeSparseVector()
	testAggregateVector(t, A, B)
}

// Assign tests.

func TestSparseAssign(t *testing.T) {
	A := makeSparseVector()
	testAssign(t, A)
}

func TestSparseAssignArray(t *testing.T) {
	A := makeSparseVector()
	testAssignArray(t, A)
}

func TestSparseAssignFunc(t *testing.T) {
	A := makeSparseVector()
	testAssignFunc(t, A)
}

func TestSparseAssignVector(t *testing.T) {
	A := makeSparseVector()
	B := makeSparseVector()
	testAssignVector(t, A, B)
}

func TestSparseAssignVectorFunc(t *testing.T) {
	A := makeSparseVector()
	B := makeSparseVector()
	testAssignVectorFunc(t, A, B)
}

func TestSparseAssignProcedure(t *testing.T) {
	A := makeSparseVector()
	testAssignProcedure(t, A)
}

func TestSparseAssignProcedureFunc(t *testing.T) {
	A := makeSparseVector()
	testAssignProcedureFunc(t, A)
}
