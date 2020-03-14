package src

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSameisSymmetric(root.Left, root.Right)
}

func isSameisSymmetric(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil && q != nil {
		return false
	}

	if p != nil && q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameisSymmetric(p.Left, q.Right) && isSameisSymmetric(p.Right, q.Left)
}
