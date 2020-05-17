package src

import "math"

/*783. 二叉搜索树节点最小距离*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
	pre := -1
	res := math.MaxInt32
	midorder(root, &pre, &res)
	return res
}

func midorder(root *TreeNode, pre *int, res *int) {
	if root == nil {
		return
	}
	midorder(root.Left, pre, res)
	if *pre != -1 {
		temp := int(math.Abs(float64(root.Val - *pre)))
		if temp < *res {
			*res = temp
		}
	}
	*pre = root.Val
	midorder(root.Right, pre, res)
}
