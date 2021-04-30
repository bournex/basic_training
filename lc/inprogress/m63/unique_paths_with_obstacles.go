package m63

// 思路
//	参考m63，对于障碍物上方的节点，仅计算其右侧的路径数，对于障碍物左侧的节点，仅计算其下方的路径数
//	提交失败了四次，分别是：
//	1、没有考虑终点为障碍物的情形
//	2、没有考虑起点为障碍物的情形
//	3、没有考虑m-1行、n-1列有障碍物的情形
//	4、同3，最右列和最底行存在障碍物时，小于障碍物索引的列或行都不可达的情形
//	需要加强对应用提醒中的特殊用例的考量
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
	for i := n - 2; i >= 0; i-- {
		if obstacleGrid[m-1][i] == 0 {
			status[i] = 1
		} else {
			break
		}
	}

	if obstacleGrid[m-2][n-1] != 1 {
		status[n-1] = 1
	}

	for i := m - 2; i >= 0; i-- {
		if obstacleGrid[i][n-1] == 1 {
			status[n-1] = 0
		}

		for j := n - 2; j >= 0; j-- {
			status[j] = status[j] + status[j+1]
		}
	}

	return status[0]
}
