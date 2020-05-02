package src

/*589. N叉树的前序遍历*/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var res []int
	res = append(res, root.Val)
	for _, child := range root.Children {
		c := preorder(child)
		res = append(res, c...)
	}
	return res
}
