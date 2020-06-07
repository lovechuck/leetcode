package src

/*
938. 二叉搜索树的范围和

给定二叉搜索树的根结点 root，返回 L 和 R（含）之间的所有结点的值的和。

二叉搜索树保证具有唯一的值。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, L int, R int) int {
	sum := 0
	if root.Val >= L && root.Val <= R {
		sum += root.Val
	}
	if root.Left != nil {
		sum += rangeSumBST(root.Left, L, R)
	}
	if root.Right != nil {
		sum += rangeSumBST(root.Right, L, R)
	}
	return sum
}
