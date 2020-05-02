package src

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth2(root.Left)
	right := maxDepth2(root.Right)
	if left >= right {
		return left + 1
	} else {
		return right + 1
	}
}
