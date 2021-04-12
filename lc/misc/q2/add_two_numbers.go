package q2

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{Val: 0}
	carrier := 0
	n1, n2 := 0, 0
	curr := res
	for l1 != nil || l2 != nil || carrier != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}

		p := new(ListNode)
		p.Val = (n1 + n2 + carrier) % 10
		curr.Next = p
		curr = p
		carrier = (n1 + n2 + carrier) / 10
	}
	return res.Next
}
