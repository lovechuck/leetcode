package src

/*
953. 验证外星语词典

某种外星语也使用英文小写字母，但可能顺序 order 不同。字母表的顺序（order）是一些小写字母的排列。

给定一组用外星语书写的单词 words，以及其字母表的顺序 order，只有当给定的单词在这种外星语中按字典序排列时，返回 true；否则，返回 false。
*/
func isAlienSorted(words []string, order string) bool {
	hash := make(map[uint8]int)
	for i := 0; i < len(order); i++ {
		hash[order[i]] = i
	}
	for i := 0; i < len(words)-1; i++ {
		if len(words[i]) <= len(words[i+1]) {
			j := 0
			for j < len(words[i]) {
				if words[i][j] == words[i+1][j] {
					j++
				} else {
					break
				}
			}
			if j != len(words[i]) {
				if hash[words[i][j]] > hash[words[i+1][j]] {
					return false
				}
			}
		} else {
			j := 0
			for j < len(words[i+1]) {
				if words[i][j] == words[i+1][j] {
					j++
				} else {
					break
				}
			}
			if j == len(words[i+1]) {
				return false
			} else {
				if hash[words[i][j]] > hash[words[i+1][j]] {
					return false
				}
			}
		}

	}
	return true
}
