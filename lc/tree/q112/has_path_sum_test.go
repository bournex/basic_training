package q112

import (
	"fmt"
	"testing"

	"github.com/bournex/basic_training/lc/tree"
)

func TestHasPathSum(t *testing.T) {
	examples := []struct {
		Input  *tree.TreeNode
		Target int
		Expect bool
	}{
		{
			nil,
			4,
			false,
		},
		{
			tree.Build([]int{1, 2, 3}),
			4,
			true,
		},
		{
			tree.Build([]int{1, 2, 3}),
			5,
			false,
		},
		{
			tree.Build([]int{5, 4, 8, 11, tree.NIL, 13, 4, 7, 2, tree.NIL, tree.NIL, tree.NIL, tree.NIL, tree.NIL, 1}),
			22,
			true,
		},
		{
			tree.Build([]int{1}),
			1,
			true,
		},
		{
			tree.Build([]int{1}),
			2,
			false,
		},
	}

	for i, v := range examples {
		t.Run(fmt.Sprintf("case%d", i+1), func(t *testing.T) {
			res := hasPathSum(v.Input, v.Target)

			if res != v.Expect {
				t.Errorf("expect %t, got %t", v.Expect, res)
			}
		})
	}
}
