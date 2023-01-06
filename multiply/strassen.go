package multiply

import (
	gb "gooseberrymat"
	st "gooseberrymat/structure"
)

// Those can apply Stessen algorithm must can be divided four pices.
// TODO need test.
func Stressen(gd, mul *st.Grid) *st.Grid {
	res := &st.Grid{
		Width:  mul.Width,
		Height: gd.Height,
		Val:    st.Init2dSlice(mul.Width, gd.Height),
	}

	var strassenAlgorithm func(a, b [][]int) [][]int
	strassenAlgorithm = func(a, b [][]int) [][]int {
		aWidth, aHeight := len(a[0]), len(a)
		bWidth, bHeight := len(b[0]), len(b)
		if aHeight%2 != 0 || aWidth%2 != 0 || bHeight%2 != 0 || bWidth%2 != 0 ||
			aHeight < gb.STRESSEN_BORDER || aWidth < gb.STRESSEN_BORDER || bWidth < gb.STRESSEN_BORDER {
			return st.BaseMultiply(a, b)
		}
		A11, A12, A21, A22 := st.Sub(a)
		B11, B12, B21, B22 := st.Sub(b)

		S1 := st.Minus(B12, B22)
		S2 := st.Add(A11, A12)
		S3 := st.Add(A21, A22)
		S4 := st.Minus(B21, B11)
		S5 := st.Add(A11, A22)
		S6 := st.Add(B11, B22)
		S7 := st.Minus(A12, A22)
		S8 := st.Add(B21, B22)
		S9 := st.Minus(A11, A21)
		S10 := st.Add(B11, B12)

		P1 := strassenAlgorithm(A11, S1)
		P2 := strassenAlgorithm(S2, B22)
		P3 := strassenAlgorithm(S3, B11)
		P4 := strassenAlgorithm(A22, S4)
		P5 := strassenAlgorithm(S5, S6)
		P6 := strassenAlgorithm(S7, S8)
		P7 := strassenAlgorithm(S9, S10)

		C11 := st.Minus(st.Add(P5, P4), st.Add(P2, P6))
		C12 := st.Add(P1, P2)
		C21 := st.Add(P3, P4)
		C22 := st.Minus(st.Add(P5, P1), st.Minus(P3, P7))

		C := st.Combine(C11, C12, C21, C22)
		return C
	}
	return res
}
