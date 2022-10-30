package gooseberrymat

func (gd *Grid) Multiply(mul *Grid) *Grid {
	return gd.SimpleMulti(mul)
}

func (gd *Grid) SimpleMulti(mul *Grid) *Grid {
	if gd.Width != mul.Height {
		return nil
	}
	res := &Grid{
		Width:  mul.Width,
		Height: gd.Height,
		Val:    Init2dSlice(mul.Width, gd.Height),
	}
	operationCount := gd.Width
	for i, line := range res.Val {
		for j := range line {
			count := 0
			for m := 0; m < operationCount; m++ {
				count += gd.Val[i][m] * mul.Val[m][j]
			}
			res.Val[i][j] = count
		}
	}
	return res
}
