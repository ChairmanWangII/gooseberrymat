package structure

import "fmt"

type DiagonalMatrix struct {
	Shape     *Shape
	Val       []int
	Direction bool // If the diagonal matrix is left-top to right-bottom, the direction is TRUE.
}

// TODO now this function can only add dianonal matrixes with same direction.
func (dg *DiagonalMatrix) Add(addend *DiagonalMatrix) *DiagonalMatrix {
	sum := make([]int, dg.Shape.Length)
	for i := range sum {
		sum[i] = dg.Val[i] + addend.Val[i]
	}
	return &DiagonalMatrix{
		Shape:     dg.Shape,
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
		Val:   Init2dSlice(dg.Shape.Length, dg.Shape.Length),
		Shape: dg.Shape,
	}
	for i := 0; i < dg.Shape.Length; i++ {
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
