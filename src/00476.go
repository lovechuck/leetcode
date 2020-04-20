package src

/*补数*/
func findComplement(num int) int {
	xor := 1
	for num >= xor {
		xor = xor << 1
	}
	return num ^ (xor - 1)
}
