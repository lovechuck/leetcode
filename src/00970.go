package src

/*
970. 强整数

给定两个正整数 x 和 y，如果某一整数等于 x^i + y^j，其中整数 i >= 0 且 j >= 0，那么我们认为该整数是一个强整数。

返回值小于或等于 bound 的所有强整数组成的列表。

你可以按任何顺序返回答案。在你的回答中，每个值最多出现一次。
*/
func powerfulIntegers(x int, y int, bound int) []int {
	hash := make(map[int]int)
	var ans []int
	for i := 0; i <= 20 && Pow(x, i) <= bound; i++ {
		for j := 0; j <= 20 && Pow(y, j) <= bound; j++ {
			key := Pow(x, i) + Pow(y, j)
			if key <= bound {
				if _, ok := hash[key]; !ok {
					hash[key] = 1
					ans = append(ans, key)
				}
			}
		}
	}
	return ans
}

func Pow(x int, y int) int {
	sum := 1
	for i := 0; i < y; i++ {
		sum *= x
	}
	return sum
}

func Test_powerfulIntegers() {
	powerfulIntegers(2, 3, 10)
}
