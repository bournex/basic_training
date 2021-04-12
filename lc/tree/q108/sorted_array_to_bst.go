package q108

import "github.com/bournex/basic_training/lc/tree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *tree.TreeNode {
	n := len(nums)
	if n == 1 {
		return &tree.TreeNode{Val: nums[0]}
	}

	return to_bst(nums, 0, n-1)
}

func to_bst(nums []int, left, right int) *tree.TreeNode {
	if left > right {
		return nil
	}

	mid := left + (right-left+1)>>1
	return &tree.TreeNode{nums[mid], to_bst(nums, left, mid-1), to_bst(nums, mid+1, right)}
}
