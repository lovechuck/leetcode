package src

import "fmt"

/*
606. 根据二叉树创建字符串
你需要采用前序遍历的方式，将一个二叉树转换成一个由括号和整数组成的字符串。

空节点则用一对空括号 "()" 表示。而且你需要省略所有不影响字符串与原始二叉树之间的一对一映射关系的空括号对。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	if t.Left == nil && t.Right == nil {
		return fmt.Sprintf("%d", t.Val)
	}

	res := fmt.Sprintf("%d", t.Val)
	left := tree2str(t.Left)
	res += fmt.Sprintf("(%s)", left)

	if t.Right != nil {
		right := tree2str(t.Right)
		res += fmt.Sprintf("(%s)", right)
	}

	return res
}
