package src

/*
914. 卡牌分组

给定一副牌，每张牌上都写着一个整数。

此时，你需要选定一个数字 X，使我们可以将整副牌按下述规则分成 1 组或更多组：

每组都有 X 张牌。
组内所有的牌上都写着相同的整数。
仅当你可选的 X >= 2 时返回 true。
*/
func hasGroupsSizeX(deck []int) bool {
	count := make([]int, 10000)
	for _, i := range deck {
		count[i]++
	}

	gc := -1
	for _, i := range count {
		if i > 0 {
			if gc == -1 {
				gc = i
			} else {
				gc = gcd(i, gc)
			}
		}
	}

	return gc >= 2
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
