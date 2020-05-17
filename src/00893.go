package src

/*
你将得到一个字符串数组 A。

如果经过任意次数的移动，S == T，那么两个字符串 S 和 T 是特殊等价的。

一次移动包括选择两个索引 i 和 j，且 i ％ 2 == j ％ 2，交换 S[j] 和 S [i]。

现在规定，A 中的特殊等价字符串组是 A 的非空子集 S，这样不在 S 中的任何字符串与 S 中的任何字符串都不是特殊等价的。

返回 A 中特殊等价字符串组的数量。
*/

/*
理解
1. A 字符串数组
2. A 中的每个元素 通过交换相等的为一组
3. 交换方式为  i ％ 2 == j ％ 2，每个元素的奇偶位数交换
4. 返回 组的个数
*/

func numSpecialEquivGroups(A []string) int {
	set := make(map[[52]int]bool)
	for _, v := range A {
		set[getKey(v)] = true
	}
	return len(set)
}

func getKey(s string) [52]int {
	res := [52]int{}
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			res[s[i]-'a']++
		} else {
			res[s[i]-'a'+26]++
		}
	}

	return res
}
