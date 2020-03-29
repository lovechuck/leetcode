package src

/*170 两数之和 III - 数据结构设计
 */

type two struct {
	i    int
	data []int
}

func (t *two) add(num int) {
	t.data = append(t.data, num)
	t.i++
}

func (t *two) find(target int) bool {
	mmap := make(map[int]int)
	for i := 0; i < t.i; i++ {
		left := target - t.data[i]
		if _, ok := mmap[left]; ok {
			return true
		}
		mmap[t.data[i]] = i
	}

	return false
}
