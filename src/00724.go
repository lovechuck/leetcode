package src

/*724. 寻找数组的中心索引*/

func pivotIndex(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	left := 0
	for i, num := range nums {
		if left*2+num == sum {
			return i
		}
		left += num
	}
	return -1
}
