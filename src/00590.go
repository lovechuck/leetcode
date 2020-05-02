package src

/*590. N叉树的后序遍历*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var res []int
	for _, child := range root.Children {
		c := postorder(child)
		res = append(res, c...)
	}
	res = append(res, root.Val)
	return res
}
