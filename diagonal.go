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
