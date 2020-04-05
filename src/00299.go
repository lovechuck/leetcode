package src

import (
	"fmt"
)

/*299. 猜数字游戏*/

func getHint(secret string, guess string) string {
	ac := 0
	bc := 0
	var b []int
	mmap := make(map[uint8]int)
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			ac++
		} else {
			if _, ok := mmap[secret[i]]; ok {
				mmap[secret[i]] = mmap[secret[i]] + 1
			} else {
				mmap[secret[i]] = 1
			}
			b = append(b, i)
		}
	}
	for i := 0; i < len(b); i++ {
		if _, ok := mmap[guess[b[i]]]; ok {
			if mmap[guess[b[i]]] != 0 {
				mmap[guess[b[i]]] = mmap[guess[b[i]]] - 1
				bc++
			}
		}
	}

	return fmt.Sprintf("%dA%dB", ac, bc)
}
