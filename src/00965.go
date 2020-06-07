package src

/*
965. 单值二叉树


如果二叉树每个节点都具有相同的值，那么该二叉树就是单值二叉树。

只有给定的树是单值二叉树时，才返回 true；否则返回 false。

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isRoot(root, root.Val)
}

func isRoot(root *TreeNode, val int) bool {
	if root.Val != val {
		return false
	}

	if root.Left != nil {
		isLeft := isRoot(root.Left, val)
		if !isLeft {
			return false
		}
	}
	if root.Right != nil {
		isRight := isRoot(root.Right, val)
		if !isRight {
			return false
		}
	}

	return true
}
