package src

/*206. 反转链表*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var pre *ListNode
	current := head
	for current != nil {
		temp := current.Next
		current.Next = pre
		pre = current
		current = temp
	}
	return pre
}

func Test_reverseList() {

	a1 := []int{1, 2, 3, 4, 5}
	l1 := getList(a1)
	l2 := reverseList(l1)
	printLstNode(l2)
}
