package question21

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	examples := []struct {
		Input  [2]*ListNode
		Expect *ListNode
	}{
		{
			[2]*ListNode{
				nil,
				nil,
			},
			nil,
		},
		{
			[2]*ListNode{
				{1, &ListNode{3, &ListNode{5, nil}}},
				{1, &ListNode{2, &ListNode{6, nil}}},
			},
			&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{5, &ListNode{6, nil}}}}}},
		},
		{
			[2]*ListNode{
				{1, &ListNode{3, &ListNode{5, nil}}},
				nil,
			},
			&ListNode{1, &ListNode{3, &ListNode{5, nil}}},
		},
	}

	for _, v := range examples {
		l1 := copyList(v.Input[0])
		l2 := copyList(v.Input[1])
		res := mergeTwoLists(l1, l2)
		head := res
		p := v.Expect
		for res != nil && p != nil {
			res = res.Next
			p = p.Next
		}

		if res != p {
			t.Errorf("input %s, %s, expect %s, got %s", formatList(v.Input[0]), formatList(v.Input[1]), formatList(v.Expect), formatList(head))
		}

	}
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
