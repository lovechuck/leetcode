package src

/*
905. 按奇偶排序数组

给定一个非负整数数组 A，返回一个数组，在该数组中， A 的所有偶数元素之后跟着所有奇数元素。
*/

func sortArrayByParity(A []int) []int {
	left := 0
	right := len(A) - 1
	for left < right {
		for A[left]%2 == 0 && left < right {
			left++
		}
		for A[right]%2 == 1 && left < right {
			right--
		}
		temp := A[right]
		A[right] = A[left]
		A[left] = temp
		left++
		right--
	}
	return A
}
