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
