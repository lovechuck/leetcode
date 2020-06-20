package src

/*
1071. 字符串的最大公因子
*/

func gcd1(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	num := gcd1(len(str1), len(str2))
	return str1[0:num]
}
