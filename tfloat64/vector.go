
package tfloat64

import "bitbucket.org/rwl/colt"

type Vector struct {
	colt.VectorData
}

func newVector(impl colt.VectorData) Vector {
	return Vector{impl}
}
