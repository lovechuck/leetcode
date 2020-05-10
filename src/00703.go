package src

import (
	"sort"
)

/*
703. 数据流中的第K大元素
*/
type KthLargest struct {
	K    int
	data []int
}

func Constructor_KthLargest(k int, nums []int) KthLargest {
	sort.Ints(nums)
	return KthLargest{
		K:    k,
		data: nums,
	}
}

func (this *KthLargest) Add(val int) int {
	if this.data == nil || len(this.data) == 0 {
		this.data = append(this.data, val)
		return val
	}
	this.data = append(this.data, val)
	for k := 0; k < this.K; k++ {
		last := len(this.data) - 1 - k
		if last-1 >= 0 && this.data[last] < this.data[last-1] {
			swag := this.data[last]
			this.data[last] = this.data[last-1]
			this.data[last-1] = swag
		} else {
			break
		}
	}
	return this.data[len(this.data)-this.K]
}
