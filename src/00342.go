package src

import "math"

/*
4的幂
int 范围类最大的4的幂 4^15
*/

func isPowerOfFour(n int) bool {
	mmap := make(map[int]bool)
	for i := 0; i <= 15; i++ {
		k := int(math.Pow(4, float64(i)))
		mmap[k] = true
	}
	_, ok := mmap[n]
	return ok
}
