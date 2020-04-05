package src

/*303. 区域和检索 - 数组不可变*/

type NumArray struct {
	data []int
}

func Constructor(nums []int) NumArray {
	return NumArray{data: nums}
}

func (this *NumArray) SumRange(i int, j int) int {
	sum := 0
	for k := i; k <= j; k++ {
		sum += this.data[k]
	}
	return sum
}

func Test_NumArray() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	sum := Constructor(nums)
	sum.SumRange(0, 2)
}
