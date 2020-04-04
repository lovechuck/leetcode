package src

/*225. 用队列实现栈*/

type MyStack struct {
	data []int
	len  int
}

/** Initialize your data structure here. */
func Constructor2() MyStack {
	return MyStack{
		data: nil,
		len:  0,
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	if this.len == 0 {
		this.data = []int{}
		this.data = append(this.data, x)
	} else {
		this.data = append(this.data, x)
		for i := 0; i < this.len; i++ {
			this.data = append(this.data, this.data[i])
		}
		this.data = this.data[this.len:]
	}
	this.len++
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	top := this.data[0]
	this.data = this.data[1:]
	this.len--
	return top
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.data[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.len == 0
}
