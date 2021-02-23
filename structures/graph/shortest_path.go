package graph

// 最短路径
func (g *graph) ShortestPath(v1, v2 string) []string {
	return g.bfs(v1, v2)
}
