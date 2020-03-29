package src

/*155. 最小栈*/

type MinStack struct {
	len     int
	data    []int
	minData []int
}

/** initialize your data structure here. */
func Constructor1() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if this.len == 0 {
		this.minData = append(this.minData, x)
	} else {
		if this.minData[this.len-1] >= x {
			this.minData = append(this.minData, x)
		} else {
			this.minData = append(this.minData, this.minData[this.len-1])
		}
	}
	this.len++
}

func (this *MinStack) Pop() {
	if this.len == 1 {
		this.data = nil
		this.minData = nil
	} else {
		this.data = this.data[0 : this.len-1]
		this.minData = this.minData[0 : this.len-1]
	}
	this.len--
}

func (this *MinStack) Top() int {
	return this.data[this.len-1]
}

func (this *MinStack) GetMin() int {
	return this.minData[this.len-1]
}
