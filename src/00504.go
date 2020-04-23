package src

import "fmt"

/*504. 七进制数*/
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	res := ""
	prefix := ""
	if num < 0 {
		prefix = "-"
		num = -num
	}
	for num > 0 {
		temp := num % 7
		res = fmt.Sprintf("%d%s", temp, res)
		num = num / 7
	}
	return prefix + res
}
