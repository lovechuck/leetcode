package src

import "sort"

/*
1005. K 次取反后最大化的数组和

给定一个整数数组 A，我们只能用以下方法修改该数组：我们选择某个个索引 i 并将 A[i] 替换为 -A[i]，然后总共重复这个过程 K 次。（我们可以多次选择同一个索引 i。）

以这种方式修改数组后，返回数组可能的最大和。
*/

func largestSumAfterKNegations(A []int, K int) int {
	sort.Ints(A)
	if A[0] >= 0 {
		if K%2 == 0 {
			sum := 0
			for _, i := range A {
				sum += i
			}
			return sum
		} else {
			sum := 0
			for _, i := range A {
				sum += i
			}
			return sum - 2*A[0]
		}
	} else {
		j := 0
		i := 0
		for ; i < K; i++ {
			if j < len(A) && A[j] <= 0 {
				A[j] = -A[j]
				j++
			} else {
				break
			}
		}
		if i == K {
			sum := 0
			for _, i := range A {
				sum += i
			}
			return sum
		} else {
			if (K-j)%2 == 0 {
				sum := 0
				for _, i := range A {
					sum += i
				}
				return sum
			} else {
				if j >= len(A) {
					sum := 0
					for _, i := range A {
						sum += i
					}
					return sum - 2*A[j-1]
				} else {
					sum := 0
					for _, i := range A {
						sum += i
					}
					if A[j-1] < A[j] {
						return sum - 2*A[j-1]
					} else {
						return sum - 2*A[j]
					}
				}
			}
		}
	}
}

func Test_largestSumAfterKNegations() {
	largestSumAfterKNegations([]int{4, 2, 3}, 1)
}
