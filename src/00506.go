package src

import "fmt"

/*506. 相对名次*/
func findRelativeRanks(nums []int) []string {
	max := 0
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	rank := make([]int, max+1)
	for i, num := range nums {
		rank[num] = i + 1
	}
	res := make([]string, len(nums))
	index := 0
	for i := len(rank) - 1; i >= 0; i-- {
		if rank[i] > 0 {
			if index == 0 {
				res[rank[i]-1] = "Gold Medal"
			} else if index == 1 {
				res[rank[i]-1] = "Silver Medal"
			} else if index == 2 {
				res[rank[i]-1] = "Bronze Medal"
			} else {
				res[rank[i]-1] = fmt.Sprintf("%d", index+1)
			}
			index++
			if index >= len(nums) {
				break
			}
		}
	}
	return res
}

func Test_findRelativeRanks() {
	nums := []int{10, 3, 8, 9, 4}
	fmt.Sprintln(findRelativeRanks(nums))
}
