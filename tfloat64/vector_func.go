package tfloat64

func (v *Vector) Square() *Vector {
	return v.AssignFunc(Square)
}

func (v *Vector) Mult(y Vec) (*Vector, error) {
	return v.AssignVectorFunc(y, Mult)
}

func (v *Vector) Inv() *Vector {
	return v.AssignFunc(Inv)
}

func (v *Vector) Neg() *Vector {
	return v.AssignFunc(Square)
}

func (v *Vector) Sign() *Vector {
	return v.AssignFunc(Sign)
}

func (v *Vector) Compare(y Vec) *Vector {
	x, _ := v.AssignVectorFunc(y, Compare)
	return x
}

func (v *Vector) Div(y Vec) *Vector {
	x, _ := v.AssignVectorFunc(y, Div)
	return x
}

func (v *Vector) DivNeg(y Vec) *Vector {
	x, _ := v.AssignVectorFunc(y, DivNeg)
	return x
}

/*func (v *Vector) Equals(y Vec) *Vector {
	x, _ := v.AssignVectorFunc(y, Equals)
	return x
}*/

func (v *Vector) Greater(y Vec) *Vector {
	x, _ := v.AssignVectorFunc(y, Greater)
	return x
}

/*func (v *Vector) IsEqual(y Vec) *Vector {
	x, _ := v.AssignVectorFunc(y, IsEqual)
	return x
}

func (v *Vector) IsLess(y Vec) *Vector {
	x, _ := v.AssignVectorFunc(y, IsLess)
	return x
}

func (v *Vector) IsGreater(y Vec) *Vector {
	x, _ := v.AssignVectorFunc(y, IsGreater)
	return x
}*/

func (v *Vector) Less(y Vec) *Vector {
	x, _ := v.AssignVectorFunc(y, Less)
	return x
}

func (v *Vector) Lg(y Vec) *Vector {
	x, _ := v.AssignVectorFunc(y, Lg)
	return x
}

func (v *Vector) Minus(y Vec) *Vector {
	x, _ := v.AssignVectorFunc(y, Minus)
	return x
}

func (v *Vector) MultNeg(y Vec) *Vector {
	x, _ := v.AssignVectorFunc(y, MultNeg)
	return x
}

func (v *Vector) MultSquare(y Vec) *Vector {
	x, _ := v.AssignVectorFunc(y, MultSquare)
	return x
}

func (v *Vector) Plus(y Vec) *Vector {
	x, _ := v.AssignVectorFunc(y, Plus)
	return x
}

func (v *Vector) PlusAbs(y Vec) *Vector {
	x, _ := v.AssignVectorFunc(y, PlusAbs)
	return x
}

/*func (v *Vector) Between(y Vec) *Vector {
	x, _ := v.AssignVectorFunc(y, Between)
	return x
}*/

func (v *Vector) CompareTo(y float64) *Vector {
	return v.AssignFunc(CompareTo(y))
}

func (v *Vector) Divide(y float64) *Vector {
	return v.AssignFunc(Divide(y))
}

/*func (v *Vector) EqualTo(y float64) *Vector {
	return v.AssignFunc(EqualTo(y))
}*/

func (v *Vector) GreaterThan(y float64) *Vector {
	return v.AssignFunc(GreaterThan(y))
}

func (v *Vector) LessThan(y float64) *Vector {
	return v.AssignFunc(LessThan(y))
}

func (v *Vector) Remainder(y float64) *Vector {
	return v.AssignFunc(Remainder(y))
}

func (v *Vector) LgVal(y float64) *Vector {
	return v.AssignFunc(LgVal(y))
}

func (v *Vector) Max(y float64) *Vector {
	return v.AssignFunc(Max(y))
}

func (v *Vector) Min(y float64) *Vector {
	return v.AssignFunc(Min(y))
}

func (v *Vector) Subtract(y float64) *Vector {
	return v.AssignFunc(Subtract(y))
}

func (v *Vector) Mod(y float64) *Vector {
	return v.AssignFunc(Mod(y))
}

func (v *Vector) Multiply(y float64) *Vector {
	return v.AssignFunc(Multiply(y))
}

func (v *Vector) Add(y float64) *Vector {
	return v.AssignFunc(Add(y))
}

func (v *Vector) Pow(y float64) *Vector {
	return v.AssignFunc(Pow(y))
}

func (v *Vector) Random() *Vector {
	return v.AssignFunc(Random())
}

func (v *Vector) IsBetween(y, z, value float64) *Vector {
	return v.AssignProcedure(IsBetween(y, z), value)
}

func (v *Vector) IsEqualTo(y, value float64) *Vector {
	return v.AssignProcedure(IsEqualTo(y), value)
}

func (v *Vector) IsGreaterThan(y, value float64) *Vector {
	return v.AssignProcedure(IsGreaterThan(y), value)
}

func (v *Vector) IsLessThan(y, value float64) *Vector {
	return v.AssignProcedure(IsLessThan(y), value)
}

func (v *Vector) MinusMult(other Vec, y float64) *Vector {
	x, _ := v.AssignVectorFunc(other, MinusMult(y))
	return x
}

func (v *Vector) MultSecond(other Vec, y float64) *Vector {
	x, _ := v.AssignVectorFunc(other, MultSecond(y))
	return x
}

func (v *Vector) PlusMultSecond(other Vec, y float64) *Vector {
	x, _ := v.AssignVectorFunc(other, PlusMultSecond(y))
	return x
}

func (v *Vector) PlusMultFirst(other Vec, y float64) *Vector {
	x, _ := v.AssignVectorFunc(other, PlusMultFirst(y))
	return x
}

