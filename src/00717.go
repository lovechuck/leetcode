package src

/*717. 1比特与2比特字符*/

func isOneBitCharacter(bits []int) bool {
	len2 := 0
	for len2 < len(bits)-1 {
		len2 += bits[len2] + 1
	}
	return len2 == len(bits)-1
}
