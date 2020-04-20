package src

import "strings"

/*500. 键盘行*/

func findWords(words []string) []string {
	map1 := map[rune]int{
		'q': 1,
		'w': 1,
		'e': 1,
		'r': 1,
		't': 1,
		'y': 1,
		'u': 1,
		'i': 1,
		'o': 1,
		'p': 1,
	}
	map2 := map[rune]int{
		'a': 1,
		's': 1,
		'd': 1,
		'f': 1,
		'g': 1,
		'h': 1,
		'j': 1,
		'k': 1,
		'l': 1,
	}
	map3 := map[rune]int{
		'z': 1,
		'x': 1,
		'c': 1,
		'v': 1,
		'b': 1,
		'n': 1,
		'm': 1,
	}

	var res []string
	for i := 0; i < len(words); i++ {
		if oneline(map1, words[i]) || oneline(map2, words[i]) || oneline(map3, words[i]) {
			res = append(res, words[i])
		}
	}

	return res
}

func oneline(mm map[rune]int, str string) bool {
	if len(str) <= 1 {
		return true
	}
	str = strings.ToLower(str)
	for _, i := range str {
		if _, ok := mm[i]; !ok {
			return false
		}
	}
	return true
}
