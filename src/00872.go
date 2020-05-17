package src

/*
872. 叶子相似的树

举个例子，如上图所示，给定一颗叶值序列为 (6, 7, 4, 9, 8) 的树。

如果有两颗二叉树的叶值序列是相同，那么我们就认为它们是 叶相似 的。
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	a1 := leafs(root1)
	a2 := leafs(root2)
	if len(a1) != len(a2) {
		return false
	}
	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}

func leafs(root1 *TreeNode) []int {
	if root1 == nil {
		return nil
	}
	var res []int
	if root1.Right == nil && root1.Left == nil {
		res = append(res, root1.Val)
	}
	if root1.Left != nil {
		left := leafs(root1.Left)
		res = append(res, left...)
	}
	if root1.Right != nil {
		right := leafs(root1.Right)
		res = append(res, right...)
	}
	return res
}
