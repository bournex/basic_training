package skiplist

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/bournex/basic_training/structures/base"
)

type VSkipList interface {
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
	//
	ToString() string
}

func MakeSkipList(maxlvl int) VSkipList {
	rand.Seed(time.Now().Unix())
	p := new(skipList)
	p.constructSkipList(maxlvl)
	return p
}

type skipNode struct {
	key   base.Comparable
	value interface{}
	level int // 从1开始计数，用作循环数时需要减一
	next  []*skipNode
}

// 总体思路
// 每次插入节点时，随机生成节点的级别，级别代表指针数量，然后从高到低逐级加入到链表中
// 查找时首先从高级别进行检索，逐级缩小检索范围，类似n分查找
// 跳表达不到严格的平衡，最差情况会退化成纯链表的O(n)复杂度，但其具备一定的平衡性，且实现比一般平衡树简单
// 当最大级别为m，元素数量最大为n时，如果有2^m > n，则可能会出现最优性能O(lgn)
type skipList struct {
	// 跳表的最高层级，不宜过高，指针占用空间过多
	maxlvl int
	// 跳表链表头
	head []*skipNode
	// 各级链表中元素数量
	cnt []int
}

func (sl *skipList) constructSkipList(maxlvl int) {
	sl.head = make([]*skipNode, maxlvl)
	sl.cnt = make([]int, maxlvl)
	sl.maxlvl = maxlvl
}

func (sl *skipList) randomLevel(maxlvl int) int {
	lvl := maxlvl
	n := rand.Intn(1 << maxlvl)
	for n > 0 {
		lvl--
		n = n >> 1
	}
	if lvl == 0 {
		lvl = 1
	}
	return lvl
}

func (sl *skipList) randomNode(level int, key base.Comparable, value interface{}) *skipNode {
	// 随机生成0-9的数字，其中：
	// 0 	代表4级节点，包含4个指针
	// 12	代表3级节点，包含3个指针
	// 345	代表2级节点，包含2个指针
	// 6789	代表1级节点，包含1个指针

	p := new(skipNode)
	p.level = level
	p.next = make([]*skipNode, p.level)
	p.key = key
	p.value = value

	return p
}

func (sl *skipList) Insert(key base.Comparable, value interface{}) {
	lvl := sl.randomLevel(sl.maxlvl)

	ps := make([]*skipNode, lvl)
	for i := 0; i < cap(ps); i++ {
		ps[i] = sl.head[i]
	}

	for i := lvl - 1; i >= 0; i-- {
		pn := ps[i]
		var previous *skipNode

		for pn != nil {
			c := pn.key.Compare(key)
			if c > 0 {
				ps[i] = previous
				break
			} else if c == 0 {
				pn.value = value
				return
			}

			previous = pn
			pn = pn.next[i]
		}

		ps[i] = previous

		if ps[i] != nil && i > 0 {
			// 即将查找下一级时，不必从头开始查，从上一级的待插入位置开始查即可
			ps[i-1] = ps[i]
		}
	}

	p := sl.randomNode(lvl, key, value)

	for i := 0; i < len(ps); i++ {
		if ps[i] != nil {
			p.next[i] = ps[i].next[i]
			ps[i].next[i] = p
		} else {
			if sl.head[i] != nil {
				p.next[i] = sl.head[i]
			}
			sl.head[i] = p
		}
	}
}

func (sl *skipList) Find(key base.Comparable) interface{} {

	index := sl.maxlvl - 1

	ps := make([]*skipNode, sl.maxlvl)
	ps[index] = sl.head[index]

	for i := index; i >= 0; i-- {
		pn := ps[i]
		var previous *skipNode

		for pn != nil {
			c := pn.key.Compare(key)
			if c > 0 {
				ps[i] = previous
				break
			} else if c == 0 {
				return pn.value
			}

			previous = pn
			pn = pn.next[i]
		}

		ps[i] = previous

		if ps[i] != nil && i > 0 {
			// 即将查找下一级时，不必从头开始查，从上一级的待插入位置开始查即可
			ps[i-1] = ps[i]
		}
	}

	return nil
}

// 脱链时先找到在每一级上，小于节点值的上界指针，遍历完成删除
func (sl *skipList) Delete(key base.Comparable) {

	index := sl.maxlvl - 1

	var previous *skipNode
	var pn *skipNode

	for i := index; i >= 0; i-- {
		if previous == nil {
			pn = sl.head[i]
		} else {
			pn = previous
		}

		for pn != nil {
			c := pn.key.Compare(key)
			if c > 0 {
				break
			} else if c == 0 {
				if previous != nil {
					previous.next[i] = previous.next[i].next[i]
				} else {
					sl.head[i] = pn.next[i]
				}
				break
			}
			previous = pn
			pn = pn.next[i]
		}
	}
}

func (sl *skipList) ToString() string {
	var result string
	for i := 0; i < sl.maxlvl; i++ {
		p := sl.head[i]
		result += "level " + strconv.Itoa(i+1) + " : "
		for p != nil {
			result += "->"
			result += p.key.ToString()
			p = p.next[i]
		}
		result += "\n"
	}

	return result
}
