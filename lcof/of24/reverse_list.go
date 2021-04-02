package of24

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

func reverseList(head *ListNode) *ListNode {
	prev := (*ListNode)(nil)
	curr := head
	for curr != nil {
		curr = head.Next
		head.Next = prev
		prev = head
		head = curr
	}
	return prev
}
