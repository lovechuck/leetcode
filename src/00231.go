package src

/*231. 2的幂*/

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}

	return n&(n-1) == 0
}
