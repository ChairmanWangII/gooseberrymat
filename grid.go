package gooseberrymat

type Grid struct {
	Width  int
	Height int
	Val    [][]int
}

func (gd *Grid) IsDiagonal() bool {
	if gd.Height != gd.Width {
		return false
	}
	for i, line := range gd.Val {
		for j, v := range line {
			if i != j && v != 0 {
				return false
			}
		}
	}
	return true
}

func (gd *Grid) IsQuadrable() bool {
	if gd.Width != gd.Height {
		return false
	}
	if gd.Width^(gd.Width-1) != 0 {
		return false
	}
	return true
}

func (gd *Grid) Transpose() *Grid {
	length := gd.Height
	for i := range gd.Val {
		for j := i + 1; j < length; j++ {
			gd.Val[i][j], gd.Val[j][i] = gd.Val[j][i], gd.Val[i][j]
		}
	}
	return gd
}

// Judge if two matrixes are cophenetic.
func (gd *Grid) IsCophenetic(grid *Grid) bool {
	if gd.Height == grid.Height && gd.Width == grid.Width {
		return true
	}
	return false
}

func (gd *Grid) Add(addend *Grid) *Grid {
	res := &Grid{
		Width:  gd.Width,
		Height: gd.Height,
		Val:    gd.Val,
	}
	for i := 0; i < gd.Height; i++ {
		for j := 0; j < gd.Width; j++ {
			res.Val[i][j] += addend.Val[i][j]
		}
	}
	return res
}
