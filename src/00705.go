package src

/*705. 设计哈希集合*/
type MyHashSet1 struct {
	Data map[int]bool
}

/** Initialize your data structure here. */
func Constructor_MyHashSet() MyHashSet1 {
	return MyHashSet1{Data: map[int]bool{}}
}

func (this *MyHashSet1) Add(key int) {
	this.Data[key] = true
}

func (this *MyHashSet1) Remove(key int) {
	this.Data[key] = false
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet1) Contains(key int) bool {
	return this.Data[key]
}
