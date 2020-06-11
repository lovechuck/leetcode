package src

/*
1047. 删除字符串中的所有相邻重复项

给出由小写字母组成的字符串 S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。

在 S 上反复执行重复项删除操作，直到无法继续删除。

在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。

栈 更合适
*/
func removeDuplicates(S string) string {
	if len(S) == 1 {
		return S
	}
	var ans []byte
	i := 0
	for i < len(S)-1 {
		if S[i] != S[i+1] {
			ans = append(ans, S[i])
			i++
		} else {
			i += 2
		}
	}
	if i == len(S)-1 {
		ans = append(ans, S[i])
	}
	if len(ans) == len(S) {
		return S
	}
	return removeDuplicates(string(ans))
}

func Test_removeDuplicates() {
	removeDuplicates("azxxzy")
}
