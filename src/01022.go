package src

/*
1022. 从根到叶的二进制数之和
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) int {
	res := 0
	ans := genSRootLeaf(root)
	for _, an := range ans {
		res += getSum2(an)
	}
	return res
}

func getSum2(an []int) int {
	sum := 0
	for _, i := range an {
		sum = sum<<1 + i
	}
	return sum
}

func genSRootLeaf(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		temp := []int{
			root.Val,
		}
		return [][]int{temp}
	}
	var ans [][]int
	if root.Left != nil {
		left := genSRootLeaf(root.Left)
		for _, ints := range left {
			temp := []int{
				root.Val,
			}
			temp = append(temp, ints...)
			ans = append(ans, temp)
		}
	}
	if root.Right != nil {
		right := genSRootLeaf(root.Right)
		for _, ints := range right {
			temp := []int{
				root.Val,
			}
			temp = append(temp, ints...)
			ans = append(ans, temp)
		}
	}
	return ans
}
