package src

/*
961. 重复 N 次的元素

在大小为 2N 的数组 A 中有 N+1 个不同的元素，其中有一个元素重复了 N 次。

返回重复了 N 次的那个元素。

示例 1：

输入：[1,2,3,3]
输出：3
*/
func repeatedNTimes(A []int) int {
	hash := make(map[int]int)
	for _, i := range A {
		if _, ok := hash[i]; ok {
			hash[i] = hash[i] + 1
			if hash[i] >= len(A)/2 {
				return i
			}
		} else {
			hash[i] = 1
		}
	}
	return -1
}
