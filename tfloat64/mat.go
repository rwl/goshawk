package tfloat64

import common "github.com/rwl/goshawk"

type Mat interface {
	common.Mat

	GetQuick(int, int) float64
	SetQuick(int, int, float64)

	Like(int, int) Mat
}
