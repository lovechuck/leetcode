package src

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return sortedArrayToBSTRoot(nums, 0, len(nums)-1)
}

func sortedArrayToBSTRoot(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	var root = &TreeNode{}
	root.Val = nums[mid]
	root.Left = sortedArrayToBSTRoot(nums, left, mid-1)
	root.Right = sortedArrayToBSTRoot(nums, mid+1, right)
	return root
}
