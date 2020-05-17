package src

/*
830. 较大分组的位置

在一个由小写字母构成的字符串 S 中，包含由一些连续的相同字符所构成的分组。

例如，在字符串 S = "abbxxxxzyy" 中，就含有 "a", "bb", "xxxx", "z" 和 "yy" 这样的一些分组。

我们称所有包含大于或等于三个连续字符的分组为较大分组。找到每一个较大分组的起始和终止位置。

最终结果按照字典顺序输出。
*/
func largeGroupPositions(S string) [][]int {
	var res [][]int
	var pre rune = -1
	preIndex := -1
	count := 1
	for index, i := range S {
		if i != pre {
			if count >= 3 {
				temp := []int{preIndex, index - 1}
				res = append(res, temp)
			}
			count = 1
			pre = i
			preIndex = index
		} else {
			count++
		}
	}
	if count >= 3 {
		temp := []int{preIndex, len(S) - 1}
		res = append(res, temp)
	}
	return res
}
