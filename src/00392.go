package src

/*392. 判断子序列*/

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	i := 0
	j := 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			j++
			if i == len(s) {
				return true
			}
		} else {
			j++
		}
	}

	return false
}
