package gooseberrymat

func (gd *Grid) Multiply(mul *Grid) *Grid {
	return gd.SimpleMultiply(mul)
}

func (gd *Grid) SimpleMultiply(mul *Grid) *Grid {
	res := &Grid{
		Width:  mul.Width,
		Height: gd.Height,
		Val:    baseMultiply(gd.Val, mul.Val),
	}
	return res
}

// Those can apply Stessen algorithm must can be divided four pices
func (gd *Grid) Stressen(mul *Grid) *Grid {
	res := &Grid{
		Width:  mul.Width,
		Height: gd.Height,
		Val:    Init2dSlice(mul.Width, gd.Height),
	}

	var stressenAlgorithm func(a, b [][]int) [][]int
	stressenAlgorithm = func(a, b [][]int) [][]int {
		aWidth, aHeight := len(a[0]), len(a)
		bWidth, bHeight := len(b[0]), len(b)
		if aHeight%2 != 0 || aWidth%2 != 0 || bHeight%2 != 0 || bWidth%2 != 0 {
			return baseMultiply(a, b)
		}
		A11 := a[:aHeight/2+1][:aWidth/2+1]
		A12 := a[:aHeight/2+1][aWidth/2+1:]
		A21 := a[aHeight/2+1:][:aWidth/2+1]
		A22 := a[aHeight/2+1:][aWidth/2+1:]

		B11 := a[:bHeight/2+1][:bWidth/2+1]
		B12 := a[:bHeight/2+1][bWidth/2+1:]
		B21 := a[bHeight/2+1:][:bWidth/2+1]
		B22 := a[bHeight/2+1:][bWidth/2+1:]

		S1 := minus(B12, B22)
		S2 := add(A11, A12)
		S3 := add(A21, A22)
		S4 := minus(B21, B11)
		S5 := add(A11, A22)
		S6 := add(B11, B22)
		S7 := minus(A12, A22)
		S8 := add(B21, B22)
		S9 := minus(A11, A21)
		S10 := add(B11, B12)

		P1 := stressenAlgorithm(A11, S1)
		P2 := stressenAlgorithm(S2, B22)
		P3 := stressenAlgorithm(S3, B11)
		P4 := stressenAlgorithm(A22, S4)
		P5 := stressenAlgorithm(S5, S6)
		P6 := stressenAlgorithm(S7, S8)
		P7 := stressenAlgorithm(S9, S10)

		C11 := minus(add(P5, P4), add(P2, P6))
		C12 := add(P1, P2)
		C21 := add(P3, P4)
		C22 := minus(add(P5, P1), minus(P3, P7))

		C1 := append(C11, C12...)
		C2 := append(C21, C22...)

		C := append(C1, C2...)

		return C
	}

	return res
}
