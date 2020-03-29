package src

/*171. Excel表列序号*/

func titleToNumber(s string) int {
	ss := 0
	for _, v := range s {
		ss = ss*26 + int(v-'A'+1)
	}

	return ss
}
