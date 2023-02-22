package structure

import (
	"fmt"
	"testing"
)

func getTestGrid() [][]int {
	return [][]int{
		{1, 1, 2, 2, 0, 0, 0, 0},
		{1, 1, 2, 2, 0, 0, 0, 0},
		{1, 1, 1, 1, 3, 3, 1, 1},
		{1, 1, 1, 1, 3, 3, 1, 1},
		{1, 1, 1, 1, 5, 5, 0, 0},
		{1, 1, 1, 1, 5, 5, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
	}
}

func getTestDiagonal() *DiagonalMatrix {
	return &DiagonalMatrix{
		Shape: &Shape{
			Length: 5,
			Height: 5,
		},
		Val:       []int{1, 3, 5, 8, 1},
		Direction: true,
	}
}

func getTestTrigram() *Trigram {
	tn1 := &TrigramNode{Row: 1, Col: 1, Val: 2}
	tn2 := &TrigramNode{Row: 2, Col: 1, Val: 3}
	tn3 := &TrigramNode{Row: 3, Col: 4, Val: 6}
	tn4 := &TrigramNode{Row: 4, Col: 4, Val: 1}
	tn5 := &TrigramNode{Row: 4, Col: 5, Val: 7}
	return &Trigram{
		Shape: &Shape{
			Height: 5,
			Length: 5,
		},
		NotNull: 5,
		Val: []*TrigramNode{
			tn1, tn2, tn3, tn4, tn5,
		},
	}
}

func TestQuadTreeConstructor(t *testing.T) {
	gd := &Grid{
		Val: getTestGrid(),
	}
	qt := gd.ToQuadTree()
	t.Log(qt)
	fmt.Println(qt)
}

func TestParseToGrid(t *testing.T) {
	gd := &Grid{
		Val: getTestGrid(),
	}
	qt := gd.ToQuadTree()
	gd = qt.ToGrid()
	t.Log(gd)
}

func TestTransposeQuadTree(t *testing.T) {
	gd := &Grid{
		Val: getTestGrid(),
	}
	t.Log("print origin matrix")
	t.Log("---------------")
	t.Log(gd.Val)
	qt := gd.ToQuadTree()
	gd = qt.ToGrid()
	t.Log("print changed matrix")
	t.Log("---------------")
	t.Log(gd.Val)

	qtt := qt.Transpose()
	res := qtt.ToGrid()
	t.Log(res.Val)
	t.Log("print transposed matrix")
	t.Log("---------------")
	t.Log(getTestGrid())
}

func TestTrigramTidy(t *testing.T) {
	matrix := [][]int{
		{1, 2, 0},
		{2, 2, 0},
		{0, 0, 1},
	}
	var tg *Trigram
	tg = tg.From2dSliceToTrigram(matrix)
	t.Log("this is a matrix")
	t.Log(tg.Val)
	tg.shuffle()
	t.Log("shuffled matrix")
	t.Log(tg.Val)
	tg.tidy()
	t.Log("tidied matrix")
	t.Log(tg.Val)
}

func Test2dSlice(t *testing.T) {
	matrix := Init2dSlice(3, 5)
	for _, line := range matrix {
		fmt.Println(line)
	}
}
func TestNilType(t *testing.T) {
	testGrid := &Grid{
		Shape: &Shape{
			Height: 12,
		},
	}
	if testGrid.Val == nil {
		fmt.Println("approve")
	}
}

func TestQuadTreePrettyPrint(t *testing.T) {
	gd := &Grid{
		Val: getTestGrid(),
		Shape: &Shape{
			Length: 8,
			Height: 8,
		},
	}
	tg := gd.ToQuadTree()
	str := tg.PrettyPrint()
	t.Log(str)
	fmt.Println(str)
}

func TestDiagonalPrettyPrint(t *testing.T) {
	str := getTestDiagonal().PrettyPrint()
	t.Log(str)
	fmt.Println(str)
}

func TestTrigramPerttyPrint(t *testing.T) {
	str := getTestTrigram().PrettyPrint()
	fmt.Println(str)
}

func TestUsingTcell(t *testing.T) {
}

func TestGridPrint(t *testing.T) {
	gd := &Grid{
		Val: getTestGrid(),
		Shape: &Shape{
			Length: 8,
			Height: 8,
		},
	}
	gd.Val[5][5] = 17
	fmt.Println(gd.print())

}

func TestOLLPrettyPrint(t *testing.T) {
	ol := &OrthogonalLinkedList{
		Shape: &Shape{
			Height: 5,
			Length: 5,
		},
		NotNull: 0,
		Col:     make([]*OrthogonalLinkedNode, 5),
		Row:     make([]*OrthogonalLinkedNode, 5),
	}
	ol.Append(&TrigramNode{
		Row: 2,
		Col: 3,
		Val: 5,
	})
	ol.Append(&TrigramNode{
		Row: 0,
		Col: 0,
		Val: 9,
	})
	ol.Append(&TrigramNode{
		Row: 3,
		Col: 2,
		Val: 4,
	})
	ol.Append(&TrigramNode{
		Row: 2,
		Col: 2,
		Val: 7,
	})
	ol.Append(&TrigramNode{
		Row: 4,
		Col: 4,
		Val: 6,
	})
	grid := ol.ToGrid()
	fmt.Println(grid.print())

	str := ol.PrettyPrint()
	fmt.Println(str)
}
