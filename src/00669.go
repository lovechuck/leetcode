package src

/*669. 修剪二叉搜索树*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}
	var res *TreeNode
	if root.Val >= L && root.Val <= R {
		res = &TreeNode{
			Val: root.Val,
		}
		res.Left = trimBST(root.Left, L, R)
		res.Right = trimBST(root.Right, L, R)
	} else {
		res = trimBST(root.Left, L, R)
		if res != nil {
			if res.Right == nil {
				res.Right = trimBST(root.Right, L, R)
			} else {
				right := res.Right
				for right.Right != nil {
					right = right.Right
				}
				right = trimBST(root.Right, L, R)
			}

		} else {
			res = trimBST(root.Right, L, R)
		}
	}

	return res
}
