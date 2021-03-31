package attach

// 基础：判断8x8棋盘中，是否有相互攻击的车
// 进阶：判断8x8棋盘中，是否有相互攻击的车或皇后

// 定义输入
//  数值为0表示棋盘上没有棋子
//  数值为1表示棋盘上有车
//  数值为2表示棋盘上有皇后

// 思考
//  如果车互相没有攻击，说明任意的车不在同一行或列
//  由于任意行列都可能存在车，则至少需要遍历棋盘一次
//  从一侧顺序遍历，已遍历过的一侧则不需要再检验
//  此问题属于八皇后的变体问题

// 基础
// 根据以上思考，判断是否有车互相攻击代码如下

// Attack 车攻击判断
func attackCart(matrix [][]int) bool {
	h := cap(matrix)
	w := cap(matrix[0])

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if matrix[i][j] == 1 {
				// 有车
				for k := j + 1; k < w; k++ {
					if matrix[i][k] == 1 {
						return true
					}
				}
				for k := j - 1; k >= 0; k-- {
					if matrix[i][k] == 1 {
						return true
					}
				}
				for k := i + 1; k < h; k++ {
					if matrix[k][j] == 1 {
						return true
					}
				}
			}
		}
	}

	return false
}

// 使用水平和垂直数组作为状态记录
// 空间换时间，使复杂度从N²下降到N
func attackCart1(matrix [][]int) bool {
	h := cap(matrix)
	w := cap(matrix[0])

	vertical := make([]int, w)
	horizontal := make([]int, h)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if matrix[i][j] == 1 && (horizontal[i] == 1 || vertical[j] == 1) {
				return true
			}
			horizontal[i], vertical[j] = matrix[i][j], matrix[i][j]
		}
	}

	return false
}

func attackQueen(matrix [][]int) bool {
	h := cap(matrix)
	w := cap(matrix[0])

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if matrix[i][j] != 0 {
				// i,j处有车
				for k := j + 1; k < w; k++ {
					if matrix[i][k] != 0 {
						return true
					}
				}
				for k := j - 1; k >= 0; k-- {
					if matrix[i][k] != 0 {
						return true
					}
				}
				for k := i + 1; k < h; k++ {
					if matrix[k][j] != 0 {
						return true
					}
				}

				if matrix[i][j] == 2 {
					// 特别地，如果有i，j处是Queen，则判断斜向
					for m, n := i+1, j+1; m < h && n < w; {
						if matrix[m][n] != 0 {
							return true
						}
						m++
						n++
					}
					for m, n := i+1, j-1; m < h && n >= 0; {
						if matrix[m][n] != 0 {
							return true
						}
						m++
						n--
					}
				}
			}
		}
	}
	return false
}

func attackQueen1(matrix [][]int) bool {
	h := cap(matrix)
	w := cap(matrix[0])

	vertical := make([]int, w)      // 垂直
	horizontal := make([]int, h)    // 水平
	slash := make([]int, w+h-1)     // 左上到右下
	backslash := make([]int, w+h-1) // 右上到左下

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if matrix[i][j] == 1 {
				if horizontal[j] == 1 || vertical[i] == 1 || slash[i+j] == 1 || backslash[h+i-j-1] == 1 {
					return true
				} else {
					horizontal[j], vertical[i] = matrix[i][j], matrix[i][j]
					slash[i+j], backslash[h+i-j-1] = matrix[i][j], matrix[i][j]
				}
			}
		}
	}
	return false
}
