package m63

// 思路
//	参考m63，对于障碍物上方的节点，仅计算其右侧的路径数，对于障碍物左侧的节点，仅计算其下方的路径数
//	提交失败了四次，分别是：
//	1、没有考虑终点为障碍物的情形
//	2、没有考虑起点为障碍物的情形
//	3、没有考虑m-1行、n-1列有障碍物的情形
//	4、同3，最右列和最底行存在障碍物时，小于障碍物索引的列或行都不可达的情形
//	需要加强对应用提醒中的特殊用例的考量
// 下面这个实现的思路，是反向的递推
// 对比官方答案，官方定义dp[i][j]是到达点(i,j)的步数
// 而我的计算方式是dp[i][j]表达的是点(i,j)到达(m,n)终点的步数
// 思考两种思路的差异
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid) // 行数
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0]) // 列数
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		arr[i] = make([]int, n)
	}

	for i := m - 1; i >= 0; i-- {
		if obstacleGrid[i][n-1] == 0 {
			arr[i][n-1] = 1
		} else {
			break
		}
	}

	for i := n - 1; i >= 0; i-- {
		if obstacleGrid[m-1][i] == 0 {
			arr[m-1][i] = 1
		} else {
			break
		}
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			if obstacleGrid[i+1][j] == 1 {
				arr[i][j] = arr[i][j+1]
			} else if obstacleGrid[i][j+1] == 1 {
				arr[i][j] = arr[i+1][j]
			} else {
				arr[i][j] = arr[i+1][j] + arr[i][j+1]
			}
		}
	}

	return arr[0][0]
}

// 按照官方思路实现
// dp + 滚动数组
func uniquePathsWithObstaclesSingleLine(obstacleGrid [][]int) int {
	m := len(obstacleGrid) // 行数
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0]) // 列数
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	status := make([]int, n)
	// 初始化最后一行到status
	for i := 0; i < n && obstacleGrid[0][i] == 0; i++ {
		status[i] = 1
	}

	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 || status[0] == 0 {
			status[0] = 0
		}

		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				status[j] = status[j] + status[j-1]
			} else {
				status[j] = 0
			}
		}
	}

	return status[n-1]
}
