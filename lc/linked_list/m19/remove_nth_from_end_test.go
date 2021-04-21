package q19

import (
	"fmt"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	examples := []struct {
		Input  *ListNode
		N      int
		Expect *ListNode
	}{
		{
			initList([]int{1, 2}),
			2,
			initList([]int{2}),
		},
		{
			initList([]int{1}),
			1,
			nil,
		},
		{
			initList([]int{1, 2, 3, 4, 5}),
			2,
			initList([]int{1, 2, 3, 5}),
		},
	}

	for _, v := range examples {
		input := copyList(v.Input)
		res := removeNthFromEnd(input, v.N)

		p1 := res
		p2 := v.Expect
		for p1 != nil && p2 != nil && p1.Val == p2.Val {
			if p1.Val != p2.Val {
				t.Errorf("input %s, expect %s, got %s", formatList(v.Input), formatList(v.Expect), formatList(res))
			}
			p1 = p1.Next
			p2 = p2.Next
		}
		if !(p1 == nil && p2 == nil) {
			t.Errorf("input %s, expect %s, got %s", formatList(v.Input), formatList(v.Expect), formatList(res))
		}
	}
}

func initList(s []int) *ListNode {
	var head *ListNode
	for i := len(s) - 1; i >= 0; i-- {
		head = &ListNode{
			s[i],
			head,
		}
	}
	return head
}

func copyList(p *ListNode) *ListNode {
	np := &ListNode{}
	head := np

	for p != nil {
		np.Next = &ListNode{
			Val:  p.Val,
			Next: p.Next,
		}
		p = p.Next
		np = np.Next
	}
	return head.Next
}

func formatList(p *ListNode) string {
	s := "["
	for p != nil {
		s += fmt.Sprintf("%d -> ", p.Val)
		p = p.Next
	}
	s += "nil]"
	return s
}
