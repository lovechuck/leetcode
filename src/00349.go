package src

/*349. 两个数组的交集*/
func intersect(nums1 []int, nums2 []int) []int {
	mmap := make(map[int]int)
	for _, i := range nums1 {
		if _, ok := mmap[i]; ok {
			mmap[i] = mmap[i] + 1
		} else {
			mmap[i] = 1
		}
	}
	var result []int
	for _, i := range nums2 {
		if _, ok := mmap[i]; ok {
			if mmap[i] > 0 {
				mmap[i] = mmap[i] - 1
				result = append(result, i)
			}
		}
	}
	return result
}
