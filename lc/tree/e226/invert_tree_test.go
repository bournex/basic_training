package e226

import (
	"fmt"
	"testing"

	"github.com/bournex/basic_training/lc/tree"
)

func TestInvertTree(t *testing.T) {
	testCases := []struct {
		Input  *tree.TreeNode
		Expect *tree.TreeNode
	}{
		{
			tree.Build([]int{1, 2, 3, 4, 5, 6, 7}),
			tree.Build([]int{1, 3, 2, 7, 6, 5, 4}),
		},
		{
			tree.Build([]int{1, tree.NIL, 2, tree.NIL, tree.NIL, tree.NIL, 3}),
			tree.Build([]int{1, 2, tree.NIL, 3, tree.NIL}),
		},
		{
			tree.Build([]int{1, tree.NIL, 2}),
			tree.Build([]int{1, 2}),
		},
	}

	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			input := tree.CopyTree(v.Input)
			res := invertTree(input)
			if !tree.IsSameTree(res, v.Expect) {
				t.Errorf("input %+v, expect %+v, got %+v", v.Input, v.Expect, res)
			}
		})
	}
}
