package rbt

import "github.com/bournex/basic_training/structures/base"

// Public----------------
type VRbt interface {
	// 成功返回nil
	// 失败返回error信息
	Insert(key base.Comparable, value interface{})
	// 成功返回指定value,nil
	// 未找到返回nil,nil
	// 失败返回nil,error信息
	Find(key base.Comparable) interface{}
	// 成功返回nil
	// 失败返回error信息
	Delete(key base.Comparable) error
	// 树转换为字符串
	ToString() string
}

func MakeRbt() VRbt {
	return new(redBlackTree)
}

// PRIVATE---------------
const (
	RED   = 0
	BLACK = 1
)

type rbtNode struct {
	left  *rbtNode
	right *rbtNode
	color int
	key   base.Comparable
	value interface{}
}

type redBlackTree struct {
	root  *rbtNode
	depth int
}

func (rbt *redBlackTree) rotateLeft(po *rbtNode) *rbtNode {
	tmp := po.right.left
	po.right.left = po
	pn := po.right
	po.right = tmp
	pn.color = po.color
	po.color = RED

	return pn
}

// 左侧二连红时，使用右旋
func (rbt *redBlackTree) rotateRight(pr *rbtNode) *rbtNode {
	pn := pr.left
	tmp := pn.right
	pr.left = tmp
	pn.right = pr
	pn.color = pr.color
	pn.left.color = BLACK
	pn.right.color = BLACK

	return pn
}

func (rbt *redBlackTree) flipColor(p *rbtNode) *rbtNode {
	p.left.color = BLACK
	p.right.color = BLACK
	p.color = RED
	return p
}

func (rbt *redBlackTree) Insert(key base.Comparable, value interface{}) {
	// 原则如下
	// 1、任意节点只能有一条红色链接
	// 2、红色链接只能出现在节点的左子节点，或与父节点的链接
	// 3、新加入节点的颜色都是红色的
	rbt.root = rbt.insertImpl(rbt.root, key, value)
	rbt.root.color = BLACK
}

func (p *rbtNode) IsRed() bool {
	return p.color == RED
}

func (p *rbtNode) IsBlack() bool {
	return p.color == BLACK
}

func (rbt *redBlackTree) insertImpl(p *rbtNode, key base.Comparable, value interface{}) *rbtNode {
	if p == nil {
		return &rbtNode{key: key, value: value, color: RED}
	}

	c := p.key.Compare(key)
	if c > 0 {
		p.left = rbt.insertImpl(p.left, key, value)
	} else if c < 0 {
		p.right = rbt.insertImpl(p.right, key, value)
	} else {
		p.value = value
		return p
	}

	if p.left.IsRed() && p.left.left.IsRed() {
		// 左二连红，违反原则1，需要右旋
		p = rbt.rotateRight(p)
	} else if p.left.IsBlack() && p.right.IsRed() {
		// 右红左黑，违反原则2，需要左旋
		p = rbt.rotateLeft(p)
	}
	if p.left.IsRed() && p.right.IsRed() {
		// 两侧为空，违反原则1，需要上翻
		p = rbt.flipColor(p)
	}
	return p
}

func (rbt *redBlackTree) Find(key base.Comparable) interface{} {
	if rbt.root == nil {
		return nil
	}

	tmp := rbt.root
	for tmp != nil {
		c := tmp.key.Compare(key)
		if c > 0 {
			tmp = tmp.left
		} else if c < 0 {
			tmp = tmp.right
		} else {
			return tmp.value
		}
	}
	return nil
}

func (rbt *redBlackTree) Delete(key base.Comparable) error {
	// TODO
	return nil
}

func (rbt *redBlackTree) ToString() string {
	// TODO
	return ""
}
