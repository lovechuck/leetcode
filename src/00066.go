package src

/*66. 加一*/

func plusOne(digits []int) []int {
	result := make([]int, len(digits)+1)
	add := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+add == 10 {
			result[i+1] = 0
			add = 1
		} else {
			result[i+1] = digits[i] + add
			add = 0
		}
	}
	if add == 1 {
		result[0] = 1
		return result
	}

	return result[1:]
}
