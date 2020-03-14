package src

import "fmt"

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	rows := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		temp := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 {
				temp[j] = 1
			} else if j == i {
				temp[j] = 1
			} else {
				temp[j] = rows[i-1][j-1] + rows[i-1][j]
			}
		}
		rows[i] = temp
	}
	return rows
}

func Test_generate() {
	arr := generate(5)
	fmt.Print(arr)
}
