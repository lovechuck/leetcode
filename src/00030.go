package src

import "fmt"

/*
给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。

注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。

示例 1：

输入：
  s = "barfoothefoobarman",
  words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。
示例 2：

输入：
  s = "wordgoodgoodgoodbestword",
  words = ["word","good","best","word"]
输出：[]
*/
func findSubstring(s string, words []string) []int {
	if len(words) == 0 || s == "" {
		return nil
	}
	single := len(words[0])
	all := single * len(words)
	var result []int
	wmap := make(map[string]int)
	for i := 0; i < len(words); i++ {
		if _, ok := wmap[words[i]]; ok {
			wmap[words[i]] = wmap[words[i]] + 1
		} else {
			wmap[words[i]] = 1
		}
	}
	for i := 0; i < len(s)+1-all; i++ {
		tmap := make(map[string]int)
		for j := i; j < i+all; j += single {
			stmp := s[j : j+single]
			if _, ok := tmap[stmp]; ok {
				tmap[stmp] = tmap[stmp] + 1
			} else {
				tmap[stmp] = 1
			}
		}
		if cmpMap(tmap, wmap) {
			result = append(result, i)
		}
	}
	return result
}
func cmpMap(m1, m2 map[string]int) bool {
	for k1, v1 := range m1 {
		if v2, has := m2[k1]; has {
			if v1 != v2 {
				return false
			}
		} else {
			return false
		}
	}
	for k2, v2 := range m2 {
		if v1, has := m1[k2]; has {
			if v1 != v2 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func Test_findSubstring() {

	str := "wordgoodgoodgoodbestword"
	arr := []string{"word", "good", "best", "good"}
	fmt.Println(findSubstring(str, arr))

}
