package src

import "fmt"

/*
696. 计数二进制子串
给定一个字符串 s，计算具有相同数量0和1的非空(连续)子字符串的数量，并且这些子字符串中的所有0和所有1都是组合在一起的。

重复出现的子串要计算它们出现的次数。
*/

func countBinarySubstrings(s string) int {
	if s == "" {
		return 0
	}
	var cur int32
	lenCur := 0
	lenPre := 0
	sum := 0
	for i, k := range s {
		if i == 0 {
			cur = k
		}
		if cur == k {
			lenCur++
		} else {
			if lenPre != 0 {
				if lenPre < lenCur {
					sum += lenPre
				} else {
					sum += lenCur
				}
			}
			lenPre = lenCur
			cur = k
			lenCur = 1
		}
	}
	if lenPre != 0 {
		if lenPre < lenCur {
			sum += lenPre
		} else {
			sum += lenCur
		}
	}
	return sum
}

func Test_countBinarySubstrings() {
	fmt.Println(countBinarySubstrings("00110"))
}
