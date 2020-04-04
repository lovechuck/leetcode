package src

/*266. 回文排列  */

/*
描述
给定一个字符串，判断字符串是否存在一个排列是回文排列。

样例

给定s = "code", 返回 False.
给定s = "aab", 返回 True.
给定s = "carerac", 返回 True.
*/

func isOrNot(code string) bool {
	mmap := make(map[int32]int)
	for _, i := range code {
		if _, ok := mmap[i]; ok {
			mmap[i] = mmap[i] + 1
		} else {
			mmap[i] = 1
		}
	}
	for _, i := range mmap {
		if i == 1 {
			return false
		}
	}

	return true
}
