package gooseberrymat

type DiagonalMatrix struct {
	Length int
	Val    []int
}

func (dg *DiagonalMatrix) Add(addend *DiagonalMatrix) *DiagonalMatrix {
	sum := make([]int, dg.Length)
	for i := range sum {
		sum[i] = dg.Val[i] + addend.Val[i]
	}
	return &DiagonalMatrix{
		Length: dg.Length,
		Val:    sum,
	}
}

// Shortest function in this package.
func (dg *DiagonalMatrix) Transpose() *DiagonalMatrix {
	return dg
}

// It is a bad idea to use 'Constructor' in Go
func (dg *DiagonalMatrix) Constructor(matrix *Grid) *DiagonalMatrix {
	dm := &DiagonalMatrix{
		Length: matrix.Height,
		Val:    make([]int, matrix.Height),
	}
	for i := 0; i < dg.Length; i++ {
		dm.Val[i] = matrix.Val[i][i]
	}
	return dm
}
