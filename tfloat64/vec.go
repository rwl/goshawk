package tfloat64

import "github.com/rwl/goshawk/common"

// Interface for all vector backends.
type Vec interface {
	common.Vec

	// Returns the matrix cell value at coordinate "index".
	//
	// Provided with invalid parameters this method may cause a panic or
	// return invalid values without causing an error. You should only
	// use this method when you are absolutely sure that the coordinate
	// is within bounds.
	// Precondition (unchecked): index < 0 || index >= Size().
	GetQuick(int) float64

	// Sets the matrix cell at coordinate "index" to the specified value.
	//
	// Provided with invalid parameters this method may cause a panic or
	// access illegal indexes without causing an error. You should only use
	// this method when you are absolutely sure that the coordinate is
	// within bounds.
	// Precondition (unchecked): index < 0 || index >= Size().
	SetQuick(int, float64)

	//Like() Vec
	Like(int) Vec
	LikeMatrix(int, int) Mat

	ReshapeMatrix(int, int) (*Matrix, error)
	ReshapeCube(int, int, int) (*Cube, error)

	// Construct and returns a new selection view using the offsets of
	// the visible elements.
	ViewSelectionLike(offsets []int) Vec
	ViewVec() Vec
}
