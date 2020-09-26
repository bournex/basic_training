package stack

import "errors"

const (
	chunk_size = 64 // 栈每次扩容的大小
	chunk_mask = (chunk_size - 1)
)

var (
	STACK_FULL  = errors.New("stack is full")
	STACK_EMPTY = errors.New("stack is empty")
)

// 不定长栈实现
// 实现思路，考虑可变长的情况，使用局部线性的chunk作为数据存储块
// chunk之间通过指针链接，在堆上分配chunk空间
type Stack interface {
	Push(interface{}) error
	Pop() (interface{}, error)
	Length() int
}

// maxlen - 栈大小，0为无限制
func MakeStack(maxlen int) Stack {

	s := new(stack)
	s.head, s.tail = nil, nil
	s.limit = maxlen
	s.len = 0
	return s
}

type chunk struct {
	prev   *chunk
	next   *chunk
	values []interface{}
}

type stack struct {
	head   *chunk // 头chunk
	tail   *chunk // 尾chunk
	limit  int    // 元素数量上限
	len    int    // 当前栈中元素数量
	chunks int    // chunk数量
}

// O(1) 压栈
func (s *stack) Push(val interface{}) error {
	if s.limit != 0 && s.len == s.limit {
		return STACK_FULL
	}

	if s.len&chunk_mask == 0 || s.tail == nil {
		// len已经是chunk_size的整数倍，说明当前栈的chunk已满，需要扩容
		// tail为空说明栈为空，需要创建第一个chunk
		c := new(chunk)
		c.prev = s.tail
		c.next = nil
		c.values = make([]interface{}, chunk_size)
		c.values[0] = val
		s.tail = c
		s.len++
		s.chunks++
		return nil
	}

	s.tail.values[s.len&chunk_mask] = val
	s.len++
	return nil
}

// O(1) 弹栈
func (s *stack) Pop() (val interface{}, err error) {
	if s.tail == nil {
		return nil, STACK_EMPTY
	}

	if s.len&chunk_mask == 1 {
		val = s.tail.values[0]
		s.tail = s.tail.prev
		s.len--
		s.chunks--
		return
	}

	s.len--
	return s.tail.values[s.len&chunk_mask], nil
}

func (s *stack) Length() int {
	return s.len
}
