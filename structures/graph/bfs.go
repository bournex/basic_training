package graph

// 广度优先遍历
// breadth first search
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
