package h23

import (
	"fmt"
	"testing"

	ll "github.com/bournex/basic_training/lc/linked_list"
)

func TestMergeKList(t *testing.T) {
	testCases := []struct {
		Input  []*ll.ListNode
		Expect *ll.ListNode
	}{
		{
			[]*ll.ListNode{nil, nil},
			nil,
		},
		{
			[]*ll.ListNode{ll.Build([]int{1, 2, 3})},
			ll.Build([]int{1, 2, 3}),
		},
		{
			[]*ll.ListNode{ll.Build([]int{1, 3, 5}), ll.Build([]int{2, 4, 6})},
			ll.Build([]int{1, 2, 3, 4, 5, 6}),
		},
		{
			[]*ll.ListNode{ll.Build([]int{1}), ll.Build([]int{1})},
			ll.Build([]int{1, 1}),
		},
		{
			[]*ll.ListNode{ll.Build([]int{1, 2, 3}), ll.Build([]int{1, 3, 6}), ll.Build([]int{2, 5, 9})},
			ll.Build([]int{1, 1, 2, 2, 3, 3, 5, 6, 9}),
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := mergeKLists(v.Input)
			p1 := res
			p2 := v.Expect

			for p1 != nil && p2 != nil {
				if p1.Val != p2.Val {
					t.Errorf("expect %s, got %s", ll.Format(v.Expect), ll.Format(res))
					return
				}

				p1 = p1.Next
				p2 = p2.Next
			}

			if !(p1 == nil && p2 == nil) {
				t.Errorf("expect %s, got %s", ll.Format(v.Expect), ll.Format(res))
				return
			}
		})
	}
}
