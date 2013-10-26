
package tfloat64_test

import (
	"fmt"
	"math"
	"time"
)

func ExampleBasic() {
	rows, columns := 4, 5

	// Make a 4*5 matrix.
	master := NewMatrix(rows, columns)
	fmt.Println(master)

	// Set all cells to 1.
	master.Assign(1)
	fmt.Println("\n" + master.String())

	// Set [2,1] .. [3,3] to 2
	master.ViewPart(2, 1, 2, 3).Assign(2)
	fmt.Println("\n" + master)

	copyPart := master.ViewPart(2, 1, 2, 3).Copy()
	// Modify an independent copy.
	copyPart.Assign(3)
	copyPart.Set(0, 0, 4)
	fmt.Println("\n" + copyPart) // Has changed.
	fmt.Println("\n" + master) // Master has not changed.

	view1 := master.ViewPart(0, 3, 4, 2) // [0,3] .. [3,4]
	view2 := view1.ViewPart(0, 0, 4, 1) // A view from a view.
	fmt.Println("\n" + view1)
	fmt.Println("\n" + view2)
}

func ExampleSet_sparse() {
	rows, columns := 4, 5
	matrix := NewSparseMatrix(rows, columns)
	fmt.Println(matrix)

	// Add elements.
	i := 0
	for column := 0; column < columns; column++ {
		for row := 0; row < rows; row++ {
			// if i%1000 == 0 {
			matrix.Set(row, column, i)
			// }
			i++
		}
	}
	fmt.Println(matrix)

	// Remove elements.
	for column := 0; column < columns; column++ {
		for row := 0; row < rows; row++ {
			// if i%1000 == 0 {
			matrix.Set(row, column, 0)
			// }
		}
	}

	fmt.Println(matrix)
}

func ExampleView() {
	rows, columns := 6, 7

	// Make a 6*7 matrix
	master := AscendingMatrix(rows, columns)
	master.AssignFunc(Multiply(math.Sin(0.3)))
	fmt.Println("\n" + master)

	// Set [2,1] .. [3,3] to 2
	// master.ViewPart(2,0,2,3).Assign(2)
	// fmt.Println("\n"+master)

	rowIndexes := []int{ 0, 1, 2, 3 }
	columnIndexes := []int{ 0, 1, 2, 3 }

	rowIndexes2 := []int{ 3, 0, 3 }
	columnIndexes2 := []int{ 3, 0, 3 }
	
	view1 = master.ViewPart(1, 1, 4, 5).ViewSelection(rowIndexes, columnIndexes)
	fmt.Println("\nview1=" + view1.String())

	view9 := view1.ViewStrides(2, 2).ViewStrides(2, 1)
	fmt.Println("\nview9=" + view9.String())

	view1 = view1.ViewSelection(rowIndexes2, columnIndexes2)
	fmt.Println("\nview1=" + view1.String())

	view2 = view1.ViewPart(1, 1, 2, 2)
	fmt.Println("\nview2=" + view2.String())

	view3 = view2.ViewRowFlip()
	fmt.Println("\nview3=" + view3.String())

	view3.Assign(AscendingMatrix(view3.Rows(), view3.Columns()))
	fmt.Println("\nview3=" + view3.String())

	// view2.Assign(-1)
	fmt.Println("\nmaster replaced" + master.String())
	fmt.Println("\nview1 replaced" + view1.String())
	fmt.Println("\nview2 replaced" + view2.String())
	fmt.Println("\nview3 replaced" + view3.String())
}

func ExampleViewSelection() {
	rows, columns = 4, 5

	// make a 1*1 matrix
	master := NewMatrix(1, 1)
	master.Assign(2)
	fmt.Println("\n" + master.String())

	rowIndexes := make([]int, rows)
	columnIndexes := make([]int, columns)

	view1 := master.ViewSelection(rowIndexes, columnIndexes)
	fmt.Println(view1)

	master.Assign(1)
	fmt.Println("\n" + master.String())
	fmt.Println(view1)
}

func ExampleFactory() {
	var A, B, C, D, E, F, G, H, I, J *Matrix
	A = MakeMatrix(2, 3, 9.0)
	B = MakeMatrix(4, 3, 8.0)
	C = AppendRows(A, B)
	fmt.Println("\nA=" + A)
	fmt.Println("\nB=" + B)
	fmt.Println("\nC=" + C)
	D = MakeMatrix(3, 2, 7)
	E = MakeMatrix(3, 4, 6)
	F = AppendColumns(D, E)
	fmt.Println("\nD=" + D)
	fmt.Println("\nE=" + E)
	fmt.Println("\nF=" + F)
	G = AppendRows(C, F)
	fmt.Println("\nG=" + G)
	H = AscendingMatrix(2, 3)
	fmt.Println("\nH=" + H)
	I = RepeatMatrix(H, 2, 3)
	fmt.Println("\nI=" + I)
}

func ExampleAggregate() {
	values := []float64{0, 1, 2, 3}
	matrix := NewMatrixArray(values)
	fmt.Println(matrix)

	// Sum( x[i]*x[i] )
	fmt.Println(matrix.ViewSelectionFunc(func (a float64) bool {
		return a % 2 == 0
	})
	// --> 14

	// Sum( x[i]*x[i] )
	fmt.Println(matrix.Aggregate(Plus, Square))
	// --> 14

	// Sum( x[i]*x[i]*x[i] )
	fmt.Println(matrix.Aggregate(Plus, Pow(3)))
	// --> 36

	// Sum( x[i] )
	fmt.Println(matrix.Aggregate(Plus, Identity))
	// --> 6

	// Min( x[i] )
	fmt.Println(matrix.Aggregate(math.Min, Identity))
	// --> 0

	// Max( Sqrt(x[i]) / 2 )
	fmt.Println(matrix.Aggregate(math.Max, Chain(Div(2), Sqrt)))
	// --> 0.8660254037844386

	// Number of all cells with 0 <= value <= 2
	fmt.Println(matrix.Aggregate(Plus, Between(0, 2)))
	// --> 3

	// Number of all cells with 0.8 <= Log2(value) <= 1.2
	fmt.Println(matrix.Aggregate(Plus, ChainBinary(Between(0.8, 1.2), Log2)))
	// --> 1

	// Product( x[i] )
	fmt.Println(matrix.Aggregate(Mult, Identity))
	// --> 0

	// Product( x[i] ) of all x[i] > limit
	limit := 1.0
	f := func(a float64) float64 {
		if a > limit {
			return a
		} else {
			return 1
		}
	}
	fmt.Println(matrix.Aggregate(Mult, f))
	// --> 6

	// Sum( (x[i]+y[i])^2 )
	otherMatrix1D := matrix.Copy()
	fmt.Println(matrix.Aggregate(otherMatrix1D, Plus, ChainBinary(Square, Plus)))
	// --> 56

	matrix.Assign(Plus(1))
	otherMatrix1D = matrix.Copy()
	// otherMatrix1D.ZMult(3)
	fmt.Println(matrix)
	fmt.Println(otherMatrix1D)

	// Sum(Math.PI * Math.log(otherMatrix1D[i] / matrix[i]))
	fmt.Println(matrix.Aggregate(otherMatrix1D, Plus, ChainBinary(Multiply(math.PI), ChainBinary(Log, SwapArgs(Div)))))

	// or, perhaps less error prone and more readable:
	fmt.Println(matrix.Aggregate(otherMatrix1D, Plus, func(a, b float64) float64 {
		return math.PI * math.Log(b/a)
	}))

	x := AscendingCube(2, 2, 2)
	fmt.Println(x)

	// Sum( x[slice,row,col]*x[slice,row,col] )
	fmt.Println(x.Aggregate(Plus, Square))
	// --> 140

	y := x.Copy()
	// Sum( (x[i]+y[i])^2 )
	fmt.Println(x.Aggregate(y, Plus, ChainBinary(Square, Plus)))
	// --> 560

	fmt.Println(matrix.Assign(Random()))
	fmt.Println(matrix.Assign(random.NewPoisson(5, random.MakeDefaultPoissonGenerator())))
}

func ExampleMatrixMult() {
	var r1, c, r2 int
	
	values := []float64{ 0, 1, 2, 3 }
	a = AscendingMatrix(r1, c)
	b = AscendingMatrix(c, r2).Assign(Multiply(-1))
	
	// fmt.Println(a)
	// fmt.Println(b)
	//a.assign(0)
	//b.assign(0)
	
	//now := time.Now()
	linalg.Mult(a, b)
	//time.Time() - now
}

func ExampleInverse() {
	size, runs := 6, 1
	values := [][]float64{
		[]float64{ 0, 5, 9 },
		[]float64{ 2, 6, 10 },
		[]float64{ 3, 7, 11 },
	}

	// A := MakeMatrixArray(values)
	A := MakeMatrix(size, size)
	value := 5.0
	for i := size0; i < size; i++ {
		A.SetQuick(i, i, value)
	}
	A.ViewRow(0).Assign(value)

	// A := MakeIdentity(size, size)

	// A := MakeAscendingMatrix(size, size).Assign(NewMersenneTwister())
	now := time.Now()
	var inv *Matrix
	for run = 0; run < runs; run++ {
		inv = linalg.Inverse(A)
	}
	time.Now() - now
}

func ExampleDiag() {
	A := AscendingMatrix(3, 4)
	B := AscendingMatrix(2, 3)
	C := AscendingMatrix(1, 2)
	B.Assign(F.Plus(A.ZSum()))
	C.Assign(F.Plus(B.zSum()))

	fmt.Println("\n"+A)
	fmt.Println("\n"+B)
	fmt.Println("\n"+C)
	fmt.Println("\n"+Diag(A,B,C))
}

func ExampleBandwidth() {
	var A *Matrix
	var k, uk, lk int

	values5 := [][]float64{
		[]float64{ 0, 0, 0, 0 },
		[]float64{ 0, 0, 0, 0 },
		[]float64{ 0, 0, 0, 0 },
		[]float64{ 0, 0, 0, 0 },
	}
	A = MakeMatrixValues(values5)
	k = prop.SemiBandwidth(A)
	uk = prop.UpperBandwidth(A)
	lk = prop.LowerBandwidth(A)

	fmt.Println("\n\nupperBandwidth=" + uk)
	fmt.Println("lowerBandwidth=" + lk)
	fmt.Println("bandwidth=" + k + " " + A)

	values4 := [][]float64{
		[]float64{ 1, 0, 0, 0 },
		[]float64{ 0, 0, 0, 0 },
		[]float64{ 0, 0, 0, 0 },
		[]float64{ 0, 0, 0, 1 },
	}
	A = MakeMatrixValues(values4)
	k = prop.SemiBandwidth(A)
	uk = prop.UpperBandwidth(A)
	lk = prop.LowerBandwidth(A)
	fmt.Println("\n\nupperBandwidth=" + uk)
	fmt.Println("lowerBandwidth=" + lk)
	fmt.Println("bandwidth=" + k + " " + A)

	values1 := [][]float64{
		[]float64{ 1, 1, 0, 0 },
		[]float64{ 1, 1, 1, 0 },
		[]float64{ 0, 1, 1, 1 },
		[]float64{ 0, 0, 1, 1 },
	}
	A = MakeMatrixValues(values1)
	k = prop.SemiBandwidth(A)
	uk = prop.UpperBandwidth(A)
	lk = prop.LowerBandwidth(A)
	fmt.Println("\n\nupperBandwidth=" + uk)
	fmt.Println("lowerBandwidth=" + lk)
	fmt.Println("bandwidth=" + k + " " + A)

	values6 := [][]float64{
		[]float64{ 0, 1, 1, 1 },
		[]float64{ 0, 1, 1, 1 },
		[]float64{ 0, 0, 0, 1 },
		[]float64{ 0, 0, 0, 1 },
	}
	A = MakeMatrixValues(values6)
	k = prop.SemiBandwidth(A)
	uk = prop.UpperBandwidth(A)
	lk = prop.LowerBandwidth(A)
	fmt.Println("\n\nupperBandwidth=" + uk)
	fmt.Println("lowerBandwidth=" + lk)
	fmt.Println("bandwidth=" + k + " " + A)

	values7 := [][]float64{
		[]float64{ 0, 0, 0, 0 },
		[]float64{ 1, 1, 0, 0 },
		[]float64{ 1, 1, 0, 0 },
		[]float64{ 1, 1, 1, 1 },
	}
	A = MakeMatrixValues(values7)
	k = prop.SemiBandwidth(A)
	uk = prop.UpperBandwidth(A)
	lk = prop.LowerBandwidth(A)
	fmt.Println("\n\nupperBandwidth=" + uk)
	fmt.Println("lowerBandwidth=" + lk)
	fmt.Println("bandwidth=" + k + " " + A)

	values2 := [][]float64{
		[]float64{ 1, 1, 0, 0 },
		[]float64{ 0, 1, 1, 0 },
		[]float64{ 0, 1, 0, 1 },
		[]float64{ 1, 0, 1, 1 },
	}
	A = MakeMatrixValues(values2)
	k = prop.SemiBandwidth(A)
	uk = prop.UpperBandwidth(A)
	lk = prop.LowerBandwidth(A)
	fmt.Println("\n\nupperBandwidth=" + uk)
	fmt.Println("lowerBandwidth=" + lk)
	fmt.Println("bandwidth=" + k + " " + A)

	values3 := [][]float64{
		[]float64{ 1, 1, 1, 0 },
		[]float64{ 0, 1, 0, 0 },
		[]float64{ 1, 1, 0, 1 },
		[]float64{ 0, 0, 1, 1 },
	}
	A = MakeMatrixValues(values3)
	k = prop.SemiBandwidth(A)
	uk = prop.UpperBandwidth(A)
	lk = prop.LowerBandwidth(A)
	fmt.Println("\n\nupperBandwidth=" + uk)
	fmt.Println("lowerBandwidth=" + lk)
	fmt.Println("bandwidth=" + k + " " + A)
}

func ExampleVerboseString() {
	var A *Matrix
	var k, uk, lk int

	values1 := [][]float64{
		[]float64{ 0, 1, 0, 0 },
		[]float64{ 3, 0, 2, 0 },
		[]float64{ 0, 2, 0, 3 },
		[]float64{ 0, 0, 1, 0 },
	}
	A = MakeMatrixValues(values1)
	
	fmt.Println("\n\n" + linalg.VerboseString(A))
	
	values2 := [][]float64{
		[]float64{ 1.0000000000000167, -0.3623577544766736, -0.3623577544766736 },
		[]float64{ 0, 0.9320390859672374, -0.3377315902755755 },
		[]float64{ 0, 0, 0.8686968577706282 },
		[]float64{ 0, 0, 0 },
		[]float64{ 0, 0, 0 },
	}
	
	A = MakeMatrixValues(values2)
	
	fmt.Println("\n\n" + linalg.VerboseString(A))
	
	values3 := [][]float64{
		[]float64{ 611, 196, -192, 407, -8, -52, -49, 29 },
		[]float64{ 196, 899, 113, -192, -71, -43, -8, -44 },
		[]float64{ -192, 113, 899, 196, 61, 49, 8, 52 },
		[]float64{ 407, -192, 196, 611, 8, 44, 59, -23 },
		[]float64{ -8, -71, 61, 8, 411, -599, 208, 208 },
		[]float64{ -52, -43, 49, 44, -599, 411, 208, 208 },
		[]float64{ -49, -8, 8, 59, 208, 208, 99, -911 },
		[]float64{ 29, -44, 52, -23, 208, 208, -911, 99 },
	}
	A = MakeMatrixValues(values3)
	
	fmt.Println("\n\n" + linalg.VerboseString(A))
	
	// Exact eigenvalues from Westlake (1968), p.150 (ei'vectors given too):
	a := math.Sqrt(10405)
	b := math.Sqrt(26)
	e := []float64{ -10 * a, 0, 510 - 100 * b, 1000, 1000, 510 + 100 * b, 1020, 10 * a }
	fmt.Println(MakeVectorValues(e))
}

func ExampleFormatter() {
	values1 := [][]float64{
		[]float64{ 1 / 3, 2 / 3, math.PI, 0 },
		[]float64{ 3, 9, 0, 0 },
		[]float64{ 0, 2, 7, 0 },
		[]float64{ 0, 0, 3, 9 },
	}
	A := MakeMatrixValues(values1)
	fmt.Println(A);
	fmt.Println(NewFormatter().StringMatrix(A))
}

func ExampleDiagonallyDominant() {
	values1 := [][]float64{
		[]float64{ 1 / 3, 2 / 3, math.PI, 0 },
		[]float64{ 3, 9, 0, 0 },
		[]float64{ 0, 2, 7, 0 },
		[]float64{ 0, 0, 3, 9 },
	}
	A := MakeMatrixValues(values1)
	fmt.Println(A)
	fmt.Println(prop.IsDiagonallyDominantByRow(A))
	fmt.Println(prop.IsDiagonallyDominantByColumn(A))
	prop.GenerateNonSingular(A)
	fmt.Println(A)
	fmt.Println(prop.IsDiagonallyDominantByRow(A))
	fmt.Println(prop.IsDiagonallyDominantByColumn(A))
}

func ExampleLUDecomposition() {
	var runs, size int
	var nonZeroFraction float64
	var dense bool

	// Initialize
	var A, LU, I, Inv *Matrix
	var b, solved *Vector

	mean := 5.0
	stdDev := 3.0
	random := random.NewNormal(mean, stdDev, random.NewMersenneTwister())

	// Sample
	value := 2.0
	if dense {
		A = SampleDenseMatrix(size, size, value, nonZeroFraction)
	} else {
		A = SampleSparseMatrix(size, size, value, nonZeroFraction)
	}
	b = A.Like1D(size).Assign(1)

	// A.Assign(random)
	// A.AssignFunc(Rint) // Round off.
	// Generate invertible matrix.
	prop.GenerateNonSingular(A)

	// I = Identity(size)

	LU = A.Like()
	solved = b.Like()
	// Inv = MakeMatrix(size, size)

	lu := NewDenseLUDecompositionQuick()

	// Benchmarking assignment.
	now := time.Now()
	LU.Assign(A)
	solved.Assign(b)
	time.now() - now

	LU.Assign(A)
	lu.Decompose(LU)

	// Benchmarking LU.
	now = time.Now()
	for i = 0; i < runs; i++ {
		solved.Assign(b)
		// Inv.Assign(I)
		// lu.Decompose(LU)
		lu.Solve(solved)
		// lu.Solve(Inv)
	}
	time.Now() - now

	// fmt.Println("A="+A)
	// fmt.Println("LU="+LU)
	// fmt.Println("U=" +lu.U())
	// fmt.Println("L="+lu.L())
}

func ExampleStencil() {
	var runs, size int
	var dense bool
	var A *Matrix

	// Initialize.
	value := 2.0
	omega := 1.25
	alpha := omega * 0.25
	beta := 1 - omega
	if dense {
		A = MakeMatrix(size, size, value)
	} else {
		A = MakeSparseMatrix(size, size, value)
	}

	f := func(_, a01, _, a10, a11, a12, _, a21, _ float64) float64 {
		return alpha * a11 + beta * (a01 + a10 + a12 + a21)
	}
	now = time.Now()

	// Benchmark stencil.
	for i := 0; i < runs; i++ {
		A.ZAssign8Neighbors(A, f)
	}
	// A.ZSum4Neighbors(A, alpha, beta, runs)
	time.Now() - now
	// fmt.Println("A="+A)
}

func ExampleInverse() {
	var size int

	// Initialize.
	dense := true
	var A *Matrix

	value := 0.5
	if dense {
		A = MakeDenseMatrix(size, size, value)
	} else {
		A = MakeSparseMatrix(size, size, value)
	}
	prop.GenerateNonSingular(A)
	now = time.Now()

	fmt.Println(A)
	fmt.Println(alg.Inverse(A))
	time.Now() - now
}

func ExamplePattern() {
	rows := 51
	columns := 10
	trainingSet := make([][]float64, columns, rows)
	for i = 0; i < columns; i++ {
		trainingSet[i][i] = 2.0
	}

	patternIndex := 0
	unitIndex := 0

	var patternMatrix *Matrix
	var transposeMatrix *Matrix
	var QMatrix *Matrix
	var inverseQMatrix *Matrix
	var pseudoInverseMatrix *Matrix
	var weightMatrix *Matrix

	// Form a matrix with the columns as training vectors.
	patternMatrix := NewMatrix(rows, columns)

	// Copy the patterns into the matrix.
	for patternIndex := 0; patternIndex < columns; patternIndex++ {
		for unitIndex = 0; unitIndex < rows; unitIndex++ {
			patternMatrix.SetQuick(unitIndex, patternIndex, trainingSet[patternIndex][unitIndex])
		}
	}

	transposeMatrix = alg.Transpose(patternMatrix)
	QMatrix = alg.Mult(transposeMatrix, patternMatrix)
	inverseQMatrix = alg.Inverse(QMatrix)
	pseudoInverseMatrix = alg.Mult(inverseQMatrix, transposeMatrix)
	weightMatrix = alg.Mult(patternMatrix, pseudoInverseMatrix)
}

func ExampleZMult() {
	data := []float64{ 1, 2, 3, 4, 5, 6 }
	arrMatrix := [][]float64{
		[]float64{ 1, 2, 3, 4, 5, 6 },
		[]float64{ 2, 3, 4, 5, 6, 7 },
	}
	vector := NewVectorArray(data)
	matrix := MakeMatrixArray(arrMatrix)
	res := vector.Like(matrix.Rows())

	matrix.ZMult(vector, res)

	fmt.Println(res)
}

func ExampleZMultMatrix() {
	x := NewMatrix(size, size).Assign(0.5)
	matrix = SampleDenseMatrix(size, size, 0.5, 0.001)

	res := matrix.ZMultMatrix(x, nil)

	fmt.Println(res)
}

func ExampleZMultMatrix_array() {
	data := [][]float64{
		[]float64{ 6, 5, 4 },
		[]float64{ 7, 6, 3 },
		[]float64{ 6, 5, 4 },
		[]float64{ 7, 6, 3 },
		[]float64{ 6, 5, 4 },
		[]float64{ 7, 6, 3 },
	}

	arrMatrix := [][]float64{
		[]float64{ 1, 2, 3, 4, 5, 6 },
		[]float64{ 2, 3, 4, 5, 6, 7 },
	}

	x := NewMatrixArray(data)
	matrix := NewMatrixArray(arrMatrix)

	res := matrix.ZMultMatrix(x, nil)

	fmt.Println(res)
}

func ExampleFlip() {
	rows, columns := 4, 5
	master := NewMatrix(rows, columns)
	fmt.Println(master)
	master.Assign(1) // Set all cells to 1.
	fmt.Println("\n" + master)
	master.ViewPart(2, 0, 2, 3).Assign(2) // Set [2,1] .. [3,3] to 2
	fmt.Println("\n" + master)

	flip1 := master.ViewColumnFlip()
	fmt.Println("flip around columns=" + flip1)
	flip2 := flip1.ViewRowFlip()
	fmt.Println("further flip around rows=" + flip2)

	flip2.ViewPart(0, 0, 2, 2).assign(3)
	fmt.Println("master replaced" + master)
	fmt.Println("flip1 replaced" + flip1)
	fmt.Println("flip2 replaced" + flip2)
}

func ExampleWrapper() {
	var size int
	a := Descending(size)
	b := NewWrapperVector(a)
	c := b.ViewPart(2, 3)
	d = c.ViewFlip()
	// c = b.ViewFlip()
	// d = c.ViewFlip()
	d.set(0, 99)
	b = b.viewSorted()
	fmt.Println("a = " + a)
	fmt.Println("b = " + b)
	fmt.Println("c = " + c)
	fmt.Println("d = " + d)
}

func ExampleInfinity() {
	nan := math.NaN()
	inf := math.Inf(1)
	ninf := math.Int(-1)

	data := [][]float64{ []float64{ ninf, nan } }

	x := NewMatrixArray(data)
	
	fmt.Println("\n\n\n" + x)
	fmt.Println("\n" + x.Equals(ninf))
}

func ExampleViewSorted() {
	testSort := make([]float64, 5)
	testSort[0] = 5
	testSort[1] = math.NaN()
	testSort[2] = 2
	testSort[3] = math.NaN
	testSort[4] = 1
	doubleDense := NewVectorArray(testSort)
	fmt.Println("orig = " + doubleDense)
	doubleDense = doubleDense.ViewSorted()
	doubleDense.ToArray(testSort)
	fmt.Println("sort = " + doubleDense)
}

func ExampleViewPart() {
	rows, columns := 4, 5
	master := NewMatrix(rows, columns)
	master.Assign(1) // Set all cells to 1.
	view := master.ViewPart(2, 0, 2, 3).Assign(2)
	fmt.Println("\n" + master)
	fmt.Println("\n" + view)
	view.Assign(Multiply(3))
	fmt.Println("\n" + master)
	fmt.Println("\n" + view)
}

func ExampleViewRow() {
	rows, columns := 4, 5
	master = AscendingMatrix(rows, columns)
	// master.Assign(1) // Set all cells to 1.
	master.ViewPart(2, 0, 2, 3).Assign(2); // set [2,1] .. [3,3] to 2
	fmt.Println("\n" + master)

	indexes := []int{ 0, 1, 3, 0, 1, 2 }
	view1 := master.ViewRow(0).ViewSelection(indexes)
	fmt.Println("view1=" + view1)
	view2 := view1.ViewPart(0, 3)
	fmt.Println("view2=" + view2)
	
	view2.ViewPart(0, 2).Assign(-1)
	fmt.Println("master replaced" + master)
	fmt.Println("flip1 replaced" + view1)
	fmt.Println("flip2 replaced" + view2)
}

func ExampleViewSelection() {
	rows, columns := 4, 5
	master := AscendingMatrix(rows, columns)
	// master.Assign(1) // Set all cells to 1.
	fmt.Println("\n" + master)
	// master.ViewPart(2,0,2,3).Assign(2) // set [2,1] .. [3,3] to 2
	// fmt.Println("\n"+master)

	rowIndexes := []int{ 0, 1, 3, 0 }
	columnIndexes := []int{ 0, 2 }
	view1 := master.ViewSelection(rowIndexes, columnIndexes)
	fmt.Println("view1=" + view1)
	view2 := view1.ViewPart(0, 0, 2, 2)
	fmt.Println("view2=" + view2)

	view2.Assign(-1)
	fmt.Println("master replaced" + master)
	fmt.Println("flip1 replaced" + view1)
	fmt.Println("flip2 replaced" + view2)
}

func ExampleViewDice() {
	rows, columns := 4, 5
	master := AscendingMatrix(rows, columns)
	// master.Assign(1) // Set all cells to 1.
	fmt.Println("\n" + master)
	// master.ViewPart(2,0,2,3).Assign(2) // set [2,1] .. [3,3] to 2
	// fmt.Println("\n"+master)

	view1 := master.ViewDice()
	fmt.Println("view1=" + view1)
	view2 := view1.ViewDice()
	fmt.Println("view2=" + view2)

	view2.Assign(-1)
	fmt.Println("master replaced" + master)
	fmt.Println("flip1 replaced" + view1)
	fmt.Println("flip2 replaced" + view2)
}

func ExampleViewRowFlip() {
	rows, columns := 4, 5
	master := AscendingMatrix(rows, columns)
	// master.Assign(1) // Set all cells to 1.
	fmt.Println("\n" + master)
	// master.ViewPart(2,0,2,3).Assign(2); // set [2,1] .. [3,3] to 2
	// fmt.Println("\n"+master)

	view1 := master.ViewRowFlip()
	fmt.Println("view1=" + view1)
	view2 := view1.ViewRowFlip()
	fmt.Println("view2=" + view2)

	view2.Assign(-1)
	fmt.Println("master replaced" + master)
	fmt.Println("flip1 replaced" + view1)
	fmt.Println("flip2 replaced" + view2)
}
