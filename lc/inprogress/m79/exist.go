package m79

func exist(board [][]byte, word string) bool {
	m := len(board)    // 高
	n := len(board[0]) // 宽

	if len(word) > m*n { // 加速
		return false
	}

	placeholder := make([][]byte, m)
	for i := 0; i < m; i++ {
		placeholder[i] = make([]byte, n)
	}

	idx := 0 // 表示在word上的索引
	var backtrace func(i, j int) bool
	backtrace = func(i, j int) bool {
		if word[idx] != board[i][j] || placeholder[i][j] == 1 {
			return false
		}
		idx++
		defer func() {
			idx--
		}()
		if idx == len(word) {
			return true
		}

		placeholder[i][j] = 1
		defer func() {
			placeholder[i][j] = 0
		}()

		var exist bool
		if i < m-1 {
			exist = backtrace(i+1, j) || exist
		}
		if i > 0 {
			exist = backtrace(i-1, j) || exist
		}
		if j < n-1 {
			exist = backtrace(i, j+1) || exist
		}
		if j > 0 {
			exist = backtrace(i, j-1) || exist
		}
		return exist
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtrace(i, j) {
				return true
			}
		}
	}
	return false
}
