package src

/*409. 最长回文串*/

func longestPalindrome409(s string) int {
	mmap := make(map[int32]int)
	for _, i := range s {
		if _, ok := mmap[i]; ok {
			mmap[i] = mmap[i] + 1
		} else {
			mmap[i] = 1
		}
	}
	sum := 0
	for _, item := range mmap {
		if item%2 == 1 {
			sum += item - 1
		} else {
			sum += item
		}
	}
	if sum < len(s) {
		sum += 1
	}
	return sum
}
