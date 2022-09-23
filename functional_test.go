package gooseberrymat

import (
	"testing"
)

func TestQuadTreeConstructor(t *testing.T) {
	testGrid := [][]int{
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
	}
	var qt *QuadTree
	qt = qt.Constructor(&Grid{
		Val: testGrid,
	})
	t.Log(qt)
}

func TestParseToGrid(t *testing.T) {
	testGrid := [][]int{
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
	}
	var qt *QuadTree
	qt = qt.Constructor(&Grid{
		Val: testGrid,
	})
	gd := qt.ParseToGrid()
	t.Log(gd)
}

func TestTransposeQuadTree(t *testing.T) {
	testGrid := [][]int{
		{1, 1, 2, 2, 0, 0, 0, 0},
		{1, 1, 2, 2, 0, 0, 0, 0},
		{1, 1, 1, 1, 3, 3, 1, 1},
		{1, 1, 1, 1, 3, 3, 1, 1},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
	}
	var qt *QuadTree
	qt = qt.Constructor(&Grid{
		Val: testGrid,
	})
	qtt := qt.Transpose()
	res := qtt.ParseToGrid()
	t.Log(res.Val)
	t.Log(testGrid)
}
