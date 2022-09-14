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

func (qt *QuadTree) Constructor(matrix *Grid) *QuadTree {
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
		Root:   dfs(matrix.Val, 0, len(matrix.Val)),
		Length: len(matrix.Val),
	}

}

func (qt *QuadTree) Width() int {
	return qt.Length
}

func (qt *QuadTree) Height() int {
	return qt.Length
}

func (qt *QuadTree) ParseToGrid() *Grid {
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
	return &Grid{
		Val:    grid,
		Height: length,
		Width:  length,
	}
}

func (qt *QuadTree) Add(addend *QuadTree) *QuadTree {
	var nodeAdd func(*Node, *Node) *Node
	nodeAdd = func(n1, n2 *Node) *Node {
		if n1.IsLeaf && n2.IsLeaf {
			return &Node{
				Val:         n1.Val + n2.Val,
				IsLeaf:      true,
				TopLeft:     nil,
				TopRight:    nil,
				BottomLeft:  nil,
				BottomRight: nil,
			}
		} else if !(n1.IsLeaf || n2.IsLeaf) {
			return &Node{
				Val:         0,
				IsLeaf:      false,
				TopLeft:     nodeAdd(n1.TopLeft, n2.TopLeft),
				TopRight:    nodeAdd(n1.TopRight, n2.TopRight),
				BottomLeft:  nodeAdd(n1.BottomLeft, n2.BottomLeft),
				BottomRight: nodeAdd(n1.BottomRight, n2.BottomRight),
			}
		} else {
			var leafNode, nonleafNode *Node
			if n1.IsLeaf {
				leafNode = n1
				nonleafNode = n2
			} else {
				leafNode = n2
				nonleafNode = n1
			}
			virtualNode := &Node{
				Val:         leafNode.Val,
				IsLeaf:      true,
				TopLeft:     nil,
				TopRight:    nil,
				BottomLeft:  nil,
				BottomRight: nil,
			}
			return &Node{
				Val:         0,
				IsLeaf:      false,
				TopLeft:     nodeAdd(nonleafNode.TopLeft, virtualNode),
				TopRight:    nodeAdd(nonleafNode.TopRight, virtualNode),
				BottomLeft:  nodeAdd(nonleafNode.BottomLeft, virtualNode),
				BottomRight: nodeAdd(nonleafNode.BottomRight, virtualNode),
			}
		}
	}
	return &QuadTree{
		Root:   nodeAdd(qt.Root, addend.Root),
		Length: qt.Length,
	}
}
