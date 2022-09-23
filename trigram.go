package gooseberrymat

type Trigram struct {
	Width  int
	Height int
	Val    []*TrigramNode
}

type TrigramNode struct {
	Row int
	Col int
	Val int
}

func (tg *Trigram) Constructor(gd *Grid) *Trigram {
	for i, line := range gd.Val {
		for j, v := range line {
			tg.Val = append(tg.Val, &TrigramNode{
				Row: i,
				Col: j,
				Val: v,
			})
		}
	}
	return tg
}

// TODO need reconstrct
func (tg *Trigram) ParseToGrid() *Grid {
	line := make([]int, tg.Width)
	gd := make([][]int, 0)
	for i := 0; i < tg.Height; i++ {
		gd = append(gd, line)
	}
	grid := &Grid{
		Width:  tg.Width,
		Height: tg.Height,
		Val:    gd,
	}
	for _, v := range tg.Val {
		grid.Val[v.Row][v.Col] = v.Val
	}
	return grid
}

func (tg *Trigram) Transpose() *Trigram {
	for _, p := range tg.Val {
		p.Col, p.Row = p.Row, p.Col
	}
	return tg
}
