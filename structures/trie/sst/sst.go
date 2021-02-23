package sst

// 字符串查找树
// String Search Trie
// 总体思路：
//		近似数的形式，每一个节点维护一个字母表的链接数组，因为任意字母的下一个字母也可能是任意一个字母
//		所以，如果字母表有m个，总共有n个节点的话，会占用m*n的空间
// 		除此之外，sst可以在常数时间内完成子串的查找
type VSst interface {
	Insert(string, interface{})
	Find(string) interface{}
	Delete(string)
}

// MakeSst - 创建sst
func MakeSst() VSst {
	return new(sst)
}

type sstNode struct {
	next  []*sstNode
	value interface{}
}

type sst struct {
	root *sstNode
}

func (s *sst) Insert(key string, value interface{}) {
	s.root = s.insert(s.root, key, value, 0)
}

func (s *sst) insert(p *sstNode, key string, value interface{}, d int) *sstNode {
	if p == nil {
		p = new(sstNode)
		p.next = make([]*sstNode, 256)
	}

	if len(key) == d {
		p.value = value
		return p
	}
	p.next[key[d]] = s.insert(p.next[key[d]], key, value, d+1)
	return p
}

func (s *sst) Find(key string) interface{} {
	return s.find(s.root, key, 0)
}

func (s *sst) find(p *sstNode, key string, d int) interface{} {
	if p == nil {
		return nil
	}

	if len(key) == d {
		return p.value
	}

	return s.find(p.next[key[d]], key, d+1)
}

func (s *sst) Delete(key string) {
	s.root = s.delete(s.root, key, 0)
}

func (s *sst) delete(p *sstNode, key string, d int) *sstNode {
	// 如果待查找的串没有在树中，则直接返回，不对树做任何处理
	// 如果待查找的串是末尾串，则删除末尾节点，并递归回退删除路径节点，删除路径节点前，需检查其next无值，且value为空

	if p == nil {
		return nil
	}
	if len(key) == d {
		p.value = nil
	} else {
		// 先递归遍历到key的末端，后序实现删除操作
		p.next[key[d]] = s.delete(p.next[key[d]], key, d+1)
	}

	for _, v := range p.next {
		if v != nil {
			return p
		}
	}
	// 节点如果有值，说明是某个key的末节点，不能删除，所以此处需要判定
	if p.value == nil {
		return nil
	}
	return p
}
