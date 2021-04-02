package of24

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	examples := []struct {
		Input  *ListNode
		Expect *ListNode
	}{
		{
			&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			&ListNode{3, &ListNode{2, &ListNode{1, nil}}},
		},
		{
			nil,
			nil,
		},
		{
			&ListNode{1, nil},
			&ListNode{1, nil},
		},
	}
	for _, v := range examples {
		input := copyList(v.Input)
		res := reverseList(v.Input)

		expect := v.Expect
		rescopy := copyList(res)
		for res != nil && expect != nil {
			if res.Val != expect.Val {
				t.Errorf("input %s, expect %s, got %s", formatList(input), formatList(v.Expect), formatList(rescopy))
			}
			res = res.Next
			expect = expect.Next
		}

		if res != nil || expect != nil {
			t.Errorf("input %s, expect %s, got %s", formatList(input), formatList(v.Expect), formatList(rescopy))
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
