package src

func strStr(haystack string, needle string) int {
	var i = 0
	var j = 0
	for i < len(haystack) {
		if j == len(needle) {
			return i - len(needle)
		}
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}
	if j == len(needle) {
		return i - len(needle)
	}

	return -1
}
