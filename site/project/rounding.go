package project

import (
	"strconv"
)

const precision = 6

type roundedFloat float64

type roundedFloatWithStr struct {
	val roundedFloat
	str string
}

func roundFloat(raw float64) roundedFloatWithStr {
	str := strconv.FormatFloat(float64(raw), 'f', precision, 64)
	val, _ := strconv.ParseFloat(str, 64)
	return roundedFloatWithStr{val: roundedFloat(val), str: str}
}

func (self roundedFloatWithStr) Lt(other roundedFloatWithStr) bool {
	return self.val < other.val
}

func (self roundedFloatWithStr) Gt(other roundedFloatWithStr) bool {
	return self.val > other.val
}
