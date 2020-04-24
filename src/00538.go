package src

/*
538. 把二叉搜索树转换为累加树
给定一个二叉搜索树（Binary Search Tree），把它转换成为累加树（Greater Tree)，使得每个节点的值是原来的节点值加上所有大于它的节点值之和。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	sum := 0
	return convertBSTS(root, &sum)
}

func convertBSTS(root *TreeNode, sum *int) *TreeNode {
	if root == nil {
		return nil
	}
	convertBSTS(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	convertBSTS(root.Left, sum)
	return root
}
