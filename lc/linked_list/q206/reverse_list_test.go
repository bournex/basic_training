package q206

import (
	"fmt"
	"testing"
)

type Case struct {
	Input  *ListNode
	Expect *ListNode
}

func genCases() []Case {
	exampels := []Case{
		{
			nil,
			nil,
		},
		{
			&ListNode{1, nil},
			&ListNode{1, nil},
		},
		{
			&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			&ListNode{3, &ListNode{2, &ListNode{1, nil}}},
		},
	}

	return exampels
}

func TestReverseListRecursively(t *testing.T) {
	cases := genCases()
	for _, v := range cases {
		input := copyList(v.Input)
		res := reverseListRecursively(input)
		exp := copyList(v.Expect)
		for res != nil && exp != nil {
			if res.Val != exp.Val {
				t.Errorf("input %s, expect %s, got %s", formatList(v.Input), formatList(v.Expect), formatList(input))
			}
			res = res.Next
			exp = exp.Next
		}

		if !(res == nil && exp == nil) {
			t.Errorf("input %s, expect %s, got %s", formatList(v.Input), formatList(v.Expect), formatList(input))
		}
	}
}

func TestReverseListIteratively(t *testing.T) {
	cases := genCases()
	for _, v := range cases {
		input := copyList(v.Input)
		res := reverseListIteratively(input)
		exp := copyList(v.Expect)
		for res != nil && exp != nil {
			if res.Val != exp.Val {
				t.Errorf("input %s, expect %s, got %s", formatList(v.Input), formatList(v.Expect), formatList(input))
			}
			res = res.Next
			exp = exp.Next
		}

		if !(res == nil && exp == nil) {
			t.Errorf("input %s, expect %s, got %s", formatList(v.Input), formatList(v.Expect), formatList(input))
		}
	}
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
