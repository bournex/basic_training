package question24

import (
	"fmt"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	examples := []struct {
		Input  *ListNode
		Expect *ListNode
	}{
		{
			&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			&ListNode{2, &ListNode{1, &ListNode{3, nil}}},
		},
		{
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
			&ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, nil}}}},
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
		res := swapPairs(input)

		var a, b int
		for p := res; p != nil; p = p.Next {
			a++
		}
		for p := v.Expect; p != nil; p = p.Next {
			b++
		}

		if a != b {
			t.Errorf("input %s, expect %s, got %s", formatList(v.Input), formatList(v.Expect), formatList(res))
		}

		expect := copyList(v.Expect)
		for res != nil {
			if res.Val != expect.Val {
				t.Errorf("input %s, expect %s, got %s", formatList(v.Input), formatList(v.Expect), formatList(res))
			}

			expect = expect.Next
			res = res.Next
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
