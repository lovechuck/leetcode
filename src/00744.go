package src

/*744. 寻找比目标字母大的最小字母*/
func nextGreatestLetter(letters []byte, target byte) byte {
	if target < letters[0] || target > letters[len(letters)-1] {
		return letters[0]
	}
	for _, letter := range letters {
		if letter > target {
			return letter
		}
	}
	return letters[0]
}
