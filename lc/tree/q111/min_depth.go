package q111

import (
	"math"

	"github.com/bournex/basic_training/lc/tree"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 思路
//	测试用例里有个梗，根节点如果只有一个子树，另一个子树为空，不能按照根节点的深度计算
//	所以需要在递归里判断左右子树为空的情况
func minDepth(root *tree.TreeNode) int {
	return min_depth(root, 0)
}

func min_depth(root *tree.TreeNode, depth int) int {
	if root == nil {
		return depth
	}

	depth++
	left := min_depth(root.Left, depth)
	right := min_depth(root.Right, depth)

	if root.Right == nil {
		return left
	} else if root.Left == nil {
		return right
	}

	return int(math.Min(float64(left), float64(right)))
}
