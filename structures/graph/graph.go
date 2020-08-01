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
	Connected(v1, v2 string) bool
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
	// BFS判断最短路径（之一）
	return nil
}

func (g *graph) bfs(target string) []string {
	return nil
}
