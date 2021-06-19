package m98

import (
	"fmt"
	"testing"

	"github.com/bournex/basic_training/lc/tree"
)

func TestIsValidBst(t *testing.T) {
	testCases := []struct {
		Input  *tree.TreeNode
		Expect bool
	}{
		{
			tree.Build([]int{2, 1, 3}),
			true,
		},
		{
			tree.Build([]int{1, 2, 3}),
			false,
		},
		{
			tree.Build([]int{5, 4, 6, tree.NIL, tree.NIL, 3, 7}),
			false,
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			res := isValidBST(v.Input)
			if res != v.Expect {
				t.Errorf("expect %t, got %t", v.Expect, res)
			}
		})
	}
}
