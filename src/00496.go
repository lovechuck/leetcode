package src

/*496. 下一个更大元素 I*/

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := []int{}
	mmap := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
			mmap[stack[len(stack)-1]] = nums2[i]
			stack = stack[0 : len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}
	for i := 0; i < len(stack); i++ {
		mmap[stack[i]] = -1
	}
	var res []int
	for i := 0; i < len(nums1); i++ {
		if _, ok := mmap[nums1[i]]; ok {
			res = append(res, mmap[nums1[i]])
		} else {
			res = append(res, -1)
		}
	}
	return res
}
