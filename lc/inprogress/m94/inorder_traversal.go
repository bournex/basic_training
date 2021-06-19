package m94

import (
	"github.com/bournex/basic_training/lc/tree"
)

func inorderTraversal(root *tree.TreeNode) []int {
	ans := make([]int, 0)
	stack := make([]*tree.TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		for top.Left != nil {
			stack = append(stack, top.Left)
			top = top.Left
		}

		ans = append(ans, top.Val)
		stack = stack[:len(stack)-1]
		for {
			if top.Right != nil {
				// 压栈右子树
				stack = append(stack, top.Right)
				break
			} else if len(stack) > 0 {
				top = stack[len(stack)-1]
				ans = append(ans, top.Val)
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
	}

	return ans
}
