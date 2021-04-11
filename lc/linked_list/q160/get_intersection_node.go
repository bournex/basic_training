package q160

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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	la, lb := 0, 0
	pa, pb := headA, headB
	for pa != nil || pb != nil {
		if pa != nil {
			la++
			pa = pa.Next
		}
		if pb != nil {
			lb++
			pb = pb.Next
		}
	}
	pa, pb = headA, headB
	for la != lb {
		if la > lb {
			pa = pa.Next
			la--
		}
		if la < lb {
			lb--
			pb = pb.Next
		}
	}

	for pa != pb {
		pa = pa.Next
		pb = pb.Next
		if pa == nil || pb == nil {
			return nil
		}
	}
	return pa
}
