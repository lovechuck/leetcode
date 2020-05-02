package src

/*575. 分糖果*/

func distributeCandies(candies []int) int {
	set := make(map[int]bool)
	sum := 0
	for _, candy := range candies {
		if _, ok := set[candy]; !ok {
			set[candy] = true
			sum++
		}
	}
	if len(candies)/2 > sum {
		return sum
	}
	return len(candies) / 2
}
