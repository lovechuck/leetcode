package src

/*344. 反转字符串*/

func reverseString(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		temp := s[right]
		s[right] = s[left]
		s[left] = temp
		left++
		right--
	}
}
