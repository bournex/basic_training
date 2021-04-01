package question19

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
		p1 := copyList(v.Input)
		p2 := removeNthFromEnd(v.Input, v.N)

		p3 := p1
		p4 := p2
		for p1 != nil && p2 != nil && p1.Val == p2.Val {
			p1 = p1.Next
			p2 = p2.Next
		}
		t.Errorf("input %s, expect %s, got %s", formatList(p3), formatList(v.Expect), formatList(p4))

		if p2 != p1 {
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
