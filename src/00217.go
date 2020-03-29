package src

/*217. 存在重复元素*/
/*
给定一个整数数组，判断是否存在重复元素。

如果任何值在数组中出现至少两次，函数返回 true。如果数组中每个元素都不相同，则返回 false。
*/

func containsDuplicate(nums []int) bool {
	mmap := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := mmap[nums[i]]; ok {
			return true
		} else {
			mmap[nums[i]] = true
		}
	}
	return false
}
