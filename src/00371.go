package src

/*371. 两整数之和*/
func getSum(a int, b int) int {
	for b != 0 {
		temp := a ^ b
		b = (a & b) << 1
		a = temp
	}
	return a
}
