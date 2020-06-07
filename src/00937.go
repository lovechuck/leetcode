package src

import (
	"sort"
	"strings"
	"unicode"
)

/*
937. 重新排列日志文件

你有一个日志数组 logs。每条日志都是以空格分隔的字串。

对于每条日志，其第一个字为字母与数字混合的 标识符。

除标识符之外，所有字均由小写字母组成的，称为 字母日志
除标识符之外，所有字均由数字组成的，称为 数字日志
题目所用数据保证每个日志在其标识符后面至少有一个字。

请按下述规则将日志重新排序：

所有 字母日志 都排在 数字日志 之前。
字母日志 在内容不同时，忽略标识符后，按内容字母顺序排序；在内容相同时，按标识符排序；
数字日志 应该按原来的顺序排列。
返回日志的最终顺序。
*/

func leter(s string) (bool, string, string) {
	i := strings.Index(s, " ")
	prefix := s[0:i]
	suffix := s[i:]
	for _, i := range suffix {
		if i != ' ' && unicode.IsLetter(i) {
			return true, prefix, suffix
		}
	}
	return false, prefix, suffix
}

func reorderLogFiles(logs []string) []string {
	var leters []string
	var nums []string
	for _, log := range logs {
		isok, _, _ := leter(log)
		if isok {
			leters = append(leters, log)
		} else {
			nums = append(nums, log)
		}
	}
	sort.Slice(leters, func(j, i int) bool {
		_, b, c := leter(leters[i])
		_, e, f := leter(leters[j])

		if c == f {
			return b < e
		} else {
			return c < f
		}
	})
	for _, num := range nums {
		leters = append(leters, num)
	}
	return leters
}

func Test_reorderLogFiles() {
	arr := []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}
	reorderLogFiles(arr)
}
