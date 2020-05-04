package src

/*
653. 两数之和 IV - 输入 BST
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	hash := make(map[int]int)
	return findTargetHash(root, k, &hash)
}

func findTargetHash(root *TreeNode, k int, hash *map[int]int) bool {
	if root == nil {
		return false
	}
	left := k - root.Val
	if _, ok := (*hash)[left]; ok {
		return true
	}
	(*hash)[root.Val] = root.Val
	return findTargetHash(root.Left, k, hash) || findTargetHash(root.Right, k, hash)
}
