package tfloat64

import "github.com/rwl/goshawk/common"

func NewVector(size int) *Vector {
	return &Vector{
		&DenseVec{
			common.NewCoreVec(false, size, 0, 1),
			make([]float64, size),
		},
	}
}

func NewVectorArray(a []float64) *Vector {
	v := NewVector(len(a))
	v.AssignArray(a)
	return v
}

//C = A||B; Constructs a new matrix which is the concatenation of two other
// matrices. Example: 0 1 append 3 4 --> 0 1 3 4.
func AppendVectors(A, B Vec) *Vector {
	v := NewVector(A.Size() + B.Size())
	v.ViewPart(0, A.Size()).AssignVector(A)
	v.ViewPart(A.Size(), B.Size()).AssignVector(B)
	return v
}

// Constructs a matrix with cells having ascending values. For debugging
// purposes. Example: 0 1 2
func Ascending(size int) *Vector {
	vector := NewVector(size)
	v := 0.0
	for i := 0; i < size; i++ {
		vector.SetQuick(i, v)
		v += 1.0
	}
	return vector
}

// Constructs a matrix with cells having descending values. For debugging
// purposes. Example: 2 1 0
func Decending(size int) *Vector {
	vector := NewVector(size)
	v := 0.0
	for i := size - 1; i >= 0; i-- {
		vector.SetQuick(i, v)
		v += 1.0
	}
	return vector
}

// Constructs a matrix which is the concatenation of all given parts. Cells
// are copied.
func NewVectorParts(parts []Vec) *Vector {
	if len(parts) == 0 {
		return NewVector(0)
	}

	size := 0
	for i := 0; i < len(parts); i++ {
		size += parts[i].Size()
	}

	vector := NewVector(size)
	size = 0
	for _, part := range parts {
		vector.ViewPart(size, part.Size()).AssignVector(part)
		size += part.Size()
	}

	return vector
}

// Constructs a matrix with the given shape, each cell initialized with the
// given value.
func NewVectorInitial(size int, initial float64) *Vector {
	return NewVector(size).Assign(initial)
}

// Constructs a matrix with uniformly distributed values in (0,1)
// (exclusive).
func NewRandomVector(size int) *Vector {
	v := NewVector(size)
	v.AssignFunc(Random())
	return v
}

// C = A||A||..||A; Constructs a new matrix which is concatenated
// 'repeat' times. Example:
// 	 0 1
// 	 repeat(3) -->
// 	 0 1 0 1 0 1
func RepeatVector(A Vec, repeat int) *Vector {
	size := A.Size()
	v := NewVector(repeat*size)
	for i := repeat - 1; i >= 0; i-- {
		v.ViewPart(size*i, size).AssignVector(A)
	}
	return v
}

// Constructs a matrix with the given shape, each cell initialized with
// zero.
func Zeros(size int) *Vector {
	return NewVector(size)
}

// Constructs a matrix with the given shape, each cell initialized with
// one.
func Ones(size int) *Vector {
	return NewVector(size).Assign(1.0)
}
