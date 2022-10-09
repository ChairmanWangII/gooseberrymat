package gooseberrymat

type Add interface {
	Add()
}

type Multiply interface {
	Multiply()
}

type Shape interface {
	Width() int
	Height() int
}

type ParseToGrid interface {
	ParseToGrid() [][]int
}

type Multipliable interface {
	Multipliable() bool
}

// If two matrixes has same width and height, we call them cophenetic.
type IsCophenetic interface {
	IsCophenetic() bool
}
type Transpose interface {
	Transpose()
}

// Used to print data structure with format.
type PrettyPrint interface {
	PrettyPrint()
}
