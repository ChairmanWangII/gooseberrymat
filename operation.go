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
