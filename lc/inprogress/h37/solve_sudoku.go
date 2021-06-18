package h37

// https://leetcode-cn.com/problems/sudoku-solver/
// 思考
//	有效的数独一题里采用回溯法，用3*n的空间在横竖和小正方形空间内进行剪枝
//	这里的思路基本一致，
func solveSudoku(board [][]byte) {
	var vertical [9][9]bool
	var horizantal [9][9]bool
	var cube [9][9]bool

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '0' // 1-9
				vertical[i][num-1] = true
				horizantal[j][num-1] = true
				cube[i/3+3*(j/3)][num-1] = true
			}
		}
	}

	var backtrace func(int, int) bool
	backtrace = func(v, h int) bool {
		box := v/3 + 3*(h/3)
		if board[v][h] == '.' {
			for i := 0; i < 9; i++ {
				if vertical[v][i] || horizantal[h][i] || cube[box][i] {
					// 在该行、该列、该box中已经出现过该数字，则跳过该数字
					continue
				} else {
					var ret bool
					vertical[v][i], horizantal[h][i], cube[box][i] = true, true, true
					board[v][h] = '1' + byte(i)
					if v < 8 {
						ret = backtrace(v+1, h)
					} else if h < 8 {
						ret = backtrace(0, h+1)
					} else {
						return true
					}

					if ret {
						return true
					} else {
						board[v][h] = '.'
						vertical[v][i], horizantal[h][i], cube[box][i] = false, false, false
					}
				}
			}
		} else {
			var ret bool
			if v < 8 {
				ret = backtrace(v+1, h)
			} else if h < 8 {
				ret = backtrace(0, h+1)
			} else {
				return true
			}
			return ret
		}
		return false
	}
	backtrace(0, 0)
}
