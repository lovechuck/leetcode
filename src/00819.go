package src

import (
	"strings"
	"unicode"
)

/*819. 最常见的单词*/
func mostCommonWord(paragraph string, banned []string) string {
	hash := map[string]bool{}
	for _, s := range banned {
		hash[s] = true
	}

	set := map[string]int{}
	paragraph += "."
	max := 0
	res := ""
	var word []rune
	for _, i := range paragraph {
		if unicode.IsLetter(i) {
			word = append(word, i)
		} else {
			if word != nil && len(word) > 0 {
				key := strings.ToLower(string(word))
				if _, ok := hash[key]; !ok {
					if _, ok := set[key]; !ok {
						set[key] = 1
					} else {
						set[key] = set[key] + 1
					}
					if set[key] > max {
						max = set[key]
						res = key
					}
				}
				word = []rune{}
			}
		}
	}
	return res
}

func Test_mostCommonWord() {
	mostCommonWord("Bob. hIt, baLl", []string{"bob", "hit"})
}
