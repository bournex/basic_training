package question19

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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tail, prev := head, head
	for ; tail != nil && n > 0; n-- {
		tail = tail.Next
	}

	for tail != nil && tail.Next != nil {
		prev = prev.Next
		tail = tail.Next
	}

	if prev != nil && prev.Next != nil {
		prev.Next = prev.Next.Next
	}
	return head
}
