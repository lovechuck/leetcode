package src

/*563. 二叉树的坡度*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	_, t := findSumTilt(root)
	return t
}

func findSumTilt(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val, 0
	}
	left, ltilt := 0, 0
	if root.Left != nil {
		left, ltilt = findSumTilt(root.Left)
	}
	right, rtilt := 0, 0
	if root.Right != nil {
		right, rtilt = findSumTilt(root.Right)
	}

	tilt := 0
	sum := left + right + root.Val
	if left > right {
		tilt = left - right
	} else {
		tilt = right - left
	}
	tilt += ltilt + rtilt

	return sum, tilt
}
