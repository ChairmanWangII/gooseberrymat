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
