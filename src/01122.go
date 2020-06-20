package src

import "sort"

/*
1122. 数组的相对排序

给你两个数组，arr1 和 arr2，

arr2 中的元素各不相同
arr2 中的每个元素都出现在 arr1 中
对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2 中出现过的元素需要按照升序放在 arr1 的末尾。
*/
func relativeSortArray(arr1 []int, arr2 []int) []int {
	hash := make(map[int]int)
	for _, i2 := range arr2 {
		hash[i2] = 1
	}
	hash2 := make(map[int]int)
	var left []int
	for _, i := range arr1 {
		if _, ok := hash[i]; ok {
			if _, ok := hash2[i]; ok {
				hash2[i] = hash2[i] + 1
			} else {
				hash2[i] = 1
			}
		} else {
			left = append(left, i)
		}
	}
	var ans []int
	for _, i2 := range arr2 {
		if _, ok := hash2[i2]; ok {
			for i := 0; i < hash2[i2]; i++ {
				ans = append(ans, i2)
			}
		}
	}
	sort.Ints(left)
	ans = append(ans, left...)
	return ans
}
