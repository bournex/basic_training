package q108

import (
	"testing"

	"github.com/bournex/basic_training/lc/tree"
)

func TestSortedArrayToBst(t *testing.T) {
	examples := []struct {
		Input  []int
		Expect *tree.TreeNode
	}{
		{
			[]int{1, 2, 3},
			tree.Build([]int{2, 1, 3}),
		},
		{
			[]int{1, 2, 3, 4},
			tree.Build([]int{3, 2, 4, 1}),
		},
	}

	for _, v := range examples {
		res := sortedArrayToBST(v.Input)
		if !tree.IsSameTree(res, v.Expect) {
			t.Errorf("input %+v failed", v.Input)
		}
	}
}
