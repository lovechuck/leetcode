package src

/*
给定两个字符串 s 和 t，它们只包含小写字母。

字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。

请找出在 t 中被添加的字母。
*/
func findTheDifference(s string, t string) byte {
	mmap := make(map[byte]int)
	sb := []byte(s)
	tb := []byte(t)
	for _, i := range sb {
		if _, ok := mmap[i]; ok {
			mmap[i] = mmap[i] + 1
		} else {
			mmap[i] = 1
		}
	}
	for _, i := range tb {
		if _, ok := mmap[i]; ok {
			if mmap[i] == 0 {
				return i
			} else {
				mmap[i] = mmap[i] - 1
			}
		} else {
			return i
		}
	}

	return 0
}
