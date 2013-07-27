
package tfloat64

// Constructs and returns a new flip view. What used to be index
// 0 is now index size()-1, ..., what used to be index
// size()-1 is now index 0. The returned view is backed by
// this matrix, so changes in the returned view are reflected in this
// matrix, and vice-versa.
func (v *Vector) ViewFlip() *Vector {
	view := v.ViewVector()
	view.VectorFlip()
	return view
}
