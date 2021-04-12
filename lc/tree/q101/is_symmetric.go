package q101

import "github.com/bournex/basic_training/lc/tree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 官方还提到一个BFS解法，但是需要O(n)的空间复杂度
func isSymmetric(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}
	return isSubSymmetric(root.Left, root.Right)
}

func isSubSymmetric(left, right *tree.TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left == nil && right != nil) || (left != nil && right == nil) {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return isSubSymmetric(left.Left, right.Right) && isSubSymmetric(left.Right, right.Left)
}

// 下面是之前实现的缺陷方案
//	思路是中序遍历到数组，查看数组是否对称
//	缺陷在于数组长度未知的情况下，需要处理非完全二叉树的情形
//	并且空间复杂度为O(n)

/*
func isSymmetric(root *tree.TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}

	values := traverse(root)
	left := 0
	right := len(values) - 1
	for left < right {
		if values[left] != values[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func traverse(root *tree.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := make([]int, 0)
	left := traverse(root.Left)
	if len(left) != 0 {
		res = append(res, left...)
	}
	res = append(res, root.Val)
	right := traverse(root.Right)
	if len(right) != 0 {
		res = append(res, right...)
	}

	return res
}
*/
