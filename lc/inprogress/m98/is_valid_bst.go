package m98

import "github.com/bournex/basic_training/lc/tree"

// 思路
//	拷贝的m94的答案，稍作修改，主要是考察中序遍历
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *tree.TreeNode) bool {
	ans := make([]int, 0)
	stack := make([]*tree.TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		for top.Left != nil {
			stack = append(stack, top.Left)
			top = top.Left
		}

		if len(ans) > 0 {
			if ans[len(ans)-1] < top.Val {
				ans = append(ans, top.Val)
			} else {
				return false
			}
		} else {
			ans = append(ans, top.Val)
		}

		stack = stack[:len(stack)-1]
		for {
			if top.Right != nil {
				// 压栈右子树
				stack = append(stack, top.Right)
				break
			} else if len(stack) > 0 {
				top = stack[len(stack)-1]
				if len(ans) > 0 {
					if ans[len(ans)-1] < top.Val {
						ans = append(ans, top.Val)
					} else {
						return false
					}
				} else {
					ans = append(ans, top.Val)
				}
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
	}

	return true
}
