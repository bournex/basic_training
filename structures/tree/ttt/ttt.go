package ttt

import "github.com/bournex/basic_training/structures/base"

// 2-3树实现
// two-three tree
type VTtt interface {
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

func MakeTtt() VTtt {
	return new(Ttt)
}

const (
	TwoNode   = 1
	ThreeNode = 2
)

type tNode struct {
	left  *tNode
	mid   *tNode
	right *tNode

	key1   base.Comparable
	value1 interface{}

	key2   base.Comparable
	value2 interface{}

	typ int // TwoNode or ThreeNode
}

type Ttt struct {
	// 2-3树满足以下性质
	// 任意节点可能包含一个key，也可能包含两个key，
	// 包含一个key的节点称为2节点（因为有两个孩子），包含两个key的节点称为3节点
	// 对于2节点，左子树所有节点值小于当前节点，右子树所有节点值大于当前节点
	// 对于3节点，左子树所有节点值小于3节点中最小值，右子树所有节点值大于3节点中最大值，中间子树所有节点值介于3节点两值之间
	root *tNode
}

// 考虑所有可能导致树结构发生变化的情形
// 1.当向3节点中插入节点时，分3种情况
//		① 当3节点为根节点时，插入使根节点变为"4节点“，此时分裂根节点为3个2节点，中间值2几点指向另外两个2几点
// 		② 当3节点父节点为2节点时，将临时4节点中，分裂为左右两个2节点，中间值上移到父节点变成3节点
// 		③ 当3节点父节点为3节点时，与②操作近似，父节点变为临时4几点，递归对父节点执行此操作
// 2.当删除一个2节点值时

func (t *Ttt) Insert(key base.Comparable, value interface{}) {
	t.root = t.insertImpl(t.root, key, value)
}

func (t *Ttt) mergeNode(p, q *tNode) *tNode {
	// 将n2中的key value合并到n1上
	// 如果n1是3节点，则返回新产出的节点
	// 如果n1是2节点，则返回nil
	// 注意待merge的q一定是2节点
	if q == nil {
		return p
	}

	if p.typ == TwoNode {
		c := p.key1.Compare(q.key1)
		if c > 0 {
			p.key2, p.value2 = p.key1, p.value1
			p.key1, p.value1 = q.key1, q.value1
			p.mid = q.right
			p.left = q.left
		} else if c < 0 {
			p.key2, p.value2 = q.key1, q.value2
			p.mid = q.left
			p.right = q.right
		}
		return nil
	} else if p.typ == ThreeNode {
		p.typ = TwoNode
		c1 := p.key1.Compare(q.key1)
		c2 := p.key2.Compare(q.key1)

		r := &tNode{key1: p.key2, value1: p.value2, typ: TwoNode}

		if c1 > 0 {
			// q<p<r,rise p up
			r.left, r.right = p.mid, p.right
			p.left, p.right = q, r
			p.mid = nil
			return p
		} else if c1 < 0 && c2 > 0 {
			// p<q<r,rise q up
			p.right, r.left = q.left, q.right
			q.left, q.right = p, r
			p.mid = nil
			return q
		} else if c2 < 0 {
			// p<r<q,rise r up
			p.right = p.mid
			p.mid = nil
			r.left, r.right = p, q
			return r
		}
	}
	return nil
}

func (t *Ttt) insertImpl(p *tNode, key base.Comparable, value interface{}) *tNode {
	if p == nil {
		return &tNode{key1: key, value1: value}
	}

	var (
		tmp *tNode
	)

	if p.key1 != nil {
		c := p.key1.Compare(key)
		if c > 0 {
			// key比key1、key2都小
			tmp = t.insertImpl(p.left, key, value)
		} else if c < 0 {
			if p.typ == TwoNode {
				tmp = t.insertImpl(p.right, key, value)
			}
		} else {
			p.value1 = value
			return p
		}
	}
	if p.key2 != nil {
		// key比key1、key2都大
		c := p.key2.Compare(key)
		if c < 0 {
			tmp = t.insertImpl(p.right, key, value)
		} else if c > 0 {
			if p.typ == TwoNode {
				tmp = t.insertImpl(p.left, key, value)
			} else {
				tmp = t.insertImpl(p.mid, key, value)
			}
		} else {
			p.value2 = value
			return p
		}
	}

	return t.mergeNode(p, tmp)
}

func (t *Ttt) Find(key base.Comparable) interface{} {
	p := t.root

	for p != nil {
		c1 := p.key1.Compare(key)
		if c1 == 0 {
			return p.value1
		} else if c1 > 0 {
			p = p.left
		} else {
			if p.typ == ThreeNode {
				c2 := p.key2.Compare(key)
				if c2 == 0 {
					return p.value2
				} else if c2 < 0 {
					p = p.right
				} else {
					p = p.mid
				}
			}
		}
	}

	return nil
}

func (t *Ttt) Delete(key base.Comparable) {

}

func (t *Ttt) ToString() string {
	return ""
}
