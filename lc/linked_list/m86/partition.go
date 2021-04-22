package m86

import ll "github.com/bournex/basic_training/lc/linked_list"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 思路
//	拆解遍历输入链表，构建三个链表，分别是包含小于、等于、大于x的值的链表
//	过程中维护三个链表的头尾节点指针
//
//	一开始理解错题意了，以为是小的在左边，大的在右边
//	实际是小的在左边，大与等于的按原序排在后面即可，如果是左右分大小的话，用三个链表就可以了
//

func partition(head *ll.ListNode, x int) *ll.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		head1 = &ll.ListNode{Val: 0, Next: nil}
		head2 = &ll.ListNode{Val: 0, Next: nil}
		tail1 = head1
		tail2 = head2
	)

	dummy := &ll.ListNode{Val: 0, Next: head}
	move := dummy
	for move.Next != nil {
		if move.Next.Val < x {
			tail1.Next = move.Next
			move.Next = move.Next.Next
			tail1.Next.Next = nil
			tail1 = tail1.Next
		} else {
			tail2.Next = move.Next
			move.Next = move.Next.Next
			tail2.Next.Next = nil
			tail2 = tail2.Next
		}
	}

	tail1.Next = head2.Next

	return head1.Next
}

func partition_threeway(head *ll.ListNode, x int) *ll.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		head1 = &ll.ListNode{Val: 0, Next: nil}
		head2 = &ll.ListNode{Val: 0, Next: nil}
		head3 = &ll.ListNode{Val: 0, Next: nil}
		tail1 = head1
		tail2 = head2
		tail3 = head3
	)

	dummy := &ll.ListNode{Val: 0, Next: head}
	move := dummy
	for move.Next != nil {
		if move.Next.Val < x {
			tail1.Next = move.Next
			move.Next = move.Next.Next
			tail1.Next.Next = nil
			tail1 = tail1.Next
		} else if move.Next.Val > x {
			tail3.Next = move.Next
			move.Next = move.Next.Next
			tail3.Next.Next = nil
			tail3 = tail3.Next
		} else {
			tail2.Next = move.Next
			move.Next = move.Next.Next
			tail2.Next.Next = nil
			tail2 = tail2.Next
		}
	}

	tail1.Next = head2.Next
	tail2.Next = head3.Next

	return head1.Next
}
