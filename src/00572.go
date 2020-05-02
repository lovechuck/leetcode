package src

/*572. 另一个树的子树*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	return isSame(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSame(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	} else if s == nil && t != nil {
		return false
	} else if s != nil && t == nil {
		return false
	} else {
		if s.Val != t.Val {
			return false
		}
		return isSame(s.Left, t.Left) && isSame(s.Right, t.Right)
	}
}
