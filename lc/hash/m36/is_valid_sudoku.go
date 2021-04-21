package m36

// 思考
//	此题是看答案看出来的
//	原以为是类似八皇后的回溯算法，或者也想到了使用数组保存行列状态
//	但是感觉太复杂，或空间使用过多，没有想到更优的解决办法
//	看了下答案，发现想多了

func isValidSudoku(board [][]byte) bool {
	v := len(board)
	h := len(board[0])

	var vf [9][9]bool
	var hf [9][9]bool
	var box [9][9]bool

	for i := 0; i < v; i++ {
		for j := 0; j < h; j++ {
			if board[i][j] == '.' {
				continue
			}

			x := (i/3)*3 + j/3
			if vf[i][board[i][j]-'0'-1] || hf[j][board[i][j]-'0'-1] || box[x][board[i][j]-'0'-1] {
				return false
			}

			vf[i][board[i][j]-'0'-1] = true
			hf[j][board[i][j]-'0'-1] = true
			box[x][board[i][j]-'0'-1] = true
		}
	}

	return true
}
