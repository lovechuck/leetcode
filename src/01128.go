package src

import "fmt"

/*
1128. 等价多米诺骨牌对的数量

给你一个由一些多米诺骨牌组成的列表 dominoes。

如果其中某一张多米诺骨牌可以通过旋转 0 度或 180 度得到另一张多米诺骨牌，我们就认为这两张牌是等价的。

形式上，dominoes[i] = [a, b] 和 dominoes[j] = [c, d] 等价的前提是 a==c 且 b==d，或是 a==d 且 b==c。
*/
func numEquivDominoPairs(dominoes [][]int) int {
	ans := 0
	hash := make(map[string]int)
	for i := 0; i < len(dominoes); i++ {
		key := fmt.Sprintf("%d_%d", dominoes[i][0], dominoes[i][1])
		key1 := fmt.Sprintf("%d_%d", dominoes[i][1], dominoes[i][0])
		if key != key1 {
			if _, ok := hash[key]; ok {
				ans += hash[key]
			}
			if _, ok := hash[key1]; ok {
				ans += hash[key1]
			}
		} else {
			if _, ok := hash[key]; ok {
				ans += hash[key]
			}
		}
		if _, ok := hash[key]; ok {
			hash[key] = hash[key] + 1
		} else {
			hash[key] = 1
		}
	}
	return ans
}

func Test_numEquivDominoPairs() {
	arr := [][]int{
		[]int{1, 2}, []int{1, 2}, []int{1, 1}, []int{1, 2}, []int{2, 2},
	}
	numEquivDominoPairs(arr)
}
