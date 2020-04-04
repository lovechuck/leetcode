package src

import "math"

/*最接近的二叉搜索树值  */
func closestValue(root *TreeNode, target float64) int {
	return findCloset(root, target)
}

func findCloset(root *TreeNode, target float64) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	current := math.Abs(float64(root.Val) - target)
	if float64(root.Val) > target {
		if root.Left == nil {
			return root.Val
		} else {
			left := math.Abs(float64(root.Left.Val) - target)
			if left >= current {
				return root.Val
			} else {
				return findCloset(root.Left, target)
			}
		}
	} else {
		if root.Right == nil {
			return root.Val
		} else {
			right := math.Abs(float64(root.Right.Val) - target)
			if right >= current {
				return root.Val
			} else {
				return findCloset(root.Right, target)
			}
		}
	}
}
