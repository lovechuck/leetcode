package src

/*706. 设计哈希映射*/
type MyHashMap struct {
	Data map[int]int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{Data: map[int]int{}}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	this.Data[key] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	if v, ok := this.Data[key]; ok {
		return v
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	this.Data[key] = -1
}
