package q101

import (
	"fmt"
	"testing"

	"github.com/bournex/basic_training/lc/tree"
)

func TestIsSymmetric(t *testing.T) {
	examples := []struct {
		Input  *tree.TreeNode
		Expect bool
	}{
		{
			// [1,2,2,2,null,2]
			tree.Build([]int{1, 2, 2, 2, tree.INT_MIN, 2}),
			// &tree.TreeNode{1, &tree.TreeNode{2, &tree.TreeNode{Val: 2}, nil}, &tree.TreeNode{2, &tree.TreeNode{Val: 2}, nil}},
			false,
		},
		{
			&tree.TreeNode{1, &tree.TreeNode{Val: 2}, &tree.TreeNode{Val: 2}},
			true,
		},
		{
			&tree.TreeNode{1, &tree.TreeNode{Val: 2}, &tree.TreeNode{Val: 3}},
			false,
		},
		{
			tree.Build([]int{1, 2}),
			false,
		},
		{
			&tree.TreeNode{1, &tree.TreeNode{2, nil, &tree.TreeNode{Val: 3}}, &tree.TreeNode{2, nil, &tree.TreeNode{Val: 3}}},
			false,
		},
	}

	for i, v := range examples {
		t.Run(fmt.Sprintf("case%d", i+1), func(t *testing.T) {
			res := isSymmetric(v.Input)
			if res != v.Expect {
				t.Errorf("expect %t, got %t", v.Expect, res)
			}
		})
	}
}
