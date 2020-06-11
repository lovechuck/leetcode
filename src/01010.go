package src

/*
1010. 总持续时间可被 60 整除的歌曲
*/
func numPairsDivisibleBy60(time []int) int {
	sum := 0
	hash := make(map[int]int, 60)
	for _, i := range time {
		key := i % 60
		if _, ok := hash[key]; ok {
			hash[key] = hash[key] + 1
		} else {
			hash[key] = 1
		}
	}
	for i := 1; i <= 29; i++ {
		sum += hash[i] * hash[60-i]
	}
	sum += hash[0]*(hash[0]-1)/2 + hash[30]*(hash[30]-1)/2
	return sum
}
