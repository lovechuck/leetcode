package src

import "math"

/*507. 完美数*/

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	a := int(math.Sqrt(float64(num)))
	sum := 0
	for i := 2; i <= a; i++ {
		if num%i == 0 {
			sum += i
			sum += num / i
		}
	}
	return sum+1-num == 0
}
