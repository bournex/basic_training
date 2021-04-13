package q111

import (
	"fmt"
	"testing"

	"github.com/bournex/basic_training/lc/tree"
)

func TestMinDepth(t *testing.T) {
	examples := []struct {
		Input  *tree.TreeNode
		Expect int
	}{
		{
			tree.Build([]int{1, 2, 3, tree.NIL, tree.NIL, 4, 5}),
			2,
		},
		{

			tree.Build([]int{1}),
			1,
		},

		{
			tree.Build([]int{1, 2, tree.NIL, 3, 4}),
			3,
		},
		{
			nil,
			0,
		},
	}

	for i, v := range examples {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			res := minDepth(v.Input)
			if res != v.Expect {
				t.Errorf("expect %d, got %d", v.Expect, res)
			}
		})
	}
}
