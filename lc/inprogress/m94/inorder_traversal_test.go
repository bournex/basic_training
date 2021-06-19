package m94

import (
	"fmt"
	"testing"

	"github.com/bournex/basic_training/lc/tree"
)

func TestInorderTraversal(t *testing.T) {
	testCases := []struct {
		Input  *tree.TreeNode
		Expect []int
	}{
		{
			tree.Build([]int{1, tree.NIL, 2, tree.NIL, tree.NIL, 3}),
			[]int{1, 3, 2},
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			res := inorderTraversal(v.Input)
			if len(res) != len(v.Expect) {
				t.Errorf("expect %d len, got %d len", len(v.Expect), len(res))
				return
			}

			for j, v1 := range res {
				if v1 != v.Expect[j] {
					t.Errorf("index %d, expect %d, got %d", j, v.Expect[j], res[j])
					return
				}
			}
		})
	}
}
