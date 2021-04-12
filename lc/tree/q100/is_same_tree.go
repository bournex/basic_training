package q100

import "github.com/bournex/basic_training/lc/tree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *tree.TreeNode, q *tree.TreeNode) bool {
	if p == nil || q == nil {
		if (p == nil && q != nil) || (p != nil && q == nil) {
			return false
		}
		return true
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
