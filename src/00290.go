package src

import "strings"

/*290. 单词规律*/

func wordPattern(pattern string, str string) bool {
	prr := strings.Split(pattern, "")
	arr := strings.Split(str, " ")
	if len(arr) != len(prr) {
		return false
	}
	pmap := make(map[string]string)
	amap := make(map[string]string)
	for i := 0; i < len(arr); i++ {
		if _, ok := pmap[prr[i]]; ok {
			if pmap[prr[i]] != arr[i] {
				return false
			}
		} else {
			pmap[prr[i]] = arr[i]
		}
		if _, ok := amap[arr[i]]; ok {
			if amap[arr[i]] != prr[i] {
				return false
			}
		} else {
			amap[arr[i]] = prr[i]
		}
	}

	return true
}
