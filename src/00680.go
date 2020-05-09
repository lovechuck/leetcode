package src

/*680. 验证回文字符串 Ⅱ*/

func validPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	res := true
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			res = false
			break
		}
	}
	if res {
		return res
	}

	res = true
	start := i + 1
	end := j
	for start < end {
		if s[start] == s[end] {
			start++
			end--
		} else {
			res = false
			break
		}
	}
	if res {
		return res
	}
	res = true
	start = i
	end = j - 1
	for start < end {
		if s[start] == s[end] {
			start++
			end--
		} else {
			return false
		}
	}
	return true
}
