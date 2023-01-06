package structure

import "fmt"

// If the diagonal matrix is left-top to right-bottom, the direction is TRUE.
type DiagonalMatrix struct {
	Length    int
	Val       []int
	Direction bool
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

func (dg *DiagonalMatrix) PrettyPrint() {
	var mark string
	if dg.Direction == true {
		mark = " -> "
	} else {
		mark = " <- "
	}
	var str string
	for _, element := range dg.Val {
		fmt.Sprint(str, element, mark)

	}

}
