package project

import (
	"strconv"
)

const precision = 6

type roundedFloat float64

func roundFloat(val float64) roundedFloat {
	temp := roundedFloat(val)
	out, _ := strconv.ParseFloat(temp.String(), 64)
	return roundedFloat(out)
}

func (f roundedFloat) String() string {
	return strconv.FormatFloat(float64(f), 'f', precision, 64)
}
