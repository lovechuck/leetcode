package src

/*
922. 按奇偶排序数组 II
*/
func sortArrayByParityII(A []int) []int {
	i := 0
	j := 1
	for i < len(A) && j < len(A) {
		for i < len(A) && A[i]%2 == 0 {
			i += 2
		}
		for j < len(A) && A[j]%2 == 1 {
			j += 2
		}
		if i > len(A) || j > len(A) {
			break
		}
		temp := A[i]
		A[i] = A[j]
		A[j] = temp
		i += 2
		j += 2
	}
	return A
}
