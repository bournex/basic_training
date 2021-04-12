package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	INT_MAX = int(^uint(0) >> 1)
	INT_MIN = ^INT_MAX
)

func BuildTree1(array ...[]int) *TreeNode {
	sumArray := make([]int, 0)
	for _, v := range array {
		sumArray = append(sumArray, v...)
	}

	return Build(sumArray)
}

func Build(array []int) *TreeNode {
	n := len(array)

	if n <= 0 {
		return nil
	}

	array = append([]int{0}, array...)

	i := 1
	nodes := make([]*TreeNode, n+1)
	nodes[i] = &TreeNode{Val: array[1]}
	for k := i; k<<1 <= n; k++ {
		if k<<1 <= n && array[k<<1] != INT_MIN {
			nodes[k].Left = &TreeNode{Val: array[k<<1]}
			nodes[k<<1] = nodes[k].Left
		}
		if k<<1+1 <= n && array[k<<1+1] != INT_MIN {
			nodes[k].Right = &TreeNode{Val: array[(k<<1)+1]}
			nodes[(k<<1)+1] = nodes[k].Right
		}
	}

	return nodes[1]
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		if (p == nil && q != nil) || (p != nil && q == nil) {
			return false
		}
		return true
	}

	if p.Val != q.Val {
		return false
	}

	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}
