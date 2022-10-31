package gooseberrymat

func Init2dSlice(width, height int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < height; i++ {
		res = append(res, make([]int, width))
	}
	return res
}

// Core algorithm for multiply.
func baseMultiply(a, b [][]int) [][]int {
	aWidth, bHeight := len(a[0]), len(b)
	res := Init2dSlice(aWidth, bHeight)
	operteCount := len(a)
	for i, line := range res {
		for j := range line {
			count := 0
			for m := 0; m < operteCount; m++ {
				count += a[i][m] * b[m][j]
			}
			res[i][j] = count
		}
	}
	return res
}

func add(a, b [][]int) [][]int {
	res := Init2dSlice(len(a[0]), len(a))
	for i, line := range res {
		for j := range line {
			res[i][j] = a[i][j] + b[i][j]
		}
	}
	return res
}

func minus(a, b [][]int) [][]int {
	res := Init2dSlice(len(a[0]), len(a))
	for i, line := range res {
		for j := range line {
			res[i][j] = a[i][j] - b[i][j]
		}
	}
	return res
}
