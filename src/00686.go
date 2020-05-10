package src

import (
	"fmt"
	"strings"
)

/*686. 重复叠加字符串匹配*/
func repeatedStringMatch(A string, B string) int {
	count := 1
	a := A
	for len(A) < len(B) {
		A += a
		count += 1
	}
	if strings.Index(A, B) > -1 {
		return count
	}
	for i := 0; i <= count; i++ {
		A += a
		if strings.Index(A, B) > -1 {
			return count + 1
		}
	}
	return -1
}

func Test_repeatedStringMatch() {
	A := "abc"
	B := "cabcabca"
	s := repeatedStringMatch(A, B)
	fmt.Println(s)
}
