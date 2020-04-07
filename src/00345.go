package src

/*345. 反转字符串中的元音字母*/
func reverseVowels(s string) string {
	arr := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	mmap := make(map[rune]bool)
	for i := 0; i < len(arr); i++ {
		mmap[arr[i]] = true
	}

	str := []rune(s)
	left := 0
	right := len(str) - 1
	for left < right {
		for left < right {
			if _, ok := mmap[str[left]]; !ok {
				left++
			} else {
				break
			}
		}
		for left < right {
			if _, ok := mmap[str[right]]; !ok {
				right--
			} else {
				break
			}
		}
		if left < right {
			temp := str[right]
			str[right] = str[left]
			str[left] = temp
			left++
			right--
		}
	}

	return string(str)
}

func Test_reverseVowels() {
	reverseVowels(".,")
}
