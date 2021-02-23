package graph

// 深度优先遍历
// depth first search
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
