package src

/*

有 N 个花园，按从 1 到 N 标记。在每个花园中，你打算种下四种花之一。

paths[i] = [x, y] 描述了花园 x 到花园 y 的双向路径。

另外，没有花园有 3 条以上的路径可以进入或者离开。

你需要为每个花园选择一种花，使得通过路径相连的任何两个花园中的花的种类互不相同。

以数组形式返回选择的方案作为答案 answer，其中 answer[i] 为在第 (i+1) 个花园中种植的花的种类。花的种类用  1, 2, 3, 4 表示。保证存在答案。
*/
func gardenNoAdj(N int, paths [][]int) []int {
	ans := make([]int, N)
	hash := make(map[int][]int)
	for _, path := range paths {
		hash[path[0]] = append(hash[path[0]], path[1])
		hash[path[1]] = append(hash[path[1]], path[0])
	}
	for i := 0; i < N; i++ {
		temp := []int{1, 2, 3, 4}
		for _, t := range hash[i+1] {
			if ans[t-1] != 0 {
				temp[ans[t-1]-1] = 0
			}
		}
		for _, i3 := range temp {
			if i3 != 0 {
				ans[i] = i3
				break
			}
		}
	}
	return ans
}

func Test_gardenNoAdj() {
	gardenNoAdj(4, [][]int{[]int{1, 2}, []int{3, 4}})
}
