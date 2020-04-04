package src

import (
	"strconv"
	"strings"
)

/**/

func isSymmetry(num int) bool {
	mmap := map[string]string{
		"0": "0",
		"1": "1",
		"8": "8",
		"6": "9",
		"9": "6",
	}

	words := strings.Split(strconv.Itoa(num), "")
	var str []string
	for i := len(words)/2 - 1; i >= 0; i-- {
		if _, ok := mmap[words[i]]; !ok {
			return false
		}
		str = append(str, mmap[words[i]])
	}
	for i := 0; i < len(words)/2; i++ {
		if str[i] != words[i] {
			return false
		}
	}

	return true
}
