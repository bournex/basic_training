package of6

import (
	"fmt"
	"testing"
)

func TestReversePrint(t *testing.T) {
	examples := []struct {
		Input  *ListNode
		Expect []int
	}{
		{
			&ListNode{1, nil},
			[]int{1},
		},
		{
			nil,
			[]int{},
		},
		{
			&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			[]int{3, 2, 1},
		},
	}
	for _, v := range examples {
		res := reversePrint(v.Input)
		if len(res) != len(v.Expect) {
			t.Errorf("input %s, expect %+v, got %+v", formatList(v.Input), v.Expect, res)
		}

		for i := 0; i < len(res); i++ {
			if res[i] != v.Expect[i] {
				t.Errorf("input %s, expect %+v, got %+v", formatList(v.Input), v.Expect, res)
			}
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
