package src

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 || lists == nil {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		return mergeTwoLists(lists[0], lists[1])
	}

	mid := len(lists) / 2
	left := lists[0:mid]
	right := lists[mid:]

	return mergeTwoLists(mergeKLists(left), mergeKLists(right))
}
