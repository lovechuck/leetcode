package src

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isDept(root) != -1
}

func isDept(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := isDept(root.Left)
	if left == -1 {
		return -1
	}
	right := isDept(root.Right)
	if right == -1 {
		return -1
	}
	if mabs(left, right) < 2 {
		return mmax(left, right) + 1
	} else {
		return -1
	}
}

func mabs(a int, b int) int {
	if a >= b {
		return a - b
	} else {
		return b - a
	}
}

func mmax(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
