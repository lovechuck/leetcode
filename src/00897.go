package src

/*
897. 递增顺序查找树

给你一个树，请你 按中序遍历 重新排列树，使树中最左边的结点现在是树的根，并且每个结点没有左子结点，只有一个右子结点。
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	ans := inorder(root)
	var res = &TreeNode{
		Val:   ans[0],
		Left:  nil,
		Right: nil,
	}
	cur := res
	for i := 1; i < len(ans); i++ {
		cur.Right = &TreeNode{
			Val: ans[i],
		}
		cur = cur.Right
	}
	return res
}
func inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	if root.Left != nil {
		left := inorder(root.Left)
		res = append(res, left...)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		left := inorder(root.Right)
		res = append(res, left...)
	}
	return res
}
