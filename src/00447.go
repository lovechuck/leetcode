package src

/*447. 回旋镖的数量*/
func numberOfBoomerangs(points [][]int) int {
	sum := 0
	for i := 0; i < len(points); i++ {
		count := 0
		mmap := make(map[int]int)
		for j := 0; j < len(points); j++ {
			temp := dis(points[i], points[j])
			if _, ok := mmap[temp]; ok {
				mmap[temp] = mmap[temp] + 1
			} else {
				mmap[temp] = 1
			}
		}
		for _, item := range mmap {
			if item != 1 {
				count += item * (item - 1)
			}
		}
		sum += count
	}
	return sum
}

func dis(a []int, b []int) int {
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += (b[i] - a[i]) * (b[i] - a[i])
	}
	return sum
}
