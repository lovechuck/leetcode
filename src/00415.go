package src

import (
	"fmt"
	"strconv"
	"strings"
)

/*
415. 字符串相加
给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

注意：

num1 和num2 的长度都小于 5100.
num1 和num2 都只包含数字 0-9.
num1 和num2 都不包含任何前导零。
你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。
*/
func addStrings(num1 string, num2 string) string {
	min := num1
	max := num2
	if len(num1) > len(num2) {
		min = num2
		max = num1
	}
	if min == "" {
		return max
	}

	min_nums := strings.Split(min, "")
	max_nums := strings.Split(max, "")
	jinwei := 0
	res := ""
	for i := len(min_nums) - 1; i >= 0; i-- {
		mini, _ := strconv.Atoi(min_nums[i])
		maxi, _ := strconv.Atoi(max_nums[len(max_nums)-len(min_nums)+i])
		temp := mini + maxi + jinwei
		if temp > 9 {
			res = fmt.Sprintf("%d", temp-10) + res
			jinwei = 1
		} else {
			res = fmt.Sprintf("%d", temp) + res
			jinwei = 0
		}
	}
	for i := len(max_nums) - len(min_nums) - 1; i >= 0; i-- {
		maxi, _ := strconv.Atoi(max_nums[i])
		temp := maxi + jinwei
		if temp > 9 {
			res = fmt.Sprintf("%d", temp-10) + res
			jinwei = 1
		} else {
			res = fmt.Sprintf("%d", temp) + res
			jinwei = 0
		}
	}
	if jinwei == 1 {
		res = fmt.Sprintf("%d", jinwei) + res
	}
	return res
}

func Test_addStrings() {
	fmt.Print(addStrings("408", "5"))
}
