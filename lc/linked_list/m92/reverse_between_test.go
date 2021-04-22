package m92

import (
	"fmt"
	"testing"

	ll "github.com/bournex/basic_training/lc/linked_list"
)

func TestReverseBetween(t *testing.T) {
	testCases := []struct {
		Input  *ll.ListNode
		Left   int
		Right  int
		Expect *ll.ListNode
	}{
		{
			ll.Build([]int{1, 2, 3}),
			1,
			3,
			ll.Build([]int{3, 2, 1}),
		},
		{
			ll.Build([]int{1, 2}),
			1,
			2,
			ll.Build([]int{2, 1}),
		},
		{
			ll.Build([]int{1, 2, 3, 4}),
			2,
			3,
			ll.Build([]int{1, 3, 2, 4}),
		},
		{
			ll.Build([]int{1, 2, 3, 4, 5}),
			2,
			4,
			ll.Build([]int{1, 4, 3, 2, 5}),
		},
		{
			ll.Build([]int{1, 2, 3, 4, 5}),
			4,
			5,
			ll.Build([]int{1, 2, 3, 5, 4}),
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			input := ll.Copy(v.Input)
			res := reverseBetween(input, v.Left, v.Right)

			p1 := res
			p2 := v.Expect
			for p1 != nil && p2 != nil {

				if p1.Val != p2.Val {
					t.Errorf("input %s, left %d, right %d, expect %s, got %s", ll.Format(v.Input), v.Left, v.Right, ll.Format(v.Expect), ll.Format(res))
					return
				}

				p1 = p1.Next
				p2 = p2.Next
			}

			if !(p1 == nil && p2 == nil) {
				t.Errorf("input %s, left %d, right %d, expect %s, got %s", ll.Format(v.Input), v.Left, v.Right, ll.Format(v.Expect), ll.Format(res))
			}
		})
	}
}
