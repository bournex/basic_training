package m86

import (
	"fmt"
	"testing"

	ll "github.com/bournex/basic_training/lc/linked_list"
)

func TestPartition(t *testing.T) {
	testCases := []struct {
		Input  *ll.ListNode
		X      int
		Expect *ll.ListNode
	}{
		{
			ll.Build([]int{1, 2, 4, 3, 2, 5}),
			3,
			ll.Build([]int{1, 2, 2, 4, 3, 5}),
		},
		{
			ll.Build([]int{5, 4, 3, 2, 1}),
			5,
			ll.Build([]int{4, 3, 2, 1, 5}),
		},
		{
			ll.Build([]int{5, 4, 3, 2, 1}),
			1,
			ll.Build([]int{5, 4, 3, 2, 1}),
		},
		{
			ll.Build([]int{5, 4, 3, 2, 1}),
			4,
			ll.Build([]int{3, 2, 1, 5, 4}),
		},
		{
			ll.Build([]int{2, 1}),
			2,
			ll.Build([]int{1, 2}),
		},
		{
			ll.Build([]int{2, 1}),
			1,
			ll.Build([]int{2, 1}),
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			input := ll.Copy(v.Input)
			res := partition(input, v.X)
			p1 := res
			p2 := v.Expect
			for p1 != nil && p2 != nil {

				if p1.Val != p2.Val {
					t.Errorf("input %s, k %d, expect %s, got %s", ll.Format(v.Input), v.X, ll.Format(v.Expect), ll.Format(res))
					return
				}

				p1 = p1.Next
				p2 = p2.Next
			}

			if !(p1 == nil && p2 == nil) {
				t.Errorf("input %s, k %d, expect %s, got %s", ll.Format(v.Input), v.X, ll.Format(v.Expect), ll.Format(res))
			}
		})
	}
}
