package goshawk

import "math"

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
