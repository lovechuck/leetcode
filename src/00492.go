package src

import "math"

/*492. 构造矩形*/

func constructRectangle(area int) []int {
	w := math.Sqrt(float64(area))
	k := int(w)
	j := 1
	i := k
	for ; i > 0; i-- {
		if area%i == 0 {
			j = area / i
			break
		}
	}
	if i < j {
		return []int{j, i}
	}
	return []int{i, j}
}
