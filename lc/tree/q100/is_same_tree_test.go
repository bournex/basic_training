package q100

import "testing"

func TestIsSameTree(t *testing.T) {
	examples := []struct {
		InputA *TreeNode
		InputB *TreeNode
		Expect bool
	}{
		{
			nil,
			nil,
			true,
		},
		{
			&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}},
			&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}},
			true,
		},
		{
			&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}},
			&TreeNode{1, &TreeNode{3, nil, nil}, &TreeNode{2, nil, nil}},
			false,
		},
	}

	for _, v := range examples {
		res := isSameTree(v.InputA, v.InputB)
		if res != v.Expect {
			t.Errorf("input")
		}
	}
}
