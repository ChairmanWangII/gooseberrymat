package multiply

import "gooseberrymat/structure"

func Multiply(gd, mul *structure.Grid) *structure.Grid {
	return SimpleMultiply(gd, mul)
}

func SimpleMultiply(gd, mul *structure.Grid) *structure.Grid {
	res := &structure.Grid{
		Width:  mul.Width,
		Height: gd.Height,
		Val:    structure.BaseMultiply(gd.Val, mul.Val),
	}
	return res
}
