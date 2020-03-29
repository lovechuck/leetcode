package src

/*204. 计数质数*/

func countPrimes(n int) int {
	if n == 1 {
		return 0
	}
	count := 0
	mmap := make(map[int]bool, n)
	for i := 2; i < n; i++ {
		if mmap[i] {
			continue
		}
		for j := i + i; j < n; j = j + i {
			mmap[j] = true
		}
		count++
	}
	return count
}
