package src

/*
给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的 绝对值 至多为 k。
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	mmap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := mmap[nums[i]]; ok {
			j := mmap[nums[i]]
			if i-j <= k {
				return true
			} else {
				mmap[nums[i]] = i
			}
		} else {
			mmap[nums[i]] = i
		}
	}
	return false
}
