package src

/*
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

案例:

s = "leetcode"
返回 0.

s = "loveleetcode",
返回 2.
*/
func firstUniqChar(s string) int {
	mmap := make(map[int32]int)
	for _, i := range s {
		if _, ok := mmap[i]; ok {
			mmap[i] = mmap[i] + 1
		} else {
			mmap[i] = 1
		}
	}
	for k, i := range s {
		if _, ok := mmap[i]; ok {
			if mmap[i] == 1 {
				return k
			}
		}
	}

	return -1
}
