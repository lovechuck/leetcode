package src

import (
	"fmt"
	"math"
)

/*
 530. 二叉搜索树的最小绝对差
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var pre int = -1
	var res int = math.MaxInt32

	helperT(root, &pre, &res)
	return res
}

func helperT(root *TreeNode, pre *int, res *int) {
	if root == nil {
		return
	}
	helperT(root.Left, pre, res)
	if *pre != -1 {
		*res = int(math.Min(float64(*res), math.Abs(float64(root.Val-*pre))))
	}
	*pre = root.Val
	helperT(root.Right, pre, res)
}

func Test_getMinimumDifference() {
	root := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}

	res := getMinimumDifference(root)
	fmt.Println(res)
}
