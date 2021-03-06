package q110

import (
	"fmt"
	"testing"

	"github.com/bournex/basic_training/lc/tree"
)

func TestIsBalanced(t *testing.T) {
	examples := []struct {
		Input  *tree.TreeNode
		Expect bool
	}{
		{
			nil,
			true,
		},
		{
			tree.Build([]int{1}),
			true,
		},

		{
			tree.Build([]int{1, 2, tree.INT_MIN, 3}),
			false,
		},

		{
			tree.Build([]int{1, 2, 3}),
			true,
		},
	}

	for i, v := range examples {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			res := isBalanced(v.Input)
			if res != v.Expect {
				t.Errorf("expect %t, got %t", v.Expect, res)
			}
		})
	}
}
