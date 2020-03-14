package src

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}
	pre := []int{1, 1}
	for i := 1; i <= rowIndex; i++ {
		temp := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 {
				temp[j] = 1
			} else if j == i {
				temp[j] = 1
			} else {
				temp[j] = pre[j-1] + pre[j]
			}
		}
		pre = temp
	}
	return pre
}
