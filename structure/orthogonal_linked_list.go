package gooseberrymat

type OrthogonalLinkedList struct {
	Width   int
	Height  int
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
	matrix := Init2dSlice(ol.Width, ol.Height)
	for _, line := range ol.Col {
		for line != nil {
			matrix[line.Col][line.Row] = line.Val
			line = line.RightNode
		}
	}
	return &Grid{
		Val:    matrix,
		Height: ol.Height,
		Width:  ol.Width,
	}
}
