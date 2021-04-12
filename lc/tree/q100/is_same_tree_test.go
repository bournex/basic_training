package q100

import (
	"fmt"
	"testing"
)

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
			buildTree(1, [][3]int{{1, 2, 3}}), //&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}},
			buildTree(1, [][3]int{{1, 3, 2}}), //&TreeNode{1, &TreeNode{3, nil, nil}, &TreeNode{2, nil, nil}},
			false,
		},
	}

	for i, v := range examples {
		t.Run(fmt.Sprintf("case%d", i+1), func(t *testing.T) {
			res := isSameTree(v.InputA, v.InputB)
			if res != v.Expect {
				t.Errorf("input")
			}
		})
	}
}

// n - root节点值
// array - 树关系，每个子数组
func buildTree(n int, array [][3]int) *TreeNode {
	var root *TreeNode
	nodes := make(map[int]*TreeNode)
	for _, v := range array {
		parentVal := v[0]
		var parent *TreeNode
		if p, ok := nodes[parentVal]; !ok {
			p = &TreeNode{Val: parentVal}
			nodes[parentVal] = p
			parent = p
		} else {
			parent = p
		}
		if parentVal == n {
			root = parent
		}

		if v[1] != -1 {
			if p, ok := nodes[v[1]]; !ok {
				p = &TreeNode{Val: v[1]}
				nodes[v[1]] = p
				parent.Left = p
			} else {
				parent.Left = p
			}
		}

		if v[2] != -1 {
			if p, ok := nodes[v[2]]; !ok {
				p = &TreeNode{Val: v[2]}
				nodes[v[2]] = p
				parent.Right = p
			} else {
				parent.Right = p
			}
		}
	}
	return root
}
