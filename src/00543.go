package src

import "fmt"

/*543. 二叉树的直径*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	ans, _ := gendiameterOfBinaryTree(root)
	return ans
}

func gendiameterOfBinaryTree(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	if root.Left == nil && root.Right == nil {
		return 0, 0
	}
	lans, left := 0, 0
	if root.Left != nil {
		lans, left = gendiameterOfBinaryTree(root.Left)
		left = left + 1
	}
	rans, right := 0, 0
	if root.Right != nil {
		rans, right = gendiameterOfBinaryTree(root.Right)
		right = right + 1
	}

	ans := lans
	if lans < rans {
		ans = rans
	}
	cans := left + right
	if cans > ans {
		ans = cans
	}
	cur := left
	if left < right {
		cur = right
	}

	return ans, cur
}

func Test_diameterOfBinaryTree() {
	root := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Print(diameterOfBinaryTree(&root))
}
