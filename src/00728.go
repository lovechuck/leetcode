package src

/*728. 自除数*/
func selfDividingNumbers(left int, right int) []int {
	var res []int
	for i := left; i <= right; i++ {
		if selfDivide(i) {
			res = append(res, i)
		}
	}
	return res
}
func selfDivide(num int) bool {
	if num == 0 {
		return false
	}
	if num >= 1 && num <= 9 {
		return true
	}
	cat := num
	for cat > 0 {
		a := cat % 10
		if a == 0 || num%a != 0 {
			return false
		}
		cat /= 10
	}
	return true
}
