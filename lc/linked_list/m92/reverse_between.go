package m92

import ll "github.com/bournex/basic_training/lc/linked_list"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 思考
//	这道题看起来不算难，但是很考察对链表的操作
//	重点在于
//	1.逆序部分的处理
//		如果输入规模已限制的话，那么这里用递归实现
//		如果输入规模较大的话，会栈溢出，可以通过额外的slice空间来实现
//	2.逆序后头尾与原链表的衔接
//		这块将新的翻转部分的头和尾部待衔接的链表剩余部分保存在递归外
func reverseBetween(head *ll.ListNode, left int, right int) *ll.ListNode {
	if head == nil || head.Next == nil || left == right {
		return head
	}

	idx := 1
	sub_head, sub_tail := (*ll.ListNode)(nil), (*ll.ListNode)(nil)

	var rotate func(*ll.ListNode) *ll.ListNode
	rotate = func(ln *ll.ListNode) *ll.ListNode {
		if idx < right {
			idx++
			rotate(ln.Next).Next = ln
		} else {
			sub_head = ln
			sub_tail = ln.Next
		}

		return ln
	}

	dummy := &ll.ListNode{Val: 0, Next: head}
	move := dummy

	for move.Next != nil {
		if left == idx {
			tailPrev := move.Next
			rotate(move.Next)
			move.Next = sub_head
			tailPrev.Next = sub_tail
			return dummy.Next
		} else {
			move = move.Next
		}
		idx++
	}

	return head
}
