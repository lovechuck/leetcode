package src

/*693. 交替位二进制数*/

func hasAlternatingBits(n int) bool {
	temp := n ^ (n >> 1)
	return temp&(temp+1) == 0
}
