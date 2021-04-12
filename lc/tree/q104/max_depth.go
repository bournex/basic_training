package q104

import "github.com/bournex/basic_training/lc/tree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	return max_depth(root, depth)
}

func max_depth(root *tree.TreeNode, depth int) int {
	if root == nil {
		return depth
	}

	depth++
	leftMax := max_depth(root.Left, depth)
	rightMax := max_depth(root.Right, depth)
	if leftMax > rightMax {
		return leftMax
	} else {
		return rightMax
	}
}
