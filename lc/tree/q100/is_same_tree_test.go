package q100

import (
	"fmt"
	"testing"

	"github.com/bournex/basic_training/lc/tree"
)

func TestIsSameTree(t *testing.T) {
	examples := []struct {
		InputA *tree.TreeNode
		InputB *tree.TreeNode
		Expect bool
	}{
		{
			nil,
			nil,
			true,
		},
		{
			tree.Build([]int{1, 2, 3}),
			tree.Build([]int{1, 2, 3}),
			true,
		},
		{
			tree.Build([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}), //&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}},
			tree.Build([]int{1, 3, 2, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}), //&TreeNode{1, &TreeNode{3, nil, nil}, &TreeNode{2, nil, nil}},
			false,
		},
	}

	for i, v := range examples {
		t.Run(fmt.Sprintf("case%d", i+1), func(t *testing.T) {
			res := isSameTree(v.InputA, v.InputB)
			if res != v.Expect {
				t.Errorf("input ")
			}
		})
	}
}
