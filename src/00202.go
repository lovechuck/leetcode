package src

/*202. 快乐数*/

func isHappy(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	slow := square(n)
	fast := square(n)
	fast = square(fast)
	for slow != fast {
		slow = square(slow)
		fast = square(fast)
		fast = square(fast)
		if fast == 1 {
			return true
		}
	}
	if fast == 1 {
		return true
	}
	return false
}

func square(n int) int {
	ans := 0
	for n > 0 {
		ans += (n % 10) * (n % 10)
		n = n / 10
	}
	return ans
}
