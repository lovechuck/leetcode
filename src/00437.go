package src

/*437. 路径总和 III*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return helper(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func helper(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	sum = root.Val - sum
	if sum == 0 {
		// 子节点为 0 的情况
		return 1 + helper(root.Left, sum) + helper(root.Right, sum)
	} else {
		return helper(root.Left, sum) + helper(root.Right, sum)
	}
}
