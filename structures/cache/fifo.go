package cache

import "fmt"

type VFifo interface {
	Init(int64)
	Get(interface{}) interface{}
	Set(interface{}, interface{})
	ToString() string
}

func MakeFifo(max int64) VFifo {
	f := new(fifo)
	f.Init(max)
	return f
}

type fifoNode struct {
	key   interface{}
	value interface{}
}

// FIFO实现
// 固定队列长度，先进入的最先被淘汰
// 使用线型空间表达，环形队列形式，head总是指向下一个要写的位置
// 当写发生时，如果head处有值，则调整为要写入的值，即覆盖
type fifo struct {
	max    int64 // 最大容量
	head   int64 // 队列头，插入位置
	buffer []*fifoNode
	access map[interface{}]*fifoNode
}

func (f *fifo) Init(max int64) {
	f.max = max
	f.head = 0
	f.buffer = make([]*fifoNode, f.max)
	f.access = make(map[interface{}]*fifoNode)
}

func (f *fifo) Set(key, value interface{}) {
	if v, ok := f.access[key]; ok {
		v.value = value
		return
	}

	node := f.buffer[f.head]
	if node != nil {
		delete(f.access, node.key)
		node.key = key
		node.value = value

	} else {
		node = &fifoNode{
			key:   key,
			value: value,
		}
		f.buffer[f.head] = node
	}

	f.access[key] = node
	// 更新head
	if f.head+1 == f.max {
		f.head = 0
	} else {
		f.head++
	}
}

func (f *fifo) Get(key interface{}) interface{} {
	if v, ok := f.access[key]; ok {
		return v.value
	}
	return nil
}

func (f *fifo) ToString() string {
	var (
		ret  string
		move int64
	)
	move = f.head - 1
	if move < 0 {
		move = f.max - 1
	}

	for {
		node := f.buffer[move]
		if node != nil {
			ret += fmt.Sprintf("%+v -> %+v\n", node.key, node.value)
			move--
			if move < 0 {
				move = f.max - 1
			}
		} else {
			break
		}
		if move == f.head && f.buffer[f.head] != nil {
			ret += fmt.Sprintf("%+v -> %+v\n", f.buffer[f.head].key, f.buffer[f.head].value)
			break
		}
	}
	return ret
}
