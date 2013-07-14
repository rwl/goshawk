package tfloat64

import (
	"math"
	"math/rand"
)

type Float64Func func(float64) float64

type Float64Float64Func func(float64, float64) float64

type Float64Procedure func(float64) bool

type Float64Float64Procedure func(float64, float64) float64

type IntIntFloat64Func func(int, int, float64) float64

type VectorProcedure func(VectorData) bool

// Function that returns a * a.
func Square(a float64) float64 {
	return a * a
}

// Function that returns a * b.
func Mult(a, b float64) float64 {
	return a * b
}

// Function that returns its argument.
func Identity(a float64) float64 {
	return a
}

// Function that returns 1.0 / a.
func Inv(a float64) float64 {
	return 1.0 / a
}

// Function that returns -a.
func Neg(a float64) float64 {
	return -a
}

// Function that returns a rounded to the nearest whole number.
/*func Rint(a float64) float64 {
	return math.Rint(a)
}*/

// Function that returns a < 0 ? -1 : a > 0 ? 1 : 0.
func Sign(a float64) float64 {
	if a < 0 {
		return -1.0
	} else if a > 0 {
		return 1.0
	} else {
		return 0.0
	}
}

// Function that returns a < b ? -1 : a > b ? 1 : 0.
func Compare(a, b float64) float64 {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}
}

// Function that returns a / b.
func Div(a, b float64) float64 {
	return a / b
}

// Function that returns -(a / b).
func DivNeg(a, b float64) float64 {
	return -(a / b)
}

// Function that returns a == b ? 1 : 0.
func Equals(a, b float64) float64 {
	if a == b {
		return 1
	} else {
		return 0
	}
}

// Function that returns a > b ? 1 : 0.
func Greater(a, b float64) float64 {
	if a > b {
		return 1
	} else {
		return 0
	}
}

// Function that returns a == b.
func IsEqual(a, b float64) bool {
	return a == b
}

// Function that returns a < b.
func IsLess(a, b float64) bool {
	return a < b
}

// Function that returns a > b.
func IsGreater(a, b float64) bool {
	return a > b
}

// Function that returns a < b ? 1 : 0.
func Less(a, b float64) float64 {
	if a < b {
		return 1
	} else {
		return 0
	}
}

// Function that returns math.Log(a) / math.Log(b).
func Lg(a, b float64) float64 {
	return math.Log(a) / math.Log(b)
}

// Function that returns a - b.
func Minus(a, b float64) float64 {
	return a - b
}

// Function that returns -(a * b).
func MultNeg(a, b float64) float64 {
	return -(a * b)
}

// Function that returns a * b^2.
func MultSquare(a, b float64) float64 {
	return a * b * b
}

// Function that returns a + b.
func Plus(a, b float64) float64 {
	return a + b
}

// Function that returns math.Abs(a) + math.Abs(b).
func PlusAbs(a, b float64) float64 {
	return math.Abs(a) + math.Abs(b)
}

// Constructs a function that returns (from<=a && a<=to) ? 1 : 0.
// "a" is a variable, "from" and "to" are fixed.
func Between(from, to float64) Float64Func {
	return func(a float64) float64 {
		if from <= a && a <= to {
			return 1
		} else {
			return 0
		}
	}
}

// Constructs a unary function from a binary function with the first
// operand (argument) fixed to the given constant "c". The second
// operand is variable (free).
func BindArg1(f Float64Float64Func, c float64) Float64Func {
	return func (v float64) float64 {
		return f(c, v)
	}
}

// Constructs a unary function from a binary function with the second
// operand (argument) fixed to the given constant "c". The first
// operand is variable (free).
func BindArg2(f Float64Float64Func, c float64) Float64Func {
	return func(v float64) float64 {
		return f(v, c)
	}
}

// Constructs the function f( g(a), h(b) ).
func Chain(f Float64Float64Func, g, h Float64Func) Float64Float64Func {
	return func(a, b float64) float64 {
		return f(g(a), h(b))
	}
}

// Constructs the function g( h(a,b) ).
func ChainBinary(g Float64Func, h Float64Float64Func) Float64Float64Func {
	return func(a, b float64) float64 {
		return g(h(a, b))
	}
}

// Constructs the function g( h(a) ).
func ChainUnary(g, h Float64Func) Float64Func {
	return func(a float64) float64 {
		return g(h(a))
	}
}

// Constructs a function that returns a < b ? -1 : a > b ? 1 : 0.
// a is a variable, b is fixed.
func CompareTo(b float64) Float64Func {
	return func(a float64) float64 {
		if a < b {
			return -1
		} else if a > b {
			return 1
		} else {
			return 0
		}
	}
}

// Constructs a function that returns the constant c.
func Constant(c float64) Float64Func {
	return func(_ float64) float64 {
		return c
	}
}

// Constructs a function that returns a / b. a is a
// variable, b is fixed.
func Divide(b float64) Float64Func {
	return Multiply(1 / b)
}

// Constructs a function that returns a == b ? 1 : 0. a is
// a variable, b is fixed.
func EqualTo(b float64) Float64Func {
	return func(a float64) float64 {
		if a == b {
			return 1
		} else {
			return 0
		}
	}
}

// Constructs a function that returns a > b ? 1 : 0. a is
// a variable, b is fixed.
func GreaterThan(b float64) Float64Func {
	return func(a float64) float64 {
		if a > b {
			return 1
		} else {
			return 0
		}
	}
}

// Constructs a function that returns math.Remainder(a, b).
// a is a variable, b is fixed.
func Remainder(b float64) Float64Func {
	return func(a float64) float64 {
		return math.Remainder(a, b)
	}
}

// Constructs a function that returns from<=a && a<=to. a
// is a variable, from and to are fixed.
func IsBetween(from, to float64) Float64Procedure {
	return func(a float64) bool {
		return from <= a && a <= to
	}
}

// Constructs a function that returns a == b. a is a
// variable, b is fixed.
func IsEqualTo(b float64) Float64Procedure {
	return func(a float64) bool {
		return a == b
	}
}

// Constructs a function that returns a > b. a is a
// variable, b is fixed.
func IsGreaterThan(b float64) Float64Procedure {
	return func(a float64) bool {
		return a > b
	}
}

// Constructs a function that returns a < b. a is a
// variable, b is fixed.
func IsLessThan(b float64) Float64Procedure {
	return func(a float64) bool {
		return a < b
	}
}

// Constructs a function that returns a < b ? 1 : 0. a is
// a variable, b is fixed.
func LessThan(b float64) Float64Func {
	return func(a float64) float64 {
		if a < b {
			return 1
		} else {
			return 0
		}
	}
}

// Constructs a function that returns math.Log(a) / math.Log(b).
// a is a variable, b is fixed.
func LgVal(b float64) Float64Func {
	logInv := 1 / math.Log(b)
	return func(a float64) float64 {
		return math.Log(a) * logInv
	}
}

// Constructs a function that returns math.Max(a,b). a is
// a variable, b is fixed.
func Max(b float64) Float64Func {
	return func(a float64) float64 {
		return math.Max(a, b)
	}
}

// Constructs a function that returns math.Min(a, b). a is
// a variable, b is fixed.
func Min(b float64) Float64Func {
	return func(a float64) float64 {
		return math.Min(a, b)
	}
}

// Constructs a function that returns a - b. a is a
// variable, b is fixed.
func Subtract(b float64) Float64Func {
	return Add(-b)
}

// Constructs a function that returns a - b*constant. a
// and b are variables, constant is fixed.
func MinusMult(constant float64) Float64Float64Func {
	return PlusMultSecond(-constant)
}

// Constructs a function that returns a % b. a is a
// variable, b is fixed.
func Mod(b float64) Float64Func {
	return func(a float64) float64 {
		return math.Mod(a, b)
	}
}

// Constructs a function that returns a * b. a is a
// variable, b is fixed.
func Multiply(b float64) Float64Func {
	return func(a float64) float64 {
		return a * b
	}
}

// Constructs a function that returns a + b. a is a
// variable, b is fixed.
func Add(b float64) Float64Func {
	return func(a float64) float64 {
		return a + b
	}
}

// Constructs a function that returns b*constant.
func MultSecond(constant float64) Float64Float64Func {
	return func(_, b float64) float64 {
		return b * constant
	}
}

// Constructs a function that returns a + b*constant. a
// and b are variables, constant is fixed.
func PlusMultSecond(constant float64) Float64Float64Func {
	return func(a, b float64) float64 {
		return a + b * constant
	}
}

// Constructs a function that returns a*constant + b. a
// and b are variables, constant is fixed.
func PlusMultFirst(constant float64) Float64Float64Func {
	return func(a, b float64) float64 {
		return a * constant + b
	}
}

// Constructs a function that returns math.Pow(a, b). a is
// a variable, b is fixed.
func Pow(b float64) Float64Func {
	return func(a float64) float64 {
		return math.Pow(a, b)
	}
}

// Constructs a function that returns a new uniform random number in the
// open unit interval (0.0, 1.0) (excluding 0.0 and 1.0).
func Random() Float64Func {
	return func(_ float64) float64 {
		return rand.Float64()
	}
}

// Constructs a function that returns the number rounded to the given
// precision; Examples:
//
//	precision = 0.01 rounds 0.012 --> 0.01, 0.018 --> 0.02
//	precision = 10   rounds 123   --> 120 , 127   --> 130
/*func Round(precision float64) Float64Func {
	return func(a float64) float64 {
		return math.Rint(a / precision) * precision
	}
}*/

// Constructs a function that returns f(b, a), i.e.
// applies the function with the first operand as second operand and the
// second operand as first operand.
func swapArgs(f Float64Float64Func) Float64Float64Func {
	return func(a, b float64) float64 {
		return f(b, a)
	}
}
