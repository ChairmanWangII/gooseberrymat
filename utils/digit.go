package utils

func GetDigits(num int) int {
	str := string(rune(num))
	return len(str)
}
