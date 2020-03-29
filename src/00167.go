package src

/*167. 两数之和 II - 输入有序数组*/

func twoSum(numbers []int, target int) []int {
	if numbers == nil || len(numbers) == 1 {
		return nil
	}
	i := 0
	j := len(numbers) - 1
	for i < j {
		if numbers[i]+numbers[j] < target {
			i++
		} else if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		} else {
			j--
		}
	}
	return nil
}
