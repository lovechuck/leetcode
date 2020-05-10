package src

/*303. 区域和检索 - 数组不可变*/

type NumArray struct {
	data []int
}

func Constructor_NumArray(nums []int) NumArray {
	return NumArray{data: nums}
}

func (this *NumArray) SumRange(i int, j int) int {
	sum := 0
	for k := i; k <= j; k++ {
		sum += this.data[k]
	}
	return sum
}
