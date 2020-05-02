package src

/*
559. N叉树的最大深度
*/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	max := 0
	for _, child := range root.Children {
		tmp := maxDepth(child)
		if max < tmp {
			max = tmp
		}
	}
	return max + 1
}
