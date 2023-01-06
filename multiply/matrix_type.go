package multiply

import st "gooseberrymat/structure"

type MatrixType struct {
	Gd   *st.Grid
	Oll  *st.OrthogonalLinkedList
	Qt   *st.QuadTree
	Tg   *st.Trigram
	Dg   *st.DiagonalMatrix
	Zero bool
}

func (mt *MatrixType) Multiply(mul *MatrixType) *MatrixType {
	return BaseMultiply(mt, mul)
}

func Grid2MatrixType(gd *st.Grid) *MatrixType {
	return &MatrixType{
		Gd: gd,
	}
}

func OrthogonalLinkedList2MatrixType(oll *st.OrthogonalLinkedList) *MatrixType {
	return &MatrixType{
		Oll: oll,
	}
}

func QuadTree2MatrixType(qt *st.QuadTree) *MatrixType {
	return &MatrixType{
		Qt: qt,
	}
}

func Trigram2MatrixType(tg *st.Trigram) *MatrixType {
	return &MatrixType{
		Tg: tg,
	}
}

func Diagonal2MatrixType(dg *st.DiagonalMatrix) *MatrixType {
	return &MatrixType{
		Dg: dg,
	}
}
