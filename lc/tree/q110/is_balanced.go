package q110

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
//	递归遍历，参数为下级节点和下级节点深度，返回当前节点下最深子节点，以及当前节点下是否平衡
// 总结
//  自底向上、自顶向下递归？
func isBalanced(root *tree.TreeNode) bool {
	_, balance := is_balanced(root, 0)
	return balance
}

func is_balanced(root *tree.TreeNode, depth int) (int, bool) {
	if root == nil {
		return depth, true
	}

	depth += 1
	n1, b1 := is_balanced(root.Left, depth)
	n2, b2 := is_balanced(root.Right, depth)

	return int(math.Max(float64(n1), float64(n2))), int(math.Abs(float64(n1-n2))) <= 1 && b1 && b2
}
