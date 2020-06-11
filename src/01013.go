package src

/*
1013. 将数组分成和相等的三个部分
*/
func canThreePartsEqualSum(A []int) bool {
	sum := 0
	for _, i := range A {
		sum += i
	}
	if sum%3 != 0 {
		return false
	}
	k3 := sum / 3
	haha := A[0]
	temp := 0
	i := 1
	for i < len(A) {
		if haha != k3 {
			haha += A[i]
		} else {
			temp++
			if temp == 2 {
				break
			}
			haha = A[i]
		}
		i++
	}
	if i == len(A) {
		return false
	}
	sum = 0
	for i < len(A) {
		sum += A[i]
		i++
	}
	if sum == k3 {
		return true
	}
	return false
}

func Test_canThreePartsEqualSum() {
	canThreePartsEqualSum([]int{1, -1, 1, -1})
}
