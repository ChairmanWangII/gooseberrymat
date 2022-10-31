package gooseberrymat

type Matrix struct {
	Error         []error
	DataStructure interface{}
	SparsityRate  float64
}

func (mt *Matrix) Multiply(mul *Matrix) *Matrix {
	data := mt.DataStructure
	switch data.(type) {
	case *Grid:

	default:
	}
	return nil
}
