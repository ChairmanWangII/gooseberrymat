package structure

func Init2dSlice(length, height int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < height; i++ {
		res = append(res, make([]int, length))
	}
	return res
}

// Core algorithm for multiply.
func BaseMultiply(a, b [][]int) [][]int {
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

func Add(a, b [][]int) [][]int {
	res := Init2dSlice(len(a[0]), len(a))
	for i, line := range res {
		for j := range line {
			res[i][j] = a[i][j] + b[i][j]
		}
	}
	return res
}

func Minus(a, b [][]int) [][]int {
	res := Init2dSlice(len(a[0]), len(a))
	for i, line := range res {
		for j := range line {
			res[i][j] = a[i][j] - b[i][j]
		}
	}
	return res
}

func Sub(mat [][]int) ([][]int, [][]int, [][]int, [][]int) {
	width, height := len(mat[0])/2, len(mat)/2
	topLeft, topRight, bottomLeft, bottomRight := Init2dSlice(width, height), Init2dSlice(width, height), Init2dSlice(width, height), Init2dSlice(width, height)
	for i, line := range mat {
		for j, val := range line {
			if i < height && j < width {
				topLeft[i][j] = val
			} else if i < height && j >= width {
				topRight[i][j-width] = val
			} else if i >= height && j < width {
				bottomLeft[i-height][j] = val
			} else {
				bottomRight[i-height][j-width] = val
			}
		}
	}
	return topLeft, topRight, bottomLeft, bottomRight
}

func Combine(topLeft, topRight, bottomLeft, bottomRight [][]int) [][]int {
	width, height := len(topLeft[0]), len(topLeft)
	res := Init2dSlice(width*2, height*2)
	for i, line := range topLeft {
		for j := range line {
			res[i][j] = topLeft[i][j]
			res[i][j+width] = topRight[i][j]
			res[i+height][j] = bottomLeft[i][j]
			res[i+height][j+width] = bottomRight[i][j]
		}
	}
	return res
}
