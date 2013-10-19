package tfloat64

import "github.com/rwl/goshawk/common"

type Mat interface {
	common.Mat

	GetQuick(int, int) float64
	SetQuick(int, int, float64)

	Like(int, int) Mat
	LikeVector(size int) Vec
}
