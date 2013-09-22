package tfloat64

import "github.com/rwl/goshawk/common"

type Cub interface {
	common.Cub

	GetQuick(int, int, int) float64
	SetQuick(int, int, int, float64)
}
