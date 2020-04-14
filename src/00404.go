package src

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	return sumOfLeftLeavesC(root, false)
}
func sumOfLeftLeavesC(root *TreeNode, flag bool) int {
	if root == nil {
		return 0
	}

	leave := 0
	if flag && root.Left == nil && root.Right == nil {
		leave = root.Val
	}

	left := sumOfLeftLeavesC(root.Left, true)
	right := sumOfLeftLeavesC(root.Right, false)

	return leave + left + right
}
