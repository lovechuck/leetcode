package src

import (
	"fmt"
	"strconv"
)

/*67. 二进制求和*/

func addBinary(a string, b string) string {
	alen := len(a)
	blen := len(b)
	maxlen := alen
	minlen := blen
	if alen < blen {
		maxlen = blen
		minlen = alen
		tep := a
		a = b
		b = tep
	}
	result := make([]int, maxlen+1)
	add := 0
	for i := minlen - 1; i >= 0; i-- {
		ai := 0
		if a[i+maxlen-minlen] == '1' {
			ai = 1
		}
		bi := 0
		if b[i] == '1' {
			bi = 1
		}
		if ai+bi+add >= 2 {
			result[i+maxlen-minlen+1] = ai + bi + add - 2
			add = 1
		} else {
			result[i+maxlen-minlen+1] = ai + bi + add
			add = 0
		}
	}
	for i := maxlen - minlen - 1; i >= 0; i-- {
		ai := 0
		if a[i] == '1' {
			ai = 1
		}
		if ai+add >= 2 {
			result[i+1] = ai + add - 2
			add = 1
		} else {
			result[i+1] = ai + add
			add = 0
		}
	}
	if add == 1 {
		result[0] = 1
		return convertString(result)
	}

	return convertString(result[1:])
}

func convertString(a []int) string {
	s := ""
	for i := 0; i < len(a); i++ {
		s += strconv.Itoa(a[i])
	}
	return s
}

func Test_addBinary() {
	fmt.Println(addBinary("100", "110010"))
	fmt.Println(addBinary("11", "1"))
}
