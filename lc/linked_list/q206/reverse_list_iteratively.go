package q206

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 迭代解法
//	双指针
func reverseListIteratively(head *ListNode) *ListNode {
	var prev *ListNode
	curr, next := head, head

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
