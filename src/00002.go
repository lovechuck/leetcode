package src

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
 	给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

	如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

	您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

	示例：

	输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
	输出：7 -> 0 -> 8
	原因：342 + 465 = 807
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func getList(data []int) *ListNode {
	l1 := &ListNode{
		Val:  0,
		Next: nil,
	}
	head := l1
	for k, v := range data {
		head.Val = v
		if k != len(data)-1 {
			head.Next = &ListNode{
				Val:  0,
				Next: nil,
			}
		}
		head = head.Next
	}
	return l1
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{
		Val:  0,
		Next: nil,
	}
	var head = result
	var i = 0
	for {

		value := 0
		if l1 == nil {
			value = l2.Val + i
			l2 = l2.Next
		} else if l2 == nil {
			value = l1.Val + i
			l1 = l1.Next
		} else {
			value = l1.Val + l2.Val + i
			l1 = l1.Next
			l2 = l2.Next
		}
		if value > 9 {
			head.Val = value - 10
			i = 1
		} else {
			head.Val = value
			i = 0
		}

		if l1 == nil && l2 == nil {
			if i != 0 {
				head.Next = &ListNode{
					Val:  1,
					Next: nil,
				}
			}
			break
		}

		head.Next = &ListNode{
			Val:  0,
			Next: nil,
		}
		head = head.Next
	}

	return result
}

func printLstNode(l1 *ListNode) {
	for {
		if l1 == nil {
			break
		}
		fmt.Print(l1.Val)
		l1 = l1.Next
	}
}

func Test_addTwoNumbers() {
	a1 := []int{2, 4, 3}
	a2 := []int{5, 6, 4}
	l1 := getList(a1)
	l2 := getList(a2)

	result1 := addTwoNumbers(l1, l2)
	printLstNode(result1)

	a11 := []int{5}
	a21 := []int{5}
	l11 := getList(a11)
	l21 := getList(a21)

	result11 := addTwoNumbers(l11, l21)
	printLstNode(result11)

	a111 := []int{1, 8}
	a211 := []int{9}
	l111 := getList(a111)
	l211 := getList(a211)

	result111 := addTwoNumbers(l111, l211)
	printLstNode(result111)

	a1111 := []int{1}
	a2111 := []int{9, 9}
	l1111 := getList(a1111)
	l2111 := getList(a2111)

	result1111 := addTwoNumbers(l1111, l2111)
	printLstNode(result1111)
}
