package question24

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	// 处理空链表和只有一个元素的链表
	if head == nil || head.Next == nil {
		return head
	}

	return nil
}
