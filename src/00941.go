package src

/*941. 有效的山脉数组*/
func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}
	i := 0
	j := len(A) - 1
	for i+1 < len(A) && A[i] < A[i+1] {
		i++
	}
	for j >= 1 && A[j] < A[j-1] {
		j--
	}
	if i == j && i > 0 && j < len(A)-1 {
		return true
	}
	return false
}
