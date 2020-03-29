package src

/*258. 各位相加*/
func addDigits(num int) int {
	if num == 0 {
		return 0
	}
	res := num % 9
	if res == 0 {
		return 9
	}
	return res
}
