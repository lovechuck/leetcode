package src

/*
给定一个赎金信 (ransom) 字符串和一个杂志(magazine)字符串，判断第一个字符串 ransom 能不能由第二个字符串 magazines 里面的字符构成。如果可以构成，返回 true ；否则返回 false。

(题目说明：为了不暴露赎金信字迹，要从杂志上搜索各个需要的字母，组成单词来表达意思。杂志字符串中的每个字符只能在赎金信字符串中使用一次。)
*/

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	mmap := make(map[int32]int)
	for _, i := range magazine {
		if _, ok := mmap[i]; ok {
			mmap[i] = mmap[i] + 1
		} else {
			mmap[i] = 1
		}
	}

	for _, i := range ransomNote {
		if _, ok := mmap[i]; ok {
			if mmap[i] > 0 {
				mmap[i] = mmap[i] - 1
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
