package h51

// 优化回溯方案、空间预分配方案
func solveNQueens(n int) [][]string {
	v := make([]bool, n) // 垂直
	// h := make([]bool, n)      // 水平
	x1 := make([]bool, 2*n-1) // 左上到右下
	x2 := make([]bool, 2*n-1) // 右上到左下

	ans := make([][]byte, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]byte, n)
	}

	result := make([][]string, 0)

	cnt := 0

	// 回溯修改为按行安排Queen，因为每一行只能有一个Queen，所以水平的剪枝就不需要了
	var backtrace func(line int)
	backtrace = func(line int) {

		for i := 0; i < n; i++ {
			ans[line][i] = '.'
		}

		for i := 0; i < n; i++ {
			if !(v[i] || x1[i-line+n-1] || x2[line+i]) {
				ans[line][i] = 'Q'
				cnt++
				v[i], x1[i-line+n-1], x2[line+i] = true, true, true
				if cnt == n {
					// 到达棋盘右下角，且已经安排完所有的Queen，记录当前棋盘
					temp := make([]string, 0, n)
					for _, v := range ans {
						temp = append(temp, string(v))
					}
					result = append(result, temp)
				} else {
					backtrace(line + 1) // 去下一行
				}
				v[i], x1[i-line+n-1], x2[line+i] = false, false, false
				cnt--
				ans[line][i] = '.'
			}
		}
	}
	backtrace(0)

	return result
}

// 思路
//	最早的解决方案，回溯+剪枝，但是剪枝过程有冗余
func solveNQueens1(n int) [][]string {
	v := make([]bool, n)      // 垂直
	h := make([]bool, n)      // 水平
	x1 := make([]bool, 2*n-1) // 左上到右下
	x2 := make([]bool, 2*n-1) // 右上到左下

	ans := make([][]byte, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]byte, n)
	}

	result := make([][]string, 0)

	cnt := 0

	var backtrace func(i, j int)
	backtrace = func(i, j int) {

		if !(v[j] || h[i] || x1[j-i+n-1] || x2[i+j]) {
			ans[i][j] = 'Q'
			cnt++
			v[j], h[i], x1[j-i+n-1], x2[i+j] = true, true, true, true
			if j+1 < n {
				backtrace(i, j+1) // 后移
			} else if i+1 < n {
				backtrace(i+1, 0) // 换行
			} else if cnt == n {
				// 到达棋盘右下角，且已经安排完所有的Queen，记录当前棋盘
				var temp []string
				for _, v := range ans {
					temp = append(temp, string(v))
				}
				result = append(result, temp)
			}
			cnt--
			v[j], h[i], x1[j-i+n-1], x2[i+j] = false, false, false, false
		}

		ans[i][j] = '.'
		if j+1 < n {
			backtrace(i, j+1)
		} else if i+1 < n {
			backtrace(i+1, 0)
		} else if cnt == n {
			// 到达棋盘右下角，且已经安排完所有的Queen，记录当前棋盘
			var temp []string
			for _, v := range ans {
				temp = append(temp, string(v))
			}
			result = append(result, temp)
		}
	}
	backtrace(0, 0)

	return result
}
