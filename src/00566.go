package src

import "fmt"

/*566. 重塑矩阵*/

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if nums == nil {
		return nil
	}
	a := len(nums)
	b := len(nums[0])

	if a*b != r*c {
		return nums
	}
	var res [][]int
	j := 0
	var temp []int
	for _, num := range nums {
		for _, r := range num {
			if j == 0 {
				temp = []int{}
			}
			temp = append(temp, r)
			j++
			if j == c {
				j = 0
				res = append(res, temp)
			}
		}
	}

	return res
}

func Test_matrixReshape() {
	nums := [][]int{
		[]int{1, 2},
		[]int{3, 4},
	}
	res := matrixReshape(nums, 1, 4)
	fmt.Println(res)
}
