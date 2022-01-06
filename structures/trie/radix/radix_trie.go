package radixtrie

import (
	"errors"
)

var inner *radix_trie

var (
	ErrRT_MissMatch = errors.New("miss match")
)

// tree node type
type node_type int

const (
	NT_Root     = iota + 1 // root
	NT_Instance            // instance
	NT_Variable            // variable
)

type tagdata struct {
	fullpath string
	prev     *node
	next     *node
	custom   interface{}
}

type node struct {
	tp       node_type // node type
	instance []*node   // instance children
	variable []*node   // variable children
	path     string    // current node path segment
	indices  string    // instance children prefix characters, accelerating comparation
	fullpath string    // fullpath, only appear at the end node of a path
}

type MatchResult struct {
	pattern  string
	variable map[string]string
}

type radix_trie struct {
	trees map[string]*node
}

func Match(method, path string) ([]string, error) {
	rt := inner
	if n := rt.trees[method]; n != nil {
		if res := n.match(path); res != nil {
			return res, nil
		} else {
			// method no found
			return nil, ErrRT_MissMatch
		}
	}
	return nil, ErrRT_MissMatch
}

// match pattern and extract variables
func MatchVariable(method, path string) ([]MatchResult, error) {
	rt := inner
	if n := rt.trees[method]; n != nil {
		if res := n.match(path); res != nil {
			var mr []MatchResult
			for _, pattern := range res {
				mr = append(mr, MatchResult{
					pattern:  pattern,
					variable: extract(pattern, path),
				})
				return mr, nil
			}
		} else {
			// method no found
			return nil, ErrRT_MissMatch
		}
	}
	return nil, ErrRT_MissMatch
}

// extract variables from path
func extract(pattern, path string) map[string]string {
	variables := make(map[string]string)
	var i, j int
	for i < len(pattern) && j < len(path) {
		if pattern[i] == path[j] {
			i, j = i+1, j+1
			continue
		}

		if pattern[i] != '{' {
			// illegal pattern or pattern & path miss match
			return nil
		}

		var key, value string

		// find key
		var idx int
		for k := i; k < len(pattern); k++ {
			if pattern[k] == '}' {
				idx = k
				break
			}
		}
		if idx != 0 && idx+1 == len(pattern) {
			// end of pattern
			key = pattern[i+1 : idx]
		} else if idx > 0 {
			key = pattern[i+1 : idx]
		} else {
			// missing }, illegal pattern
			return nil
		}
		i = idx + 1

		// find value
		idx = 0
		for k := j; k < len(path); k++ {
			if path[k] == '/' {
				idx = k
				break
			}
		}
		if idx == 0 {
			// end of path
			value = path[j:]
		} else if idx > 0 {
			value = path[j:idx]
		}
		j = idx

		variables[key] = value
	}

	return variables
}

func (rt radix_trie) dup() *radix_trie {
	newtrie := &radix_trie{
		trees: make(map[string]*node, len(rt.trees)),
	}
	for k, v := range rt.trees {
		newtrie.trees[k] = v.dup()
	}
	return newtrie
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func longestCommonPrefix(a, b string) int {
	i := 0
	max := min(len(a), len(b))
	for i < max && a[i] == b[i] && a[i] != '{' && b[i] != '{' {
		i++
	}
	return i
}

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

type stackNode struct {
	n *node  // current node
	i int    // variable children node index
	p string // path when we reached stackNode.n
	c bool   // is matching instance child
}

func (sn *stackNode) next() *node {
	if sn.i+1 < len(sn.n.variable) {
		sn.i++
		return sn.n.variable[sn.i]
	}
	return nil
}

type vstack []*stackNode

// push stack
func (v *vstack) push(sn *stackNode) {
	*v = append(*v, sn)
}

// pop stack
func (v *vstack) pop() *stackNode {
	sn := (*v)[len(*v)-1]
	*v = (*v)[:len(*v)-1]
	return sn
}

// get stack top
func (v vstack) top() *stackNode {
	if len(v) == 0 {
		return nil
	}
	return v[len(v)-1]
}

func (n *node) addRoute(path string) {
	fullPath := path

	if n.path == "" && n.indices == "" {
		n.insertChild(path, fullPath)
		n.tp = NT_Root
		return
	}

walk:
	for {
		i := longestCommonPrefix(path, n.path)

		if i < len(n.path) {
			child := node{
				path:     n.path[i:],
				tp:       NT_Instance,
				indices:  n.indices,
				instance: n.instance,
				variable: n.variable,
				fullpath: n.fullpath,
			}

			n.instance = []*node{&child}
			n.variable = nil
			n.indices = string([]byte{n.path[i]})
			n.path = path[:i]
			n.fullpath = ""
		}

		if i < len(path) {
			path = path[i:]
			idxc := path[0]

			if idxc != '{' {
				for i, c := range []byte(n.indices) {
					if c == idxc {
						n = n.instance[i]
						continue walk
					}
				}

				n.indices += string([]byte{idxc})
				child := &node{tp: NT_Instance}
				n.instance = append(n.instance, child)
				n = child
			}
			n.insertChild(path, fullPath)
			return
		}

		// n.path equal to path，n is the end point of insert
		n.fullpath = fullPath
		return
	}
}

func (n *node) insertChild(path, fullPath string) {
	for {
		// find {} variable
		wildcard, i := findVariable(path)
		if i < 0 { // no variable
			break
		}

		// cut prefix const string off
		if i > 0 {
			n.path = path[:i] // save prefix
			path = path[i:]   //
		}

		// create variable node
		child := &node{
			tp:   NT_Variable,
			path: wildcard,
		}

		n.variable = append(n.variable, child)
		n.tp = NT_Instance
		n = child

		// if variable node is not the end of path, create const string node for the rest part
		if len(wildcard) < len(path) {
			path = path[len(wildcard):]
			child := &node{
				tp:   NT_Instance,
				path: path,
			}
			// after {}, it must be /
			n.instance = append(n.instance, child)
			n.indices += string([]byte{path[0]})
			n = child
			continue
		}

		// reached the end of path
		break
	}

	//
	n.path = path
	n.fullpath = fullPath
}

// Non-recursive backtracking match
func (n *node) match(path string) []string {
	// create a stack for saving intermediate nodes
	var result []string
	stack := vstack(make([]*stackNode, 0, 20))

walk:
	for {
		prefix := n.path
		var idxc byte
		if n.tp == NT_Variable {
			i := 0
			// skip variable part
			for i < len(path) && path[i] != '/' {
				i++
			}

			// the tail of path, fully matched
			if i == len(path) && len(n.fullpath) != 0 {
				result = append(result, n.fullpath)
				goto loopback
			}

			// cut path off
			path = path[i:]
			if len(path) == 0 {
				// path end early，current branch match failed
				goto loopback
			}
			idxc = path[0]
		} else if len(path) > len(prefix) {
			if path[:len(prefix)] == prefix {
				// prefix matched
				path = path[len(prefix):]
				idxc = path[0]
			} else {
				// current branch match failed, go back
				goto loopback
			}
		} else if path == prefix {
			// the remaining constant string fully matched
			if len(n.fullpath) != 0 {
				result = append(result, n.fullpath)
				goto loopback
			}
		}

		// current node matched, look into children nodes
		// try instance
		for i, c := range []byte(n.indices) {
			if c == idxc {
				// save current node as context, go deep
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

		// try variable
		if len(n.variable) > 0 {
			// save current node as context, go deep
			stack.push(&stackNode{
				n: n,
				i: 0,
				p: path,
			})
			n = n.variable[0]
			continue walk
		}

		// current path match finished
	loopback:
		top := stack.top()
		for top != nil {
			// instance branch match failed, try variables
			if top.c && len(top.n.variable) > 0 {
				top.i = 0
				top.c = false
				n = top.n.variable[0]
				continue walk
			}

			// try next variable node
			if next := top.next(); next != nil {
				n = next
				path = top.p
				continue walk
			}

			// no variable node left, pop stack
			stack.pop()
			top = stack.top()
		}
		// stack is nil, match finished
		break
	}
	return result
}

// duplicate node
func (n *node) dup() *node {
	d := &node{
		tp:       n.tp,
		path:     n.path,
		fullpath: n.fullpath,
		indices:  n.indices,
	}
	for _, v := range n.instance {
		d.instance = append(d.instance, v.dup())
	}
	for _, v := range n.variable {
		d.variable = append(d.variable, v.dup())
	}
	return d
}
