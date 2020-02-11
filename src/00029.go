package src

func divide(dividend int, divisor int) int {
	res, c := 0, 1
	var flag bool
	switch {
	case dividend > 0 && divisor < 0:
		{
			divisor = -divisor
			flag = false
		}
	case dividend < 0 && divisor > 0:
		{

			dividend = -dividend
			flag = false

		}
	case dividend < 0 && divisor < 0:
		{

			dividend = -dividend
			divisor = -divisor
			flag = true

		}
	default:
		flag = true
	}
	for dividend >= divisor<<1 {
		divisor <<= 1
		c <<= 1
	}
	for c > 0 && dividend > 0 {
		if dividend >= divisor {
			dividend -= divisor
			res += c
		}
		c >>= 1
		divisor >>= 1
	}
	if !flag {
		res = -res
	}
	if !(res >= int(-0x80000000) && res <= int(0x7FFFFFFF)) {
		return int(0x7FFFFFFF)
	}
	return res

}
