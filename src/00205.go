package src

/*205. 同构字符串*/

func isIsomorphic(s string, t string) bool {
	a := getIsomorphic(s)
	b := getIsomorphic(t)
	return a == b
}

func getIsomorphic(s string) string {
	result := ""
	current := 0
	mmap := make(map[int32]string)
	for _, r := range s {
		if _, ok := mmap[r]; ok {
			result += mmap[r]
		} else {
			current++
			mmap[r] = string(current)
			result += mmap[r]
		}
	}

	return result
}
