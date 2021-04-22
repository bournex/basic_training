package m82

import (
	ll "github.com/bournex/basic_training/lc/linked_list"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ll.ListNode) *ll.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ll.ListNode{Val: 0, Next: head}
	left := dummy
	prev := head
	next := head.Next

	for prev != nil && next != nil {
		if prev.Val == next.Val {
			next = next.Next
		} else if prev.Next != next {
			prev.Val = next.Val // 拷贝next的值到prev
			prev.Next = next.Next
			next.Next = nil
			next = prev.Next
		} else {
			left = prev
			prev = next
			next = prev.Next
		}
	}

	if prev != nil && prev.Next != next {
		left.Next = nil
	}

	return dummy.Next
}
