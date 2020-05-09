package src

/*
671. 二叉树中第二小的节点
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	min := findSecondMinimumValueX(root, root.Val)
	return min
}

func findSecondMinimumValueX(root *TreeNode, min int) int {
	if root == nil {
		return -1
	}
	if root.Val != min {
		return root.Val
	}
	left := findSecondMinimumValueX(root.Left, min)
	right := findSecondMinimumValueX(root.Right, min)
	if left == -1 {
		return right
	}
	if right == -1 {
		return left
	}
	if left < right {
		return left
	}
	return right
}
