package structure

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Quad tree is a structure used to represent a type of matrix
// whose width equals to height and length is an exponential multiple of two.
// Especially for those matrixes has lots of same data.
type QuadTree struct {
	Root  *QuadTreeNode
	Shape *Shape
}

const (
	BoxVer       = "│"
	BoxHor       = "─"
	BoxVerRight  = "├"
	BoxDownLeft  = "┐"
	BoxDownRight = "┌"
	BoxDownHor   = "┬"
	BoxUpRight   = "└"
	// Gutter is number of spaces between two adjacent child nodes.
	Gutter = 2
)

type QuadTreeNode struct {
	Val         int
	IsLeaf      bool
	TopLeft     *QuadTreeNode
	TopRight    *QuadTreeNode
	BottomLeft  *QuadTreeNode
	BottomRight *QuadTreeNode
}

func (qt *QuadTree) Width() int {
	return qt.Shape.Length
}

func (qt *QuadTree) Height() int {
	return qt.Shape.Length
}

func (qt *QuadTree) ToGrid() *Grid {
	length := qt.Shape.Length
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
	var dfs func(*QuadTreeNode, int, int, int, int)
	dfs = func(n *QuadTreeNode, top, bottom, left, right int) {
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
		Val:   grid,
		Shape: qt.Shape,
	}
}

// The add operation for quad tree is a bit more complex than AND or OR operation.
func (qt *QuadTree) Add(addend *QuadTree) *QuadTree {
	var nodeAdd func(*QuadTreeNode, *QuadTreeNode) *QuadTreeNode
	nodeAdd = func(n1, n2 *QuadTreeNode) *QuadTreeNode {
		if n1.IsLeaf && n2.IsLeaf {
			return &QuadTreeNode{
				Val:         n1.Val + n2.Val,
				IsLeaf:      true,
				TopLeft:     nil,
				TopRight:    nil,
				BottomLeft:  nil,
				BottomRight: nil,
			}
		} else if !(n1.IsLeaf || n2.IsLeaf) {
			return &QuadTreeNode{
				Val:         0,
				IsLeaf:      false,
				TopLeft:     nodeAdd(n1.TopLeft, n2.TopLeft),
				TopRight:    nodeAdd(n1.TopRight, n2.TopRight),
				BottomLeft:  nodeAdd(n1.BottomLeft, n2.BottomLeft),
				BottomRight: nodeAdd(n1.BottomRight, n2.BottomRight),
			}
		} else {
			var leafNode, nonleafNode *QuadTreeNode
			if n1.IsLeaf {
				leafNode = n1
				nonleafNode = n2
			} else {
				leafNode = n2
				nonleafNode = n1
			}
			virtualNode := &QuadTreeNode{
				Val:         leafNode.Val,
				IsLeaf:      true,
				TopLeft:     nil,
				TopRight:    nil,
				BottomLeft:  nil,
				BottomRight: nil,
			}
			return &QuadTreeNode{
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
		Root:  nodeAdd(qt.Root, addend.Root),
		Shape: qt.Shape,
	}
}

// For quad tree, Transpose operation is easier using DFS.
func (qt *QuadTree) Transpose() *QuadTree {
	var dfs func(*QuadTreeNode) *QuadTreeNode
	dfs = func(qtn *QuadTreeNode) *QuadTreeNode {
		qtn.BottomLeft, qtn.TopRight = qtn.TopRight, qtn.BottomLeft
		if !qtn.IsLeaf {
			dfs(qtn.BottomLeft)
			dfs(qtn.BottomRight)
			dfs(qtn.TopLeft)
			dfs(qtn.TopRight)
		}
		return qtn
	}
	return &QuadTree{
		Shape: qt.Shape,
		Root:  dfs(qt.Root),
	}
}

// Return the list of children.
func (qt *QuadTreeNode) Children() []*QuadTreeNode {
	var res []*QuadTreeNode
	leaves := []*QuadTreeNode{qt.TopLeft, qt.TopRight, qt.BottomLeft, qt.BottomRight}
	for _, leaf := range leaves {
		if leaf != nil {
			res = append(res, leaf)
		}
	}
	return res
}

// Returns the horizontal formatted tree.
func (qt *QuadTree) PrettyPrint() string {
	var str string
	for _, line := range split(qt.Root) {
		// Ignore runes before root node.
		line = string([]rune(line)[2:])
		str += strings.TrimRight(line, " ") + "\n"
	}
	return str
}

func split(root *QuadTreeNode) (s []string) {
	data := fmt.Sprintf("%s %v ", BoxHor, root.Val)
	l := len(root.Children())
	if l == 0 {
		s = append(s, data)
		return
	}

	w := utf8.RuneCountInString(data)
	for i, c := range root.Children() {
		for j, line := range split(c) {
			if i == 0 && j == 0 {
				if l == 1 {
					s = append(s, data+BoxHor+line)
				} else {
					s = append(s, data+BoxDownHor+line)
				}
				continue
			}

			var box string
			if i == l-1 && j == 0 {
				// first line of the last child
				box = BoxUpRight
			} else if i == l-1 {
				box = " "
			} else if j == 0 {
				box = BoxVerRight
			} else {
				box = BoxVer
			}
			s = append(s, strings.Repeat(" ", w)+box+line)
		}
	}
	return
}
