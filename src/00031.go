package src

/*
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
*/

func nextPermutation(nums []int) {
	if len(nums) == 1 {
		return
	}

	i := len(nums) - 2
	for ; i >= 0; i-- {
		if nums[i+1] > nums[i] {
			break
		}
	}
	if i >= 0 {
		j := len(nums) - 1
		for ; j > i; j-- {
			if nums[j] > nums[i] {
				temp := nums[j]
				nums[j] = nums[i]
				nums[i] = temp
				break
			}
		}
	}

	i = i + 1
	j := len(nums) - 1
	for i < j {
		temp := nums[j]
		nums[j] = nums[i]
		nums[i] = temp
		i++
		j--
	}
}

func Test_nextPermutation() {
	arr := []int{1, 2, 3}
	nextPermutation(arr)
}
