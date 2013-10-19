package common

import (
	"math"
)

const (
	VectorThreshold = 32768
	MatrixThreshold = 65536
	CubeThreshold = 65536
)

func ContainsInt(a []int, x int) bool {
	contains := false
	for _, val := range a {
		if val == x {
			contains = true
			break
		}
	}
	return contains
}

func ContainsFloat(a []float64, x, prec float64) bool {
	contains := false
	for _, val := range a {
		diff := math.Abs(val - x)
		if diff <= prec {
			contains = true
			break
		}
	}
	return contains
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
