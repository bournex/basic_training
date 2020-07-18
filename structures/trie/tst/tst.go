package tst

// VTst - Ternary Search Tree
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
