package src

import "math"

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	x0 := float64(x)
	x1 := (float64(x0) + float64(x)/x0) / 2

	for math.Abs(x0-x1) >= 1 {
		x0 = x1
		x1 = (x0 + float64(x)/x0) / 2
	}

	return int(x1)
}
