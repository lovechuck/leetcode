package src

/*532. 数组中的K-diff数对*/

func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	mmp := make(map[int]int)
	mmap := make(map[int]int)
	res := 0
	for _, num := range nums {
		if _, ok := mmap[num+k]; ok {
			if _, ok := mmp[mmap[num+k]+num]; !ok {
				mmp[mmap[num+k]+num] = 1
				res++
			}
		}
		if _, ok := mmap[num-k]; ok {
			if _, ok := mmp[mmap[num-k]+num]; !ok {
				mmp[mmap[num-k]+num] = 1
				res++
			}
		}
		if _, ok := mmap[num]; !ok {
			mmap[num] = num
		}
	}
	return res
}
