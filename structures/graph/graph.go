package graph

type Relation struct {
	V1, V2 string // vertex
	E      int    // edge
}

// 一般常用邻接矩阵或邻接表来表达图，区别在于，当图比较稀疏时，使用邻接表表达图的空间利用率较高，当图比较稠密时
// 则比较适用邻接矩阵来表达，这里采用邻接表来实现图

// VGraph 图接口
type VGraph interface {
	Init([]*Relation)
	// DFS遍历判定连通性
	Connected(v1, v2 string) bool
	// BFS寻找最短路径
	ShortestPath(v1, v2 string) []string
}

// MakeGraph 创建图
func MakeGraph() VGraph {
	return new(graph)
}

type rnode struct {
	distance int
	vertex   string
}

type graph struct {
	nt map[string][]*rnode
}

func (g *graph) Init(rel []*Relation) {
	g.nt = make(map[string][]*rnode)
	for _, v := range rel {
		if g.nt[v.V1] == nil {
			g.nt[v.V1] = make([]*rnode, 0)
		}
		g.nt[v.V1] = append(g.nt[v.V1], &rnode{distance: v.E, vertex: v.V2})
		if g.nt[v.V2] == nil {
			g.nt[v.V2] = make([]*rnode, 0)
		}
		g.nt[v.V2] = append(g.nt[v.V2], &rnode{distance: v.E, vertex: v.V1})
	}
}

func (g *graph) Connected(v1, v2 string) bool {
	// DFS快速判定是否连通
	if g.nt[v1] == nil || g.nt[v2] == nil {
		return false
	}

	accessed := make(map[string]bool)
	path := g.dfs(v2, v1, accessed, 0)

	return len(path) != 0
}

func (g *graph) dfs(target, previous string, accessed map[string]bool, depth int) []string {
	if target == previous {
		// 如果找到target节点，则根据深度创建slice，保存路径
		path := make([]string, depth+1)
		path[depth] = target
		return path
	}

	var path []string
	// 将当前根节点标注为已访问
	accessed[previous] = true

	for _, v := range g.nt[previous] {
		// 避免环，判断是否已经访问过
		if _, exist := accessed[v.vertex]; exist {
			continue
		}

		// 深度优先向下访问
		path = g.dfs(target, v.vertex, accessed, depth+1)
		if path != nil {
			path[depth] = previous
			break
		}
	}
	return path
}

func (g *graph) ShortestPath(v1, v2 string) []string {
	return g.bfs(v1, v2)
}

func (g *graph) bfs(start, end string) []string {
	// 如果起终点相同，则返回起点
	if start == end {
		return []string{start}
	}

	// 如果起终点所在图不连通或有任意点不在当前图中，则返回不连通
	if g.nt[start] == nil || g.nt[end] == nil {
		return nil
	}

	// 用于bfs的FIFO
	c := make(chan string, len(g.nt))
	m := make(map[string]bool)
	edge := make(map[string]string)
	c <- start
	found := false

	for len(c) != 0 && !found {
		select {
		case v := <-c:

			for _, next := range g.nt[v] {
				if _, exist := m[next.vertex]; exist {
					continue
				}

				m[next.vertex] = true
				edge[next.vertex] = v

				if next.vertex == end {
					// 找到目标，不用再入
					found = true
					break
				}
				c <- next.vertex
			}

		default:
			close(c)
		}
	}

	if found {
		res := make([]string, 0)
		for end != start {
			res = append(res, end)
			end = edge[end]
		}
		res = append(res, start)

		l := len(res)
		for i := 0; i < l/2; i++ {
			tmp := res[i]
			res[i] = res[l-i-1]
			res[l-i-1] = tmp
		}
		return res
	}
	return nil
}
