package src

/*762. 二进制表示中质数个计算置位*/

func countPrimeSetBits(L int, R int) int {
	sum := 0
	for i := L; i <= R; i++ {
		if isPrime(getCount(i)) {
			sum++
		}
	}
	return sum
}

func getCount(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 2
		num = num / 2
	}
	return sum
}

func isPrime(num int) bool {
	if num == 2 || num == 3 || num == 5 || num == 7 ||
		num == 11 || num == 13 || num == 17 || num == 19 || num == 23 || num == 29 || num == 31 {
		return true
	}
	return false
}
