package src

/*594. 最长和谐子序列*/

func findLHS(nums []int) int {
	hash := make(map[int]int)
	for _, num := range nums {
		if _, ok := hash[num]; ok {
			hash[num] = hash[num] + 1
		} else {
			hash[num] = 1
		}
	}
	max := 0
	for k, v := range hash {
		if _, ok := hash[k+1]; ok {
			temp := v + hash[k+1]
			if max < temp {
				max = temp
			}
		}
	}
	return max
}
