package structure

import (
	"fmt"
	"gooseberrymat/utils"
)

type Grid struct {
	Shape *Shape
	Val   [][]int
}

func (gd *Grid) Add(addend *Grid) *Grid {
	res := &Grid{
		Shape: gd.Shape,
		Val:   gd.Val,
	}
	for i := 0; i < gd.Shape.Height; i++ {
		for j := 0; j < gd.Shape.Length; j++ {
			res.Val[i][j] += addend.Val[i][j]
		}
	}
	return res
}

func (gd *Grid) Transpose() *Grid {
	height, length := gd.Shape.Height, gd.Shape.Length
	tGrid := make([][]int, length)
	for i := range tGrid {
		tGrid[i] = make([]int, height)
		for j := range tGrid[i] {
			tGrid[i][j] = -1
		}
	}
	for i, row := range gd.Val {
		for j, v := range row {
			tGrid[j][i] = v
		}
	}
	return &Grid{
		Shape: gd.Shape,
		Val:   tGrid,
	}
}

// Though I don't know what to do...
func (gd *Grid) IsToeplitzMatrix() bool {
	for i := 1; i < gd.Shape.Height; i++ {
		for j := 1; j < gd.Shape.Length; j++ {
			if gd.Val[i][j] != gd.Val[i-1][j-1] {
				return false
			}
		}
	}
	return true
}

func (gd *Grid) IsDiagonal() bool {
	if gd.Shape.Height != gd.Shape.Length {
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
	if gd.Shape.Height != gd.Shape.Length {
		return false
	}
	if gd.Shape.Length^(gd.Shape.Length-1) != 0 {
		return false
	}
	return true
}

// Judge if two matrixes are cophenetic.
func (gd *Grid) IsCophenetic(grid *Grid) bool {
	if gd.Shape.Height == grid.Shape.Height && gd.Shape.Length == grid.Shape.Length {
		return true
	}
	return false
}

func (gd *Grid) ToOrthgonal() *OrthogonalLinkedList {
	orth := &OrthogonalLinkedList{
		Shape:      gd.Shape,
		NotNull:    0,
		Horizontal: make([]*OrthogonalLinkedNode, gd.Shape.Height),
		Vertical:   make([]*OrthogonalLinkedNode, gd.Shape.Length),
	}

	downList := make([]*OrthogonalLinkedNode, orth.Shape.Length)
	for i, line := range gd.Val {
		head := orth.Horizontal[i]
		for j := range line {
			if gd.Val[i][j] != 0 {
				node := &OrthogonalLinkedNode{
					Col:       i,
					Row:       j,
					Val:       gd.Val[i][j],
					DownNode:  nil,
					RightNode: nil,
				}
				if head == nil {
					head = node
				} else {
					head.RightNode = node
					head = head.RightNode
				}
				if orth.Vertical[j] == nil {
					orth.Vertical[j] = node
					downList[j] = node
				} else {
					downList[j].DownNode = node
					downList[j] = downList[j].DownNode
				}
			}

		}
	}
	return orth
}

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
		Root:  dfs(gd.Val, 0, len(gd.Val)),
		Shape: gd.Shape,
	}
}

// It is a bad idea to use 'Constructor' in Go
func (gd *Grid) ToDiag() *DiagonalMatrix {
	dm := &DiagonalMatrix{
		Shape: gd.Shape,
		Val:   make([]int, gd.Shape.Height),
	}
	for i := 0; i < dm.Shape.Length; i++ {
		dm.Val[i] = gd.Val[i][i]
	}
	return dm
}

func (gd *Grid) ToTrigram() *Trigram {
	tg := &Trigram{
		Shape: gd.Shape,
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
	return tg
}

// Print the grid.
// Automatically print 0 as null and align the different digit length.
func (gd *Grid) print() string {
	widthList := make([]int, gd.Shape.Length)
	for _, line := range gd.Val {
		for j, v := range line {
			if utils.GetDigitLength(v) > widthList[j] {
				widthList[j] = utils.GetDigitLength(v)
			}
		}
	}
	s := ""
	for i := range gd.Val {
		for j, n := range gd.Val[i] {
			if n == 0 {
				s += " "
			} else {
				s += fmt.Sprintf("%d", n)
			}
			blanks := widthList[j] - utils.GetDigitLength(n) + 1
			for blanks > 0 {
				s += " "
				blanks--
			}
		}
		s += "\n"
	}
	return s
}
