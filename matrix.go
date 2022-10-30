package gooseberrymat

type Matrix struct {
	Error         []error
	DataStructure interface{}
}

func (mt *Matrix) Multiply(mul *Matrix) *Matrix {
	data := mt.DataStructure
	switch data.(type) {
	case *Grid:

	default:
	}
}
