package src

/*720. 词典中最长的单词*/
func longestWord(words []string) string {
	hash := make(map[string]string)
	for _, word := range words {
		hash[word] = word
	}
	leng := -1
	sword := ""
	for _, word := range words {
		find := true
		for i := len(word) - 1; i >= 1; i-- {
			key := word[:i]
			if _, ok := hash[key]; !ok {
				find = false
				break
			}
		}
		if find {
			if leng < len(word) {
				leng = len(word)
				sword = word
			} else if leng == len(word) {
				if sword > word {
					sword = word
				}
			}
		}
	}
	return sword
}

func Test_longestWord() {
	arr := []string{"w", "wo", "wor", "worl", "world"}
	longestWord(arr)
}
