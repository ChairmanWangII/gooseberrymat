package structure

import "gooseberrymat/utils"

// Orthogonal linked list is a data structure
// based on linked list to store sparse matrix.
type OrthogonalLinkedList struct {
	Shape   *Shape
	NotNull int
	Col     []*OrthogonalLinkedNode
	Row     []*OrthogonalLinkedNode
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
	for _, line := range ol.Col {
		for line != nil {
			matrix[line.Col][line.Row] = line.Val
			line = line.RightNode
		}
	}
	return &Grid{
		Val:   matrix,
		Shape: ol.Shape,
	}
}

// TODO need test
func (ol *OrthogonalLinkedList) Append(tn *TrigramNode) {
	oln := &OrthogonalLinkedNode{
		Val:       tn.Val,
		Col:       tn.Col,
		Row:       tn.Row,
		DownNode:  nil,
		RightNode: nil,
	}
	// If unadded node is the only node.
	if ol.Col[tn.Col] == nil {
		ol.Col[tn.Col] = oln
	} else {
		head := ol.Col[tn.Col]
		if head.Val > tn.Val {
			ol.Col[tn.Col] = oln
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
	if ol.Row[tn.Row] == nil {
		ol.Row[tn.Row] = oln
	} else {
		head := ol.Row[tn.Row]
		if head.Val > tn.Val {
			ol.Row[tn.Row] = oln
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
}

// TODO untest function, and I'm not sure if the iterate direction right or not.
func (ol *OrthogonalLinkedList) Transpose() {
	ol.Col, ol.Row = ol.Row, ol.Col
	for _, node := range ol.Col {
		for node != nil {
			node.DownNode, node.RightNode = node.RightNode, node.DownNode
			node.Col, node.Row = node.Row, node.Col
			node = node.RightNode
		}
	}
	for _, node := range ol.Row {
		for node != nil {
			node.DownNode, node.RightNode = node.RightNode, node.DownNode
			node = node.DownNode
		}
	}
}

// TODO
func (ol *OrthogonalLinkedList) PrettyPrint() string {
	// To pprint othogonal linked list, we need to know the width
	// of every element.
	widthList := make([]int, ol.Shape.Length)
	for _, i := range ol.Col {
		if utils.GetDigits(i.Val) > widthList[i.Col] {
			widthList[i.Col] = utils.GetDigits(i.Val)
		}
	}

	return ""
}
