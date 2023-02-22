package utils

func GetDigitLength(num int) int {
	count := 1
	for num > 9 {
		count++
		num /= 10
	}
	return count
}
