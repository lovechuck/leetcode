package src

import "math"

/*812. 最大三角形面积*/

func largestTriangleArea(points [][]int) float64 {
	n := len(points)
	var ans float64 = 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				temp := area(points[i], points[j], points[k])
				if ans < temp {
					ans = temp
				}
			}
		}
	}
	return ans
}
func area(P []int, Q []int, R []int) float64 {
	return 0.5 * math.Abs(float64(P[0]*Q[1]+Q[0]*R[1]+R[0]*P[1]-P[1]*Q[0]-Q[1]*R[0]-R[1]*P[0]))
}
