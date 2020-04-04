package src

import (
	"strconv"
)

/*257. 二叉树的所有路径*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	result := []string{}

	findpath(root, "", &result)

	return result
}

func findpath(root *TreeNode, path string, result *[]string) {
	if root.Left == nil && root.Right == nil {
		path += strconv.Itoa(root.Val)
		*result = append(*result, path)
		return
	}
	if root.Left != nil {
		left := path + strconv.Itoa(root.Val) + "->"
		findpath(root.Left, left, result)
	}
	if root.Right != nil {
		right := path + strconv.Itoa(root.Val) + "->"
		findpath(root.Right, right, result)
	}
}
