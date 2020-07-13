package misc

import "fmt"

type VNode interface {
	Left() VNode
	Right() VNode
	ToString() string
}

func PreOrder(p VNode) {
	if p == nil {
		return
	}

	fmt.Println("pre-order traversing node", p.ToString())
	PreOrder(p.Left())
	PreOrder(p.Right())
}

func InOrder(p VNode) {
	if p == nil {
		return
	}

	PreOrder(p.Left())
	fmt.Println("in-order traversing node", p.ToString())
	PreOrder(p.Right())
}

func PostOrder(p VNode) {
	if p == nil {
		return
	}

	PreOrder(p.Left())
	PreOrder(p.Right())
	fmt.Println("post-order traversing node", p.ToString())
}
