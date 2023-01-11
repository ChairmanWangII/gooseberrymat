package multiply

import (
	st "gooseberrymat/structure"
)

func BaseMultiply(multiplier, multiplicand *MatrixType) *MatrixType {
	// If any of the matrix have error, return all the error.
	if len(multiplier.Err) != 0 && len(multiplicand.Err) != 0 {
		return &MatrixType{
			Err: append(multiplier.Err, multiplicand.Err...),
		}
	}

	// It's plenty of tricks if you multiply with a diagonal matrix.
	if multiplier.Dg != nil && multiplicand.Dg != nil {
		// Two diagonal matrixes have different direction.
		if multiplier.Dg.Direction != multiplicand.Dg.Direction {
			length := multiplier.Dg.Length
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
			// Two diagonal matrixes have same direction.
		} else {
			res := make([]int, multiplier.Dg.Length)
			for i := range res {
				res[i] = multiplier.Dg.Val[i] * multiplicand.Dg.Val[i]
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
	res := Grid2MatrixType(SimpleMultiply(multiplicand.Gd, multiplier.Gd))
	return res
}

func SimpleMultiply(gd, mul *st.Grid) *st.Grid {
	res := &st.Grid{
		Width:  mul.Width,
		Height: gd.Height,
		Val:    st.BaseMultiply(gd.Val, mul.Val),
	}
	return res
}
