package src

/*
896. 单调数列
*/
func isMonotonic(A []int) bool {
	if len(A) < 3 {
		return true
	}

	flag := 0
	for i := 1; i < len(A); i++ {
		c := 0
		if A[i-1] < A[i] {
			c = -1
		} else if A[i-1] > A[i] {
			c = 1
		}
		if c != 0 {
			if c != flag && flag != 0 {
				return false
			}
			flag = c
		}
	}
	return true
}
