package src

import (
	"fmt"
	"unicode"
)

/*125. 验证回文串*/

func isPalindrome2(s string) bool {
	i := 0
	j := len(s) - 1
	for i <= j {
		vi, vj := rune(s[i]), rune(s[j])
		if !isValidD(vi) {
			i++
			continue
		}
		if !isValidD(vj) {
			j--
			continue
		}
		if unicode.ToLower(vi) != unicode.ToLower(vj) {
			return false
		} else {
			i++
			j--
			continue
		}
	}

	return true
}

func isValidD(r rune) bool {
	return unicode.IsDigit(r) || unicode.IsLetter(r)
}

func Test_isPalindrome2() {
	fmt.Println(isPalindrome2("A man, a plan, a canal: Panama"))
}
