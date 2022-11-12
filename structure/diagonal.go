package structure

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

func (dg *DiagonalMatrix) ToGrid() *Grid {
	grid := &Grid{
		Val:    Init2dSlice(dg.Length, dg.Length),
		Width:  dg.Length,
		Height: dg.Length,
	}
	for i := 0; i < dg.Length; i++ {
		grid.Val[i][i] = dg.Val[i]
	}
	return grid
}
