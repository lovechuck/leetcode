package src

import (
	"math"
	"unicode"
)

/*748. 最短完整词*/

func shortestCompletingWord(licensePlate string, words []string) string {
	result := ""
	leng := math.MaxInt32
	res, sum := countlicensePlate(licensePlate)
	for _, word := range words {
		if checklicensePlate(res, sum, word) {
			if len(word) < leng {
				leng = len(word)
				result = word
			}
		}
	}
	return result
}

func checklicensePlate(res map[rune]int, sum int, word string) bool {
	temp, ts := countlicensePlate(word)
	if ts < sum {
		return false
	}
	for r, i := range res {
		if temp[r] < i {
			return false
		}
	}
	return true
}

func countlicensePlate(licensePlate string) (map[rune]int, int) {
	res := make(map[rune]int)
	sum := 0
	for _, i := range licensePlate {
		if unicode.IsLetter(i) {
			key := unicode.ToLower(i)
			if _, ok := res[key]; ok {
				res[key] = res[key] + 1
			} else {
				res[key] = 1
			}
			sum++
		}
	}
	return res, sum
}

func Test_shortestCompletingWord() {
	shortestCompletingWord("1s3 PSt", []string{"step", "steps", "stripe", "stepple"})
}
