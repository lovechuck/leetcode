package src

import "fmt"

/*53. 最大子序和*/

func maxSubArray(nums []int) int {
	ans := nums[0]
	sum := 0

	for i := 0; i < len(nums); i++ {
		if sum > 0 {
			sum = sum + nums[i]
		} else {
			sum = nums[i]
		}
		if sum > ans {
			ans = sum
		}
	}

	return sum
}

func Test_maxSubArray() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(arr))
}
