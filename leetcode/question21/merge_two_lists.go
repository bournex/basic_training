package question21

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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		if l1 == nil {
			return l2
		} else if l2 == nil {
			return l1
		} else {
			return nil
		}
	}

	head := new(ListNode)
	curr := head

	for l1 != nil || l2 != nil {
		if l1 == nil {
			curr.Next = l2
			break
		} else if l2 == nil {
			curr.Next = l1
			break
		}

		if l1.Val >= l2.Val {
			tmp := l2
			l2 = l2.Next
			curr.Next = tmp
		} else if l1.Val < l2.Val {
			tmp := l1
			l1 = l1.Next
			curr.Next = tmp
		}
		curr = curr.Next
	}

	return head.Next
}
