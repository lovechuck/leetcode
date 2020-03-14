package src

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	mmap := make(map[int]int)
	mmap[1] = 1
	mmap[2] = 2
	return climbStairsMap(n, mmap)
}

func climbStairsMap(n int, mmap map[int]int) int {
	if _, ok := mmap[n]; ok {
		return mmap[n]
	}

	result := climbStairsMap(n-1, mmap) + climbStairsMap(n-2, mmap)
	mmap[n] = result
	return result
}
