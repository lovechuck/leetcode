package src

/*234. 回文链表*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	var values []int
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}

	i := 0
	j := len(values) - 1
	for i < j {
		if values[i] != values[j] {
			return false
		}
		i++
		j--
	}

	return true
}
