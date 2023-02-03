package structure

type Matrix struct {
	Error         []error
	DataStructure interface{}
	SparsityRate  float64
}

type Shape struct {
	Height int
	Length int
}
