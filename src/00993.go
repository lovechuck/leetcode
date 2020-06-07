package src

/*
993. 二叉树的堂兄弟节点

在二叉树中，根节点位于深度 0 处，每个深度为 k 的节点的子节点位于深度 k+1 处。

如果二叉树的两个节点深度相同，但父节点不同，则它们是一对堂兄弟节点。

我们给出了具有唯一值的二叉树的根节点 root，以及树中两个不同节点的值 x 和 y。

只有与值 x 和 y 对应的节点是堂兄弟节点时，才返回 true。否则，返回 false。
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
	a, b := finDept(root, 0, x)
	c, d := finDept(root, 0, y)
	return a == c && b != d
}

func finDept(root *TreeNode, dept int, x int) (int, int) {
	if root == nil {
		return -1, -1
	}
	if root.Val == x {
		return -1, -1
	}

	if root.Left != nil {
		if root.Left.Val == x {
			return dept + 1, root.Val
		} else {
			a, b := finDept(root.Left, dept+1, x)
			if b != -1 {
				return a, b
			}
		}
	}
	if root.Right != nil {
		if root.Right.Val == x {
			return dept + 1, root.Val
		} else {
			a, b := finDept(root.Right, dept+1, x)
			if b != -1 {
				return a, b
			}
		}
	}
	return -1, -1
}
