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

// TODO this function need to be done.
func (gd *Grid) ToOrthgonal() *OrthogonalLinkedList {
	orth := &OrthogonalLinkedList{
		Width:   gd.Width,
		Height:  gd.Height,
		NotNull: 0,
		Col:     make([]*OrthogonalLinkedNode, gd.Height),
		Row:     make([]*OrthogonalLinkedNode, gd.Width),
	}

	return orth
}

// As usual, in this package, Constructor is how a grid matrix parse to this data type.
func (gd *Grid) ToQuadTree() *QuadTree {
	var dfs func([][]int, int, int) *QuadTreeNode
	dfs = func(grid [][]int, leftBound, rightBound int) *QuadTreeNode {
		for _, row := range grid {
			for _, v := range row[leftBound:rightBound] {
				if v != grid[0][leftBound] {
					rowMid, colMid := len(grid)/2, (leftBound+rightBound)/2
					return &QuadTreeNode{
						0,
						false,
						dfs(grid[:rowMid], leftBound, colMid),
						dfs(grid[:rowMid], colMid, rightBound),
						dfs(grid[rowMid:], leftBound, colMid),
						dfs(grid[rowMid:], colMid, rightBound),
					}
				}
			}
		}
		return &QuadTreeNode{Val: grid[0][leftBound], IsLeaf: true}
	}
	return &QuadTree{
		Root:   dfs(gd.Val, 0, len(gd.Val)),
		Length: len(gd.Val),
	}
}

// It is a bad idea to use 'Constructor' in Go
func (gd *Grid) ToDiag() *DiagonalMatrix {
	dm := &DiagonalMatrix{
		Length: gd.Height,
		Val:    make([]int, gd.Height),
	}
	for i := 0; i < dm.Length; i++ {
		dm.Val[i] = gd.Val[i][i]
	}
	return dm
}

func (gd *Grid) ToTrigram() *Trigram {
	tg := &Trigram{
		Width:  gd.Width,
		Height: gd.Height,
	}
	for i, line := range gd.Val {
		for j, v := range line {
			tg.Val = append(tg.Val, &TrigramNode{
				Row: i,
				Col: j,
				Val: v,
			})
		}
	}
	tg.Length = len(gd.Val)
	return tg
}

// Though I don't know what to do...
func (gd *Grid) IsToeplitzMatrix() bool {
	for i := 1; i < gd.Height; i++ {
		for j := 1; j < gd.Width; j++ {
			if gd.Val[i][j] != gd.Val[i-1][j-1] {
				return false
			}
		}
	}
	return true
}
