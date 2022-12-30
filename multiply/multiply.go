package multiply

import (
	st "gooseberrymat/structure"
)

func BaseMultiply(multiplier, multiplicand *MatrixType) (res *MatrixType) {

	muSlice := []*MatrixType{multiplicand, multiplier}
	for _, mu := range muSlice {
		if mu.Dg != nil {
			mu.Gd = mu.Dg.ToGrid()
		} else if mu.Oll != nil {
			mu.Gd = mu.Oll.ToGrid()
		} else if mu.Qt != nil {
			mu.Gd = mu.Qt.ToGrid()
		} else if mu.Tg != nil {
			mu.Gd = mu.Tg.ToGrid()
		}
	}
	res = &MatrixType{Gd: SimpleMultiply(multiplicand.Gd, multiplier.Gd)}
	return
}

func SimpleMultiply(gd, mul *st.Grid) *st.Grid {
	res := &st.Grid{
		Width:  mul.Width,
		Height: gd.Height,
		Val:    st.BaseMultiply(gd.Val, mul.Val),
	}
	return res
}
