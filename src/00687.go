package src

/*687. 最长同值路径*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {
	_, max := diglongestUnivaluePath(root)
	return max
}

func diglongestUnivaluePath(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	left, maxLeft := diglongestUnivaluePath(root.Left)
	right, maxRight := diglongestUnivaluePath(root.Right)
	curleft, curRight := 0, 0
	if root.Left != nil && root.Val == root.Left.Val {
		curleft = left + 1
	}
	if root.Right != nil && root.Val == root.Right.Val {
		curRight = right + 1
	}
	cur := curleft + curRight
	max := maxLeft
	if max < maxRight {
		max = maxRight
	}
	if max < cur {
		max = cur
	}
	if curleft < curRight {
		return curRight, max
	}
	return curleft, max
}
