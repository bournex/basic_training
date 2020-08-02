package tst

// VTst - Ternary Search Tree
// 思路
// 		属于sst的变体，当数据量上升时、或字母表扩大时，sst空间占用过大
//		tst是三向字符串查找树，基本思路与三向快排类似，将当前节点区分为等于、大于、小于当前字母的三叉树形式
//		对于等于当前节点字母的，继续向下查找，如果大于或小于当前节点值，则向左右子树继续查找，此时待查单词的
//		索引不递增
//		tst树的深度比sst增加，查找比较次数增多，但空间使用骤降。
type VTst interface {
	Insert(string, interface{})
	Find(string) interface{}
	Delete(string)
}

// MakeTst - 创建Tst
func MakeTst() VTst {
	return new(tst)
}

type tstNode struct {
	left  *tstNode
	next  *tstNode
	right *tstNode
	k     byte
	value interface{}
}

type tst struct {
	root *tstNode
}

func (t *tst) Insert(key string, value interface{}) {
	if key == "" || value == nil {
		return
	}
	t.root = t.insert(t.root, key, value, 0)
}

func (t *tst) insert(p *tstNode, key string, value interface{}, d int) *tstNode {
	if p == nil {
		p = new(tstNode)
		p.k = key[d]
	}

	if len(key)-1 > d {
		if p.k > key[d] {
			p.left = t.insert(p.left, key, value, d)
		} else if p.k < key[d] {
			p.right = t.insert(p.right, key, value, d)
		} else {
			p.next = t.insert(p.next, key, value, d+1)
		}
	} else {
		// d == len(key)-1
		// 由于d作为key数组索引，最大只能等于len(key)-1，说明到达key的末端
		p.value = value
	}

	return p
}

func (t *tst) Find(key string) interface{} {
	if key == "" {
		return nil
	}

	return t.find(t.root, key, 0)
}

func (t *tst) find(p *tstNode, key string, d int) interface{} {
	if p == nil {
		return nil
	}

	if len(key)-1 == d {
		// 这里有两种语义，p.value有值和无值
		return p.value
	}
	if p.k > key[d] {
		return t.find(p.left, key, d)
	} else if p.k < key[d] {
		return t.find(p.right, key, d)
	}
	return t.find(p.next, key, d+1)
}

func (t *tst) Delete(key string) {
	if key == "" {
		return
	}
	t.root = t.delete(t.root, key, 0)
}

func (t *tst) delete(p *tstNode, key string, d int) *tstNode {
	if p == nil {
		return p
	}

	if len(key)-1 == d {
		// 逻辑正确的情况下，不会出现即没有value，又没有left、right、next的情形
		p.value = nil
	} else {
		if p.k > key[d] {
			p.left = t.delete(p.left, key, d)
		} else if p.k < key[d] {
			p.right = t.delete(p.right, key, d)
		} else {
			p.next = t.delete(p.next, key, d+1)
		}
	}

	// 如果没有子节点，则删除该节点，否则保留该节点
	if p.left == nil && p.right == nil && p.next == nil {
		return nil
	}
	return p
}
