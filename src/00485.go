package src

/*最大连续1的个数*/

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			sum++
			if sum > max {
				max = sum
			}
		} else {
			sum = 0
		}
	}
	return max
}
