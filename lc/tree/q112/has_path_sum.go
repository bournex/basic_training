package q112

import "github.com/bournex/basic_training/lc/tree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *tree.TreeNode, targetSum int) bool {

	return has_path_sum(root, targetSum, 0)
}

func has_path_sum(root *tree.TreeNode, target int, current int) bool {
	var (
		bLeft  bool
		bRight bool
	)
	if root != nil {
		if root.Left == nil && root.Right == nil {
			return current+root.Val == target
		} else {
			if root.Left != nil {
				bLeft = has_path_sum(root.Left, target, current+root.Val)
			}

			if root.Right != nil {
				bRight = has_path_sum(root.Right, target, current+root.Val)
			}
		}
	}
	return bLeft || bRight
}
