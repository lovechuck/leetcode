package src

/*141. 环形链表*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	step1 := head.Next
	step2 := head.Next.Next

	for step1 != step2 {
		if step1 == nil || step2 == nil || step2.Next == nil {
			return false
		}
		step1 = step1.Next
		step2 = step2.Next.Next
	}

	return true
}
