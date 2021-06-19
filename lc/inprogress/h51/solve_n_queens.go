package h51

func solveNQueens(n int) [][]string {
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

		// if i == n && j == n && cnt == n {
		// 	// 到达棋盘右下角，且已经安排完所有的Queen，记录当前棋盘
		// 	var temp []string
		// 	for _, v := range ans {
		// 		temp = append(temp, string(v))
		// 	}
		// 	result = append(result, temp)
		// }
	}
	backtrace(0, 0)

	return result
}
