package q234

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

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	_, res := isPalindromeRecurrsively(head, head)
	return res
}

// [1,1]
// [1,2,1]
// [1,2,2,1]
func isPalindromeRecurrsively(slow, fast *ListNode) (*ListNode, bool) {
	if fast == nil {
		return slow, true
	}
	if fast.Next == nil {
		return slow.Next, true
	}

	curr, b2 := isPalindromeRecurrsively(slow.Next, fast.Next.Next)
	if !b2 {
		return nil, b2
	}
	return curr.Next, slow.Val == curr.Val
}
