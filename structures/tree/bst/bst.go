package bst

import (
	"github.com/bournex/basic_training/structures/base"
	"github.com/bournex/basic_training/structures/queue"
)

// 二叉查找树实现

// ------------------------PUBLIC--------------

type VBst interface {
	// 成功返回nil
	// 失败返回error信息
	Insert(key base.Comparable, value interface{})
	// 成功返回指定value,nil
	// 未找到返回nil,nil
	// 失败返回nil,error信息
	Find(key base.Comparable) interface{}
	// 成功返回nil
	// 失败返回error信息
	Delete(key base.Comparable)
	// 树转换为字符串
	ToString() string
}

func MakeBst() VBst {
	return new(binarySearchTree)
}

//--------------------PRIVATE---------------
type bstNode struct {
	left  *bstNode
	right *bstNode
	key   base.Comparable
	value interface{}
}

func (node *bstNode) Compare(key base.Comparable) int {
	return node.key.Compare(key)
}

type binarySearchTree struct {
	root  *bstNode
	depth int
}

func (bst *binarySearchTree) Insert(key base.Comparable, value interface{}) {
	bst.root = bst.insertImpl(bst.root, key, value)
}

func (bst *binarySearchTree) insertImpl(p *bstNode, key base.Comparable, value interface{}) *bstNode {
	if p == nil {
		return &bstNode{key: key, value: value}
	}

	c := p.Compare(key)
	if c > 0 {
		p.left = bst.insertImpl(p.left, key, value)
	} else if c < 0 {
		p.right = bst.insertImpl(p.right, key, value)
	} else {
		p.value = value
	}
	return p
}

func (b *binarySearchTree) Find(key base.Comparable) interface{} {
	tmp := b.root
	for tmp != nil {
		c := tmp.Compare(key)
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

func (b *binarySearchTree) Delete(key base.Comparable) {
	// 分四种情况
	// 待删除节点左右子树都为空，直接删除该节点
	// 待删除节点有左子树，右子树为空；将左子树挂到父节点
	// 待删除节点有右子树，左子树为空；将右子树挂到父节点
	// 待删除节点有左、右子树均不为空；在右子树中找到最小值（或左子树的最大值），上升到待删除节点作为新节点
	b.root = b.deleteImpl(b.root, key)
}

func (b *binarySearchTree) deleteImpl(p *bstNode, key base.Comparable) *bstNode {
	if p == nil {
		return p
	}

	c := p.Compare(key)
	if c > 0 {
		p.left = b.deleteImpl(p.left, key)
	} else if c < 0 {
		p.right = b.deleteImpl(p.right, key)
	} else {
		if p.left == nil && p.right == nil {
			p = nil
		} else if p.left == nil && p.right != nil {
			p = p.right
		} else if p.left != nil && p.right == nil {
			p = p.left
		} else {
			tmp := b.min(p.right)
			tmp.right = b.removeMin(p.right)
			tmp.left = p.left
			p = tmp
		}
	}
	return p
}

func (b *binarySearchTree) min(p *bstNode) *bstNode {
	tmp := p
	for tmp != nil {
		if tmp.left == nil {
			break
		}
	}
	return tmp
}

func (b *binarySearchTree) removeMin(p *bstNode) *bstNode {
	// 递归向左遍历
	// 当到达树末梢时，必然有p.left为nil，此时p为最小节点，p的右节点可能有子树，将右子树替代p
	if p.left != nil {
		p.left = b.removeMin(p.left)
	}
	return p.right
}

func (b *binarySearchTree) ToString() string {
	var result string
	q := queue.MakeQueue(100)

	if b.root == nil {
		return "empty bst"
	}
	q.Enqueue(b.root)
	for {
		tmp, _ := q.Dequeue()
		p, b := tmp.(*bstNode)
		if b {
			if p.left != nil {
				q.Enqueue(p.left)
			}
			if p.right != nil {
				q.Enqueue(p.right)
			}
		}
		result += p.key.ToString()
		result += "_"

		if q.Length() == 0 {
			break
		}
	}

	return result
}
