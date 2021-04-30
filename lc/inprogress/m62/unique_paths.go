package m62

import "container/list"

// 思路
//	DFS或BFS
//	曾尝试用DFS或BFS，但是都超时了，应该是花了太多时间在列表处理或出入栈
//	考虑到任意一点的路数等于右侧点到终点的路数加上下方点到终点的路数只和，所以从右下角向左上角做逐级加和即可
//	官方答案里说这是DP，从右下角的延展来看，像是一个倒放的杨辉三角
//	感觉图问题大多可以转化为DP问题，或者说这种每一步骤都存在多种选择的情况时，这类问题就类似DP问题
func uniquePaths(m int, n int) int {
	return reverseFillForm(m, n)
}

func uniquePathsDFS(m int, n int) int {
	arr := make([]int, m*n)

	var goNext func(int, int) int
	goNext = func(i, j int) int {

		if i+1 == m && j+1 == n {
			return 1
		}

		var down, right int
		if i+1 < m && arr[(i+1)*n+j] == 0 {
			down = goNext(i+1, j)
		}
		if j+1 < n && arr[i*n+(j+1)] == 0 {
			right = goNext(i, j+1)
		}

		return down + right
	}
	return goNext(0, 0)
}

func uniquePathsBFS(m int, n int) int {
	arr := make([]int, m*n)

	var paths int
	l := list.New()
	l.PushBack([]int{0, 0})
	for l.Len() > 0 {
		k := l.Front().Value.([]int)
		if k[0]+1 == m && k[1]+1 == n {
			paths++
		} else if arr[k[0]*n+k[1]] > 0 {

		} else {
			if k[0]+1 < m {
				l.PushBack([]int{k[0] + 1, k[1]})
			}
			if k[1]+1 < n {
				l.PushBack([]int{k[0], k[1] + 1})
			}
		}
		l.Remove(l.Front())
	}
	return paths
}

// 思路
//	反向填表

func reverseFillForm(m int, n int) int {
	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		arr[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		arr[i][n-1] = 1
	}

	for i := 0; i < n; i++ {
		arr[m-1][i] = 1
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			arr[i][j] = arr[i+1][j] + arr[i][j+1]
		}
	}

	return arr[0][0]
}
