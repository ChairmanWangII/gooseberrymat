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
	gd := &Grid{
		Val: testGrid,
	}
	qt := gd.ParseToQuadTree()
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
	gd := &Grid{
		Val: testGrid,
	}
	qt := gd.ParseToQuadTree()
	gd = qt.ParseToGrid()
	t.Log(gd)
}

func TestTransposeQuadTree(t *testing.T) {
	testGrid := [][]int{
		{1, 1, 2, 2, 0, 0, 0, 0},
		{1, 1, 2, 2, 0, 0, 0, 0},
		{1, 1, 1, 1, 3, 3, 1, 1},
		{1, 1, 1, 1, 3, 3, 1, 1},
		{1, 1, 1, 1, 5, 5, 0, 0},
		{1, 1, 1, 1, 5, 5, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
	}
	gd := &Grid{
		Val: testGrid,
	}
	t.Log("print origin matrix")
	t.Log("---------------")
	t.Log(gd.Val)
	qt := gd.ParseToQuadTree()
	gd = qt.ParseToGrid()
	t.Log("print changed matrix")
	t.Log("---------------")
	t.Log(gd.Val)

	qtt := qt.Transpose()
	res := qtt.ParseToGrid()
	t.Log(res.Val)
	t.Log("print transposed matrix")
	t.Log("---------------")
	t.Log(testGrid)
}

func TestTrigramTidy(t *testing.T) {
	matrix := [][]int{
		{1, 2, 0},
		{2, 2, 0},
		{0, 0, 1},
	}
	var tg *Trigram
	tg.ParseTwoDimensionalSliceToTrigram(matrix)
	t.Log("this is a matrix")
	t.Log(tg.Val)
	tg.shuffle()
	t.Log("shuffled matrix")
	t.Log(tg.Val)
	tg.tidy()
	t.Log("tidied matrix")
	t.Log(tg.Val)
}
