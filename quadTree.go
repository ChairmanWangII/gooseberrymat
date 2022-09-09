package gooseberrymat

type QuadTree struct {
	Root   *Node
	Length int
}

type Node struct {
	Val         int
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func (qt *QuadTree) Constructor(matrix [][]int) *QuadTree {
	var dfs func([][]int, int, int) *Node
	dfs = func(grid [][]int, leftBound, rightBound int) *Node {
		for _, row := range grid {
			for _, v := range row[leftBound:rightBound] {
				if v != grid[0][leftBound] {
					rowMid, colMid := len(grid)/2, (leftBound+rightBound)/2
					return &Node{
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
		return &Node{Val: grid[0][leftBound], IsLeaf: true}
	}
	return &QuadTree{
		Root:   dfs(matrix, 0, len(matrix)),
		Length: len(matrix),
	}

}

func (qt *QuadTree) Width() int {
	return qt.Length
}

func (qt *QuadTree) Height() int {
	return qt.Length
}

func (qt *QuadTree) ParseToGrid() [][]int {
	length := qt.Length
	grid := make([][]int, length)
	for i := range grid {
		grid[i] = make([]int, length)
	}

	fill := func(top, bottom, left, right, val int) {
		for i := top; i <= bottom; i++ {
			for j := left; j <= right; j++ {
				grid[i][j] = val
			}
		}
	}
	var dfs func(*Node, int, int, int, int)
	dfs = func(n *Node, top, bottom, left, right int) {
		if n == nil {
			return
		} else if n.IsLeaf {
			fill(top, bottom, left, right, n.Val)
			return
		} else {
			dfs(n.BottomLeft, (top+bottom)/2+1, bottom, left, (left+right)/2)
			dfs(n.BottomRight, (top+bottom)/2+1, bottom, (left+right)/2+1, right)
			dfs(n.TopLeft, top, (top+bottom)/2, left, (left+right)/2)
			dfs(n.TopRight, top, (top+bottom)/2, (left+right)/2+1, right)
		}
	}
	dfs(qt.Root, 0, length-1, 0, length-1)
	return grid
}
