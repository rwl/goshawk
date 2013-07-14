
package tfloat64

import l4g "code.google.com/p/log4go"

type Vector struct {
	VectorData
}

func newVector(impl VectorData) Vector {
	return Vector{impl}
}

func (v *Vector) checkSize(other VectorData) error {
	if v.Size() != other.Size() {
//		formatter := NewFormatter()
//		return l4g.Error("Incompatible sizes: %s and %s",
//			formatter.VectorShape(v), formatter.VectorShape(other))
		return l4g.Error("Incompatible sizes: %d and %d",
			v.Size(), other.Size())
	}
	return nil
}

func (v *Vector) Copy() *Vector {
	copy := &Vector{v.Like(v.Size())}
	copy.AssignVector(v)
	return copy
}
