package src

import (
	"math"
	"strings"
)

func myAtoi(str string) int {
	if str == "" {
		return 0
	}

	// 有空格前缀则全部替换
	for strings.HasPrefix(str, " ") {
		str = str[1:]
	}
	if str == "" {
		return 0
	}
	// 正负符号判定
	sign := 1
	for str[0] == 45 || str[0] == 43 {
		if str[0] == 45 {
			sign = -1
			str = str[1:]
			break
		}
		if str[0] == 43 {
			sign = 1
			str = str[1:]
			break
		}
	}

	if str == "" {
		return 0
	}

	for strings.HasPrefix(str, "0") {
		str = str[1:]
	}
	if str == "" {
		return 0
	}

	tmp := []int{}

	for _, v := range str {
		// ASCII 码 ， 57 -> 9 , 48 -> 0
		if v > 57 || v < 48 {
			break
		} else {
			tmp = append(tmp, int(v-48))
		}
	}

	//fmt.Println(tmp)

	// 没有数字返回0
	if len(tmp) == 0 {
		return 0
	}
	// 简单越界处理
	if len(tmp) > 10 || (len(tmp) == 9 && tmp[0] > 2) {
		if sign == 1 {
			return math.MaxInt32
		}
		if sign == -1 {
			return math.MinInt32
		}
	}

	j := len(tmp) - 1
	tmpInt := 0
	for i := 0; i < len(tmp); i++ {
		tmpInt += tmp[i] * pow10(j)
		j--
	}

	if tmpInt > math.MaxInt32 {
		if sign == -1 {
			return math.MinInt32
		} else {
			return math.MaxInt32
		}
	}

	return tmpInt * sign
}

func pow10(i int) int {
	if i == 0 {
		return 1
	}
	rst := 1
	for ; i > 0; i-- {
		rst *= 10
	}
	return rst
}
