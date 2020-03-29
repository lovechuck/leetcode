package src

/*191. 位1的个数*/

func hammingWeight(num uint32) int {
	bit := 0
	var mask uint32 = 1
	for i := 0; i < 32; i++ {
		if (mask & num) == 1 {
			bit++
		}
		num = num >> 1
	}
	return bit
}
