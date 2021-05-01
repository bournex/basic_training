package m64

// 思路
//	跟m62，m63相近的dp题目
func minPathSum(grid [][]int) int {
	m := len(grid)    // 行数
	n := len(grid[0]) // 列数

	status := make([]int, n)
	status[0] = grid[0][0]
	// 初始化最后一行到status
	for i := 1; i < n; i++ {
		status[i] = grid[0][i] + status[i-1]
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 1; i < m; i++ {
		status[0] += grid[i][0]
		for j := 1; j < n; j++ {
			status[j] = min(status[j], status[j-1]) + grid[i][j]
		}
	}

	return status[n-1]
}
