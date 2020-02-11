package src

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	result := &ListNode{
		Val:  l1.Val,
		Next: nil,
	}
	res := result
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			res.Next = l1
			res = res.Next
			l1 = l1.Next
		} else {
			res.Next = l2
			res = res.Next
			l2 = l2.Next
		}
	}
	if l1 != nil {
		res.Next = l1
	}
	if l2 != nil {
		res.Next = l2
	}
	return result.Next
}

func Test_mergeTwoLists() {
	a1 := []int{2, 4, 3}
	a2 := []int{5, 6, 4}
	l1 := getList(a1)
	l2 := getList(a2)

	l := mergeTwoLists(l1, l2)

	printLstNode(l)
}
