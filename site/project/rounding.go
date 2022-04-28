package project

import "math"

const precision = 1e-6

type roundedFloat float64

func roundFloat(val float64) roundedFloat {
	return roundedFloat(math.Round(val/precision) * precision)
}
