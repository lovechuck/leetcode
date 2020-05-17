package src

import "unicode"

/*
917. 仅仅反转字母
*/
func reverseOnlyLetters(S string) string {
	arr := []rune(S)
	left := 0
	right := len(arr) - 1
	for left < right {
		for !unicode.IsLetter(arr[left]) && left < right {
			left++
		}
		for !unicode.IsLetter(arr[right]) && left < right {
			right--
		}
		temp := arr[right]
		arr[right] = arr[left]
		arr[left] = temp
		left++
		right--
	}
	return string(arr)
}
