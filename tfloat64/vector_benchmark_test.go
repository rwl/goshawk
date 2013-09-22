
package tfloat64

import (
	"testing"
	"math"
	"math/rand"
)

var (
	MATRIX_SIZE_1D int = int(math.Pow(2, 19))
	MATRIX_SIZE_2D = []int{ int(math.Pow(2, 12)), int(math.Pow(2, 12)) }
	MATRIX_SIZE_3D = []int{ int(math.Pow(2, 7)), int(math.Pow(2, 7)), int(math.Pow(2, 7)) }
	NTHREADS = []int { 1, 2, 4, 8 }
	result float64
	i int
	array []float64
)

func randomArray() []float64 {
	a := make([]float64, MATRIX_SIZE_1D)
	for i := 0; i < MATRIX_SIZE_1D; i++ {
		a[i] = rand.Float64()
	}
	return a
}

func BenchmarkVectorAggregate(b *testing.B) {
	A := NewVectorArray(randomArray())
	b.ResetTimer()
	var sum float64
	for i := 0; i < b.N; i++ {
		sum = A.Aggregate(Plus, Square)
	}
	result = sum
}

func BenchmarkVectorAggregateView(b *testing.B) {
	A, _ := NewVector(MATRIX_SIZE_1D).ViewFlip().AssignArray(randomArray())
	b.ResetTimer()
	var sum float64
	for i := 0; i < b.N; i++ {
		sum = A.Aggregate(Plus, Square)
	}
	result = sum
}

func BenchmarkVectorAssignVector(b *testing.B) {
	A := NewVector(MATRIX_SIZE_1D)
	B := NewVectorArray(randomArray())
	A.Assign(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		A.AssignVector(B)
	}
}

func BenchmarkVectorAssignVectorView(b *testing.B) {
	A := NewVector(MATRIX_SIZE_1D)
	Av := A.ViewFlip()
	B := NewVectorArray(randomArray())
	Bv := B.ViewFlip()
	Av.Assign(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Av.AssignVector(Bv)
	}
}

func BenchmarkVectorZDotProduct(b *testing.B) {
	A := NewVectorArray(randomArray())
	B := NewVectorArray(randomArray())
	b.ResetTimer()
	var product float64
	for i := 0; i < b.N; i++ {
		product = A.ZDotProduct(B)
	}
	result = product
}

func BenchmarkVectorZDotProductView(b *testing.B) {
	A := NewVectorArray(randomArray())
	Av := A.ViewFlip()
	B := NewVectorArray(randomArray())
	Bv := B.ViewFlip()
	b.ResetTimer()
	var product float64
	for i := 0; i < b.N; i++ {
		product = Av.ZDotProduct(Bv)
	}
	result = product
}

func BenchmarkVectorAssignFunc(b *testing.B) {
	f := Multiply(2.5)
	A := NewVectorArray(randomArray())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		A.AssignFunc(f)
	}
}

func BenchmarkVectorAssignFuncView(b *testing.B) {
	f := Multiply(2.5)
	Av, _ := NewVector(MATRIX_SIZE_1D).ViewFlip().AssignArray(randomArray())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Av.AssignFunc(f)
	}
}

func BenchmarkVectorAssignVectorFunc(b *testing.B) {
	A := NewVectorArray(randomArray())
	B := NewVectorArray(randomArray())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		A.AssignVectorFunc(B, Div)
	}
}

func BenchmarkVectorAssignVectorFuncView(b *testing.B) {
	Av, _ := NewVector(MATRIX_SIZE_1D).ViewFlip().AssignArray(randomArray())
	Bv, _ := NewVector(MATRIX_SIZE_1D).ViewFlip().AssignArray(randomArray())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Av.AssignVectorFunc(Bv, Div)
	}
}

func BenchmarkVectorAggregateVectorFunc(b *testing.B) {
	A := NewVectorArray(randomArray())
	B := NewVectorArray(randomArray())
	b.ResetTimer()
	var sum float64
	for i := 0; i < b.N; i++ {
		sum, _ = A.AggregateVector(B, Plus, Mult)
	}
	result = sum
}

func BenchmarkVectorAggregateVectorFuncView(b *testing.B) {
	Av, _ := NewVector(MATRIX_SIZE_1D).ViewFlip().AssignArray(randomArray())
	Bv, _ := NewVector(MATRIX_SIZE_1D).ViewFlip().AssignArray(randomArray())
	b.ResetTimer()
	var sum float64
	for i := 0; i < b.N; i++ {
		sum, _ = Av.AggregateVector(Bv, Plus, Mult)
	}
	result = sum
}

func BenchmarkVectorAssign(b *testing.B) {
	A := NewVector(MATRIX_SIZE_1D)
	value := rand.Float64()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		A.Assign(value)
	}
}

func BenchmarkVectorAssignView(b *testing.B) {
	Av := NewVector(MATRIX_SIZE_1D).ViewFlip()
	value := rand.Float64()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Av.Assign(value)
	}
}

var testProcedure = func(element float64) bool {
	if math.Abs(element) > 0.1 {
		return true
	} else {
		return false
	}
}

func BenchmarkVectorAssignProcedure(b *testing.B) {
	A := NewVectorArray(randomArray())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		A.AssignProcedure(testProcedure, -1)
	}
}

func BenchmarkVectorAssignProcedureView(b *testing.B) {
	Av, _ := NewVector(MATRIX_SIZE_1D).ViewFlip().AssignArray(randomArray())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Av.AssignProcedure(testProcedure, -1)
	}
}

func BenchmarkVectorAssignProcedureFunc(b *testing.B) {
	A := NewVectorArray(randomArray())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		A.AssignProcedureFunc(testProcedure, Square)
	}
}

func BenchmarkVectorAssignProcedureFuncView(b *testing.B) {
	Av, _ := NewVector(MATRIX_SIZE_1D).ViewFlip().AssignArray(randomArray())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Av.AssignProcedureFunc(testProcedure, Square)
	}
}

func BenchmarkVectorCardinality(b *testing.B) {
	A := NewVectorArray(randomArray())
	b.ResetTimer()
	var card int
	for i := 0; i < b.N; i++ {
		card = A.Cardinality()
	}
	i = card
}

func BenchmarkVectorCardinalityView(b *testing.B) {
	Av, _ := NewVector(MATRIX_SIZE_1D).ViewFlip().AssignArray(randomArray())
	b.ResetTimer()
	var card int
	for i := 0; i < b.N; i++ {
		card = Av.Cardinality()
	}
	i = card
}

func BenchmarkVectorPositiveValues(b *testing.B) {
	A := NewVectorArray(randomArray())
	var indexes []int
	var values []float64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		A.PositiveValues(&indexes, &values)
	}
}

func BenchmarkVectorPositiveValuesView(b *testing.B) {
	Av, _ := NewVector(MATRIX_SIZE_1D).ViewFlip().AssignArray(randomArray())
	var indexes []int
	var values []float64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Av.PositiveValues(&indexes, &values)
	}
}

func BenchmarkVectorNegativeValues(b *testing.B) {
	A := NewVectorArray(randomArray())
	var indexes []int
	var values []float64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		A.NegativeValues(&indexes, &values)
	}
}

func BenchmarkVectorNegativeValuesView(b *testing.B) {
	Av, _ := NewVector(MATRIX_SIZE_1D).ViewFlip().AssignArray(randomArray())
	var indexes []int
	var values []float64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Av.NegativeValues(&indexes, &values)
	}
}

func BenchmarkVectorMaxLocation(b *testing.B) {
	A := NewVectorArray(randomArray())
	b.ResetTimer()
	var max float64
	var loc int
	for i := 0; i < b.N; i++ {
		max, loc = A.MaxLocation()
	}
	result = max
	i = loc
}

func BenchmarkVectorMaxLocationView(b *testing.B) {
	Av, _ := NewVector(MATRIX_SIZE_1D).ViewFlip().AssignArray(randomArray())
	b.ResetTimer()
	var max float64
	var loc int
	for i := 0; i < b.N; i++ {
		max, loc = Av.MaxLocation()
	}
	result = max
	i = loc
}

func BenchmarkVectorMinLocation(b *testing.B) {
	A := NewVectorArray(randomArray())
	b.ResetTimer()
	var max float64
	var loc int
	for i := 0; i < b.N; i++ {
		max, loc = A.MinLocation()
	}
	result = max
	i = loc
}

func BenchmarkVectorMinLocationView(b *testing.B) {
	Av, _ := NewVector(MATRIX_SIZE_1D).ViewFlip().AssignArray(randomArray())
	b.ResetTimer()
	var max float64
	var loc int
	for i := 0; i < b.N; i++ {
		max, loc = Av.MinLocation()
	}
	result = max
	i = loc
}

func BenchmarkVectorReshapeMatrix(b *testing.B) {
	A := NewVectorArray(randomArray())
	b.ResetTimer()
	rows := MATRIX_SIZE_1D / 64
	cols := 64
	for i := 0; i < b.N; i++ {
		A.ReshapeMatrix(rows, cols)
	}
}

func BenchmarkVectorReshapeMatrixView(b *testing.B) {
	Av, _ := NewVector(MATRIX_SIZE_1D).ViewFlip().AssignArray(randomArray())
	b.ResetTimer()
	rows := MATRIX_SIZE_1D / 64
	cols := 64
	for i := 0; i < b.N; i++ {
		Av.ReshapeMatrix(rows, cols)
	}
}

func BenchmarkVectorReshapeCube(b *testing.B) {
	A := NewVectorArray(randomArray())
	b.ResetTimer()
	slices := MATRIX_SIZE_1D / 64
	rows := 16
	cols := 4
	for i := 0; i < b.N; i++ {
		A.ReshapeCube(slices, rows, cols)
	}
}

func BenchmarkVectorReshapeCubeView(b *testing.B) {
	Av, _ := NewVector(MATRIX_SIZE_1D).ViewFlip().AssignArray(randomArray())
	b.ResetTimer()
	slices := MATRIX_SIZE_1D / 64
	rows := 16
	cols := 4
	for i := 0; i < b.N; i++ {
		Av.ReshapeCube(slices, rows, cols)
	}
}

func BenchmarkVectorSwap(b *testing.B) {
	A := NewVectorArray(randomArray())
	B := NewVectorArray(randomArray())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		A.Swap(B)
	}
}

func BenchmarkVectorSwapView(b *testing.B) {
	Av, _ := NewVector(MATRIX_SIZE_1D).ViewFlip().AssignArray(randomArray())
	Bv, _ := NewVector(MATRIX_SIZE_1D).ViewFlip().AssignArray(randomArray())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Av.Swap(Bv)
	}
}

func BenchmarkVectorToArray(b *testing.B) {
	A := NewVectorArray(randomArray())
	var a []float64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a = A.ToArray()
	}
	array = a
}

func BenchmarkVectorToArrayView(b *testing.B) {
	Av, _ := NewVector(MATRIX_SIZE_1D).ViewFlip().AssignArray(randomArray())
	var a []float64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a = Av.ToArray()
	}
	array = a
}

func BenchmarkVectorZDotProductRange(b *testing.B) {
	A := NewVectorArray(randomArray())
	B := NewVectorArray(randomArray())
	var product float64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		product = A.ZDotProductRange(B, 5, B.Size() - 10)
	}
	result = product
}

func BenchmarkVectorZDotProductRangeView(b *testing.B) {
	Av, _ := NewVector(MATRIX_SIZE_1D).ViewFlip().AssignArray(randomArray())
	Bv, _ := NewVector(MATRIX_SIZE_1D).ViewFlip().AssignArray(randomArray())
	var product float64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		product = Av.ZDotProductRange(Bv, 5, Bv.Size() - 10)
	}
	result = product
}

func BenchmarkVectorZSum(b *testing.B) {
	A := NewVectorArray(randomArray())
	var sum float64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum = A.ZSum()
	}
	result = sum
}

func BenchmarkVectorZSumView(b *testing.B) {
	Av, _ := NewVector(MATRIX_SIZE_1D).ViewFlip().AssignArray(randomArray())
	var sum float64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum = Av.ZSum()
	}
	result = sum
}
