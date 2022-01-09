package radixtrie

import (
	"errors"
	"sync/atomic"
)

var (
	ErrRT_TrieUninit  = errors.New("trie uninit")
	ErrRT_MissMatch   = errors.New("miss match")
	ErrRT_PathIllegal = errors.New("illegal path")
)

// 树节点
type node_type int

const (
	NT_Root     = iota + 1 // 根节点
	NT_Static              // 实例节点
	NT_Variable            // 参数节点
)

const (
	NTS_Pause = -1
	NTS_Serve = 0
	NTS_Swap  = 1
)

type tagdata struct {
	fullpath string
	next     *tagdata
	prev     *tagdata
	custom   interface{}
}

// trie树节点
type node struct {
	tp       node_type // 当前节点类型
	instance []*node   // 实例节点
	variable []*node   // 参数节点
	path     string    // 当前节点路径
	indices  string    // 子节点首字母表，只存储非参数子节点（instance）的首字母
	tag      *tagdata  // 末端节点的全路径
}

type MatchResult struct {
	pattern  string
	variable map[string]string
}

// 回溯栈节点对象
type stackNode struct {
	n *node  // 当前节点
	i int    // 子节点数组索引
	p string // path，用于跟n.variable匹配
	c bool   // 是否instance匹配失败
}

// 获取stackNode.n.variable中的下一个node
func (sn *stackNode) next() *node {
	if sn.i+1 < len(sn.n.variable) {
		sn.i++
		return sn.n.variable[sn.i]
	}
	return nil
}

type vstack []*stackNode

// 压栈
func (v *vstack) push(sn *stackNode) {
	*v = append(*v, sn)
}

// 弹栈
func (v *vstack) pop() *stackNode {
	sn := (*v)[len(*v)-1]
	*v = (*v)[:len(*v)-1]
	return sn
}

// 获取栈顶
func (v vstack) top() *stackNode {
	if len(v) == 0 {
		return nil
	}
	return v[len(v)-1]
}

// trie树
type radix_trie struct {
	trees   atomic.Value // index指向的map用于流量端查找
	taghead *tagdata
}

func NewTrie() *radix_trie {
	rt := &radix_trie{
		taghead: new(tagdata),
	}
	rt.trees.Store(new(node))
	return rt
}

// 使用nt原子替换rt（nt - New Trie、 rt - Runnig Trie）
func (rt *radix_trie) AtomicReplace(nt *radix_trie) {
	rt.trees.Store(nt.trees.Load())
}

// 如果需要对trie做批量/增量替换，可以先Copy，在副本中添加完增量路由后，原子替换原来的rt
func (rt *radix_trie) Copy() *radix_trie {
	n := rt.trees.Load().(*node)
	td := new(tagdata)
	replicate := n.dup(td)

	nt := &radix_trie{
		taghead: td,
	}
	nt.trees.Store(replicate)
	return nt
}

// 只能向一个没有在Serve状态的radix_trix添加route
func (rt *radix_trie) AddRoute(pattern ...string) *radix_trie {
	tree := rt.trees.Load().(*node)

	for _, v := range pattern {
		tree.addRoute(v, rt.taghead)
	}
	return rt
}

type PathLink interface {
	Path() string
	Next() PathLink
}

func (td tagdata) Path() string {
	return td.fullpath
}

func (td tagdata) Next() PathLink {
	if td.next != nil {
		return td.next
	}
	return nil
}

func (rt *radix_trie) First() PathLink {
	return rt.taghead.next
}

// 匹配路径
func (rt *radix_trie) Match(path string) ([]string, error) {
	tree := rt.trees.Load().(*node)

	result := tree.match(path)
	if len(result) == 0 {
		return nil, ErrRT_MissMatch
	}

	return result, nil
}

// 匹配路径并取得路径参数
func (rt *radix_trie) MatchVariable(path string) ([]MatchResult, error) {
	tree := rt.trees.Load().(*node)

	res := tree.match(path)
	if len(res) > 0 {
		var mr []MatchResult
		for _, pattern := range res {
			mr = append(mr, MatchResult{
				pattern:  pattern,
				variable: extract(pattern, path),
			})
		}
		return mr, nil
	}

	return nil, ErrRT_MissMatch
}

// 提取参数
func extract(pattern, path string) map[string]string {
	variables := make(map[string]string)
	var i, j int
	for i < len(pattern) && j < len(path) {
		if pattern[i] == path[j] {
			i, j = i+1, j+1
			continue
		}

		if pattern[i] != '{' {
			// 非法pattern或pattern与path并不匹配
			return nil
		}

		var key, value string

		// 查找key
		var idx int
		for k := i; k < len(pattern); k++ {
			if pattern[k] == '}' {
				idx = k
				break
			}
		}
		if idx != 0 && idx+1 == len(pattern) {
			// 结尾
			key = pattern[i+1 : idx]
		} else if idx > 0 {
			key = pattern[i+1 : idx]
		} else {
			// 右大括号缺失，pattern格式非法
			return nil
		}
		i = idx + 1

		// 查找value
		idx = 0
		for k := j; k < len(path); k++ {
			if path[k] == '/' {
				idx = k
				break
			}
		}
		if idx == 0 {
			// 结尾
			value = path[j:]
		} else if idx > 0 {
			value = path[j:idx]
		}
		j = idx

		variables[key] = value
	}

	return variables
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// 参数节点不计入公共前缀
func longestCommonPrefix(a, b string) int {
	i := 0
	max := min(len(a), len(b))
	for i < max && a[i] == b[i] && a[i] != '{' && b[i] != '{' {
		i++
	}
	return i
}

// 查找参数节点，获得参数节点值和之前、之后的节点路径
// 如果path中包含参数，比如/A/{param}/B
// variable - {param}
// i - 3
func findVariable(path string) (variable string, i int) {
	// Find start
	for start, c := range []byte(path) {
		if c != '{' {
			continue
		}

		for end, c := range []byte(path[start+1:]) {
			if c == '}' {
				return path[start : start+1+end+1], start
			}
		}
		break
	}
	return "", -1
}

// 添加path
//
func (n *node) addRoute(path string, tag *tagdata) {
	fullPath := path

	// Empty tree
	if n.path == "" && n.indices == "" {
		n.insertChild(path, fullPath, tag)
		n.tp = NT_Root
		return
	}

walk:
	for {
		// 找到待插入path与当前节点path的公共前缀，遇到参数节点会提前返回
		i := longestCommonPrefix(path, n.path)

		// 公共前缀长度小于当前节点n.path，裂变当前节点为父子节点
		if i < len(n.path) {
			child := node{
				path:     n.path[i:],
				tp:       NT_Static,
				indices:  n.indices,
				instance: n.instance,
				variable: n.variable,
				tag:      n.tag,
			}

			n.instance = []*node{&child}
			n.variable = nil
			n.indices = string([]byte{n.path[i]})
			n.path = path[:i]
			n.tag = nil
		}

		// 公共前缀长度小于剩余path，剩余部分
		if i < len(path) {
			path = path[i:]
			idxc := path[0]

			if idxc != '{' {
				// 尝试查找同前缀n.instance子节点
				for i, c := range []byte(n.indices) {
					if c == idxc {
						n = n.instance[i]
						continue walk
					}
				}

				// 没有同前缀instance子节点，创建新的子节点，并执行insertChild
				n.indices += string([]byte{idxc})
				child := &node{tp: NT_Static}
				n.instance = append(n.instance, child)
				n = child
			} else {
				// 尝试查找同名n.variable子节点
				variable, _ := findVariable(path)

				for _, vchild := range n.variable {
					if variable == vchild.path {
						path = path[len(variable):]
						// 此时path要么是空串，要么是"/"开头的内容
						if len(path) > 0 {
							idxc = path[0]
							for i, k := range []byte(vchild.indices) {
								if k == idxc {
									// 继续向下查找
									n = vchild.instance[i]
									continue walk
								}
							}
							vchild.indices += string([]byte{idxc})
							child := &node{tp: NT_Static}
							vchild.instance = append(vchild.instance, child)
							n = child
						} else {
							// 空串，到达末端
							n = vchild
							goto finish
						}

						continue walk
					}
				}
			}
			n.insertChild(path, fullPath, tag)
			return
		}
	finish:
		// n.path与path相等，当前n就是输入path的末端节点
		td := &tagdata{
			fullpath: fullPath,
			next:     tag.next,
		}
		if td.next != nil {
			td.next.prev = td
		}
		tag.next, td.prev = td, tag
		n.tag = td
		return
	}
}

// 在当前node下插入path
func (n *node) insertChild(path, fullPath string, tag *tagdata) {
	for {
		// 查找路径中的参数
		wildcard, i := findVariable(path)
		if i < 0 { // 路径中没有参数
			break
		}

		// 解析参数
		if i > 0 {
			n.path = path[:i] // 保存前缀常量串到当前node
			path = path[i:]   // 取得path的其他部分
		}

		// 将参数部分构造一个新的子节点
		child := &node{
			tp:   NT_Variable,
			path: wildcard,
		}

		n.variable = append(n.variable, child)
		n.tp = NT_Static
		n = child

		// 参数不是末尾，将当前节点修改为参数节点，并继续迭代
		if len(wildcard) < len(path) {
			path = path[len(wildcard):]
			child := &node{
				tp:   NT_Static,
				path: path,
			}
			// 之所以赋值给instance子节点slice，是因为规则中假定不会出现{PARAM1}{PARAM2}的情形
			// 连续参数之间至少会有斜线分割，即{PARAM1}/{PARAM2}，所以{PARAM1}下一个节点必然是instance
			n.instance = append(n.instance, child)
			n.indices += string([]byte{path[0]})
			n = child
			continue
		}

		// 到达path的末尾，结束
		break
	}

	// 将路径与全路径赋值给末端n节点
	n.path = path

	td := &tagdata{
		fullpath: fullPath,
		next:     tag.next,
	}
	if td.next != nil {
		td.next.prev = td
	}
	tag.next, td.prev = td, tag
	n.tag = td
}

func (n *node) match(path string) []string {
	// 栈长度设置为20，大部分路径深度不会达到这个数量
	var result []string
	stack := vstack(make([]*stackNode, 0, 20))

walk:
	for {
		prefix := n.path
		var idxc byte
		if n.tp == NT_Variable {
			i := 0
			// 当前节点是参数节点，跳过参数部分
			for i < len(path) && path[i] != '/' {
				i++
			}

			// 到达结尾且末端节点为合法路径，当前分支匹配结束
			if i == len(path) && n.tag != nil {
				result = append(result, n.tag.fullpath)
				goto loopback
			}

			// 截断path
			path = path[i:]
			if len(path) == 0 {
				// path匹配在非末端参数节点提前结束，当前分支匹配结束
				goto loopback
			}
			idxc = path[0]
		} else if len(path) > len(prefix) {
			if path[:len(prefix)] == prefix {
				// 前缀吻合当前节点
				path = path[len(prefix):]
				idxc = path[0]
			} else {
				// 当前路径匹配失败，需要回退
				goto loopback
			}
		} else if path == prefix {
			// 剩余路径完全匹配，当前分支匹配结束
			if n.tag != nil {
				result = append(result, n.tag.fullpath)
				goto loopback
			}
		}

		// 执行到这里，说明已经完成了当前节点的匹配，还需要在当前节点的子节点中继续匹配
		// 尝试在instance列表中匹配
		for i, c := range []byte(n.indices) {
			if c == idxc {
				// go deep
				stack.push(&stackNode{
					n: n,
					i: 0,
					p: path,
					c: true,
				})
				n = n.instance[i]
				continue walk
			}
		}

		// 尝试在variable列表中匹配
		if len(n.variable) > 0 {
			// 保存现场
			stack.push(&stackNode{
				n: n,
				i: 0,
				p: path,
			})
			n = n.variable[0]
			continue walk
		}

		// 当前路径匹配结束，回退
	loopback:
		top := stack.top()
		for top != nil {
			// instance分支匹配失败，尝试variable分支
			if top.c && len(top.n.variable) > 0 {
				top.i = 0
				top.c = false
				n = top.n.variable[0]
				continue walk
			}
			// variable下top.i分支匹配失败，尝试迭代top.i+1
			if next := top.next(); next != nil {
				n = next
				path = top.p
				continue walk
			}

			// top节点的variable已经完成遍历，出栈
			stack.pop()
			top = stack.top()
		}
		// stack为空，结束匹配
		break
	}
	return result
}

// 复制node
func (n *node) dup(head *tagdata) *node {
	d := &node{
		tp:      n.tp,
		path:    n.path,
		indices: n.indices,
	}

	if head != nil {
		if n.tag != nil {
			td := &tagdata{
				fullpath: n.tag.fullpath,
				next:     head.next,
			}
			if td.next != nil {
				td.next.prev = td
			}
			head.next, td.prev = td, head
			d.tag = td
		}
	} else {
		d.tag = n.tag
	}

	for _, v := range n.instance {
		d.instance = append(d.instance, v.dup(head))
	}
	for _, v := range n.variable {
		d.variable = append(d.variable, v.dup(head))
	}
	return d
}
