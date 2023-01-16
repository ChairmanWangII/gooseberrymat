package structure

import "fmt"

type DiagonalMatrix struct {
	Length    int
	Val       []int
	Direction bool // If the diagonal matrix is left-top to right-bottom, the direction is TRUE.
}

// TODO now this function can only add dianonal matrixes with same direction.
func (dg *DiagonalMatrix) Add(addend *DiagonalMatrix) *DiagonalMatrix {
	sum := make([]int, dg.Length)
	for i := range sum {
		sum[i] = dg.Val[i] + addend.Val[i]
	}
	return &DiagonalMatrix{
		Length:    dg.Length,
		Val:       sum,
		Direction: dg.Direction,
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

func (dg *DiagonalMatrix) PrettyPrint() string {
	var mark string
	if dg.Direction {
		mark = " ↗ "
	} else {
		mark = " ↘ "
	}
	var str string
	for _, element := range dg.Val {
		str = fmt.Sprint(str, element, mark)
	}
	return str
}
