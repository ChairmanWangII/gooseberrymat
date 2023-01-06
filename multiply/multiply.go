package multiply

import (
	st "gooseberrymat/structure"
)

func BaseMultiply(multiplier, multiplicand *MatrixType) (res *MatrixType) {
	// It's plenty of tricks if you multiply with a diagonal matrix
	if multiplier.Dg != nil && multiplicand.Dg != nil {
		if multiplier.Dg.Direction != multiplicand.Dg.Direction {
			length := multiplier.Gd.Height
			if length%2 == 0 {
				val := st.Init2dSlice(length, length)
				val[length/2][length/2] = multiplier.Dg.Val[length/2] * multiplicand.Dg.Val[length/2]
				return &MatrixType{
					Gd: &st.Grid{
						Val: val,
					},
				}
			} else {
				return &MatrixType{
					Zero: true,
				}
			}

		}
	}

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
	res = Grid2MatrixType(SimpleMultiply(multiplicand.Gd, multiplier.Gd))
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
