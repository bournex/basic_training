package cache

import "fmt"

// LRU - 最近最少用缓存
type VLru interface {
	Init(int64)
	Get(interface{}) interface{}
	Set(interface{}, interface{})
	Del(interface{})
	ToString() string
}

func MakeLru(max int64) VLru {
	if max <= 0 {
		max = 1
	}
	l := new(lru)
	l.Init(max)
	return l
}

type lruNode struct {
	prev  *lruNode    // head方向前一个节点
	next  *lruNode    // tail方向后一个节点
	key   interface{} // node key
	value interface{} // node 值
}

type lru struct {
	max  int64 // 最大容量
	size int64 // 当前使用量

	access map[interface{}]*lruNode // 快速访问hash
	head   *lruNode                 // 新插入入口
	tail   *lruNode                 // 淘汰点
}

func (l *lru) Init(max int64) {
	l.max = max
	l.size = 0
	l.access = make(map[interface{}]*lruNode)
	l.head = nil
	l.tail = nil
}

func (l *lru) Get(key interface{}) interface{} {
	if v, ok := l.access[key]; ok {
		if v.prev != nil {
			// 从当前位置脱链
			if v.next != nil {
				v.next.prev = v.prev
				v.prev.next = v.next
			} else {
				l.tail = v.prev
				v.prev.next = nil
			}
			l.head.prev = v
			v.next = l.head
			l.head = v
		} else {
			// 已经是头结点，不做任何处理
		}

		return v.value
	}

	return nil
}

func (l *lru) Set(key, value interface{}) {

	if v, ok := l.access[key]; ok {
		// 更新值
		v.value = value
	} else {
		// 写入值
		if l.size == l.max && l.tail != nil {
			// 已满，尾部脱链
			node := l.tail
			if node.prev != nil {
				node.prev.next = nil
				node.prev = nil
			} else {
				// 说明已经是队列的最后一个元素
				l.head = nil
				l.tail = nil
			}
			// 从快速访问hash移除
			delete(l.access, node.key)
		} else {
			// 增加计数
			l.size++
		}

		// 新建node
		v := &lruNode{
			next:  l.head,
			prev:  nil,
			key:   key,
			value: value,
		}
		// 添加到快速访问hash
		l.access[key] = v

		// 将淘汰链表头指向刚写入的lruNode
		if l.head != nil {
			l.head.prev = v
		}
		l.head = v
		if l.tail == nil {
			l.tail = v
		}
	}
}

func (l *lru) Del(key interface{}) {

	if v, ok := l.access[key]; ok {
		if v.prev != nil {
			if v.next != nil {
				// 中间节点
				v.next.prev = v.prev
				v.prev.next = v.next
			} else {
				// 尾节点
				l.tail = v.prev   // 修改tail指向
				v.prev.next = nil // 倒数第二个节点next置空
			}
		} else {
			// 是头
			if v.next != nil {
				// 不是最后一个元素
				v.next.prev = nil
				l.head = v.next
			} else {
				// 是最后一个元素
				l.head, l.tail = nil, nil
			}
		}
		v.prev, v.next = nil, nil
		// 递减计数
		l.size--
		// 从快速访问hash移除
		delete(l.access, v.key)
	}
}

func (l *lru) ToString() string {
	var (
		ret  string
		head = l.head
	)

	for head != nil {
		ret += fmt.Sprintf("%+v -> %+v\n", head.key, head.value)
		head = head.next
	}

	return ret
}
