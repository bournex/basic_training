package q141

import (
	"fmt"
	"testing"
)

func genCycleList(array []int, pos int) *ListNode {
	var special *ListNode
	head := new(ListNode)
	p := head
	for i, v := range array {
		if i == pos {
			special = &ListNode{v, nil}
			p.Next = special
			p = special
		} else {
			p.Next = &ListNode{v, nil}
			p = p.Next
		}
	}
	p.Next = special
	return head.Next
}
func TestHasCycle(t *testing.T) {
	examples := []struct {
		Input  *ListNode
		Expect bool
	}{
		{
			genCycleList([]int{1}, 0),
			true,
		},
		{
			genCycleList([]int{1}, -1),
			false,
		},
		{
			genCycleList([]int{1, 2, 3, 4}, 2),
			true,
		},
		{
			genCycleList([]int{1, 2, 3, 4}, -1),
			false,
		},
	}
	for _, v := range examples {
		res := hasCycle(v.Input)
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
