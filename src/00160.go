package src

/*160. 相交链表*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	aLen := 0
	hpA := headA
	for hpA != nil {
		aLen++
		hpA = hpA.Next
	}
	hpB := headB
	bLen := 0
	for hpB != nil {
		bLen++
		hpB = hpB.Next
	}
	min := headA
	max := headB
	len := bLen - aLen
	if aLen > bLen {
		min = headB
		max = headA
		len = aLen - bLen
	}
	for i := 0; i < len; i++ {
		max = max.Next
	}
	for min != nil && max != nil {
		if min == max {
			return min
		}
		min = min.Next
		max = max.Next
	}

	return nil
}
