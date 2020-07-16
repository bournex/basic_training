package misc

import "fmt"

type VNode interface {
	Left() VNode
	Right() VNode
	ToString() string
}

// 先序遍历
func PreOrder(p VNode) {
	if p == nil {
		return
	}

	fmt.Println("pre-order traversing node", p.ToString())
	PreOrder(p.Left())
	PreOrder(p.Right())
}

// 中序遍历
func InOrder(p VNode) {
	if p == nil {
		return
	}

	PreOrder(p.Left())
	fmt.Println("in-order traversing node", p.ToString())
	PreOrder(p.Right())
}

// 后序遍历
func PostOrder(p VNode) {
	if p == nil {
		return
	}

	PreOrder(p.Left())
	PreOrder(p.Right())
	fmt.Println("post-order traversing node", p.ToString())
}
