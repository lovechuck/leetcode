package src

import "fmt"

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。

*/

func letterCombinations(digits string) []string {
	mmap := make(map[rune][]string)
	mmap['2'] = []string{"a", "b", "c"}
	mmap['3'] = []string{"d", "e", "f"}
	mmap['4'] = []string{"g", "h", "i"}
	mmap['5'] = []string{"j", "k", "l"}
	mmap['6'] = []string{"m", "n", "o"}
	mmap['7'] = []string{"p", "q", "r", "s"}
	mmap['8'] = []string{"t", "u", "v"}
	mmap['9'] = []string{"w", "x", "y", "z"}
	mmap['0'] = []string{}
	mmap['1'] = []string{}

	var result []string
	if len(digits) == 0 {
		return result
	}

	for _, k := range digits {
		result = combinations(result, mmap[k])
	}

	return result
}

func combinations(before []string, letters []string) []string {
	if letters == nil || len(letters) == 0 {
		return before
	}

	if before == nil || len(before) == 0 {
		return letters
	}

	var after []string
	for _, s := range before {
		for _, letter := range letters {
			after = append(after, s+letter)
		}
	}

	return after
}

func Test_letterCombinations() {
	resp := letterCombinations("23")
	for _, s := range resp {
		fmt.Sprintln(s)
	}
}
