package src

import (
	"fmt"
	"strconv"
	"strings"
)

/*
38. 外观数列
*/

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}

	before := "11"
	for i := 3; i <= n; i++ {
		arr := strings.Split(before, "")
		current := ""
		var tmpcount int64 = 1
		for j := 0; j < len(arr)-1; j++ {
			temp := arr[j]
			if arr[j] == arr[j+1] {
				tmpcount++
			} else {
				current += strconv.FormatInt(tmpcount, 10) + temp
				tmpcount = 1
			}
		}
		current += strconv.FormatInt(tmpcount, 10) + arr[len(arr)-1]
		before = current
	}
	return before
}

func Test_countAndSay() {
	fmt.Println(countAndSay(4))
	fmt.Println(countAndSay(5))
}
