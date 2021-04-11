package q206

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

// 递归解法
func reverseListRecursively(head *ListNode) *ListNode {
	_, phead := recursively(head)
	return phead
}

func recursively(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}

	prev, nhead := recursively(head.Next)
	if prev == nil {
		return head, head
	}

	prev.Next = head
	head.Next = nil
	return head, nhead
}
