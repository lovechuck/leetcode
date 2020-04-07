package src

/*349. 两个数组的交集*/
func intersection(nums1 []int, nums2 []int) []int {
	mmap := make(map[int]int)
	for _, i := range nums1 {
		mmap[i] = 1
	}
	var result []int
	for _, i := range nums2 {
		if _, ok := mmap[i]; ok {
			if mmap[i] == 1 {
				mmap[i] = 0
				result = append(result, i)
			}
		}
	}
	return result
}
