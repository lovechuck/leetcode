package src

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Val == sum {
		return true
	}
	if root.Left != nil && root.Right == nil {
		return hasPathSum(root.Left, sum-root.Val)
	}

	if root.Left == nil && root.Right != nil {
		return hasPathSum(root.Right, sum-root.Val)
	}

	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
