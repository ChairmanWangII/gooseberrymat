package structure

import (
	"fmt"
	"math/rand"
	"sort"
)

// Trigram is the only indefinite length structure, so we need calculate its Length.
type Trigram struct {
	Width  int
	Height int
	Length int
	Val    []*TrigramNode
}

type TrigramNode struct {
	Row int
	Col int
	Val int
}

func (tg *Trigram) ToGrid() *Grid {
	val := Init2dSlice(tg.Width, tg.Height)
	grid := &Grid{
		Width:  tg.Width,
		Height: tg.Height,
		Val:    val,
	}
	for _, v := range tg.Val {
		grid.Val[v.Row][v.Col] = v.Val
	}
	return grid
}

func (tg *Trigram) Transpose() *Trigram {
	for _, p := range tg.Val {
		p.Col, p.Row = p.Row, p.Col
	}
	return tg.tidy()
}

// Tidy is used to tidy up the sequence of trigrams, so that we can simplier add trigram matrixes.
func (tg *Trigram) tidy() *Trigram {
	sort.Slice(tg.Val, func(i, j int) bool {
		if tg.Val[i].Row != tg.Val[j].Row {
			return tg.Val[i].Row < tg.Val[j].Row
		} else {
			return tg.Val[i].Col < tg.Val[j].Col
		}
	})
	return tg
}

// This function is used to test.
func (tg *Trigram) From2dSliceToTrigram(matrix [][]int) *Trigram {
	tr := &Trigram{
		Val:    make([]*TrigramNode, 0),
		Width:  len(matrix[0]),
		Height: len(matrix),
	}
	for i, line := range matrix {
		for j := range line {
			if matrix[i][j] != 0 {
				tr.Val = append(tr.Val, &TrigramNode{i, j, matrix[i][j]})
			}
		}
	}
	return tr
}

func (tg *Trigram) shuffle() {
	N := len(tg.Val)
	for i := 0; i < N; i++ {
		r := i + rand.Intn(N-i)
		tg.Val[r], tg.Val[i] = tg.Val[i], tg.Val[r]
	}
}

func (tg *Trigram) Add(addend *Trigram) *Trigram {
	tg.tidy()
	addend.tidy()
	res := &Trigram{
		Width:  tg.Width,
		Height: tg.Height,
		Val:    make([]*TrigramNode, 0),
	}
	leftLen, rightLen := tg.Length, addend.Length
	l, r := 0, 0
	for (leftLen-l-1)*(rightLen-r-1) != 0 {
		if tg.Val[l].Col == addend.Val[r].Col && tg.Val[l].Row == addend.Val[r].Row {
			res.Val = append(res.Val, &TrigramNode{
				Val: tg.Val[l].Val + addend.Val[r].Val,
				Row: tg.Val[l].Row,
				Col: tg.Val[l].Col,
			})
			l++
			r++
		}
		if tg.Val[l].Row > addend.Val[r].Row || tg.Val[l].Row == addend.Val[r].Row && tg.Val[l].Col > addend.Val[r].Col {
			res.Val = append(res.Val, addend.Val[r])
			r++
		}
		if tg.Val[l].Row < addend.Val[r].Row || tg.Val[l].Row == addend.Val[r].Row && tg.Val[l].Col < addend.Val[r].Col {
			res.Val = append(res.Val, tg.Val[l])
			l++
		}
		// Append unmerged data to result Trigram.
		if leftLen-l-1 != 0 {
			res.Val = append(res.Val, tg.Val[leftLen-l:]...)
		}
		if rightLen-r-1 != 0 {
			res.Val = append(res.Val, addend.Val[rightLen-r:]...)
		}
	}
	return res
}

func (tn *TrigramNode) format() string {
	return fmt.Sprint("(", tn.Row, tn.Col, tn.Val, ")")
}

func (tg *Trigram) PrettyPrint() string {
	var str string
	for _, item := range tg.Val {
		str = fmt.Sprint(str, item.format(), " ")
	}
	return str
}
