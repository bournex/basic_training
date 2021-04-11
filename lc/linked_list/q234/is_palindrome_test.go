package q234

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	examples := []struct {
		Input  *ListNode
		Expect bool
	}{
		{
			&ListNode{1, nil},
			true,
		},
		{
			&ListNode{1, &ListNode{1, nil}},
			true,
		},
		{
			&ListNode{1, &ListNode{2, nil}},
			false,
		},
		{
			&ListNode{1, &ListNode{2, &ListNode{1, nil}}},
			true,
		},
		{
			&ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}},
			true,
		},
	}

	for _, v := range examples {
		res := isPalindrome(v.Input)

		if res != v.Expect {
			t.Errorf("input %s, expect %t, got %t", formatList(v.Input), v.Expect, res)
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
