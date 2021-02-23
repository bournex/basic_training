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
