package src

import "fmt"

/*501. 二叉搜索树中的众数*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	mmap := make(map[int]int)
	inOrder(root, mmap)
	max := 1
	var res []int
	for k, v := range mmap {
		if v > max {
			res = []int{}
			res = append(res, k)
			max = v
		} else if v == max {
			res = append(res, k)
		}
	}
	return res
}

func inOrder(root *TreeNode, mmap map[int]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, mmap)
	if _, ok := mmap[root.Val]; ok {
		mmap[root.Val] = mmap[root.Val] + 1
	} else {
		mmap[root.Val] = 1
	}
	inOrder(root.Right, mmap)
}

func Test_findMode() {
	root := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}

	res := findMode(root)
	fmt.Println(res)
}
