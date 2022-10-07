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

// TODO need reconstrct
func (tg *Trigram) ParseToGrid() *Grid {
	gd := make([][]int, 0)
	for i := 0; i < tg.Height; i++ {
		gd = append(gd, make([]int, tg.Width))
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
