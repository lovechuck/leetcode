package src

/*
859. 亲密字符串

给定两个由小写字母构成的字符串 A 和 B ，只要我们可以通过交换 A 中的两个字母得到与 B 相等的结果，就返回 true ；否则返回 false 。
*/

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	if A == B {
		count := make([]int, 26)
		for _, i := range A {
			count[i-'a']++
			if count[i-'a'] > 1 {
				return true
			}
		}
		return false
	}
	n := len(A)
	num := 0
	var sumA uint8 = 0
	var sumB uint8 = 0
	for i := 0; i < n; i++ {
		if A[i] != B[i] {
			num++
			sumA += A[i]
			sumB += B[i]
			if num > 2 {
				return false
			}
		}
	}
	if sumA == sumB {
		return true
	}

	return false
}
