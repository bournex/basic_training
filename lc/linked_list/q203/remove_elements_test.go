package q203

import (
	"fmt"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	examples := []struct {
		Input    *ListNode
		ToDelete int
		Expect   *ListNode
	}{
		{
			nil,
			1,
			nil,
		},
		{
			&ListNode{1, nil},
			1,
			nil,
		},
		{
			&ListNode{1, nil},
			2,
			&ListNode{1, nil},
		},
		{
			&ListNode{1, &ListNode{1, &ListNode{2, nil}}},
			1,
			&ListNode{2, nil},
		},
		{
			&ListNode{1, &ListNode{1, &ListNode{2, nil}}},
			2,
			&ListNode{1, &ListNode{1, nil}},
		}, {
			&ListNode{1, &ListNode{2, &ListNode{2, &ListNode{3, nil}}}},
			2,
			&ListNode{1, &ListNode{3, nil}},
		},
	}

	for _, v := range examples {
		input := copyList(v.Input)
		res := removeElements(input, v.ToDelete)

		exp := copyList(v.Expect)
		for res != nil && exp != nil {
			if res.Val != exp.Val {
				t.Errorf("input %s, to delete %d, expect %s, got %s", formatList(v.Input), v.ToDelete, formatList(v.Expect), formatList(res))
				continue
			}
			res = res.Next
			exp = exp.Next
		}

		if !(res == nil && exp == nil) {
			t.Errorf("input %s, to delete %d, expect %s, got %s", formatList(v.Input), v.ToDelete, formatList(v.Expect), formatList(res))
			continue
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
