package src

/*
1089. 复写零


给你一个长度固定的整数数组 arr，请你将该数组中出现的每个零都复写一遍，并将其余的元素向右平移。

注意：请不要在超过该数组长度的位置写入元素。

要求：请对输入的数组 就地 进行上述修改，不要从函数返回任何东西。
*/

func duplicateZeros(arr []int) {
	al := len(arr)
	sum := 0
	i := 0
	more := 0
	for ; i < al; i++ {
		if arr[i] == 0 {
			sum += 2
		} else {
			sum++
		}
		if sum == al {
			break
		} else if sum > al {
			more = 1
			break
		}
	}
	last := len(arr) - 1
	for j := i; j >= 0; j-- {
		if arr[j] != 0 {
			arr[last] = arr[j]
			last--
			if last < 0 {
				break
			}
		} else {
			arr[last] = 0
			last--
			if last < 0 {
				break
			}
			if more == 0 {
				arr[last] = 0
				last--
				if last < 0 {
					break
				}
			} else {
				more = 0
			}

		}
	}
}

func Test_duplicateZeros() {
	duplicateZeros([]int{8, 4, 5, 0, 0, 0, 0, 7})
}
