package graph

// 连通性判断

func (g *graph) Connected(v1, v2 string) bool {
	// DFS快速判定是否连通
	if g.nt[v1] == nil || g.nt[v2] == nil {
		return false
	}

	accessed := make(map[string]bool)
	path := g.dfs(v2, v1, accessed, 0)

	return len(path) != 0
}
