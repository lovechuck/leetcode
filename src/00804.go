package src

/*804. 唯一摩尔斯密码词*/
func uniqueMorseRepresentations(words []string) int {
	data := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	hash := make(map[string]int)
	for _, word := range words {
		temp := ""
		for _, k := range word {
			i := k - 'a'
			temp += data[i]
		}
		hash[temp] = 1
	}
	return len(hash)
}
