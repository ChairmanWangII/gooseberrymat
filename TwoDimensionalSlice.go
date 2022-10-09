package gooseberrymat

func Init2dSlice(width, height int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < height; i++ {
		res = append(res, make([]int, width))
	}
	return res
}
