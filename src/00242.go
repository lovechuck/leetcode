package src

/*242 有效的字母异位词  */

func isAnagram(s string, t string) bool {
	mmap := make(map[int32]int)
	for _, i := range s {
		if _, ok := mmap[i]; ok {
			mmap[i] = mmap[i] + 1
		} else {
			mmap[i] = 1
		}
	}
	for _, i := range t {
		if _, ok := mmap[i]; ok {
			mmap[i] = mmap[i] - 1
		} else {
			return false
		}
	}
	for _, v := range mmap {
		if v != 0 {
			return false
		}
	}
	return true
}
