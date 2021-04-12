package q104

import (
	"fmt"
	"testing"

	"github.com/bournex/basic_training/lc/tree"
)

func TestMaxDepth(t *testing.T) {
	examples := []struct {
		Input  *tree.TreeNode
		Expect int
	}{
		{
			tree.Build([]int{1, 2, 3, tree.INT_MIN, tree.INT_MIN, 4, 5}),
			3,
		},
		{
			tree.Build([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}),
			4,
		},
		{
			tree.Build([]int{1}),
			1,
		},
		{
			nil,
			0,
		},
	}

	for i, v := range examples {
		t.Run(fmt.Sprintf("case%d", i+1), func(t *testing.T) {
			res := maxDepth(v.Input)
			if res != v.Expect {
				t.Errorf("expect %d, got %d", v.Expect, res)
			}
		})
	}
}
