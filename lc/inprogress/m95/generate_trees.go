package h95

import "github.com/bournex/basic_training/lc/tree"

// 思路
//	看起来跟m96题目类似，但是存在树节点拷贝的问题，如果回溯去解决的话，当所有节点都已经添加到树上之后
//	需要拷贝整棵树
//	看了一下官方解法，官方解法里没有考虑节点复用导致的问题，是自己想复杂了？貌似题目并不考虑返回数据怎么用的问题

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func generateTrees(n int) []*tree.TreeNode {

	var generate func(s, e int) []*tree.TreeNode
	generate = func(s, e int) []*tree.TreeNode {
		if s > e {
			return []*tree.TreeNode{
				nil,
			} // 这个操作确保了叶子节点也会被append
		}
		nodes := make([]*tree.TreeNode, 0)

		for i := s; i <= e; i++ {
			lefts := generate(s, i-1)
			rights := generate(i+1, e)

			for _, vl := range lefts {
				for _, vr := range rights {
					nodes = append(nodes, &tree.TreeNode{i, vl, vr})
				}
			}
		}
		return nodes
	}

	return generate(1, n)
}
