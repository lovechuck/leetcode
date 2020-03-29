package src

/*190. 颠倒二进制位*/
func reverseBits(num uint32) uint32 {
	var result uint32 = 0
	for i := 0; i < 32; i++ {
		result = result*2 + num%2
		num = num / 2
	}
	return result
}
