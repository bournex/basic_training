package h95

import (
	"fmt"
	"testing"

	"github.com/bournex/basic_training/lc/tree"
)

func TestGenerateTrees(t *testing.T) {
	testCases := []struct {
		Input  int
		Expect []*tree.TreeNode
	}{
		{
			3,
			[]*tree.TreeNode{
				tree.Build([]int{2, 1, 3}),
				tree.Build([]int{1, tree.NIL, 2, tree.NIL, tree.NIL, tree.NIL, 3}),
				tree.Build([]int{1, tree.NIL, 3, tree.NIL, tree.NIL, 2}),
				tree.Build([]int{3, 2, tree.NIL, 1}),
				tree.Build([]int{3, 1, tree.NIL, tree.NIL, 2}),
			},
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			res := generateTrees(v.Input)
			if len(res) != len(v.Expect) {
				t.Errorf("input %d, expect %d trees, got %d trees", v.Input, len(v.Expect), len(res))
			}

			for _, v1 := range res {
				var found bool
				for _, v2 := range v.Expect {
					if tree.IsSameTree(v1, v2) {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("input %d, expect %d trees, got %d trees", v.Input, len(v.Expect), len(res))
				}
			}
		})
	}
}
