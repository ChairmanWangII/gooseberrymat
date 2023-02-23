package structure

// Orthogonal linked list is a data structure
// based on linked list to store sparse matrix.

import (
	"gooseberrymat/utils"
)

// Oll use two slices as the head of linked lists.
type OrthogonalLinkedList struct {
	Shape      *Shape
	NotNull    int
	Horizontal []*OrthogonalLinkedNode
	Vertical   []*OrthogonalLinkedNode
}

type OrthogonalLinkedNode struct {
	Col       int
	Row       int
	Val       int
	DownNode  *OrthogonalLinkedNode
	RightNode *OrthogonalLinkedNode
}

func (ol *OrthogonalLinkedList) ToGrid() *Grid {
	matrix := Init2dSlice(ol.Shape.Length, ol.Shape.Height)
	for _, line := range ol.Vertical {
		for line != nil {
			if line.Val != 0 {
				matrix[line.Row][line.Col] = line.Val
			}
			line = line.RightNode
		}
	}
	return &Grid{
		Val:   matrix,
		Shape: ol.Shape,
	}
}

// Append new element to orthogonal linked list.
func (ol *OrthogonalLinkedList) Append(tn *TrigramNode) {
	oln := &OrthogonalLinkedNode{
		Val:       tn.Val,
		Col:       tn.Col,
		Row:       tn.Row,
		DownNode:  nil,
		RightNode: nil,
	}
	// If unadded node is the only node.
	if ol.Horizontal[tn.Col] == nil {
		ol.Horizontal[tn.Col] = oln
	} else {
		head := ol.Horizontal[tn.Col]
		if head.Val > tn.Val {
			ol.Horizontal[tn.Col] = oln
			oln.RightNode = head
		}
		for head != nil {
			if head.RightNode == nil {
				head.RightNode = oln
				break
			} else if head.RightNode.Val < tn.Val {
				head = head.RightNode
			} else {
				head.RightNode, oln.RightNode = oln, head.RightNode
			}
		}
	}
	if ol.Vertical[tn.Row] == nil {
		ol.Vertical[tn.Row] = oln
	} else {
		head := ol.Vertical[tn.Row]
		if head.Val > tn.Val {
			ol.Vertical[tn.Row] = oln
			oln.DownNode = head
		}
		for head != nil {
			if head.DownNode == nil {
				head.DownNode = oln
				break
			} else if head.DownNode.Val < tn.Val {
				head = head.DownNode
			} else {
				head.DownNode, oln.DownNode = oln, head.DownNode
			}
		}
	}
	ol.NotNull++
}

// TODO untest function, and I'm not sure if the iterate direction right or not.
func (ol *OrthogonalLinkedList) Transpose() *OrthogonalLinkedList {
	res := &OrthogonalLinkedList{
		Shape: &Shape{
			Height: ol.Shape.Length,
			Length: ol.Shape.Height,
		},
		NotNull:    ol.NotNull,
		Horizontal: ol.Vertical,
		Vertical:   ol.Horizontal,
	}
	for _, node := range res.Horizontal {
		for node != nil {
			node.DownNode, node.RightNode = node.RightNode, node.DownNode
			node.Col, node.Row = node.Row, node.Col
			node = node.DownNode
		}
	}
	for _, node := range res.Vertical {
		for node != nil {
			node.DownNode, node.RightNode = node.RightNode, node.DownNode
			node = node.RightNode
		}
	}
	return res
}

// TODO
func (ol *OrthogonalLinkedList) PrettyPrint() string {
	// If the matrix is too big(over 9 on height or length),
	// it is too big to read easily.
	if ol.Shape.Height > 9 || ol.Shape.Length > 9 {
		return ol.ToGrid().print()
	}
	// Init a proper canvas for orthogonal linked list.
	height := ol.Shape.Height*2 + 1
	// To pprint othogonal linked list, we need to know the width
	// of every element.
	lengthList := make([]int, ol.Shape.Length)
	for i := range lengthList {
		if i < 9 {
			lengthList[i] = 2
		} else {
			lengthList[i] = 3
		}
	}
	for _, node := range ol.Horizontal {
		for node != nil {
			digitalLength := 1
			value := node.Val
			for value != 0 {
				value /= 10
				digitalLength++
			}
			if digitalLength > lengthList[node.Col] {
				lengthList[node.Col] = digitalLength
			}
			node = node.DownNode
		}
	}
	positionList := make([]int, ol.Shape.Length)
	positionList[0] = lengthList[0]
	for i := range positionList[1:] {
		if i == 0 {
			continue
		}
		positionList[i] = lengthList[i-1] + lengthList[i]
	}
	canvas := make([][]rune, height)
	for i := range canvas {
		canvas[i] = make([]rune, positionList[ol.Shape.Length-1])
	}

	for i := range canvas[1][2:] {
		canvas[1][i] = utils.BoxHor
	}
	canvas[1][1] = utils.BoxDownRight

	for i := 2; i < ol.Shape.Height; i++ {
		canvas[i][1] = utils.BoxVer
	}

	res := ""
	for _, line := range canvas {
		str := ""
		for _, v := range line {
			str = string(v)
		}
		res += str
		res += "\n"
	}
	return res
}
