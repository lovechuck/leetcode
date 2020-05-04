package src

/*
637. 二叉树的层平均值


给定一个非空二叉树, 返回一个由每层节点平均值组成的数组.

示例 1:

输入:
    3
   / \
  9  20
    /  \
   15   7
输出: [3, 14.5, 11]
解释:
第0层的平均值是 3,  第1层是 14.5, 第2层是 11. 因此返回 [3, 14.5, 11].
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	var queue []*TreeNode
	queue = append(queue, root)
	var res []float64
	for len(queue) > 0 {
		sum := 0
		var tmp []*TreeNode
		for i := 0; i < len(queue); i++ {
			sum += queue[i].Val
			if queue[i].Left != nil {
				tmp = append(tmp, queue[i].Left)
			}
			if queue[i].Right != nil {
				tmp = append(tmp, queue[i].Right)
			}
		}
		res = append(res, float64(sum)/float64(len(queue)))
		queue = tmp
	}
	return res
}
