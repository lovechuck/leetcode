package src

import (
	"fmt"
	"strconv"
)

/*443. 压缩字符串*/
func compress(chars []byte) int {
	if chars == nil {
		return 0
	}
	if len(chars) == 1 {
		return 1
	}

	y := 1
	leng := 1
	for i := 1; i < len(chars); i++ {
		if chars[i] == chars[i-1] {
			leng++
		} else {
			if leng != 1 {
				s := []byte(strconv.Itoa(leng))
				for _, b := range s {
					chars[y] = b
					y++
				}
			}
			chars[y] = chars[i]
			y++
			leng = 1
		}
	}
	if leng != 1 {
		s := []byte(strconv.Itoa(leng))
		for _, b := range s {
			chars[y] = b
			y++
		}
	}
	return y
}

func Test_compress() {
	temp := []byte{
		'a', 'a', 'b', 'b', 'c', 'c', 'c',
	}
	fmt.Printf("%d", compress(temp))
}
