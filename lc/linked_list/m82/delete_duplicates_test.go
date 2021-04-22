package m82

import (
	"fmt"
	"testing"

	ll "github.com/bournex/basic_training/lc/linked_list"
)

func TestDeleteDuplicates(t *testing.T) {
	testCases := []struct {
		Input  *ll.ListNode
		Expect *ll.ListNode
	}{
		{
			ll.Build([]int{1, 2, 3, 3, 4, 4, 5}),
			ll.Build([]int{1, 2, 5}),
		},
		{
			ll.Build([]int{1, 2, 3, 3, 3}),
			ll.Build([]int{1, 2}),
		},
		{
			ll.Build([]int{1, 1, 1, 1, 1}),
			ll.Build([]int{}),
		},
		{
			ll.Build([]int{1, 2, 3, 3, 3, 4, 5}),
			ll.Build([]int{1, 2, 4, 5}),
		},
		{
			ll.Build([]int{}),
			ll.Build([]int{}),
		},
		{
			ll.Build([]int{1}),
			ll.Build([]int{1}),
		},
		{
			ll.Build([]int{1, 1, 1, 2, 3}),
			ll.Build([]int{2, 3}),
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			input := ll.Copy(v.Input)
			res := deleteDuplicates(input)

			p1 := res
			p2 := v.Expect
			for p1 != nil && p2 != nil {
				if p1.Val != p2.Val {
					t.Errorf("input %s, expect %s, got %s", ll.Format(v.Input), ll.Format(v.Expect), ll.Format(res))
				}
				p1 = p1.Next
				p2 = p2.Next
			}

			if !(p1 == nil && p2 == nil) {
				t.Errorf("input %s, expect %s, got %s", ll.Format(v.Input), ll.Format(v.Expect), ll.Format(res))
			}
		})
	}
}
