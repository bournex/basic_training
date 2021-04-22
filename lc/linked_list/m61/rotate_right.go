package m61

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

// 思考
//	双指针的变形题，首先O(n)获取到长度

func rotateRight(head *ll.ListNode, k int) *ll.ListNode {
	p := head
	n := 0
	for p != nil {
		n++
		p = p.Next
	}

	if n <= 1 {
		return head
	}

	n = k % n

	var prev *ll.ListNode = head
	var tail *ll.ListNode = head

	for i := 0; i < n; i++ {
		tail = tail.Next
	}

	for tail.Next != nil {
		prev = prev.Next
		tail = tail.Next
	}

	tail.Next = head
	head = prev.Next
	prev.Next = nil
	return head
}
