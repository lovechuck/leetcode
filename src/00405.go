package src

import "fmt"

/*405. 数字转换为十六进制数*/

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	if num < 0 { // 如果是负数，转换成补码形式，比如：-1 + 4294967296 = 4294967295 = 0xffffffff
		num += 4294967296
	}
	var ans string
	hash := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	for num != 0 { // 迭代提取每个位上的十六进制字符
		t := num % 16
		num /= 16
		ans = fmt.Sprintf("%s%s", hash[t], ans)
	}
	return ans
}
