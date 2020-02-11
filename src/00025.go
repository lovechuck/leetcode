package src

import "fmt"

/*
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

示例 :

给定这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	result := &ListNode{
		Next: head,
	}
	start := result
	end := result
	for end != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		//记录上一次结束 和 下一次开始
		preEnd := start.Next //翻转后为最后一位
		nextStart := end.Next

		end.Next = nil
		//翻转
		start.Next = reverseGroup(start)
		preEnd.Next = nextStart

		// 前置节点
		start = preEnd
		end = preEnd
	}

	return result.Next
}

func reverseGroup(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var tail *ListNode
	for head != nil {
		cur := head
		head = head.Next
		cur.Next = tail
		tail = cur
	}

	return tail
}

func Test_reverseKGroup() {
	data := []int{1, 2, 3, 4, 5}
	ls := getList(data)
	printLstNode(ls)
	fmt.Println("after")
	re := reverseKGroup(ls, 2)
	printLstNode(re)
	fmt.Println("done")
}
