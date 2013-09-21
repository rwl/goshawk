package tfloat64

import common "github.com/rwl/goshawk"

type Cub interface {
	common.Cub

	GetQuick(int, int, int) float64
	SetQuick(int, int, int, float64)
}
