package src

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	var a []*TreeNode
	var temp []int
	temp = append(temp, root.Val)
	result = append(result, temp)
	a = AppendTree(root)
	for len(a) > 0 {
		var b []*TreeNode
		var temp []int
		for i := 0; i < len(a); i++ {
			temp = append(temp, a[i].Val)
			b = append(b, AppendTree(a[i])...)
		}
		result = append(result, temp)
		if len(b) >= 0 {
			a = b
		} else {
			break
		}
	}

	var xx = make([][]int, len(result))
	for i := 0; i < len(result); i++ {
		xx[i] = result[len(result)-1-i]
	}
	return xx
}

func AppendTree(root *TreeNode) []*TreeNode {
	var a []*TreeNode
	if root.Left != nil {
		a = append(a, root.Left)
	}
	if root.Right != nil {
		a = append(a, root.Right)
	}

	return a
}
