package src

/*
844. 比较含退格的字符串
*/

func backspaceCompare(S string, T string) bool {
	if S == T {
		return true
	}
	return remove(S) == remove(T)
}

func remove(str string) string {
	var res []rune
	for _, i := range str {
		if i == '#' {
			if res != nil && len(res) > 0 {
				res = res[:len(res)-1]
			}
		} else {
			res = append(res, i)
		}
	}
	return string(res)
}
