package src

/*541. 反转字符串 II*/
func reverseStr(s string, k int) string {
	ss := []rune(s)
	total := len(ss)
	for i := 0; i < total; i = i + 2*k {
		temp := (total - i)
		if temp < k {
			for i, j := i, total-1; i < j; i, j = i+1, j-1 {
				ss[i], ss[j] = ss[j], ss[i]
			}
		} else if temp < 2*k && temp >= k {
			for i, j := i, i+k-1; i < j; i, j = i+1, j-1 {
				ss[i], ss[j] = ss[j], ss[i]
			}
		} else {
			for i, j := i, i+k-1; i < j; i, j = i+1, j-1 {
				ss[i], ss[j] = ss[j], ss[i]
			}
		}
	}
	return string(ss)
}
